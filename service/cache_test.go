package service

import (
	"bytes"
	"testing"
)

func TestCacheSetAndGet(t *testing.T) {
	cache := NewCache()
	cache.Set("key1", []byte("value1"))
	value, err := cache.Get("key1")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !bytes.Equal(value, []byte("value1")) {
		t.Errorf("Expected value1, got %s", value)
	}
	//if value != []byte("value1") {
	//	t.Errorf("Expected value1, got %s", value)
	//}
}

func TestCacheGetNotFound(t *testing.T) {
	cache := NewCache()
	_, err := cache.Get("nonexistent")
	if err == nil {
		t.Errorf("Expected an error for nonexistent key, got nil")
	}
}

func TestCacheDelete(t *testing.T) {
	cache := NewCache()
	cache.Set("key1", []byte("value1"))
	cache.Delete("key1")
	_, err := cache.Get("key1")
	if err == nil {
		t.Errorf("Expected an error after deleting key, got nil")
	}
}

func TestCacheClear(t *testing.T) {
	cache := NewCache()
	cache.Set("key1", []byte([]byte("value1")))
	cache.Clear()
	_, err := cache.Get("key1")
	if err == nil {
		t.Errorf("Expected an error after clearing cache, got nil")
	}
}

func TestConcurrentAccess(t *testing.T) {
	cache := NewCache()
	cache.Set("key1", []byte("value1"))

	done := make(chan bool)

	go func() {
		cache.Set("key2", []byte("value2"))
		done <- true
	}()

	go func() {
		cache.Delete("key1")
		done <- true
	}()

	<-done
	<-done

	_, err := cache.Get("key1")
	if err == nil {
		t.Errorf("Expected an error for key1 after deletion, got nil")
	}

	value, err := cache.Get("key2")
	if err != nil || !bytes.Equal(value, []byte("value2")) {
		t.Errorf("Expected value2 for key2, got %v, %v", value, err)
	}
}
