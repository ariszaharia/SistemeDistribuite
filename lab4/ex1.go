package main
import "fmt"


type Node struct {
	id       int
	isLeader bool
	in <-chan int
	out chan<- int
}

func (n *Node) Run(liderGasit chan<- int) {
	n.out <- n.id
	for primit := range n.in {
		if primit > n.id {
			n.out <- primit
		}else if primit == n.id {
			n.isLeader = true
			liderGasit <- n.id
			return
		}
	}
}

func main() {
	ids := []int{12, 45, 7, 89, 23}
	nrnodes := len(ids)

	channels := make([]chan int, nrnodes)

	for i := 0; i < nrnodes; i++ {
		channels[i] = make(chan int, nrnodes)
	}

	nodes := make([]Node, nrnodes)
	for i := 0; i < nrnodes; i++ {
		left := channels[(i-1 + nrnodes) % nrnodes]
		right := channels[i]

		nodes[i] = Node{
			id : ids[i],
			in : left,
			out : right,
		}

	}

	liderGasit := make(chan int)

	for i := 0; i < nrnodes; i++ {

			go nodes[i].Run(liderGasit)

		}

	liderID := <-liderGasit
	fmt.Printf("Lider %d\n", liderID)
}