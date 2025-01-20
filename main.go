package main

import (
	"flag"
	"fmt"
)

func main() {
	idPtr := flag.Int("i", 0, "id of todo")
	titlePtr := flag.String("t", "Untitled", "title for todo")
	descriptionPtr := flag.String("d", "", "description for todo")

	flag.Parse()

	fmt.Println(*idPtr)
	fmt.Println(*titlePtr)
	fmt.Println(*descriptionPtr)
}
