package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	ae []string
	em []string
	bf []string
	dn []string
	sa []string
	es []string
	en []string
)

func mswith(s string) {
	switch st := s[0:2]; st {
	case "AE":
		ae = append(ae, s)
		fmt.Println(ae)
	case "EM":
		em = append(em, s)
	case "BF":
		bf = append(bf, s)
	case "DN":
		dn = append(dn, s)
	case "SA":
		sa = append(sa, s)
	case "ES":
		es = append(es, s)
	case "EN":
		en = append(en, s)
	default:
		fmt.Println("default")
	}

}

func readLines(filename string) ([]string, error) {
	var lines []string
	f, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 0 {
			lines = append(lines, line)
		}
	}
	return lines, scanner.Err()

}
func main() {

	lines, err := readLines("./tst/F0.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for i, line := range lines {
		fmt.Println(i, line)
	}
}
