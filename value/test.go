package main

import "fmt"

func test() {
	a = 1
}
func main() {
	slice := []int{1, 2, 3, 4, 5, 7, 8, 9}
	for i, val := range slice {
		fmt.Println(i, val)
	}
}
