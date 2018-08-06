package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const minimumWordsWeWant = 10

// Print dictionary matching anagrams for each #letters_per_word
func main() {
	path := "words_alpha.txt"
	for i := 4; i < 32; i++ {
		processForLength(path, i)
	}
}

// Loop through dictionary, only looking at words of len == requestLength, print anagrams
func processForLength(fname string, requestLength int) {
	//fmt.Printf("reading file: %s\n", fname)
	wordCount := 0

	wordMap := make(map[uint32][]string)

	inFile, _ := os.Open(fname)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	// For each line (one word)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) != requestLength {
			continue
		}
		wordCount++
		code := getCode(word)
		wordMap[code] = append(wordMap[code], word)
	}
	//fmt.Println(wordMap)

	//fmt.Println("wordLength, wordCount, bucketCount: ", requestLength, ",", wordCount, ",", len(wordMap))

	// Now that we have grouped words into "buckets" based on our algorithm, the rest is easy
	// for each bucket, make sure they are really anagrams (our bucket algo ignores letter repeats)
	// and if #words that are anagrams >= minimumWordsWeWant we print them
	for k := range wordMap {
		// Don't bother unless we have enough words to care about
		if len(wordMap[k]) < minimumWordsWeWant {
			continue
		}

		wordList := wordMap[k]
		//fmt.Println(k, wordList)

		// Make a map where the alphagram is the key so we can easily group
		alphagramMap := make(map[string][]string)
		for _, word := range wordList {
			alphagram := getAlphagram(word)
			alphagramMap[alphagram] = append(alphagramMap[alphagram], word)
		}
		//fmt.Println(alphagramMap)

		// For each alphagram, if wordCount >= MINIMUM_WORDS, print them
		for alphagram := range alphagramMap {
			anagramCount := len(alphagramMap[alphagram])
			if anagramCount >= minimumWordsWeWant {
				fmt.Println(requestLength, anagramCount, alphagramMap[alphagram])
			}
		}

	}
}

// Basic alphagram so we can compare words to see if they are anagrams
func getAlphagram(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

// get the "Letter Bitmap" of a word for easier comparison
// All anagrams have the same letter bitmap (but not vice versa since we ignore repeats)
// Letter a is 1, b is 10, c is 100 in binary etc..
func getCode(word string) uint32 {
	charA := byte('a')
	charZ := byte('z')
	code := uint32(0)
	//fmt.Println("AZ", charA, charZ)
	for _, c := range word {
		if byte(c) >= charA && byte(c) <= charZ {
			bitIndex := byte(c) - charA
			code |= 1 << bitIndex
		}

	}
	//fmt.Printf("%b\n", code)
	return code
}
