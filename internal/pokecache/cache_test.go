package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second

	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "key1",
			val: []byte("val1"),
		},
		{
			key: "key2",
			val: []byte("val2"),
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("key=%s", c.key), func(t *testing.T) {
			cache := NewCache(interval)

			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key %s in cache", c.key)
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value %s in cache, got %s", string(c.val), string(val))
				return
			}
		})
	}
}

func TestReap(t *testing.T) {
	const interval = 1 * time.Second
	cache := NewCache(interval)

	keyOne := "key1"
	valOne := []byte("val1")
	cache.Add(keyOne, valOne)

	// Sleep for interval + some buffer time
	time.Sleep(interval + 100*time.Millisecond)

	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("expected key %s to be reaped from cache", keyOne)
	}
}

func TestReapFresh(t *testing.T) {
	const interval = 1 * time.Second
	cache := NewCache(interval)

	keyOne := "key1"
	valOne := []byte("val1")
	cache.Add(keyOne, valOne)

	// Sleep for less than the interval
	time.Sleep(interval / 2)

	_, ok := cache.Get(keyOne)
	if !ok {
		t.Errorf("expected to find key %s in cache", keyOne)
	}
}
