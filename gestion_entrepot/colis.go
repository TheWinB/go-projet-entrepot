package main

import "fmt"

// CouleurPoidsMap is a map with all the possible weight and value associated
var CouleurPoidsMap = map[string]int{
	"YELLOW": 100,
	"GREEN":  200,
	"BLUE":   500,
}

// Colis is the object 'Colis' with his param
type Colis struct {
	Objet
	Couleur   string
	ChoisiPar string
}

func (c Colis) String() string {
	return fmt.Sprintf("%v %s choisi par: %s", c.Objet, c.Couleur, c.ChoisiPar)
}
