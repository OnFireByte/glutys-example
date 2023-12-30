package math

import (
	"strconv"

	"github.com/OnFireByte/glutys-example/server/di/cache"
)

func Fib(cache cache.Cache, n int) int {
	if raw, ok := cache.Get(strconv.Itoa(n)); ok {
		v, err := strconv.Atoi(raw)
		if err == nil {
			return v
		}
	}

	if n <= 1 {
		return n
	}

	a := 0
	b := 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	cache.Set(strconv.Itoa(n), strconv.Itoa(b))

	return b
}
