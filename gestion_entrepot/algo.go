package main

import (
	"fmt"
	"sync"
)

func run(entrepot *Entrepot) (int, error) {
	tour := 1
	fmt.Println(*entrepot)
	for len(entrepot.Colis) > 0 && entrepot.Temps <= tour {
		fmt.Printf("tour %d\n", tour)
		// find objective foreach transpalette
		for _, transpalette := range entrepot.Transpalettes {
			if transpalette.AObjectif == false {
				transpalette.getObjectif(entrepot)
			}
		}

		// cal path map foreach transpalette
		var wg sync.WaitGroup
		for _, transpalette := range entrepot.Transpalettes {
			if transpalette.ADestination == false {
				wg.Add(1)
				go transpalette.generatePathMap(*entrepot, &wg)
			}
		}
		wg.Wait()

		// cal path foreach transpalette
		for _, transpalette := range entrepot.Transpalettes {
			if transpalette.ADestination == false {
				wg.Add(1)
				go transpalette.getPath(&wg)
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
