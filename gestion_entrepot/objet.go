package main

import "fmt"

// Objet is default Objet with his Position
type Objet struct {
	Nom string
	Position
}

func (o Objet) String() string {
	return fmt.Sprintf("%s: %v", o.Nom, o.Position)
}
