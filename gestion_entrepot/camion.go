package main

import "fmt"

// Constant value satuts for the truck
const (
	C_ETAT_EN_ATTENTE   = "WAITING"
	C_ETAT_EN_LIVRAISON = "GONE"
)

// Camion struct with param
type Camion struct {
	Objet
	ChargeActuel    int
	ChargeEnAttente int
	ChargeMax       int
	Etat            string
	TempsRestant    int
	TempsLivraison  int
}

func (c Camion) String() string {
	return fmt.Sprintf("%s %d(%d)/%d %s/%d", c.Nom, c.ChargeActuel, c.ChargeEnAttente, c.ChargeMax, c.Etat, c.TempsLivraison)
}

func (c Camion) shouldGo(e *Entrepot) bool {
	if c.ChargeActuel > 0 && c.ChargeEnAttente == c.ChargeActuel &&
		(c.ChargeActuel > c.ChargeMax-e.PlusGrosColis || len(e.Colis) == 0) {
		return true
	}
	return false
}

func (c *Camion) getAction(entrepot *Entrepot) string {
	if c.TempsRestant > 0 {
		c.TempsRestant--
		if c.TempsRestant == 0 {
			c.Etat = C_ETAT_EN_ATTENTE
			c.ChargeActuel = 0
			c.ChargeEnAttente = 0
		}
	} else if c.shouldGo(entrepot) {
		c.TempsRestant = c.TempsLivraison
		c.Etat = C_ETAT_EN_LIVRAISON
	}
	return fmt.Sprintf("%s %s %d/%d\n", c.Nom, c.Etat, c.ChargeActuel, c.ChargeMax)
}
