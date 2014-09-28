package main

import "fmt"

func main() {
	var number [30]int
	var count int

	for i := 0; i < 30; i++ {
		var input int
		fmt.Scanf("%d", &input)
		if input == 0 {
			break
		}
		number[i] = input
		count = i
	}

	for p := 0; p <= count; p++ {
		var prime int
		for x := 0; x <= number[p]; x++ {
			switch {
			case x == 1:
			case x == 2 || x == 3 || x == 5 || x == 7:
				prime++
			case x%2 != 0 && x%3 != 0 && x%5 != 0 && x%7 != 0:
				prime++
			}
		}
		fmt.Println(prime)
	}
}
