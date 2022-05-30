package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"log"
	"os"
	"sort"
)

func sortStr(str string) string {
	s := []rune(str)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	return string(s)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GenerateMappings() {
	// Open file
	file, err := os.Open("./data/scrabbleWords.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create mapping for each word, for each word length
	// Schema: { sorted letters : [words that can be formed from those letters]}
	lettersToWords3 := make(map[string][]string)
	lettersToWords4 := make(map[string][]string)
	lettersToWords5 := make(map[string][]string)
	lettersToWords6 := make(map[string][]string)
	lettersToWords7 := make(map[string][]string)
	lettersToWords8 := make(map[string][]string)
	lettersToWords9 := make(map[string][]string)
	lettersToWords10 := make(map[string][]string)
	lettersToWords11 := make(map[string][]string)
	lettersToWords12 := make(map[string][]string)
	lettersToWords13 := make(map[string][]string)
	lettersToWords14 := make(map[string][]string)
	lettersToWords15 := make(map[string][]string)
	maxLen := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		sorted := sortStr(word)
		var lettersToWords map[string][]string

		switch len(word) {
		case 2:
			continue
		case 3:
			lettersToWords = lettersToWords3
		case 4:
			lettersToWords = lettersToWords4
		case 5:
			lettersToWords = lettersToWords5
		case 6:
			lettersToWords = lettersToWords6
		case 7:
			lettersToWords = lettersToWords7
		case 8:
			lettersToWords = lettersToWords8
		case 9:
			lettersToWords = lettersToWords9
		case 10:
			lettersToWords = lettersToWords10
		case 11:
			lettersToWords = lettersToWords11
		case 12:
			lettersToWords = lettersToWords12
		case 13:
			lettersToWords = lettersToWords13
		case 14:
			lettersToWords = lettersToWords14
		case 15:
			lettersToWords = lettersToWords15
		default:
			log.Fatal("Found word of len ", len(word))
		}

		if _, in := lettersToWords[sorted]; !in {
			lettersToWords[sorted] = make([]string, 0)
		}
		lettersToWords[sorted] = append(lettersToWords[sorted], word)
		maxLen = max(maxLen, len(word))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Max word len:", maxLen)
	// fmt.Println("lettersToWords3 has len:", len(lettersToWords3))
	// fmt.Println("lettersToWords4 has len:", len(lettersToWords4))
	// fmt.Println("lettersToWords5 has len:", len(lettersToWords5))
	// fmt.Println("lettersToWords6 has len:", len(lettersToWords6))
	// fmt.Println("lettersToWords7 has len:", len(lettersToWords7))
	// fmt.Println("lettersToWords8 has len:", len(lettersToWords8))
	// fmt.Println("lettersToWords9 has len:", len(lettersToWords9))
	// fmt.Println("lettersToWords10 has len:", len(lettersToWords10))
	// fmt.Println("lettersToWords11 has len:", len(lettersToWords11))
	// fmt.Println("lettersToWords12 has len:", len(lettersToWords12))
	// fmt.Println("lettersToWords13 has len:", len(lettersToWords13))
	// fmt.Println("lettersToWords14 has len:", len(lettersToWords14))
	// fmt.Println("lettersToWords15 has len:", len(lettersToWords15))

	// Serialize sharded maps
	serializeMap("lettersToWords3", lettersToWords3)
	serializeMap("lettersToWords4", lettersToWords4)
	serializeMap("lettersToWords5", lettersToWords5)
	serializeMap("lettersToWords6", lettersToWords6)
	serializeMap("lettersToWords7", lettersToWords7)
	serializeMap("lettersToWords8", lettersToWords8)
	serializeMap("lettersToWords9", lettersToWords9)
	serializeMap("lettersToWords10", lettersToWords10)
	serializeMap("lettersToWords11", lettersToWords11)
	serializeMap("lettersToWords12", lettersToWords12)
	serializeMap("lettersToWords13", lettersToWords13)
	serializeMap("lettersToWords14", lettersToWords14)
	serializeMap("lettersToWords15", lettersToWords15)
}

// Serializes a map and writes it to filename
func serializeMap(filename string, m map[string][]string) {
	b := new(bytes.Buffer)

	e := gob.NewEncoder(b)
	err := e.Encode(m)
	check(err)

	err = os.WriteFile("./data/"+filename, b.Bytes(), 0644)
	check(err)
}

// Deserializes and returns a map stored in filename
func deserializeMap(filename string) (decodedMap map[string][]string) {
	// Read in data
	data, err := os.ReadFile("./data/" + filename)
	check(err)
	b := bytes.NewBuffer(data)
	d := gob.NewDecoder(b)

	// Decoding the serialized data
	err = d.Decode(&decodedMap)
	check(err)
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
