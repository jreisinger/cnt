package bytes

import (
	"bufio"
	"io"
	"log"
)

// Count counts the number of ocurrences of each byte from input.
func Count(input io.Reader) map[string]int {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanBytes) // split into bytes
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return counts
}
