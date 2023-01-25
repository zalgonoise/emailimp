// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package customerimporter

import (
	"encoding/csv"
	"errors"
	"os"
	"sort"
	"strings"
)

const colName = "email"
const colIdx = 2

var (
	ErrInvalidDomain   = errors.New("invalid domain name")
	ErrInvalidColCount = errors.New("invalid number of columns")
	ErrEmptySet        = errors.New("empty record set")
	ErrExtracting      = errors.New("failed to extract email rows from CSV records")
)

// Entry describes a parsed domain from a CSV file. It contains the domain name and a count
// for the number of customers with e-mail addresses for that same domain.
type Entry struct {
	Count  int
	Domain string
}

// Parse reads a CSV file from `path` in the filesystem to extract the number of occurrences
// for each present domain. Returns a slice of Entry and an error
func Parse(path string) ([]Entry, error) {
	records, err := parseCSV(path)
	if err != nil {
		return nil, err
	}

	entryMap, err := mapEmailRow(records)
	if err != nil {
		return nil, err
	}

	return sortResults(entryMap), nil
}

func parseCSV(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(f)
	return r.ReadAll()
}

func extractDomain(email string) (string, bool) {
	for idx, c := range email {
		if c == '@' {
			return email[idx+1:], true
		}
	}
	return "", false
}

func mapEmailRow(records [][]string) (map[string]int, error) {
	if len(records) == 0 {
		return nil, ErrEmptySet
	}
	if len(records[0]) < 3 {
		return nil, ErrInvalidColCount
	}

	var entries = map[string]int{}
	for _, r := range records {
		// skip CSV header
		if r[colIdx] == colName {
			continue
		}

		domain, ok := extractDomain(r[colIdx])
		if !ok {
			return nil, ErrInvalidDomain
		}
		if count, ok := entries[domain]; ok {
			entries[domain] = count + 1
			continue
		}

		entries[domain] = 1
	}

	return entries, nil
}

func sortResults(results map[string]int) []Entry {
	var (
		output = make([]Entry, len(results))
		idx    = 0
	)

	for domain, count := range results {
		output[idx] = Entry{
			Count:  count,
			Domain: domain,
		}
		idx++
	}

	sort.Slice(output, func(i, j int) bool {
		switch strings.Compare(output[i].Domain, output[j].Domain) {
		case -1:
			return true
		default:
			return false
		}
	})

	return output
}
