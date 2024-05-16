package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	c = SafeCounter{
		v: map[string]int{
			"key1": 2,
			"key2": 4,
			"key3": 5,
		},
		mu: sync.Mutex{},
	}
	for m := range c.v {
		go c.Inc(m)
		// go c.Inc("key2")
	}

	time.Sleep(time.Millisecond)
	// fmt.Println(c.Value("key1"))
	// fmt.Println(c.Value("key2"))
	for m := range c.v {
		fmt.Println(c.Value(m))
	}
}
