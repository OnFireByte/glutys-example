package main

import (
	"fmt"
	"os"

	"server/reqcontext"
	"server/route"

	"github.com/onfirebyte/glutys"
)

func main() {
	fmt.Println("Generating routes...")

	builder := glutys.Builder{
		GeneratePath: "server/generated/routegen",
	}
	builder.AddContextParser(reqcontext.ParseUsername)

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

	fmt.Println("Done!")
}
