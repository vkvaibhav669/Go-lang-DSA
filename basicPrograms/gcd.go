package main

import "fmt"

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {

	x, y := 55, 45

	var gcd int = gcd(x, y)
	fmt.Println(gcd)

}
