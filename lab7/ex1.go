package main

import (
	"fmt"
	"math/rand"
)


type Message struct{
	SenderID int
	Round int
	Value int
}

type Node struct {
	id int
	x int 
	s int
	n int 
	isByzantine bool
	inbox chan Message
}

func (node *Node) broadcast(nodes []*Node, round int, val int, malicious bool) {
	for _, target := range nodes {
		sendVal := val
		if malicious {
			sendVal = rand.Intn(2)
		}
		target.inbox <- Message{
			SenderID: node.id,
			Round : round,
			Value : sendVal,
		}
		
	}
}

func (node *Node) receive(count int) []Message{
	messages := make([]Message, 0, count)
	for i := 0; i < count ; i++{
		msg := <-node.inbox
		messages = append(messages, msg)

	}
	return messages
}


func main(){
	var n, s int
	if _, err := fmt.Scan(&n, &s); err != nil {
		return
	}
	if n <= 4*s {
		fmt.Println("Eroare: Conditia n > 4s nu este respectata.")
		return
	}

	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		nodes[i] = &Node{
			id:    i,
			x:     v,
			s:     s,
			n:     n,
			inbox: make(chan Message, n*2),
		}
	}

}