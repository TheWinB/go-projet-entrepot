package main

import "fmt"

// Position is position value struct
type Position struct {
	X int
	Y int
}

func (p Position) String() string {
	return fmt.Sprintf("%d/%d", p.X, p.Y)
}
