package main

import "fmt"

type Position struct {
	X int
	Y int
}

func (p Position) String() string {
	return fmt.Sprintf("%d/%d", p.X, p.Y)
}

type Objet struct {
	Nom string
	Position
}

func (o Objet) String() string {
	return fmt.Sprintf("%s: %v", o.Nom, o.Position)
}

var couleur = map[string]int{
	"YELLOW": 100,
	"GREEN":  200,
	"BLUE":   500,
}

type Colis struct {
	Objet
	Couleur string
}

func (c Colis) String() string {
	return fmt.Sprintf("%v %s", c.Objet, c.Couleur)
}

type Camion struct {
	Objet
	ChargeActuel   int
	ChargeMax      int
	Disponible     int
	TempsLivraison int
}

func (c Camion) String() string {
	return fmt.Sprintf("%v %d/%d %d/%d", c.Objet, c.ChargeActuel, c.ChargeMax, c.Disponible, c.TempsLivraison)
}

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
