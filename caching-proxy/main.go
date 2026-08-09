package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type cachedUrl struct {
	URL        string
	CacheValue []byte
	resetCache time.Time
}

var originURL string
var alreadyCached []cachedUrl

func findIndex(slice []cachedUrl, url string) int {
	for i, v := range slice {
		if v.URL == url {
			if time.Now().After(v.resetCache) {
				return -1
			}
			return i
		}
	}
	return -1
}

func handler(w http.ResponseWriter, r *http.Request) {
	if originURL == "" {
		http.Error(w, "origin flag is required", http.StatusBadRequest)
		return
	}
	valueCachedFound := findIndex(alreadyCached, r.URL.Path)

	if valueCachedFound != -1 {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Cache", "HIT")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(alreadyCached[valueCachedFound].CacheValue); err != nil {
			log.Println("error writing cached response body:", err)
		}
		return
	}

	resp, err := http.Get(originURL + r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	alreadyCached = append(alreadyCached, cachedUrl{URL: r.URL.Path, CacheValue: body, resetCache: time.Now().Add(10 * time.Second)})

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Cache", "MISS")
	w.WriteHeader(resp.StatusCode)
	if _, err := w.Write(body); err != nil {
		log.Println("error writing response body:", err)
	}
}

func main() {
	fmt.Println("Iniciando el proxy de caching")
	port := flag.Int("port", 8080, "Port to run the proxy on")
	origin := flag.String("origin", "", "Origin server URL")
	flag.Parse()
	originURL = *origin
	if originURL == "" {
		log.Fatal("flag -origin is required")
	}

	addr := fmt.Sprintf(":%d", *port)
	fmt.Printf("Starting proxy on port %d, forwarding to origin %s\n", *port, originURL)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(addr, nil))
}
