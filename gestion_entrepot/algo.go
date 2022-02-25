package main

import (
	"fmt"
	"sync"
)

func checkEnd(e *Entrepot) bool {
	if len(e.Colis) != 0 {
		return false
	}
	for _, t := range e.Transpalettes {
		if t.Objectif != "" {
			return false
		}
	}
	for _, c := range e.Camions {
		if c.Etat != C_ETAT_EN_ATTENTE {
			return false
		}
	}
	return true
}

func run(entrepot *Entrepot) error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("😱\nErreur dans l'algo:", r)
		}
	}()
	tour := 0
	end := false
	for entrepot.Temps > tour && !end {
		tour++
		// find objective foreach transpalette
		for i, t := range entrepot.Transpalettes {
			if t.Objectif == "" {
				entrepot.Transpalettes[i].getObjectif(entrepot)
			}
		}
		// cal path map foreach transpalette
		var wg sync.WaitGroup
		for i, t := range entrepot.Transpalettes {
			if !t.AChemin {
				wg.Add(1)
				entrepot.Transpalettes[i].generatePathMap(*entrepot, &wg)
			}
		}

		// cal path foreach transpalette
		for i, t := range entrepot.Transpalettes {
			if !t.AChemin {
				wg.Add(1)
				entrepot.Transpalettes[i].getPath(&wg)
			}
		}
		wg.Wait()

		tourStr := fmt.Sprintf("tour %d\n", tour)
		for i := range entrepot.Transpalettes {
			tourStr += entrepot.Transpalettes[i].getAction(entrepot, entrepot.Transpalettes[:i])
		}
		for i := range entrepot.Camions {
			tourStr += entrepot.Camions[i].getAction(entrepot)
		}
		end = checkEnd(entrepot)
		if !end {
			fmt.Print(tourStr)
		}
		fmt.Println()
	}
	if entrepot.Temps == tour {
		fmt.Println("🙂")
	} else {
		fmt.Println("😎")
	}
	return err
}
