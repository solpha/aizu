package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for n := 1; n <= 9; n++ {
			fmt.Printf("%dx%d=%d\n", i, n, i*n)
		}
	}
}
