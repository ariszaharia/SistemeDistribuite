package main
import (
	"fmt"
)

func main() {
	numere := []int{45, 12, 89, 33, 7, 24, 99, 1}
	
	for i:=0; i<len(numere); i++{
		for j:=i; j<len(numere); j++{
			if numere[i] > numere[j]{
				numere[i], numere[j] = numere[j], numere[i]
			}
		}
	}
	
	fmt.Println(numere)
}