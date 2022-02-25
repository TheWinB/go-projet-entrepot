package main

import (
	"fmt"
	"math"
	"sync"
)

// Action list
const (
	T_ACTION_GO    = "GO"
	T_ACTION_TAKE  = "TAKE"
	T_ACTION_WAIT  = "WAIT"
	T_ACTION_LEAVE = "LEAVE"
)

// Objectif list
const (
	T_OBJECTIF_CAMION = "CAMION"
	T_OBJECTIF_COLIS  = "COLIS"
)

// Transpalette is all information related to Transpalette
type Transpalette struct {
	pathMap [][]int
	Objet
	action      string
	Objectif    string
	AChemin     bool
	Colis       Colis
	Destination Position
	Chemin      []Position
}

func (t Transpalette) String() string {
	return fmt.Sprintf("Transpalette: %s, %s, %v, COLIS: %v", t.Nom, t.action, t.Position, t.Colis)
}

func (t Transpalette) checkPathMapRight(p Position, value int) bool {
	if p.X+1 < len(t.pathMap[0]) {
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
	if p.Y+1 < len(t.pathMap) {
		return t.pathMap[p.Y+1][p.X] == value
	}
	return false
}

// InitPathMap init
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
	var newstack []Position
	for len(stack) > 0 {
		newstack = make([]Position, 0)
		for _, position := range stack {
			if t.checkPathMapRight(position, 0) {
				t.pathMap[position.Y][position.X+1] = count
				newstack = append(newstack, Position{position.X + 1, position.Y})
			}
			if t.checkPathMapDown(position, 0) {
				t.pathMap[position.Y+1][position.X] = count
				newstack = append(newstack, Position{position.X, position.Y + 1})
			}
			if t.checkPathMapLeft(position, 0) {
				t.pathMap[position.Y][position.X-1] = count
				newstack = append(newstack, Position{position.X - 1, position.Y})
			}
			if t.checkPathMapUp(position, 0) {
				t.pathMap[position.Y-1][position.X] = count
				newstack = append(newstack, Position{position.X, position.Y - 1})
			}
		}
		stack = newstack
		count++
	}
}

func (t *Transpalette) getPath(wg *sync.WaitGroup) {
	defer wg.Done()

	stack := make([]Position, 0)
	count := t.pathMap[t.Destination.Y][t.Destination.X]
	position := t.Destination
	for count > 0 {
		if t.checkPathMapRight(position, count-1) && !t.checkPathMapRight(position, -1) {
			position.X++
			stack = append(stack, position)
		} else if t.checkPathMapDown(position, count-1) && !t.checkPathMapDown(position, -1) {
			position.Y++
			stack = append(stack, position)
		} else if t.checkPathMapLeft(position, count-1) && !t.checkPathMapLeft(position, -1) {
			position.X--
			stack = append(stack, position)
		} else if t.checkPathMapUp(position, count-1) && !t.checkPathMapUp(position, -1) {
			position.Y--
			stack = append(stack, position)
		}
		count--
	}
	// reverse list
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}
	t.AChemin = len(stack) > 0
	t.Chemin = stack
}

func (t *Transpalette) findBestColis(entrepot *Entrepot) {
	// look for closest colis

	// look position of first coli
	objectif := entrepot.Colis[0]
	distance := int(math.Abs(float64(t.X-objectif.X)) + math.Abs(float64(t.Y-objectif.Y)))
	for i, coli := range entrepot.Colis {
		newDistance := int(math.Abs(float64(t.X-coli.X)) + math.Abs(float64(t.Y-coli.Y)))
		if newDistance < distance && coli.ChoisiPar == "" {
			objectif = entrepot.Colis[i]
			distance = newDistance
		}
	}
	// can still be the first so check if not taken to be sure
	if objectif.ChoisiPar == "" {
		t.Objectif = T_OBJECTIF_COLIS
		t.Destination = objectif.Position
		objectif.ChoisiPar = t.Nom
		t.Colis = objectif
	} else {
		t.Objectif = ""
		t.Destination = t.Position
	}
}

func (t *Transpalette) findBestCamion(entrepot *Entrepot) {
	// look for truck that can take the charge
	charge := CouleurPoidsMap[t.Colis.Couleur]
	objectif := &entrepot.Camions[0]
	distance := int(math.Abs(float64(t.X-objectif.X)) + math.Abs(float64(t.Y-objectif.Y)))
	found := false
	for i, camion := range entrepot.Camions {
		newDistance := int(math.Abs(float64(t.X-camion.X)) + math.Abs(float64(t.Y-camion.Y)))
		if charge+camion.ChargeEnAttente <= camion.ChargeMax && newDistance < distance && camion.Etat == C_ETAT_EN_ATTENTE {
			objectif = &entrepot.Camions[i]
			distance = newDistance
			found = true
		}
	}
	// can still be first so check
	if found {
		t.Destination = objectif.Position
		t.Objectif = T_OBJECTIF_CAMION
		objectif.ChargeEnAttente += charge
	} else {
		// look for closest incoming truck
		objectif = &entrepot.Camions[0]
		distance := int(math.Abs(float64(t.X-objectif.X)) + math.Abs(float64(t.Y-objectif.Y)))
		for i, camion := range entrepot.Camions {
			newDistance := int(math.Abs(float64(t.X-camion.X)) + math.Abs(float64(t.Y-camion.Y)))
			if charge+camion.ChargeActuel <= camion.ChargeMax && newDistance < distance && camion.Etat == C_ETAT_EN_ATTENTE {
				objectif = &entrepot.Camions[i]
				distance = newDistance
			}
		}
		t.Destination = objectif.Position
		t.Objectif = T_OBJECTIF_CAMION
		objectif.ChargeEnAttente += charge
	}
}

func (t *Transpalette) getObjectif(entrepot *Entrepot) {
	if t.Colis.Nom == "" && entrepot.availableColis() != 0 {
		t.findBestColis(entrepot)
	} else if t.Colis.Nom != "" {
		t.findBestCamion(entrepot)
	} else {
		t.Objectif = ""
	}
}

func (t *Transpalette) getAction(entrepot *Entrepot, transpalettes []Transpalette) string {
	actionStr := ""
	if t.Objectif == "" {
		t.action = T_ACTION_WAIT
		actionStr = T_ACTION_LEAVE
	} else if len(t.Chemin) == 0 {
		if t.Objectif == T_OBJECTIF_COLIS {
			t.action = T_ACTION_TAKE
			actionStr = fmt.Sprintf("%s %s %s", T_ACTION_TAKE, t.Colis.Nom, t.Colis.Couleur)
			t.AChemin = false
			t.Objectif = ""
			for i, c := range entrepot.Colis {
				if c.Position == t.Colis.Position {
					entrepot.Colis[i] = entrepot.Colis[len(entrepot.Colis)-1]
					entrepot.Colis = entrepot.Colis[:len(entrepot.Colis)-1]
					break
				}
			}
		} else if t.Objectif == T_OBJECTIF_CAMION {
			t.action = T_ACTION_LEAVE
			actionStr = fmt.Sprintf("%s %s %s", T_ACTION_LEAVE, t.Colis.Nom, t.Colis.Couleur)
			for i, c := range entrepot.Camions {
				if t.Destination == c.Position {
					entrepot.Camions[i].ChargeActuel += CouleurPoidsMap[t.Colis.Couleur]
				}
			}
			t.Colis = Colis{}
			t.Objectif = ""
		}
	} else {
		wait := false
		for _, transpalette := range transpalettes {
			if len(transpalette.Chemin) > 0 &&
				(transpalette.action == T_ACTION_GO && transpalette.Chemin[0] == t.Chemin[0]) ||
				transpalette.Position == t.Chemin[0] {
				wait = true
			}
		}
		if wait {
			t.action = T_ACTION_WAIT
			actionStr = T_ACTION_WAIT
		} else {
			t.action = T_ACTION_GO
			actionStr = fmt.Sprintf("%s [%d,%d]", T_ACTION_GO, t.Chemin[0].X, t.Chemin[0].Y)
			t.Chemin = t.Chemin[1:]
		}
	}
	return fmt.Sprintf("%s %s\n", t.Nom, actionStr)
}
