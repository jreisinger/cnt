package lines

import (
	"bytes"
	"testing"
)

const lines = `line1
line2
line2
line3
line3
line3`

func TestCountLines(t *testing.T) {
	input := bytes.NewBufferString(lines)
	counts := make(map[string]int)
	counts = Count(input)
	if counts["line1"] != 1 {
		t.Errorf("count for line1 should be 1 is %d\n", counts["line1"])
	}
	if counts["line2"] != 2 {
		t.Errorf("count for line2 should be 2 is %d\n", counts["line2"])
	}
	if counts["line3"] != 3 {
		t.Errorf("count for line3 should be 3 is %d\n", counts["line3"])
	}
}
