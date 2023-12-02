package main

import (
	"fmt"
	"net/http"
	"server/generated/routegen"
)

func main() {
	http.HandleFunc("/api", routegen.RouteHandler)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
