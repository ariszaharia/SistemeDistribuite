package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan string, 1)
	go primesteMesaj(ch)
	ch <- "salut"
	time.Sleep(10 * time.Millisecond)
}

func primesteMesaj(c chan string){
	mesaj := <-c
	fmt.Println(mesaj)
}
