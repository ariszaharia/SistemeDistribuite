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



}