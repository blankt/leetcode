package leetcode

import "testing"

func TestLRUCache(t *testing.T) {
	lruCache := Constructor(3)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Put(3, 3)
	lruCache.Put(4, 4)

	if val := lruCache.Get(4); val != 4 {
		t.Errorf("Expected 1, got %d", val)
	}
	if val := lruCache.Get(3); val != 3 {
		t.Errorf("Expected 1, got %d", val)
	}
	if val := lruCache.Get(2); val != 2 {
		t.Errorf("Expected 1, got %d", val)
	}
	if val := lruCache.Get(1); val != -1 {
		t.Errorf("Expected 1, got %d", val)
	}

	lruCache.Put(5, 5)

	if val := lruCache.Get(1); val != -1 {
		t.Errorf("Expected 1, got %d", val)
	}
	if val := lruCache.Get(2); val != 2 {
		t.Errorf("Expected 1, got %d", val)
	}
	if val := lruCache.Get(3); val != 3 {
		t.Errorf("Expected 1, got %d", val)
	}
	if val := lruCache.Get(4); val != -1 {
		t.Errorf("Expected 1, got %d", val)
	}
	if val := lruCache.Get(5); val != 5 {
		t.Errorf("Expected 1, got %d", val)
	}
}
