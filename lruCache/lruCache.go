package lrucache

import "strings"

type Cache struct {
	val  string
	prev *Cache
	next *Cache
}

func LRUCache(arrStr []string, limit int) string {
	cacheMap := map[string]*Cache{}
	var (
		tailKey, headKey                             string
		counter                                      int
		prevPtr, nextPtr, tailPtr, headPtr, cachePtr *Cache
	)

	for _, str := range arrStr {
		if tailKey == str {
			continue
		}

		tailPtr = cacheMap[tailKey]
		headPtr = cacheMap[headKey]
		cachePtr = cacheMap[str]
		if cachePtr == nil {
			cachePtr = &Cache{val: str}
			cacheMap[str] = cachePtr
			counter++
		}

		if headKey == "" {
			headKey = str
			cacheMap[headKey] = cachePtr
			goto SET_TAIL_CACHE
		}

		prevPtr = cachePtr.prev
		nextPtr = cachePtr.next
		if prevPtr != nil {
			prevPtr.next = nextPtr
		}

		if nextPtr != nil {
			nextPtr.prev = prevPtr
		}

		if counter > limit && headPtr != nil {
			delete(cacheMap, headKey)
			headKey = headPtr.next.val
			cacheMap[headKey] = headPtr.next
		}

		if headKey == str {
			headKey = nextPtr.val
			cacheMap[headKey] = nextPtr
		}

	SET_TAIL_CACHE:
		if tailPtr != nil {
			tailPtr.next = cachePtr
			cachePtr.prev = tailPtr
		}

		tailKey = str
		cacheMap[tailKey] = cachePtr
		cachePtr.next = nil
	}

	result := []string{}
	cursor := cacheMap[headKey]
	for cursor != nil {
		result = append(result, cursor.val)
		cursor = cursor.next

	}

	return strings.Join(result, "-")
}
