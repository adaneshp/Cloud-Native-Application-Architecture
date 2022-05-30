package lru

import "errors"

type Cacher interface {
	Get(interface{}) (interface{}, error)
	Put(interface{}, interface{}) error
}

type lruCache struct {
	size      int
	remaining int
	cache     map[string]string
	queue     []string
}

func NewCache(size int) Cacher {
	return &lruCache{size: size, remaining: size, cache: make(map[string]string), queue: make([]string, 0)}
}

func (lru *lruCache) Get(key interface{}) (interface{}, error) {
	/*****Armin******/
	k := key.(string)

	//look for the key in the cache
	_, exist := lru.cache[k]
	if exist { //the key is found in cache
		lru.qDel(k)                      //delete the key from the queue
		lru.queue = append(lru.queue, k) //append the key to the queue, so it becomes the latest used element
		return lru.cache[k], nil
	} else { //the key is not in the cache return an error
		return nil, errors.New("the element is not found")
	}
}

func (lru *lruCache) Put(key, val interface{}) error {
	/********Ghazai******/
	//convert the input to a concrete type
	k := key.(string)
	v := val.(string)

	//look for the key in the cache
	_, exist := lru.cache[k]
	if exist {
		//uncomment the follow 2 line only if a failure put is consider as an access to the element
		//lru.qDel(k)                      //delete the key from the queue
		//lru.queue = append(lru.queue, k) //append the key to the queue, so it becomes the latest used element
		lru.cache[k] = v //update the value, in case there is a new value for that key
		return errors.New("element is already existed in cache")
	} else {

		/*****Rosa*****/
		//to check if there is space in cache
		if lru.remaining == 0 { //the cache is full
			//delete the least used element from both cache and queue
			delete(lru.cache, lru.queue[0])
			lru.qDel(lru.queue[0])
			//put the new value into the cache and queue
			lru.cache[k] = v
			lru.queue = append(lru.queue, k)

			/*****Yinfei*******/
		} else { //the cache has some remaining space
			lru.remaining = lru.remaining - 1
			//insert the element to the cache and appended into queue
			lru.cache[k] = v
			lru.queue = append(lru.queue, k)
		}
		return nil
	}
}

// Delete element from queue
func (lru *lruCache) qDel(ele string) {
	for i := 0; i < len(lru.queue); i++ {
		if lru.queue[i] == ele {
			oldlen := len(lru.queue)
			copy(lru.queue[i:], lru.queue[i+1:])
			lru.queue = lru.queue[:oldlen-1]
			break
		}
	}
}
