package main

import "fmt"

func main() {
	var moji string
	fmt.Scanf("%s", &moji)

	length := len(moji)
	fmt.Println(length)

	for i := 5; i >= 1; i = i - 1 {
		fmt.Printf("%s", moji[i-1:i])
	}
}
