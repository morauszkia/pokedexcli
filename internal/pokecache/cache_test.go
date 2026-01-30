package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second

	cases := []struct {
		key	string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("this is the test data"),
		},
		{
			key: "https://example.com/subpath",
			val: []byte("this is another test data for subpath"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func (t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Key '%s' not found", c.key)
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("Values don't match.\nExpected: %s\nFound: %s", c.val, val)
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const interval = 5 * time.Millisecond
	const waitTime = interval + 5 * time.Millisecond

	cache := NewCache(interval)

	key := "https://example.com"

	cache.Add(key, []byte("this is the test data"))

	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("Key '%s' not found. Expected to find key.", key)
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get(key)
	if ok {
		t.Errorf("Key '%s' still present. Expected to NOT find key.", key)
		return
	}
}