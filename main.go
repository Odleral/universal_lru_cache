package main

import (
	"fmt"
	"github.com/Odleral/universal_lru_cache/app"
	"time"
)

func main() {
	cache := app.NewLRUCache(app.WithTTL(2*time.Second), app.WithCapacity(100))
	cache.Add("e1", 3.14)
	cache.Add("e2", 67)
	cache.Add("user:1", "some text")
	fmt.Println(cache.Get("e1"))
	fmt.Println(cache.Get("user:1"))

	time.Sleep(2 * time.Second)

	fmt.Println(cache.Get("e2"))
}
