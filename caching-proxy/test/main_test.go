package test

import (
	"testing"
	"time"

	"github.com/xXMolinaXx/golang/caching-proxy/cache"
)

func TestCleanupCacheExpired(t *testing.T) {
	var cachedUrls []cache.CachedUrl = []cache.CachedUrl{
		{URL: "/test1", CacheValue: []byte(`{"message": "test1"}`), ResetCache: time.Now().Add(-10 * time.Second)},
		{URL: "/test2", CacheValue: []byte(`{"message": "test2"}`), ResetCache: time.Now().Add(-10 * time.Second)},
	}
	cache.CleanupCache(&cachedUrls)
	if len(cachedUrls) != 0 {
		t.Errorf("expected 0 cached URLs, got %d", len(cachedUrls))
	}
}

func TestCleanupCacheNotExpired(t *testing.T) {
	var cachedUrls []cache.CachedUrl = []cache.CachedUrl{
		{URL: "/test1", CacheValue: []byte(`{"message": "test1"}`), ResetCache: time.Now().Add(10 * time.Second)},
		{URL: "/test2", CacheValue: []byte(`{"message": "test2"}`), ResetCache: time.Now().Add(10 * time.Second)},
	}
	cache.CleanupCache(&cachedUrls)
	if len(cachedUrls) != 2 {
		t.Errorf("expected 2 cached URLs, got %d", len(cachedUrls))
	}
}

func TestCleanupCacheOneExpireOneNotExpired(t *testing.T) {
	var cachedUrls []cache.CachedUrl = []cache.CachedUrl{
		{URL: "/test1", CacheValue: []byte(`{"message": "test1"}`), ResetCache: time.Now().Add(10 * time.Second)},
		{URL: "/test2", CacheValue: []byte(`{"message": "test2"}`), ResetCache: time.Now().Add(-10 * time.Second)},
	}
	cache.CleanupCache(&cachedUrls)
	if len(cachedUrls) != 1 {
		t.Errorf("expected 1 cached URL, got %d", len(cachedUrls))
	}
}

func TestFindIndex(t *testing.T) {
	var cachedUrls []cache.CachedUrl = []cache.CachedUrl{
		{URL: "/test1", CacheValue: []byte(`{"message": "test1"}`), ResetCache: time.Now().Add(10 * time.Second)},
		{URL: "/test2", CacheValue: []byte(`{"message": "test2"}`), ResetCache: time.Now().Add(10 * time.Second)},
	}
	index := cache.FindIndex(&cachedUrls, "/test1")
	if index != 0 {
		t.Errorf("expected index 0, got %d", index)
	}

	index = cache.FindIndex(&cachedUrls, "/test2")
	if index != 1 {
		t.Errorf("expected index 1, got %d", index)
	}

	index = cache.FindIndex(&cachedUrls, "/test3")
	if index != -1 {
		t.Errorf("expected index -1, got %d", index)
	}
}
