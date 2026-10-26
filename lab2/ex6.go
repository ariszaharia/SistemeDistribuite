package main

import (
	"fmt"
	"time"
)

const finish = 10

func main(){
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go robot("ROBOT1", ch1)
	go robot("ROBOT2", ch2)

	select{
	case nume := <-ch1:
		fmt.Println(nume, "a ajuns primul la finish")

	case nume := <-ch2:
		fmt.Println(nume, "a ajuns primul la finish")
	}
}

func robot(nume string, finishLine chan<- string){
	for pozitie := 0; pozitie < finish; pozitie++ {
		time.Sleep(100 * time.Millisecond)
	}
	finishLine <- nume
}
