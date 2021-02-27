package main

import (
	"flag"
	"fmt"
)

func main() {
	source := flag.String("source", ".", "Source directory")
	target := flag.String("target", "", "Target directory")

	flag.Parse()

	fmt.Println("Source:", *source)
	fmt.Println("Target:", *target)
}
