package main

import "fmt"

func main() {
	slice := []string{"Laboratorul", "de", "sisteme", "distribuite", "1"}
	left, right := 0, len(slice) - 1

	for left < right{
		slice[left], slice[right] = slice[right], slice[left]
		left++
		right--
	}

	fmt.Println(slice)
}
