package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/xXMolinaXx/golang/caching-proxy/cache"
)

var originURL string
var alreadyCached []cache.CachedUrl

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("Received request for %s %s\nContent-Type: %s\nAccept: %s\nQuery: %s\n", r.URL.Path, r.Method, r.Header.Get("Content-Type"), r.Header.Get("Accept"), r.URL.Query().Encode())
	if originURL == "" {
		http.Error(w, "origin flag is required", http.StatusBadRequest)
		return
	}
	valueCachedFound := cache.FindIndex(&alreadyCached, r.URL.Path)

	if valueCachedFound != -1 {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Cache", "HIT")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(alreadyCached[valueCachedFound].CacheValue); err != nil {
			log.Println("error writing cached response body:", err)
		}
		return
	}
	req, err := http.NewRequest(r.Method, originURL+r.URL.Path, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
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

	alreadyCached = append(alreadyCached, cache.CachedUrl{URL: r.URL.Path, CacheValue: body, ResetCache: time.Now().Add(10 * time.Second)})

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
