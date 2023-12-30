package main

import (
	"fmt"
	"net/http"

	"github.com/OnFireByte/glutys-example/server/di/cache"
	"github.com/OnFireByte/glutys-example/server/generated/routegen"
)

func main() {
	// the order of dependencies depends on the order of AddDependencyType calls
	handler := routegen.NewHandler(
		cache.NewCacheImpl(),
	)
	http.HandleFunc("/api", handler.Handle)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
