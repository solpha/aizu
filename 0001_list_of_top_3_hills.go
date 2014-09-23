package main

import "fmt"

func main() {
	var mount [10]int
	var top, second, third int
	for i := 0; i < 10; i++ {
		var input int
		fmt.Scanf("%d", &input)
		mount[i] = input
		switch {
		case i == 0:
			top = mount[i]
		case mount[i-1] < mount[i]:
			third = second
			second = top
			top = mount[i]
		case second < mount[i]:
			third = second
			second = mount[i]
		case third < mount[i]:
			third = mount[i]
		}
	}
	fmt.Printf("%d\n%d\n%d\n", top, second, third)
}
