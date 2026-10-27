package main

import (
	"fmt"
	"time"
)

type Clock struct{ offset time.Duration }

func (c *Clock) Now() time.Time { return time.Now().Add(c.offset) }

func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2
	delta = t4.Sub(t1) - t3.Sub(t2)
	return theta, delta
}

func measure(client, server *Clock, out, back time.Duration) (time.Duration, time.Duration) {
	t1 := client.Now()
	time.Sleep(out) 
	t2 := server.Now()
	time.Sleep(5 * time.Millisecond)
	t3 := server.Now()
	time.Sleep(back)
	t4 := client.Now()
	return estimate(t1, t2, t3, t4)
}

func main() {
	const initial = 2 * time.Second
	server := &Clock{offset: 0}

	client := &Clock{offset: initial}
	theta, delta := measure(client, server, 30*time.Millisecond, 50*time.Millisecond)
	fmt.Println("theta:", theta, "delta:", delta)

	client.offset += theta
	fmt.Println("offset dupa corectie:", client.offset)
	fmt.Println("offset-ul a scazut:", client.offset > -initial && client.offset < initial)

	outs := []time.Duration{30, 55, 12, 48, 40, 22, 60, 18, 35, 50}
	backs := []time.Duration{50, 20, 14, 35, 52, 45, 25, 58, 30, 15}

	c := &Clock{offset: initial}
	var firstTheta, bestTheta, bestDelta time.Duration
	for i := range outs {
		t, d := measure(c, server, outs[i]*time.Millisecond, backs[i]*time.Millisecond)
		fmt.Println("masuratoarea", i+1, "- delta:", d, "theta:", t)
		if i == 0 {
			firstTheta = t
		}
		if i == 0 || d < bestDelta {
			bestDelta = d
			bestTheta = t
		}
	}

	fmt.Println("o singura masuratoare: offset ramas =", initial+firstTheta)
	fmt.Println("delta minim (", bestDelta, "): offset ramas =", initial+bestTheta)
}