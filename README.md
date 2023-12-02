# Универсальный LRU кэш

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)]()

Универсальный LRU кэш с регулируемой TTL и размером. Покрыт тестами и тестами производительности и потребления памяти.

## Особенности
- Возможно хранить любой тип
- Можно задать время жизни для ноды
- Можно задать размер кэша

## Использование

Получение нового кеша.
```go
    cache := app.NewLRUCache(WithTTL(2 * time.Second), WithCapacity(100))
	cache.Add("e1", 3.14)
	cache.Add("e2", 67)
	cache.Add("user:1", "some text")
	fmt.Println(cache.Get("e1"))
    fmt.Println(cache.Get("user"))
```
## Тестирование
```sh
# тестирование 
make test
# тестирование производительности
make bench
```
go test -bench=. -benchtime=100x -benchmem ./app
goos: windows
goarch: amd64                                            
pkg: github.com/Odleral/universal_lru_cache/app          
cpu: Intel(R) Xeon(R) CPU E5-2670 v3 @ 2.30GHz
|Name|times|ns/op|b/op| allocs/op|
|-----|-----|-----|------|------|
|BenchmarkCacheAdd_WithCapacity1000-24|100||924 B/op|26 allocs/op|
BenchmarkCacheAdd_WithCapacity10000-24 |100|30244 ns/op|9564 B/op|296 allocs/op|
BenchmarkCacheAdd_WithCapacity100000-24|100|300073 ns/op|95964 B/op|2996 allocs/op|
|BenchmarkCache_AddWithCapacity1000000-24                             |100|           4028643 ns/op|          959964 B/op|      29996 allocs/op|
|BenchmarkCache_Add1000-24                                            |100|             10267 ns/op|            2127 B/op|         26 allocs/op|
|BenchmarkCache_Add10000-24                                           |100|             40017 ns/op|           19076 B/op|        298 allocs/op|
|BenchmarkCache_Add100000-24                                          |100|            457972 ns/op|          246961 B/op|       3025 allocs/op|
|BenchmarkCache_Add1000000-24                                         |100|           6250085 ns/op|         2167662 B/op|      30231 allocs/op|
|BenchmarkCache_Add_WithCapacity100_And_Input1000-24                  |100|              9961 ns/op|             926 B/op|         26 allocs/op|
|BenchmarkCache_Add_WithCapacity100_And_Input10000-24                 |100|             40196 ns/op|            9603 B/op|        296 allocs/op|
|BenchmarkCache_Add_WithCapacity100_And_Input100000-24                |100|            370012 ns/op|           96549 B/op|       2997 allocs/op|
|BenchmarkList_Append1000-24                                          |100|              9986 ns/op|             924 B/op|         26 allocs/op|
|BenchmarkList_Append10000-24                                         |100|             19687 ns/op|            9564 B/op|        296 allocs/op|
|BenchmarkList_Append100000-24                                        |100|            210295 ns/op|           95964 B/op|       2996 allocs/op|
|BenchmarkList_Append1000000-24                                       |100|           1649901 ns/op|          959964 B/op|      29996 allocs/op|
## Лицензия
MIT
