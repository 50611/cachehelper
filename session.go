package lcache

import (
	lru "github.com/50611/golang-lru/v2"
	"time"
)

type ISessionT[K comparable, T any] interface {
	Add(k K, bean T)
	Get(k K, life int) (T, bool)
}

type sessionBeanT[T any] struct {
	bean T
	t    time.Time
}

func newSessionBeanT[T any](bean T) *sessionBeanT[T] {
	return &sessionBeanT[T]{bean: bean, t: time.Now()}
}

type SessionT[K comparable, V any] struct {
	sessions *lru.Cache[K, *sessionBeanT[V]]
	size     int
}

func NewSessionT[K comparable, V any](cache int) *SessionT[K, V] {
	s, _ := lru.New[K, *sessionBeanT[V]](cache)
	return &SessionT[K, V]{sessions: s, size: cache}
}

func (s *SessionT[K, V]) Get(key K, life int) (value V, ok bool) {

	bean, ok := s.sessions.Get(key)
	if ok {

		i := int(time.Now().Sub(bean.t).Seconds())
		if i > life {
			return
		}
		return bean.bean, true
	}
	return
}

func (s *SessionT[K, T]) Add(Key K, bean T) {

	s.sessions.Add(Key, newSessionBeanT(bean))

}
