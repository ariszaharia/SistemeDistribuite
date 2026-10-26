package main

import (
	"fmt"
)

func main(){
	ch := make(chan string)

	go func(){
		ch <- "salut"
		ch <- "aris"
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
