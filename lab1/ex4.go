package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	slice := []float32{1, 2, 3, 4, 5}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	index, err := strconv.Atoi(scanner.Text())

	if err != nil{
		fmt.Println("Eroare la conversie")
	}

	partea1 := slice[0:index]
	partea2 := slice[index+1: len(slice)]

	result := append(partea1, partea2...)
	fmt.Println(result)

}
