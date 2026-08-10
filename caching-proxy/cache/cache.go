package cache

import "time"

type CachedUrl struct {
	URL        string
	CacheValue []byte
	ResetCache time.Time
}

func CleanupCache(cachedUrls *[]CachedUrl) {
	now := time.Now()
	var cleanCache []CachedUrl

	for _, v := range *cachedUrls {
		if now.Before(v.ResetCache) {
			cleanCache = append(cleanCache, v)
		}
	}

	*cachedUrls = cleanCache

}
func FindIndex(cachedUrls *[]CachedUrl, url string) int {
	CleanupCache(cachedUrls)
	foundIndex := -1

	for index, v := range *cachedUrls {
		if foundIndex == -1 && v.URL == url {
			foundIndex = index
		}
	}
	return foundIndex
}
