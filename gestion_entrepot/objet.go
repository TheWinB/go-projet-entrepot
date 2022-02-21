package main

import "fmt"

type Objet struct {
	Nom string
	Position
}

func (o Objet) String() string {
	return fmt.Sprintf("%s: %v", o.Nom, o.Position)
}
