package main

import "fmt"

func main() {
	var week int
	fmt.Scanf("%d", &week)

	var interest float32 = 100000
	for i := 0; i < week; i++ {
		interest = interest * 1.05
	}
	intInterest := int(interest)
	fmt.Println(intInterest)

	if intInterest%1000 != 0 {
		x := intInterest / 1000
		y := (x + 1) * 1000
		fmt.Println(y)
	}
}
