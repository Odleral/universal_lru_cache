package app

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	cache := NewLRUCache(WithTTL(3*time.Second), WithCapacity(10))
	t.Run("Test new initialize", func(t *testing.T) {
		assert.Equal(t, 10, cache.capacity, "cache size is 5")
		assert.NotEqual(t, nil, cache.hash, "hash map not nil")
		assert.Equal(t, 3*time.Second, cache.list.TTLDuration, "cache list ttl is 3 seconds")
		assert.Equal(t, 0, cache.list.len, "cache len is 0")

		cache.Add("e1", 1)
		cache.Add("e2", 2)
		cache.Add("e3", 3)
		cache.Add("e4", 4)
		cache.Add("e5", 5)
		cache.Add("e6", 6)
		cache.Add("e7", 7)
		cache.Add("e8", 8)
		cache.Add("e9", 9)
		cache.Add("e10", 10)
		cache.Add("e11", 11)

		e1, ok := cache.Get("e1")
		assert.Equal(t, false, ok, "ok is false")
		assert.Equal(t, nil, e1, "e1 is nil")

		e2, ok := cache.Get("e2")
		assert.Equal(t, 2, e2.(int), "e2 is 2")

		cache.Add("e12", 12)

		e3, ok := cache.Get("e3")
		assert.Equal(t, false, ok, "ok is false")
		assert.Equal(t, nil, e3, "e3 is nil")

		e2, ok = cache.Get("e2")
		assert.Equal(t, 2, e2.(int), "e2 until is 2")
		assert.Equal(t, true, ok, "ok is true")
	})

	t.Run("Test cache add method", func(t *testing.T) {
		cache.Add("test", "some string")

		assert.Equal(t, 10, cache.list.len, "cache size until is 10")

		test, ok := cache.Get("test")
		assert.Equal(t, true, ok, "test key exists")
		assert.Equal(t, "some string", test.(string), "test is 'some string'")

		test2, ok := cache.Get("test2")
		assert.Equal(t, nil, test2, "test2 is nil")
		assert.Equal(t, false, ok, "ok is false")
	})
}

func insertXNodeToCache(x int, b *testing.B) {
	cache := NewLRUCache(WithTTL(2*time.Second), WithCapacity(x))
	b.ResetTimer()

	for i := 0; i < x; i++ {
		cache.Add(strconv.Itoa(i), i)
	}
}

func insertXNodeToCacheWithoutCap(x int, b *testing.B) {
	cache := NewLRUCache(WithTTL(2 * time.Second))
	b.ResetTimer()

	for i := 0; i < x; i++ {
		cache.Add(strconv.Itoa(i), i)
	}
}

func insertXNodeToCacheWithCap100(x int, b *testing.B) {
	cache := NewLRUCache(WithTTL(2*time.Second), WithCapacity(100))
	b.ResetTimer()

	for i := 0; i < x; i++ {
		cache.Add(strconv.Itoa(i), i)
	}
}

func BenchmarkCacheAdd_WithCapacity1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXNodeToCache(1000, b)
	}
}

func BenchmarkCacheAdd_WithCapacity10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXNodeToCache(10000, b)
	}
}

func BenchmarkCacheAdd_WithCapacity100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXNodeToCache(100000, b)
	}
}

func BenchmarkCache_AddWithCapacity1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXNodeToCache(1000000, b)
	}
}

func BenchmarkCache_Add1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXNodeToCacheWithoutCap(1000, b)
	}
}

func BenchmarkCache_Add10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXNodeToCacheWithoutCap(10000, b)
	}
}

func BenchmarkCache_Add100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXNodeToCacheWithoutCap(100000, b)
	}
}

func BenchmarkCache_Add1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXNodeToCacheWithoutCap(1000000, b)
	}
}

func BenchmarkCache_Add_WithCapacity100_And_Input1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXNodeToCacheWithCap100(1000, b)
	}
}

func BenchmarkCache_Add_WithCapacity100_And_Input10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXNodeToCacheWithCap100(10000, b)
	}
}

func BenchmarkCache_Add_WithCapacity100_And_Input100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertXNodeToCacheWithCap100(100000, b)
	}
}
