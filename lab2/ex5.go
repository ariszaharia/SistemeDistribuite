package main

import (
	"fmt"
)

func main(){
	chsend := make(chan string)
	chrez := make(chan string)

	go send(chsend)

	go func(){
		chrez <- receive(chsend)
	}()

	fmt.Println(<-chrez)
}

func send(ch chan<- string){
	ch <- "mesaj"
}

func receive(ch <-chan string) string{
	return <-ch
}
