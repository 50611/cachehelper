package lcache

import (
	lru "github.com/50611/golang-lru/v2"
	"sync"
)

type CacheRw struct {
	cache *lru.Cache[string, *sync.RWMutex]
}

func NewCacheRw(cnt int) *CacheRw {
	c, _ := lru.New[string, *sync.RWMutex](cnt)

	return &CacheRw{cache: c}
}

func (l *CacheRw) GetOrAdd(key string) *sync.RWMutex {

	_, v := l.cache.GetOrAdd(key, func(k string) *sync.RWMutex {
		return &sync.RWMutex{}
	})

	return v
}
