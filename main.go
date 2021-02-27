package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	source := flag.String("source", ".", "Source directory")
	target := flag.String("target", "", "Target directory")

	flag.Parse()

	fmt.Println("Source:", *source)
	fmt.Println("Target:", *target)
	fmt.Println()

	err := listDir(*source)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

}
