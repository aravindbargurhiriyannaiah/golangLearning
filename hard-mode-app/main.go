package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	arguments := os.Args
	fileName := arguments[len(arguments)-1]
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicln(err)
	} else {
		log.Printf("Contents of %v = %s", fileName, data)
	}
}
