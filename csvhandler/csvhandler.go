package csvhandler

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strings"
)

type KeyValue struct {
	Key   string
	Value int
}

func ReadAndCountDomains(f *os.File) []KeyValue {
	records := ReadCsvFile(f)
	domainCounts, err := countEmailDomains(records)
	if err != nil {
		return nil
	}
	return domainCounts
}

func ReadCsvFile(f *os.File) [][]string {
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for ", err)
	}
	return records
}

func countEmailDomains(records [][]string) ([]KeyValue, error) {
	emailDomainCount := make(map[string]int)
	for _, record := range records {
		if len(record) > 1 {
			email := record[2]
			domain := extractEmailDomain(email)
			if domain == "" {
				continue
			}
			emailDomainCount[domain]++
		}
	}
	return sortMapByDomain(emailDomainCount), nil
}

func extractEmailDomain(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) == 2 {
		return strings.ToLower(parts[1])
	}
	return ""
}

func sortMapByDomain(inputMap map[string]int) []KeyValue {
	var domainCounts []KeyValue
	for key, value := range inputMap {
		domainCounts = append(domainCounts, KeyValue{key, value})
	}
	sort.Slice(domainCounts, func(i, j int) bool {
		return domainCounts[i].Value > domainCounts[j].Value
	})
	return domainCounts
}
