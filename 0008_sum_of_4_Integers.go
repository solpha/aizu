package main

import "fmt"

func main() {
	var x, count int
	fmt.Scanf("%d", &x)
	for a := 0; a <= 9; a++ {
		if a == x {
			count++
			continue
		}
		for b := 0; b <= 9; b++ {
			if (a + b) == x {
				count++
				continue
			}
			for c := 0; c <= 9; c++ {
				if (a + b + c) == x {
					count++
					continue
				}
				for d := 0; d <= 9; d++ {
					if (a + b + c + d) == x {
						count++
						continue
					}
				}
			}
		}
	}
	fmt.Println(count)
}
