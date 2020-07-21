package words

import (
	"bytes"
	"testing"
)

const words = `word1 word2 word2 word3 word3 word3`

func TestCountWords(t *testing.T) {
	input := bytes.NewBufferString(words)
	counts := Count(input)
	if counts["word1"] != 1 {
		t.Errorf("count for word1 should be 1 is %d\n", counts["word1"])
	}
	if counts["word2"] != 2 {
		t.Errorf("count for word2 should be 2 is %d\n", counts["word2"])
	}
	if counts["word3"] != 3 {
		t.Errorf("count for word3 should be 3 is %d\n", counts["word3"])
	}
}
