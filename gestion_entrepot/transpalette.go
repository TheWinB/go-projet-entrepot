package main

import (
	"fmt"
	"sync"
)

const (
	T_ACTION_GO   = "GO"
	T_ACTION_TAKE = "TAKE"
	T_ACTION_WAIT = "WAIT"
)

type Transpalette struct {
	pathMap [][]int
	Objet
	action     string
	AObjectif  bool
	AColis     bool
	AChemin    bool
	Colis      Colis
	Desination Position
	Chemin     []Position
}

func (t Transpalette) String() string {
	return fmt.Sprintf("Transpalette: %s, %s, %v, COLIS: %v", t.Nom, t.action, t.Position, t.Colis)
}

func (t Transpalette) checkPathMapRight(p Position, value int) bool {
	if p.X+1 < len(t.pathMap[0])+1 {
		return t.pathMap[p.Y][p.X+1] == value
	}
	return false
}

func (t Transpalette) checkPathMapLeft(p Position, value int) bool {
	if p.X-1 >= 0 {
		return t.pathMap[p.Y][p.X-1] == value
	}
	return false
}

func (t Transpalette) checkPathMapUp(p Position, value int) bool {
	if p.Y-1 >= 0 {
		return t.pathMap[p.Y-1][p.X] == value
	}
	return false
}

func (t Transpalette) checkPathMapDown(p Position, value int) bool {
	if p.Y+1 < len(t.pathMap)+1 {
		return t.pathMap[p.Y+1][p.X] == value
	}
	return false
}

func (t *Transpalette) InitPathMap(entrepot Entrepot) {
	// Reset pathMap with only 0's
	t.pathMap = make([][]int, entrepot.Largeur)
	for i := 0; i < len(t.pathMap); i++ {
		t.pathMap[i] = make([]int, entrepot.Longueur)
	}
	for i := 0; i < len(entrepot.Colis); i++ {
		if entrepot.Colis[i].ChoisiPar != t.Nom {
			pos := entrepot.Colis[i].Position
			t.pathMap[pos.Y][pos.X] = -1
		}
	}
	for i := 0; i < len(entrepot.Transpalettes); i++ {
		if entrepot.Transpalettes[i].Nom != t.Nom {
			pos := entrepot.Transpalettes[i].Position
			t.pathMap[pos.Y][pos.X] = -1
		}
	}
}

func (t *Transpalette) generatePathMap(entrepot Entrepot, wg *sync.WaitGroup) {
	defer wg.Done()

	t.InitPathMap(entrepot)

	count := 1
	stack := make([]Position, 0)
	stack = append(stack, t.Position)
	for len(stack) > 0 {
		for _, position := range stack {
			if t.checkPathMapRight(position, 0) {
				t.pathMap[position.Y][position.X+1] = count
				stack = append(stack, Position{position.X + 1, position.Y})
			}
			if t.checkPathMapDown(position, 0) {
				t.pathMap[position.Y+1][position.X] = count
				stack = append(stack, Position{position.X, position.Y + 1})
			}
			if t.checkPathMapLeft(position, 0) {
				t.pathMap[position.Y][position.X-1] = count
				stack = append(stack, Position{position.X - 1, position.Y})
			}
			if t.checkPathMapUp(position, 0) {
				t.pathMap[position.Y-1][position.X] = count
				stack = append(stack, Position{position.X, position.Y - 1})
			}
			count++
		}
	}
}

func (t *Transpalette) getPath(wg *sync.WaitGroup) {
	defer wg.Done()

	stack := make([]Position, 0)
	count := t.pathMap[t.Desination.Y][t.Desination.X]
	position := t.Desination
	for count > 0 {
		if t.checkPathMapRight(position, count-1) && !t.checkPathMapRight(position, -1) {
			position.X = position.X + 1
			stack = append(stack, position)
		} else if t.checkPathMapDown(position, count-1) && !t.checkPathMapDown(position, -1) {
			position.Y = position.Y + 1
			stack = append(stack, position)
		} else if t.checkPathMapLeft(position, count-1) && !t.checkPathMapLeft(position, -1) {
			position.X = position.X - 1
			stack = append(stack, position)
		} else if t.checkPathMapUp(position, count-1) && !t.checkPathMapUp(position, -1) {
			position.Y = position.Y - 1
			stack = append(stack, position)
		}
		count--
	}
	//reverse list
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}
	t.AChemin = true
	t.Chemin = stack
}

func (t *Transpalette) findBestColis(entrepot *Entrepot) {
	// look for closest colis

	// look position of first coli
	objectif := &entrepot.Colis[0]
	distance := objectif.X + objectif.Y
	for _, coli := range entrepot.Colis {
		if coli.Position.X+coli.Position.Y < distance && coli.ChoisiPar == "" {
			objectif = &coli
			distance = coli.X + coli.Y
		}
	}
	// can still be the first so check if not taken to be sure
	if objectif.ChoisiPar == "" {
		t.AObjectif = true
		t.Desination = objectif.Position
		objectif.ChoisiPar = t.Objet.Nom
	} else {
		t.AObjectif = false
		t.Desination = t.Position
	}
}

func (t *Transpalette) findBestCamion(entrepot *Entrepot) {
	// look for truck that can take the charge
	charge := COULEUR_POIDS_MAP[t.Colis.Couleur]
	objectif := &entrepot.Camions[0]
	distance := objectif.X + objectif.Y
	found := false
	for _, camion := range entrepot.Camions {
		if charge+camion.ChargeActuel <= camion.ChargeMax && camion.Position.X+camion.Position.Y < distance && camion.Etat == C_ETAT_EN_ATTENTE {
			objectif = &camion
			distance = camion.Position.X + camion.Position.Y
			found = true
		}
	}
	// can still be first so check
	if found == true && objectif.Etat == C_ETAT_EN_ATTENTE && charge+objectif.ChargeActuel <= objectif.ChargeMax {
		t.Desination = objectif.Position
		t.AObjectif = true
	} else {
		// look for closest incoming truck
		objectif = &entrepot.Camions[0]
		distance = objectif.X + objectif.Y
		for _, camion := range entrepot.Camions {
			if charge+camion.ChargeActuel <= camion.ChargeMax && camion.Position.X+camion.Position.Y < distance && camion.Etat == C_ETAT_EN_ATTENTE {
				objectif = &camion
				distance = camion.Position.X + camion.Position.Y
			}
		}
		t.Desination = objectif.Position
		t.AObjectif = true
	}
}

func (t *Transpalette) getObjectif(entrepot *Entrepot) {
	if t.AColis == false && entrepot.availableColis() != 0 {
		t.findBestColis(entrepot)
	} else if t.AColis == true {
		t.findBestCamion(entrepot)
	}
	t.AObjectif = false
}
