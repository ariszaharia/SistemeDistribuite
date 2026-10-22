package main

import (
	"fmt"
)

func main() {
	data := map[string]int{
		"Andrei": 190,
		"Maria":  165,
		"Ion":    182,
		"Elena":  172,
		"George": 178,
		"Ana":    165,
	}

	type Persoana struct {
		Nume string
		Inaltime int
	}

	var lista []Persoana
	for nume, inaltime := range data {
		lista = append(lista, Persoana{nume, inaltime})
	}

	n := len(lista)
	for i := 0; i < n-1; i++ {
		for j := i; j < n-1; j++ {
			if lista[j].Inaltime < lista[j+1].Inaltime {
				lista[j], lista[j+1] = lista[j+1], lista[j]
			}
		}
	}

	for _, p := range lista {
		fmt.Println(p.Nume)
	}
}