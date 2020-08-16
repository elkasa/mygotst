package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func mswith(s string) {
	switch st := s[0:2]; st {
	case "AE":
		fmt.Println("ae")
	case "EM":
		fmt.Println("em")
	case "BF":
		fmt.Println("bf")
	default:
		fmt.Println("default")
	}

}

func main() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "INFO: ", log.Lshortfile)
		infof  = func(info string) {
			logger.Output(2, info)
		}
	)

	infof("Hello world")
	fmt.Print(&buf)

	files, err := ioutil.ReadDir("./tst")
	if err != nil {
		log.Fatal(err)
	}
	for i, file := range files {
		fmt.Println(file.Name())
		content, err := ioutil.ReadFile("./tst/" + file.Name())
		check(err)
		scanner := bufio.NewScanner(bytes.NewReader(content))
		j := 0
		for scanner.Scan() {
			j++
			fmt.Println(" line :  ", strconv.Itoa(j))
			line := scanner.Text()

			//fmt.Printf("%T\n", line)
			//fmt.Println(scanner.Text()) // Println will add back the final '\n'
			fmt.Println(" line size ==> :  ", strconv.Itoa(len(line)))
			if len(line) != 0 {
				mswith(line)
				fmt.Println(line)
			} else {
				continue
			}
		}
		check(err)
		//message := []byte(content)
		fmt.Println(" i :  ", strconv.Itoa(i))
		//fmt.Printf("File contents: %s", content)

	}

}
