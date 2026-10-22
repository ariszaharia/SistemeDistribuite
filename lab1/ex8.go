package main
import "fmt"

func main() {
	angajati := map[string]int{
	"Ana Popescu": 1,
	"Mihai Ionescu": 2,
	"Bogdan Radu": 1,
	"Elena Dumitrescu": 3,
	"Cristian Vasile": 2,
	"Ioana Enache": 1,
	}

	departamente := make(map[int][]string)

	for nume, id_departament := range angajati{
		departamente[id_departament] = append(departamente[id_departament], nume)
	}

	for id, lista := range departamente{
		fmt.Printf("%d: %v\n", id, lista)
	}
}