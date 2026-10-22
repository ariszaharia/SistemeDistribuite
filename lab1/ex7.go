package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Eroare", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	frecventa := make(map[string]int)

	for scanner.Scan() {
		cuvant := scanner.Text()
		frecventa[cuvant]++
	}


	for cuvant, numar := range frecventa {
		fmt.Printf("%s: %d\n", cuvant, numar)
	}
}