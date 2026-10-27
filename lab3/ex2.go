package main

import (
	"fmt"
	"sort"
	"sync"
)

type LamportClock struct{ c int }

func (l *LamportClock) Tick() int {
	l.c++
	return l.c
}

func (l *LamportClock) Send() int {
	l.c++
	return l.c
}

func (l *LamportClock) Receive(ts int) int {
	if ts > l.c {
		l.c = ts
	}
	l.c++
	return l.c
}

type Event struct{ Proc, C int }

func less(a, b Event) bool {
	if a.C != b.C {
		return a.C < b.C
	}
	return a.Proc < b.Proc
}

func process(id int, in <-chan int, out chan<- int, log chan<- Event, bad chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	clock := &LamportClock{}
	violations := 0

	for round := 0; round < 3; round++ {
		for i := 0; i < id; i++ {
			log <- Event{id, clock.Tick()}
		}

		ts := clock.Send()
		log <- Event{id, ts}
		out <- ts

		ts = <-in
		c := clock.Receive(ts)
		log <- Event{id, c}
		if c <= ts {
			violations++
		}
	}
	bad <- violations
}

func main() {
	inbox := []chan int{make(chan int, 3), make(chan int, 3), make(chan int, 3)}
	log := make(chan Event)
	bad := make(chan int, 3)
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go process(i+1, inbox[i], inbox[(i+1)%3], log, bad, &wg)
	}

	go func() {
		wg.Wait()
		close(log)
	}()

	events := []Event{}
	for e := range log {
		events = append(events, e)
	}

	sort.Slice(events, func(i, j int) bool { return less(events[i], events[j]) })

	for _, e := range events {
		fmt.Println("C =", e.C, " P", e.Proc)
	}

	total := 0
	for i := 0; i < 3; i++ {
		total += <-bad
	}
	fmt.Println("Mesaje primite cu valoare <= cea trimisa:", total)
}