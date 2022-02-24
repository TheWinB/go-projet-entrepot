package main

import (
	"fmt"
	"sync"
)

func run(entrepot *Entrepot) (int, error) {
	tour := 1
	for len(entrepot.Colis) > 0 && entrepot.Temps >= tour {
		fmt.Printf("tour %d\n", tour)
		// find objective foreach transpalette
		for i := 0; i < len(entrepot.Transpalettes); i++ {
			if entrepot.Transpalettes[i].AObjectif == false {
				entrepot.Transpalettes[i].getObjectif(entrepot)
			}
		}
		// cal path map foreach transpalette
		var wg sync.WaitGroup
		for i := 0; i < len(entrepot.Transpalettes); i++ {
			if entrepot.Transpalettes[i].AObjectif == false {
				wg.Add(1)
				entrepot.Transpalettes[i].generatePathMap(*entrepot, &wg)
			}
		}

		// cal path foreach transpalette
		for i := 0; i < len(entrepot.Transpalettes); i++ {
			if entrepot.Transpalettes[i].AObjectif == false {
				wg.Add(1)
				entrepot.Transpalettes[i].getPath(&wg)
			}
		}
		wg.Wait()

		// check colisions

		// set actions

		// print
		tour++
	}
	return 0, nil
}
