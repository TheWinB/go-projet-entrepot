package main

import "fmt"

var COULEUR_POIDS_MAP = map[string]int{
	"YELLOW": 100,
	"GREEN":  200,
	"BLUE":   500,
}

type Colis struct {
	Objet
	Couleur   string
	ChoisiPar string
}

func (c Colis) String() string {
	return fmt.Sprintf("%v %s choisi par: %s", c.Objet, c.Couleur, c.ChoisiPar)
}
