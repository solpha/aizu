package main

import "fmt"

func main() {
	var resultX, resultY [100]float32
	var count int
	for {
		var a, b, c, d, e, f, x, y float32
		fmt.Scanf("%f %f %f %f %f %f", &a, &b, &c, &d, &e, &f)
		if a == 0 {
			break
		}
		//Simultaneous equation
		y = (((a * f) - (d * c)) / ((a * e) - (b * d)))
		x = ((c - (b * y)) / a)
		resultX[count] = x
		resultY[count] = y
		count++
	}
	for i := 0; i < count; i++ {
		fmt.Printf("%.3f %.3f\n", resultX[i], resultY[i])
	}
}
