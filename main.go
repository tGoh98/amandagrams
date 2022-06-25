package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// Generate word mapping if run with cmd arg "-g"
	args := os.Args
	if len(args) == 2 && args[1] == "-g" {
		fmt.Println("Generating letterToWord data...")
		GenerateMappings()
		fmt.Println("Finished generating letterToWord data")
	} else {
		// gin.SetMode(gin.ReleaseMode) // Uncomment for release mode
		router := gin.Default()

		router.GET("/", greetings)
		router.GET("/words", getWords)
		router.GET("/randWord", randWord)

		router.Run(":8000") // https://gin-gonic.com/docs/deployment/
	}
}

// hello world for root index
func greetings(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "beat amanda",
	})
}

// randWords responds with a random word of length n
func randWord(c *gin.Context) {
	nStr := c.Query("randLen")
	n, err := strconv.Atoi(nStr)

	if err != nil || n > 15 {
		c.String(http.StatusBadRequest, "Incorrect parameter\n")
	}

	c.String(http.StatusOK, getRandWord(n))
}

// generates a random word of length n
func getRandWord(n int) string {
	ltwMap := deserializeMap("lettersToWords" + strconv.Itoa(n))

	candidateArr, in := ltwMap[sortStr(getRandLetters(n))]
	for !in {
		candidateArr, in = ltwMap[sortStr(getRandLetters(n))]
	}
	return candidateArr[rand.Intn(len(candidateArr))]
}

// generates n random letters
func getRandLetters(n int) string {
	alpha := alphabetArray()
	var sb strings.Builder

	for i := 0; i < n; i++ {
		sb.WriteString(alpha[rand.Intn(len(alpha))])
	}
	return sb.String()
}

// getWords responds with the list of all words as JSON
func getWords(c *gin.Context) {
	base := c.Query("base")
	nStr := c.Query("n")

	n, err := strconv.Atoi(nStr)

	if base == "" || err != nil {
		c.String(http.StatusBadRequest, "Incorrect parameters\n")
	}

	// c.String(http.StatusOK, "Received base: %s and n: %s\n", base, n)
	c.IndentedJSON(http.StatusOK, getAllWords(base, n))
}

// Returns a list of all words that can be formed by adding 1..n letters to base
func getAllWords(base string, n int) []string {
	// Add all words from 1..n
	allWords := make([]string, 0)
	for i := 1; i <= n; i++ {
		// TODO: proxy sharding here
		allWords = append(allWords, getWordsNAway(strings.ToUpper(base), i)...)
	}
	return allWords
}

// Returns a list of all words that can be formed by adding n letters to base
func getWordsNAway(base string, n int) []string {
	// Load up relevant data structure for n
	length := len(base) + n
	lettersToWords := deserializeMap("lettersToWords" + strconv.Itoa(length))

	// Find all valid words
	wordsNAway := make([]string, 0)
	for _, combination := range getAllCombinations(alphabetArray(), n) {
		sorted := sortStr(combination + base)
		if validWords, in := lettersToWords[sorted]; in {
			wordsNAway = append(wordsNAway, validWords...)
		}
	}
	return removeDups(wordsNAway)
}

// removes duplicates in an array
func removeDups(arr []string) []string {
	seen := make(map[string]bool)
	for _, elem := range arr {
		if _, in := seen[elem]; !in {
			seen[elem] = true
		}
	}

	res := make([]string, len(seen))
	i := 0
	for k := range seen {
		res[i] = k
		i++
	}
	return res
}

// Returns a list of all permutations of n letters, sorted alphabetically
func getAllCombinations(letters []string, n int) []string {
	// Base case
	if n <= 1 {
		return letters
	}

	// Add all letters and recurse
	allLetters := make([]string, 0)
	for _, letter := range letters {
		for _, c := range alphabetArray() {
			// Note: if really need to, can prune when n == 2. just need to pass in base and lettersToWords
			allLetters = append(allLetters, letter+c)
		}
	}
	return getAllCombinations(allLetters, n-1)
}

// Returns [A, B, C, ...]
func alphabetArray() []string {
	return strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
}
