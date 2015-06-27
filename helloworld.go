package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		fmt.Println("Hello", os.Args[1])		
	} else {
		fmt.Println("Usage:", os.Args[0], "name")
	}
}