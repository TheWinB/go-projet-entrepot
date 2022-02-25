package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func getEntrepot(e *Entrepot, s string) error {
	stringTab := strings.Fields(s)
	if len(stringTab) != 3 {
		return fmt.Errorf("Entrepot:\nNombre incorrect d'information, 3 attendu, %d reçu", len(stringTab))
	}
	i, err := strconv.Atoi(stringTab[0])
	if err != nil {
		return fmt.Errorf("Entrepot:\nLa largeur n'est pas un chiffre: %s", stringTab[0])
	}
	if i <= 0 {
		return fmt.Errorf("Entrepot:\nLa largeur doit être positive: %d", i)
	}
	e.Largeur = i
	i, err = strconv.Atoi(stringTab[1])
	if err != nil {
		return fmt.Errorf("Entrepot:\nLa longueur n'est pas un chiffre: %s", stringTab[1])
	}
	if i <= 0 {
		return fmt.Errorf("Entrepot:\nLa longueur doit être positive: %d", i)
	}
	e.Longueur = i
	i, err = strconv.Atoi(stringTab[2])
	if err != nil {
		return fmt.Errorf("Entrepot:\nLa duréé de la simulation n'est pas un chiffre: %s", stringTab[2])
	}
	if i < 10 || i > 100000 {
		return fmt.Errorf("Entrepot:\nLa duréé de la simulation doit être comprise entre 10 et 100'000: %d", i)
	}
	e.Temps = i
	return nil
}

func getPosition(x, y string) (Position, error) {
	p := Position{}
	i, err := strconv.Atoi(x)
	if err != nil {
		return p, fmt.Errorf("Position:\nLa position X n'est pas un chiffre: %s", x)
	}
	if i < 0 {
		return p, fmt.Errorf("Position:\nLa position X doit être supérieur à 0: %s", x)
	}
	p.X = i
	i, err = strconv.Atoi(y)
	if err != nil {
		return p, fmt.Errorf("Position:\nLa position Y n'est pas un chiffre: %s", y)
	}
	if i < 0 {
		return p, fmt.Errorf("Position:\nLa position Y doit être supérieur à 0: %s", y)
	}
	p.Y = i
	return p, nil
}

func getColis(e *Entrepot, tab []string) error {
	for _, v := range tab {
		stringTab := strings.Fields(v)
		if len(stringTab) != 4 {
			return nil
		}
		c := Colis{}
		c.Nom = stringTab[0]
		p, err := getPosition(stringTab[1], stringTab[2])
		if err != nil {
			return fmt.Errorf("Colis \"%q\":\n%s", c.Nom, err.Error())
		}
		if p.X >= e.Longueur {
			return fmt.Errorf("Colis \"%q\":\nla position X est plus grande que la taille de l'entrepot: %d", c.Nom, p.X)
		}
		if p.Y >= e.Largeur {
			return fmt.Errorf("Colis \"%q\":\nla position X est plus grande que la taille de l'entrepot: %d", c.Nom, p.Y)
		}
		c.Position = p
		poids, ok := CouleurPoidsMap[strings.ToUpper(stringTab[3])]
		if !ok {
			return fmt.Errorf("Colis \"%q\":\nLa couleur %s n'est pas valide", c.Nom, stringTab[3])
		}
		if poids > e.PlusGrosColis {
			e.PlusGrosColis = poids
		}
		c.Couleur = strings.ToUpper(stringTab[3])
		e.Colis = append(e.Colis, c)
	}
	return nil
}

func getTranspalette(e *Entrepot, tab []string) error {
	for _, v := range tab {
		stringTab := strings.Fields(v)
		if len(stringTab) != 3 {
			return nil
		}
		t := Transpalette{}
		t.Nom = stringTab[0]
		p, err := getPosition(stringTab[1], stringTab[2])
		if err != nil {
			return fmt.Errorf("Transpalette \"%q\":\n%s", t.Nom, err.Error())
		}
		if p.X >= e.Longueur {
			return fmt.Errorf("Transpalette \"%q\":\nla position X est plus grande que la taille de l'entrepot: %d", t.Nom, p.X)
		}
		if p.Y >= e.Largeur {
			return fmt.Errorf("Transpalette \"%q\":\nla position X est plus grande que la taille de l'entrepot: %d", t.Nom, p.Y)
		}
		t.Position = p
		e.Transpalettes = append(e.Transpalettes, t)
	}
	return nil
}

func getCamion(e *Entrepot, tab []string) error {
	for _, v := range tab {
		stringTab := strings.Fields(v)
		if len(stringTab) != 5 {
			return fmt.Errorf("Camion: le nombre d'information est incorrect: %d %s", len(stringTab), v)
		}
		c := Camion{}
		c.Nom = stringTab[0]
		p, err := getPosition(stringTab[1], stringTab[2])
		if err != nil {
			return fmt.Errorf("Camion \"%q\":\n%s", c.Nom, err.Error())
		}
		if p.X >= e.Longueur {
			return fmt.Errorf("Camion \"%q\":\nla position X est plus grande que la taille de l'entrepot: %d", c.Nom, p.X)
		}
		if p.Y >= e.Largeur {
			return fmt.Errorf("Camion \"%q\":\nla position X est plus grande que la taille de l'entrepot: %d", c.Nom, p.Y)
		}
		c.Position = p
		i, err := strconv.Atoi(stringTab[3])
		if err != nil {
			return fmt.Errorf("Camion:\nLa charge max n'est pas un chiffre: %s", stringTab[3])
		}
		if i < 0 {
			return fmt.Errorf("Camion:\nLa charge max doit être supérieur à 0: %s", stringTab[3])
		}
		c.ChargeMax = i
		i, err = strconv.Atoi(stringTab[4])
		if err != nil {
			return fmt.Errorf("Camion:\nLa durée de livraison n'est pas un chiffre: %s", stringTab[4])
		}
		if i < 0 {
			return fmt.Errorf("Camion:\nLa durée de livraison doit être supérieur à 0: %s", stringTab[4])
		}
		c.TempsLivraison = i
		c.Etat = C_ETAT_EN_ATTENTE
		e.Camions = append(e.Camions, c)
	}
	return nil
}

// Parse is the principal func to Parse the map
func Parse(fileContent string) (Entrepot, error) {
	e := Entrepot{}
	contentTab := strings.Split(fileContent, "\n")
	if len(contentTab) == 0 {
		return e, errors.New("Parseur:\nLe fichier est vide")
	}
	if err := getEntrepot(&e, contentTab[0]); err != nil {
		return e, fmt.Errorf("Parseur:\n%s", err.Error())
	}
	contentTab = contentTab[1:]
	if err := getColis(&e, contentTab); err != nil {
		return e, fmt.Errorf("Parseur:\n%s", err.Error())
	}
	contentTab = contentTab[len(e.Colis):]
	if err := getTranspalette(&e, contentTab); err != nil {
		return e, fmt.Errorf("Parseur:\n%s", err.Error())
	}
	contentTab = contentTab[len(e.Transpalettes):]
	if err := getCamion(&e, contentTab); err != nil {
		return e, fmt.Errorf("Parseur:\n%s", err.Error())
	}
	return e, nil
}
