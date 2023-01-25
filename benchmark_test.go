package customerimporter

import (
	"bytes"
	"encoding/csv"
	"testing"

	_ "embed"
)

//go:embed testdata/customers.csv
var rawData []byte

const rawPath = "./testdata/customers.csv"

func BenchmarkParse(b *testing.B) {
	var (
		entries []Entry
		err     error
	)
	for i := 0; i < b.N; i++ {
		entries, err = Parse(rawPath)
		if err != nil {
			b.Error(err)
			return
		}
	}
	_ = entries
}

func BenchmarkMapAndSort(b *testing.B) {
	buf := bytes.NewBuffer(rawData)
	r := csv.NewReader(buf)

	records, err := r.ReadAll()
	if err != nil {
		b.Errorf("failed to parse CSV data: %v", err)
		return
	}

	var output []Entry
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		entryMap, err := mapEmailRow(records)
		if err != nil {
			b.Error(err)
			return
		}
		output = sortResults(entryMap)
	}
	_ = output
}
