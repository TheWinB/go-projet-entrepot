package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func getFile() (string, error) {
	av := os.Args[1:]
	if len(av) == 0 {
		return "", errors.New("aucun nom de fichier donné")
	}
	content, err := ioutil.ReadFile(av[0])
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {
	fileContent, err := getFile()
	if err != nil {
		fmt.Println("Erreur:\n", err)
		return
	}
	entrepot, err := Parse(fileContent)
	if err != nil {
		fmt.Println("Erreur:\n", err)
		return
	}
	fmt.Println(entrepot)
}
