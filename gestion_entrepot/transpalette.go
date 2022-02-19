package main

import (
	"errors"
	"fmt"
)

type Transpalette struct {
	pathMap [][]int
	Objet
	Colis
	action string
}

func (t Transpalette) String() string {
	return fmt.Sprintf("Transpalette: %s, %s, %v, %v", t.Nom, t.action, t.Position, t.Colis)
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
	t.pathMap = make([][]int, entrepot.Longueur)
	for i := 0; i < len(t.pathMap); i++ {
		t.pathMap[i] = make([]int, entrepot.Largeur)
	}

	// place -1 for other transplattes or other collisionable objects
}

func (t *Transpalette) generatePathMap(entrepot Entrepot) {
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

func (t Transpalette) getPath(destination Position) ([]Position, error) {
	if len(t.pathMap) == 0 {
		return []Position{}, errors.New("path map was not generated")
	}
	if t.pathMap[destination.Y][destination.X] == -1 {
		return []Position{}, errors.New("there is no point in going to that position")
	}

	stack := make([]Position, 0)
	count := t.pathMap[destination.Y][destination.X]
	position := destination
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
	return stack, nil
}
