package main

import (
	"bufio"
	"os"
	"testing"
)

func TestOpenfile(t *testing.T) {
	// Try to open the file.
	_, err := os.Open("../test/named_stats.txt")
	if err != nil {
		t.Error("Error opening file")
	}

}

func TestParseOutput(t *testing.T) {
	// Test that the last line has the correct output.
	file, _ := os.Open("../test/named_stats.txt")

	scanner := bufio.NewScanner(file)

	lines := parseOutput(scanner)
	if lines[len(lines)-1] != "1551387572 Statistics_Dump Per_Zone_Query TCP/IPv4_recv_errors 1" {
		t.Error("No match on last line")
	}
}

func TestGetSpecificData(t *testing.T) {
	// Test that the function should match only the selected option.
	lines := []string{"1551387526 Statistics_Dump Incoming A 23715921",
		"1551387526 Statistics_Dump Incoming NS 742175",
		"1551387526 Statistics_Dump Incoming CNAME 1382",
		"1551387526 Statistics_Dump Incoming SOA 4576810",
		"1551387526 Statistics_Dump Incoming PTR 517203",
		"1551387572 Statistics_Dump Incoming A 23715922"}

	out := getSpecificData("Statistics_Dump Incoming A", lines)

	if len(out) != 2 {
		t.Error("should get 2 matches")
	}

}
