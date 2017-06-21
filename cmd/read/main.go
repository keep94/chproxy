package main

import (
	"fmt"
	"github.com/Symantec/chproxy/chreader"
	"log"
	"os"
)

func main() {
	entries, next, err := chreader.DefaultCH.Fetch(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Entries:")
	for _, e := range entries {
		fmt.Println(*e)
	}
	fmt.Println("Next = ", next)
}
