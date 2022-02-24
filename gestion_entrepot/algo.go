package main

import "fmt"

func run(entrepot *Entrepot) (int, error) {
	//tour := 1
	for _, transpalette := range entrepot.Transpalettes {
		transpalette.getObjectif(entrepot)
	}
	fmt.Println(*entrepot)
	// for len(entrepot.Colis) > 0 && entrepot.Temps <= tour {
	// 	fmt.Println("tour %d", tour)

	// 	// find objective foreach transpalette

	// 	// cal path map foreach transpalette
	// 	var wg sync.WaitGroup
	// 	for _, transpalette := range entrepot.Transpalettes {
	// 		wg.Add(1)
	// 		go transpalette.generatePathMap(*entrepot, &wg)
	// 	}
	// 	wg.Wait()

	// 	// cal path foreach transpalette
	// 	for _, transpalette := range entrepot.Transpalettes {
	// 		go transpalette.generatePathMap(*entrepot)
	// 	}

	// 	// check colisions

	// 	// set actions

	// 	// print
	// 	tour++
	// }
	return 0, nil
}
