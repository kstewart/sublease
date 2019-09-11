package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomWord(dictionary string) string {
	word := fmt.Sprintf("{%s}", dictionary)
	return word
}

func generateRandomTrailingNumber(size int) int {
	rand.Seed(time.Now().UnixNano())
	if size > 0 {
		return rand.Intn(size)
	} else {
		// Default
		return rand.Intn(4)
	}
}

func main() {
	adjective := getRandomWord("adjective")
	noun := getRandomWord("noun")
	trailing := generateRandomTrailingNumber(4)
	fmt.Printf("%s-%s-%04d\n", adjective, noun, trailing)
}
