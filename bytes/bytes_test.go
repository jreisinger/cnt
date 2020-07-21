package bytes

import (
	"bytes"
	"testing"
)

const mybytes = "abbccc"

func TestCountBytes(t *testing.T) {
	input := bytes.NewBufferString(mybytes)
	counts := Count(input)
	if counts["a"] != 1 {
		t.Errorf("count for a should be 1 is %d\n", counts["a"])
	}
	if counts["b"] != 2 {
		t.Errorf("count for b should be 2 is %d\n", counts["b"])
	}
	if counts["c"] != 3 {
		t.Errorf("count for c should be 3 is %d\n", counts["c"])
	}
}
