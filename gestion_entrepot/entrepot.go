package main

import "fmt"

type Entrepot struct {
	Longueur      int
	Largeur       int
	Temps         int
	Colis         []Colis
	Transpalettes []Transpalette
	Camions       []Camion
}

func (e Entrepot) String() string {
	return fmt.Sprintf(
		"----- ENTREPOT -----\n"+
			"Longueur: %d\n"+
			"Largeur: %d\n"+
			"Temps Restant: %d\n"+
			"Colis:\n"+
			"%v\n"+
			"Transpalettes\n"+
			"%v\n"+
			"Camions:\n"+
			"%v",
		e.Longueur,
		e.Largeur,
		e.Temps,
		e.Colis,
		e.Transpalettes,
		e.Camions,
	)
}
