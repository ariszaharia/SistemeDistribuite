package main

import (
	"fmt"
	"time"
)

func main(){
	done := make(chan bool)
	go func(){
		time.Sleep(2 * time.Second)
		fmt.Println("Executie terminata")
		done <- true
	}()
	fmt.Println("Inainte de functie")
	<-done
	fmt.Println("Dupa functie")
}
