package main

import (
	"awesomeProject/csvhandler"
	_ "awesomeProject/csvhandler"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	test0()
}

func test0() {
	currentDir, err := filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
	}

	dataFilePath := filepath.Join(currentDir, "data.csv")

	f, err := os.Open(dataFilePath)
	if err != nil {
		log.Fatal("Unable to read input file "+dataFilePath, err)
	}
	defer f.Close()

	domains := csvhandler.ReadAndCountDomains(f)
	fmt.Println(domains)
}
