package lru

import "testing"

func TestReadWrite(t *testing.T) {
	testlru := NewCache(3)
	testkey := "key"
	testval := "val"
	if err := testlru.Put(testkey, testval); err != nil {
		t.Error("Write test failed", err)
	}
	if val, err := testlru.Get(testkey); err != nil {
		t.Error("Read test failed", err)
	} else if val.(string) != "val" {
		t.Error("Read failed with incorrect return value", val)
	}
}

func TestWriteWithEviction(t *testing.T) {
	testlru := NewCache(3)
	testlru.Put("key1", "val1")
	testlru.Put("key2", "val2")
	testlru.Put("key3", "val3")
	testlru.Put("key4", "val4")
	if _, err := testlru.Get("key1"); err == nil {
		t.Error("LRU replacement policy test failed")
	}
}

func TestGetFailure(t *testing.T) {
	testlru := NewCache(3)
	testlru.Put("key1", "val1")
	if _, err := testlru.Get("key2"); err == nil {
		t.Error("Get an value that is not suppose to be in the cache", err)
	}
}

func TestPutFailure(t *testing.T) {
	testlru := NewCache(3)
	testlru.Put("key1", "val1")
	if err := testlru.Put("key1", "val1"); err == nil {
		t.Error("Put an exist value into the cache that do not return an error", err)
	}
}
