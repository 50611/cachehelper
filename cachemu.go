package lcache

import (
	lru "github.com/50611/golang-lru/v2"
	"sync"
)

type CacheMu struct {
	cache *lru.Cache[string, *sync.Mutex]
}

func NewCacheMu(cnt int) *CacheMu {
	c, _ := lru.New[string, *sync.Mutex](cnt)

	return &CacheMu{cache: c}
}

func (l *CacheMu) GetOrAdd(key string) *sync.Mutex {

	_, v := l.cache.GetOrAdd(key, func(k string) *sync.Mutex {
		return &sync.Mutex{}
	})

	return v
}
