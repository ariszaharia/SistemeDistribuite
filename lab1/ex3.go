package main
import "fmt"

func main() {
	numere := []int{45, 12, 89, 33, 7, 24, 99, 1}
	var numerePare []int
	for _, numar := range numere{
		if numar%2 == 0{
			numerePare = append(numerePare, numar)
		}
	}
	fmt.Println(numerePare)

}
