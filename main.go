package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
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
	// word := fmt.Sprintf("{%s}", dictionary)
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
	adjective := getRandomWord("adjectives")
	noun := getRandomWord("nouns")
	trailing := generateRandomTrailingNumber(4)
	fmt.Printf("%s-%s-%04d\n", adjective, noun, trailing)
}
