package main

import (
	"flag"
	"fmt"
)

var (
	name = flag.String("name", "world", "Name")
)

func main() {
	flag.Parse()

	fmt.Printf("Hello %s.\n", *name)
}
