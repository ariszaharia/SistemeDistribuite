package main

import (
	"fmt"
)

func main(){
	chsend := make(chan string)
	chrez := make(chan string)

	go send(chsend)

	go receive(chsend, chrez)

	fmt.Println(<-chrez)
}

func send(ch chan string){
	ch <- "mesaj"
}

func receive(ch chan string, ch1 chan string){
	ch1 <- <-ch
}
