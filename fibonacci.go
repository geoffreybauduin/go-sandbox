package main

import (
	"fmt"
	"os"
	"strconv"
)

func usage() {
	fmt.Println("Usage:", os.Args[0], "number")
}

func fibonacci(n int64) int64 {
	if n <= 2 {
		return 1
	} else {
		return fibonacci(n - 1) + fibonacci(n - 2)
	}
}

func main() {
	if len(os.Args) == 2 {
		number, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err == nil {
			fmt.Println(fibonacci(number))		
		} else {
			usage()
		}
	} else {
		usage()
	}
}