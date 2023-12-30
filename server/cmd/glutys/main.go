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

	// You must use pointer to type that you want to use, not the type itself.
	// For example, if you want to use *FooStruct, you must use (**FooStruct)(nil).
	// Or if you want to use Bar interface, you must use (*Bar)(nil)
	builder.AddDependencyType((*cache.Cache)(nil))

	builder.CreateRouter(route.RootRoute)

	builder.Build()
}
