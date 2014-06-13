package spark

import (
	"testing"
)

func TestSpark(t *testing.T) {
	var (
		// Wish I could use constant slices...
		in  = []float64{1, 2, 3, 4, 5, 6, 7, 8}
		out = "▁▂▃▄▅▆▇█"
	)
	if x := Spark(in); x != out {
		t.Errorf("Spark(%v) = %v, want %v", in, x, out)
	}
}
