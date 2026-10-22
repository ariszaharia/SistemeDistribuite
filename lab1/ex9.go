package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	s, _ := strconv.Atoi(scanner.Text())

	

	arr := make([]int, 0, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		nr, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, nr)
	}

	hashmap := make(map[int]int)

	var result [][]int	
	for index, val := range arr{
		target := s - val

		id , ok := hashmap[target]

		if ok{
			result = append(result, []int{id, index})
		}
		

		hashmap[val] = index
	}
	fmt.Println(result)






}
