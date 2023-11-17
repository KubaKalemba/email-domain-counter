package main

import (
	"awesomeProject/csvhandler"
	_ "awesomeProject/csvhandler"
	"fmt"
	"log"
	"path/filepath"
)

func main() {

	currentDir, err := filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
	}

	dataFilePath := filepath.Join(currentDir, "data.csv")

	domains := csvhandler.ReadAndCountDomains(dataFilePath)
	fmt.Println(domains)
}
