package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Message struct {
	SenderID int
	Round    int
	Value    int
}

type Node struct {
	id          int
	x           int
	s           int
	n           int
	isByzantine bool
	inbox       chan Message
}

func (node *Node) broadcast(nodes []*Node, round int, val int, malicious bool) {
	for _, target := range nodes {
		sendVal := val
		if malicious {
			sendVal = rand.Intn(2) 
		}
		target.inbox <- Message{
			SenderID: node.id,
			Round:    round,
			Value:    sendVal,
		}
	}
}

func (node *Node) computeMajority(votes []int) (int, int) {
	count1 := 0
	for _, v := range votes {
		if v == 1 {
			count1++
		}
	}
	count0 := node.n - count1
	if count1 >= (node.n/2)+1 {
		return 1, count1
	}
	return 0, count0
}

func main() {
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

	for t := 0; t <= s+1; t++ {
		for i := 0; i < n; i++ {
			nodes[i].isByzantine = false
		}
		f := rand.Intn(s + 1)
		perm := rand.Perm(n)
		for i := 0; i < f; i++ {
			nodes[perm[i]].isByzantine = true
		}

		var wg1 sync.WaitGroup
		majorities := make([]int, n)
		mults := make([]int, n)

		for i := 0; i < n; i++ {
			wg1.Add(1)
			go func(nd *Node) {
				defer wg1.Done()
				nd.broadcast(nodes, 1, nd.x, nd.isByzantine)

				votes := make([]int, 0, n)
				for k := 0; k < n; k++ {
					msg := <-nd.inbox
					votes = append(votes, msg.Value)
				}
				majorities[nd.id], mults[nd.id] = nd.computeMajority(votes)
			}(nodes[i])
		}
		wg1.Wait()

		var wg2 sync.WaitGroup
		for i := 0; i < n; i++ {
			wg2.Add(1)
			go func(nd *Node) {
				defer wg2.Done()
				if nd.id == t {
					nd.broadcast(nodes, 2, majorities[nd.id], nd.isByzantine)
				}

				msg := <-nd.inbox
				tie := msg.Value

				if 2*mults[nd.id] > nd.n+2*nd.s {
					nd.x = majorities[nd.id]
				} else {
					nd.x = tie
				}
			}(nodes[i])
		}
		wg2.Wait()
	}

	for i := 0; i < n; i++ {
		if !nodes[i].isByzantine {
			fmt.Println(nodes[i].x)
			break
		}
	}
}