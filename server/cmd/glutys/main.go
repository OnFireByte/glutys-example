package main

import (
	"fmt"

	"github.com/OnFireByte/glutys-example/server/di/cache"
	"github.com/OnFireByte/glutys-example/server/reqcontext"
	"github.com/OnFireByte/glutys-example/server/route"

	"github.com/onfirebyte/glutys"
)

func main() {
	fmt.Println("Generating routes...")

	builder := glutys.NewBuilder(
		"github.come/user/example/generated/routegen",
		"generated/routegen/route.go",
		"../client/generated/contract.ts",
	)

	builder.AddContextParser(reqcontext.ParseUsername)

	// You must use pointer to type, not the type itself
	builder.AddDependencyType((*cache.Cache)(nil))

	builder.CreateRouter(route.RootRoute)

	builder.Build()

	fmt.Println("Done!")
}
