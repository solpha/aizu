package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var result [1000]string
	for i := 0; i < n; i++ {
		var x, y, z int
		fmt.Scanf("%d %d %d", &x, &y, &z)
		switch {
		case x > y && x > z:
			if (y*y)+(z*z) == (x * x) {
				result[i] = "YES"
			} else {
				result[i] = "NO"
			}
		case y > x && y > z:
			if (x*x)+(z*z) == (y * y) {
				result[i] = "YES"
			} else {
				result[i] = "NO"
			}
		case z > x && x > y:
			if (x*x)+(y*y) == (z * z) {
				result[i] = "YES"
			} else {
				result[i] = "NO"
			}
		default:
			result[i] = "NO"
		}
	}
	for a := 0; a < n; a++ {
		fmt.Printf("%s\n", result[a])
	}
}
