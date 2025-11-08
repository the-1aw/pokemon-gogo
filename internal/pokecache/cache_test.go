package pokecache

import (
	"reflect"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	testCases := map[string]struct {
		key   string
		value []byte
	}{
		"basic use": {
			key:   "key",
			value: []byte("value"),
		},
		"url cache key": {
			key:   "http://whocares.com/foo?bar=3",
			value: []byte("somevalue"),
		},
	}
	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			cache := NewCache(25 * time.Second)
			cache.Add(test.key, test.value)
			got, ok := cache.Get(test.key)
			if !ok {
				t.Fatalf("Missing cache value")
			}
			if !reflect.DeepEqual(got, test.value) {
				t.Fatalf("expected: %#v, got%#v", test.value, got)
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	testCase := struct {
		key   string
		value []byte
	}{
		key:   "key",
		value: []byte("value"),
	}
	ttl := 25 * time.Millisecond
	cache := NewCache(ttl)
	cache.Add(testCase.key, testCase.value)
	time.Sleep(ttl * 2)
	_, ok := cache.Get(testCase.key)
	if ok {
		t.Fatal("Key should have expired after ttl")
	}
}
