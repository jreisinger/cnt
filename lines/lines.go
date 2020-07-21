package lines

import (
	"bufio"
	"io"
	"log"
)

// Count counts the number of ocurrences of each line from input.
func Count(input io.Reader) map[string]int {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return counts
}
