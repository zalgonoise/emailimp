package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	customerimporter "github.com/zalgonoise/emailimp"
)

func main() {
	filePath := flag.String("f", "", "path to the file to parse")
	flag.Parse()

	if *filePath == "" {
		log.Fatal("no input file provided")
		os.Exit(1)
	}

	entries, err := customerimporter.Parse(*filePath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	sb := &strings.Builder{}
	sb.WriteString("Listing entries:\n")
	for _, e := range entries {
		sb.WriteString(fmt.Sprintf("  - %s: %d\n", e.Domain, e.Count))
	}
	fmt.Print(sb.String())
	os.Exit(0)
}
