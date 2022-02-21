package main

import "fmt"

type Colis struct {
	Objet
	Couleur   string
	ChoisiPar string
}

func (c Colis) String() string {
	return fmt.Sprintf("%v %s", c.Objet, c.Couleur)
}
