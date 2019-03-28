package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func printBuffer(s *bufio.Scanner) {
	for s.Scan() {
		fmt.Printf("Fields are: %q", strings.Fields(s.Text()))
		fmt.Println()
	}
}

// Parse lines using matches to parts of the line.
// other option is to use regex, this might impact performance.
func parseOutput(s *bufio.Scanner) []string {
	var Times, Title, SubTitle, Variable, Value string
	var rValue []string

	// Go over each line.
	for s.Scan() {
		lineF := strings.Fields(s.Text())
		switch lineF[0] {
		case "---":
			continue
		case "+++":
			Title = strings.Join(lineF[1:len(lineF)-2], "_")
			Times = lineF[len(lineF)-1]
		case "++":
			s2 := strings.Fields(s.Text())
			SubTitle = strings.Join(s2[1:len(lineF)-2], "_")

		default:
			// add only if doesn't start with "["
			switch strings.Contains(lineF[0], "[") {
			case true:
				continue
			case false:
				Variable = strings.Join(lineF[1:], "_")
				Value = lineF[0]

			}

		}

		// If value is not empty, create and append the line.
		if Value != "" {
			line := strings.Join([]string{Times[1 : len(Times)-1], Title, SubTitle, Variable, Value}, " ")
			rValue = append(rValue, line)
		}

	}
	return rValue
}

func getSpecificData(input string, data []string) []string {

	rValue := []string{}

	for _, line := range data {
		if input == strings.Join(strings.Split(line, " ")[1:4], " ") {
			rValue = append(rValue, line)
		}
	}

	return rValue
}

// List available options to query from the rndc statistics file.
func listAvailableValues(data []string) {

	nList := []string{}

	for _, line := range data {
		cline := strings.Join(strings.Split(line, " ")[1:4], " ")
		nList = append(nList, cline)

	}

	nList = removeDuplicates(nList)

	sort.Strings(nList)

	for _, line := range nList {
		fmt.Println(line)
	}

}

// Remove duplicated lines from the string.
func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func main() {
	// Intialize variables using Args
	statisticsFile := flag.String("sFile", "/var/named/data/named.stats", "Location of the statistics file")
	options := flag.Bool("list", false, "List all the available values to query")
	inputOption := flag.String("option", "", "Selected option")

	// Call the parser for the file
	flag.Parse()

	file, err := os.Open(*statisticsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	pOutput := parseOutput(scanner)
	if *options == true {
		listAvailableValues(pOutput)
		os.Exit(0)
	}

	if *inputOption != "" {
		list := getSpecificData(*inputOption, pOutput)
		nl := list[len(list)-1]
		fmt.Println(strings.Split(nl, " ")[4])
		os.Exit(0)
	}

}
