package main

import "fmt"

type Colis struct {
	Objet
	Couleur string
}

func (c Colis) String() string {
	return fmt.Sprintf("%v %s", c.Objet, c.Couleur)
}