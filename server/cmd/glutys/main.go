package main

import (
	"os"
	"server/route"

	"github.com/onfirebyte/glutys"
)

func main() {

	builder := glutys.Builder{
		GeneratePath: "server/generated/routegen",
	}
	// builder.AddContextParser(hello.GetUserContext)

	builder.CreateRouter(route.RootRoute)

	goFileString, tsFileString := builder.Build()

	file, err := os.Create("generated/routegen/route.go")
	if err != nil {
		panic(err)
	}

	file.WriteString(goFileString)

	tsFile, err := os.Create("../client/generated/contract.ts")
	if err != nil {
		panic(err)
	}

	tsFile.WriteString(tsFileString)
}
