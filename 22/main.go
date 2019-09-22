package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var (
	raw string
)

func main() {
	loadFile()
	names := parseNames()

	sort.Strings(names)

	totalScore := 0

	for i, name := range names {

		characters := []rune(name)

		nameValue := 0

		for _, char := range characters {
			nameValue += getValue(char)
		}

		totalScore += nameValue * (i+1)

	}

	fmt.Println(totalScore)

}

func getValue(char rune) int {

	return int(char) - 64

}

func loadFile() {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadFile(dir + string(os.PathSeparator) + "names.txt")
	if err != nil {
		log.Fatal(err)
	}

	raw = string(data)

}

func parseNames() []string {

	raw = strings.ReplaceAll(raw, "\"", "")

	return strings.Split(raw, ",")

}
