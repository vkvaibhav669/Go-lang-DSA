package main

import "fmt"

func main() {
	xi := []int{3, 4, 5, 6, 77867, 567576, 5}
	for i, v := range xi {
		fmt.Println(i, v)
	}
}
