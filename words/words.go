package words

import (
	"bufio"
	"io"
	"log"
)

// Count counts the number of ocurrences of each word from input.
func Count(input io.Reader) map[string]int {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords) // split into words
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return counts
}
