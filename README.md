`cnt` module has packages for counting occurrences of lines, words and bytes.

Sample usage:

```
// Dup prints count and text of lines that appear more than once in stdin.
package main

import (
	"fmt"
	"os"

	"github.com/jreisinger/cnt/lines"
)

func main() {
	counts := lines.Count(os.Stdin)
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
```