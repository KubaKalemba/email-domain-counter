package csvhandler

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestReadAndCountDomains(t *testing.T) {

	currentDir, err := filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
	}
	dataFilePath := filepath.Join(currentDir, "test.csv")
	f, err := os.Open(dataFilePath)
	if err != nil {
		log.Fatal("Unable to read input file "+dataFilePath, err)
	}

	domains := ReadAndCountDomains(f)

	fmt.Println(domains)

	if len(domains) != 2 {
		t.Error("Inadequate number of different domains")
	}
	if domains[0].Value != 2 {
		t.Error("Inadequate count of gmail domains")
	}
	for i := 1; i < len(domains); i++ {
		if domains[i].Value > domains[i-1].Value {
			t.Error("List is not sorted")
		}
	}

}
