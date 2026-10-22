package main

import (
	"fmt"
	"time"
)


func main(){
	ex1()
	// ex2()
	// ex3()
	// ex4()
	// ex5()
	// ex6()
	// ex7()
}


func ex1(){
	ch := make(chan string)
	go func(){
		ch <- "salut"
	}()
	
	mesajprimit := <-ch

	fmt.Println(mesajprimit)


}

func ex2(){
	ch := make(chan string, 1)
	go primesteMesaj(ch)
	ch <- "salut"
	time.Sleep(10 * time.Millisecond)
	

}
func primesteMesaj(c chan string){
		mesaj := <-c
		fmt.Println(mesaj)
	}




func ex3(){
	ch := make(chan string, 2)
	ch <- "salut"
	ch <- "aris"
	close(ch)

	for elem := range ch {
		fmt.Println(elem)
	}
}


func ex4(){
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

func ex5(){
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



func ex6(){
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func(){
		time.Sleep(4 * time.Second)
		ch1 <- "f1"
	}()

	go func(){
		time.Sleep(3 * time.Second)
		ch2 <- "f2"
	}()

	select{
	case msg := <- ch1:
		fmt.Println(msg)
	
	case msg := <- ch2:
		fmt.Println(msg)
	}
}

type coordonate struct{
		length int
		height int
	}

func ex7(){

	ch1ch2 := make(chan coordonate)
	ch2ch1 := make(chan coordonate)

	done := make(chan string)

	traseu1 := []coordonate{{0, 0}, {1, 0}, {2, 0}, {3, 0}}
	traseu2 := []coordonate{{5, 5}, {5, 6}, {5, 7}}


	go robot("ROBOT1", traseu1, ch1ch2, ch2ch1, done)
	go robot("ROBOT2", traseu2, ch2ch1, ch1ch2, done)


	fmt.Println("Observer confirm:", <-done)
	fmt.Println("Observer confirma:", <-done)

}

func robot(nume string, traseu []coordonate, in chan coordonate, out chan coordonate, observer chan <- string){
	over := make(chan bool)

	go func(){
		for pozitie := range in{
			fmt.Printf("%s pozitie coleg: X:%d, Y:%d\n", nume, pozitie.height, pozitie.length)
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



