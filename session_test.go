package lcache

import (
	"testing"
)

func TestNewSessionT(t *testing.T) {
	s := NewSessionT[string, struct{}](100)
	s.Add("1", struct {
	}{})
	s.Add("2", struct{}{})
	_, ok := s.Get("1", 10)
	if ok {
		t.Log(ok)
	}
	_, ok = s.Get("3", 10)
	if ok {
		t.Log(ok)
	} else {
		t.Log(false)
	}
}
func TestNewSessionT2(t *testing.T) {
	s := NewSessionT[string, int](100)
	s.Add("1", 1)
	s.Add("2", 2)
	v, ok := s.Get("1", 10)
	if ok {
		t.Log(v)
	}
	_, ok = s.Get("3", 10)
	if ok {
		t.Log(ok)
	}
}

type User struct {
	Name string
}

func NewUser(name string) *User {
	return &User{Name: name}
}
func TestNewSessionT3(t *testing.T) {
	s := NewSessionT[string, *User](100)
	s.Add("1", NewUser("1111"))
	s.Add("2", NewUser("2222"))
	v, ok := s.Get("1", 10)
	if ok {
		t.Log(v.Name)
	}
	_, ok = s.Get("3", 10)
	if ok {
		t.Log(ok)
	}
}
