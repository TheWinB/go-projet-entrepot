package main

import "fmt"

// Entrepot is representing the map
type Entrepot struct {
	Longueur      int
	Largeur       int
	Temps         int
	Colis         []Colis
	Transpalettes []Transpalette
	Camions       []Camion
	PlusGrosColis int
}

func (e Entrepot) availableColis() int {
	res := 0

	for _, coli := range e.Colis {
		if coli.ChoisiPar == "" {
			res++
		}
	}
	return res
}

func (e Entrepot) String() string {
	str := fmt.Sprintf(
		"----- ENTREPOT -----\n"+
			"Longueur: %d\n"+
			"Largeur: %d\n"+
			"Temps Restant: %d\n",
		e.Longueur,
		e.Largeur,
		e.Temps,
	)
	str += "--\nColis:\n"
	for i, c := range e.Colis {
		str += fmt.Sprintf("%d: %v\n", i, c)
	}
	str += "--\nTranspalettes:\n"
	for i, c := range e.Transpalettes {
		str += fmt.Sprintf("%d: %v\n", i, c)
	}
	str += "--\nCamions:\n"
	for i, c := range e.Camions {
		str += fmt.Sprintf("%d: %v\n", i, c)
	}
	return str
}
