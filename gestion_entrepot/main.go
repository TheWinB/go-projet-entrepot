package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/akamensky/argparse"
)

func printIntTab(tab [][]int) {
	for _, line := range tab {
		for _, el := range line {
			fmt.Printf("%d    ", el)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n\n\n")
}

func getFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", errors.New("Impossible de lire ce fichier.")
	}
	return string(content), nil
}

func main() {
	parser := argparse.NewParser("gestion_entrepot", "")
	fileFlag := parser.String("f", "filePath", &argparse.Options{Help: "input file path"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
		return
	}

	fileContent, err := getFile(*fileFlag)
	if err != nil {
		fmt.Println("Erreur:\n", err)
		return
	}

	entrepot, err := Parse(fileContent)
	if err != nil {
		fmt.Println("Erreur:\n", err)
		return
	}
	run(&entrepot)
}
