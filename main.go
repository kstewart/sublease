package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func readLine(r io.Reader, lineNum int) (line string, lastLine int, err error) {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		lastLine++
		if lastLine == lineNum {
			// you can return sc.Bytes() if you need output in []bytes
			return sc.Text(), lastLine, sc.Err()
		}
	}
	return line, lastLine, io.EOF
}

func getRandomWord(dictionary string) string {
	randomWord := ""
	var line int
	dictFile := fmt.Sprintf("%s.txt", dictionary)

	file, err := os.Open(dictFile)
	if err != nil {
		log.Fatal(err)
	} else {
		numWords, countError := lineCounter(file)
		if countError != nil {
			fmt.Println("Cannot get line count")
		} else {
			rand.Seed(time.Now().UnixNano())
			line = rand.Intn(numWords)

			rs, resetErr := file.Seek(0, 0)
			if resetErr != nil {
				log.Fatal(resetErr)
			}

			if rs >= 0 {
				word, last, dictErr := readLine(file, line)
				if dictErr != nil {
					fmt.Printf("Cannot read line %d\n", last)
				}
				randomWord = word
			}
		}
	}

	return randomWord
}

func generateRandomTrailingNumber(size int) string {
	rand.Seed(time.Now().UnixNano())

	upper := strings.Repeat("9", size)
	upperBound, err := strconv.Atoi(upper)
	if err != nil {
		upperBound = 9999
	}
	randomNumber := rand.Intn(upperBound)
	trailing := fmt.Sprintf("%0*d", size, randomNumber)

	return trailing
}

func getSubdomainName() string {
	adjective := getRandomWord("adjectives")
	noun := getRandomWord("nouns")
	trailing := generateRandomTrailingNumber(4) // Defaulting to 4-digit random suffix for now
	subdomain := fmt.Sprintf("%s-%s-%s\n", adjective, noun, trailing)

	return subdomain
}

func main() {
	subdomain := getSubdomainName()
	fmt.Println(subdomain)
}
