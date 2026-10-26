package main

import (
	"fmt"
)

func main(){
	ch := make(chan string, 2)
	ch <- "salut"
	ch <- "aris"
	close(ch)

	for elem := range ch {
		fmt.Println(elem)
	}
}
