package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ParserError struct {
	s string
}

func (e ParserError) Error() string {
	return e.s
}

func printTab(tab []string) {
	fmt.Println("[")
	for i, v := range tab {
		fmt.Println(i, ": ", v)
	}
	fmt.Println("]")
}

func getEntrepot(e *Entrepot, s string) error {
	stringTab := strings.Fields(s)
	if len(stringTab) != 3 {
		return errors.New(fmt.Sprintf("Entrepot:\nNombre incorrect d'information, 3 attendu, %d reçu", len(stringTab)))
	}
	i, err := strconv.Atoi(stringTab[0])
	if err != nil {
		return errors.New(fmt.Sprintf("Entrepot:\nLa largeur n'est pas un chiffre: %s", stringTab[0]))
	}
	if i <= 0 {
		return errors.New(fmt.Sprintf("Entrepot:\nLa largeur doit être positive: %d", i))
	}
	e.Largeur = i
	i, err = strconv.Atoi(stringTab[1])
	if err != nil {
		return errors.New(fmt.Sprintf("Entrepot:\nLa longueur n'est pas un chiffre: %s", stringTab[1]))
	}
	if i <= 0 {
		return errors.New(fmt.Sprintf("Entrepot:\nLa longueur doit être positive: %d", i))
	}
	e.Longueur = i
	i, err = strconv.Atoi(stringTab[2])
	if err != nil {
		return errors.New(fmt.Sprintf("Entrepot:\nLa duréé de la simulation n'est pas un chiffre: %s", stringTab[2]))
	}
	if i < 10 || i > 100000 {
		return errors.New(fmt.Sprintf("Entrepot:\nLa duréé de la simulation doit être comprise entre 10 et 100'000: %d", i))
	}
	e.Temps = i
	return nil
}

func getPosition(x, y string) (Position, error) {
	p := Position{}
	i, err := strconv.Atoi(x)
	if err != nil {
		return p, errors.New(fmt.Sprintf("Position:\nLa position X n'est pas un chiffre: %s", x))
	}
	if i < 0 {
		return p, errors.New(fmt.Sprintf("Position:\nLa position X doit être supérieur à 0: %s", x))
	}
	p.X = i
	i, err = strconv.Atoi(y)
	if err != nil {
		return p, errors.New(fmt.Sprintf("Position:\nLa position Y n'est pas un chiffre: %s", y))
	}
	if i < 0 {
		return p, errors.New(fmt.Sprintf("Position:\nLa position Y doit être supérieur à 0: %s", y))
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
			return errors.New(fmt.Sprintf("Colis \"%s\":\n%s", c.Nom, err.Error()))
		}
		if p.X >= e.Longueur {
			return errors.New(fmt.Sprintf("Colis \"%s\":\nla position X est plus grande que la taille de l'entrepot: %d", c.Nom, p.X))
		}
		if p.Y >= e.Largeur {
			return errors.New(fmt.Sprintf("Colis \"%s\":\nla position X est plus grande que la taille de l'entrepot: %d", c.Nom, p.Y))
		}
		c.Position = p
		_, ok := COULEUR_POIDS_MAP[strings.ToUpper(stringTab[3])]
		if !ok {
			return errors.New(fmt.Sprintf("Colis \"%s\":\nLa couleur %s n'est pas valide", c.Nom, stringTab[3]))
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
			return errors.New(fmt.Sprintf("Transpalette \"%s\":\n%s", t.Nom, err.Error()))
		}
		if p.X >= e.Longueur {
			return errors.New(fmt.Sprintf("Transpalette \"%s\":\nla position X est plus grande que la taille de l'entrepot: %d", t.Nom, p.X))
		}
		if p.Y >= e.Largeur {
			return errors.New(fmt.Sprintf("Transpalette \"%s\":\nla position X est plus grande que la taille de l'entrepot: %d", t.Nom, p.Y))
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
			return errors.New(fmt.Sprintf("Camion: le nombre d'information est incorrect: %d %s", len(stringTab), v))
		}
		c := Camion{}
		c.Nom = stringTab[0]
		p, err := getPosition(stringTab[1], stringTab[2])
		if err != nil {
			return errors.New(fmt.Sprintf("Camion \"%s\":\n%s", c.Nom, err.Error()))
		}
		if p.X >= e.Longueur {
			return errors.New(fmt.Sprintf("Camion \"%s\":\nla position X est plus grande que la taille de l'entrepot: %d", c.Nom, p.X))
		}
		if p.Y >= e.Largeur {
			return errors.New(fmt.Sprintf("Camion \"%s\":\nla position X est plus grande que la taille de l'entrepot: %d", c.Nom, p.Y))
		}
		c.Position = p
		i, err := strconv.Atoi(stringTab[3])
		if err != nil {
			return errors.New(fmt.Sprintf("Camion:\nLa charge max n'est pas un chiffre: %s", stringTab[3]))
		}
		if i < 0 {
			return errors.New(fmt.Sprintf("Camion:\nLa charge max doit être supérieur à 0: %s", stringTab[3]))
		}
		c.ChargeMax = i
		i, err = strconv.Atoi(stringTab[4])
		if err != nil {
			return errors.New(fmt.Sprintf("Camion:\nLa durée de livraison n'est pas un chiffre: %s", stringTab[4]))
		}
		if i < 0 {
			return errors.New(fmt.Sprintf("Camion:\nLa durée de livraison doit être supérieur à 0: %s", stringTab[4]))
		}
		c.TempsLivraison = i
		c.Etat = C_ETAT_EN_ATTENTE
		e.Camions = append(e.Camions, c)
	}
	return nil
}

func Parse(fileContent string) (Entrepot, error) {
	e := Entrepot{}
	contentTab := strings.Split(fileContent, "\n")
	if len(contentTab) == 0 {
		return e, errors.New("Parseur:\nLe fichier est vide")
	}
	if err := getEntrepot(&e, contentTab[0]); err != nil {
		return e, errors.New(fmt.Sprintf("Parseur:\n%s", err.Error()))
	}
	contentTab = contentTab[1:]
	if err := getColis(&e, contentTab); err != nil {
		return e, errors.New(fmt.Sprintf("Parseur:\n%s", err.Error()))
	}
	contentTab = contentTab[len(e.Colis):]
	if err := getTranspalette(&e, contentTab); err != nil {
		return e, errors.New(fmt.Sprintf("Parseur:\n%s", err.Error()))
	}
	contentTab = contentTab[len(e.Transpalettes):]
	if err := getCamion(&e, contentTab); err != nil {
		return e, errors.New(fmt.Sprintf("Parseur:\n%s", err.Error()))
	}
	return e, nil
}
