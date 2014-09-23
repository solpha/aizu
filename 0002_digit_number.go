package main

import "fmt"

func main() {
	var sum [200]int
	var progress, count, digit int
	for i := 0; i < 200; i++ {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		if a == 0 {
			break
		}
		sum[i] = a + b
		progress = i
	}
	for n := 0; n <= progress; n++ {
		for x := 1; ; x *= 10 {
			if sum[n]/x == 0 {
				digit = count
				count = 0
				break
			}
			count++
		}
		fmt.Printf("%d\n", digit)
	}
}
