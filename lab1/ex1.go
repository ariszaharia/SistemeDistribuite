package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	nr, err := strconv.Atoi(scanner.Text()) 

	if err != nil{
		fmt.Println("Eroare la conversie")
	}

	var nrdiv int

	for i := 2; i <= nr/2; i++ {
		if nr%i == 0 {
			nrdiv++
		}
	}


	if nrdiv > 0 || nr < 1{
		fmt.Println("Nu e prim")
	} else {
		fmt.Println("Prim")
	}

//2
	numere := []int{45, 12, 89, 33, 7, 24, 99, 1}
	
	sort.Ints(numere)
	
	fmt.Println(numere)


//3

	var numerePare []int
	for _, numar := range numere{
		if numar%2 == 0{
			numerePare = append(numerePare, numar)
		}
	}
	fmt.Println(numerePare)

}