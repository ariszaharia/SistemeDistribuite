package main

import (
	"fmt"
)

func main(){
	ch := make(chan string)
	done := make(chan bool)

	go primesteMesaj(ch, done)
	ch <- "salut"

	<-done
}

func primesteMesaj(c <-chan string, done chan<- bool){
	mesaj := <-c
	fmt.Println(mesaj)
	done <- true
}
