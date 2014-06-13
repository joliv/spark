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
	in = []float64{1, 0, 0, 1}
	out = "█▁▁█"
	if x := Spark(in); x != out {
		t.Errorf("Spark(%v) = %v, want %v", in, x, out)
	}
	in = []float64{67, 71, 77, 85, 95, 104, 106, 105, 100, 89, 76, 66}
	out = "▁▂▃▄▆███▇▅▃▁"
	if x := Spark(in); x != out {
		t.Errorf("Spark(%v) = %v, want %v", in, x, out)
	}
}
