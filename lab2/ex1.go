package main

import (
	"fmt"
)

func main(){
	ch := make(chan string)
	go func(){
		ch <- "salut"
	}()

	mesajprimit := <-ch

	fmt.Println(mesajprimit)
}
