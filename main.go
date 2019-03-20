package main

import (
	"flag"
	"log"
)

func main() {
	inputFileName := flag.String("i", "", "input file(required)")
	outputFileName := flag.String("o", "", "output directory")
	flag.Parse()
	var input = *inputFileName
	if input == "" {
		log.Println("input file is empty")
	} else {
		log.Println("read file:", input)
	}
	var output = *outputFileName
	if output == "" {
		log.Println("output is stdout")
	} else {
		log.Println("write file:", output)
	}
}
