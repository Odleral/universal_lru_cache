package main

import (
	"fmt"
	"github.com/Odleral/universal_lru_cache/app"
	"time"
)

func main() {
	cache := app.NewLRUCache(app.WithTTL(2*time.Second), app.WithCapacity(100))
	cache.Add("f", 1)
	cache.Add("a", 11)

	fmt.Println(cache.Get("f"))
	time.Sleep(6 * time.Second)
	fmt.Println(cache.Get("a"))
}
