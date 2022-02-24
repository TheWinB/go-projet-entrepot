package main

import "fmt"

const (
	C_ETAT_EN_ATTENTE   = "WAITING"
	C_ETAT_EN_LIVRAISON = "GO"
)

type Camion struct {
	Objet
	ChargeActuel   int
	ChargeMax      int
	Etat           string
	TempsLivraison int
}

func (c Camion) String() string {
	return fmt.Sprintf("%s %d/%d %s/%d", c.Objet.Nom, c.ChargeActuel, c.ChargeMax, c.Etat, c.TempsLivraison)
}
