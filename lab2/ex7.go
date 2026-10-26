package main

import (
	"fmt"
	"time"
)

type coordonate struct{
	length int
	height int
}

func main(){
	ch1ch2 := make(chan coordonate)
	ch2ch1 := make(chan coordonate)

	done := make(chan string)

	traseu1 := []coordonate{{0, 0}, {1, 0}, {2, 0}, {3, 0}}
	traseu2 := []coordonate{{5, 5}, {5, 6}, {5, 7}}

	go robot("ROBOT1", traseu1, ch1ch2, ch2ch1, done)
	go robot("ROBOT2", traseu2, ch2ch1, ch1ch2, done)

	fmt.Println("Observer confirma:", <-done)
	fmt.Println("Observer confirma:", <-done)
}

func robot(nume string, traseu []coordonate, in chan coordonate, out chan coordonate, observer chan <- string){
	over := make(chan bool)

	go func(){
		for pozitie := range in{
			fmt.Printf("%s pozitie coleg: X:%d, Y:%d\n", nume, pozitie.length, pozitie.height)
		}
		over <- true
	}()

	for _, pct:= range traseu{
		out <- pct
		time.Sleep(400 * time.Millisecond)
	}

	close(out)

	<-over

	observer <- fmt.Sprintf("%s a terminat", nume)
}
