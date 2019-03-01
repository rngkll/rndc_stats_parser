package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func printBuffer(s *bufio.Scanner) {
	for s.Scan() {
		fmt.Printf("Fields are: %q", strings.Fields(s.Text()))
		fmt.Println()
	}
}

func parseOutput(s *bufio.Scanner) {
	var Times, Title, SubTitle, Variable, Value string

	for s.Scan() {
		lineF := strings.Fields(s.Text())
		switch lineF[0] {
		case "+++":
			//fmt.Println("grupo de tres", s.Text())
			Title = strings.Join(lineF[1:len(lineF)-2], "_")
			Times = lineF[len(lineF)-1]
			//Title = s.Text()
		case "++":
			//fmt.Println("grupo de dos", s.Text())
			s2 := strings.Fields(s.Text())
			SubTitle = strings.Join(s2[1:len(lineF)-2], "_")

		default:

			switch strings.Contains(lineF[0], "[") {
			case true:
				continue
			case false:
				//fmt.Println("grupo numero", s.Text())
				Variable = strings.Join(lineF[1:], "_")
				Value = lineF[0]

			}

		}
		if Value != "" {
			fmt.Println(Times, Title, SubTitle, Variable, Value)

		}
	}
}
func getHeaders(s *bufio.Scanner) {
	re := regexp.MustCompile(`^\+\+\s.*\s\+\+$`)

	out := re.FindAll(s.Bytes(), 1)

	fmt.Println(out)
}

func main() {

	// You can get individual args with normal indexing.
	testFile := os.Args[1]

	file, err := os.Open(testFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//printBuffer(scanner)
	//getHeaders(scanner)
	parseOutput(scanner)

}
