package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/akamensky/argparse"
)

func getFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", errors.New("impossible de lire ce fichier")
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
		fmt.Println("😱\nErreur:\n", err)
		return
	}

	entrepot, err := Parse(fileContent)
	if err != nil {
		fmt.Println("😱\nErreur:\n", err)
		return
	}
	if err = run(&entrepot); err != nil {
		fmt.Println("😱\nErreur:\n", err)
		return
	}
}
