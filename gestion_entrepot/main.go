package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("gestion_entrepot", "")
	fileFlag := parser.String("f", "file", &argparse.Options{Help: "input file path"})
	err := parser.Parse(os.Args)

	if err != nil {
		log.Fatalln(parser.Usage(err))
	}
	if fileFlag != nil {
		fmt.Println(*fileFlag)
	}
	return
}
