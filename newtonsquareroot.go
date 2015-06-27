package main

import (
	"fmt"
	"os"
	"strconv"
)

const delta = 1 / (1 << 100)

func abs(x float64) float64 {
	if x < 0.0 {
		return -1.0 * x
	} else {
		return x
	}
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0 ; i < 15 ; i++ {
		nz := z - (z * z - x) / (2 * z)
		if abs(z - nz) < delta {
			return nz
		}
		z = nz
	}
	return z
}

func usage() {
	fmt.Println("Usage:", os.Args[0], "number")
	fmt.Println("\tNumber must be a float or an integer")
}

func main() {
	if len(os.Args) == 2 {
		number, err := strconv.ParseFloat(os.Args[1], 64)
		if err == nil {
			fmt.Println(Sqrt(number))		
		} else {
			usage()
		}
	} else {
		usage()
	}
}