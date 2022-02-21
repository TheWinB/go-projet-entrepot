package main

import "fmt"

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
