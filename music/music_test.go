package music

import (
	"testing"
)

func Test_Transpose(t *testing.T) {
	pairs := [][]Note{
		{"A", "A#"},
		{"A#", "B"},
		{"B", "C"},
		{"D#", "E"},
		{"E", "F"},
		{"F", "F#"},
		{"G#", "A"},
	}

	for _, pair := range pairs {
		result, err := Transpose(pair[0])

		if err != nil {
			t.Errorf("incorrect err, expecting nil, got '%s'", err)
		}

		if result != pair[1] {
			t.Errorf("incorrect result, expecting '%s', got '%s'", pair[1], result)
		}
	}
}

func Test_TransposeByInterval(t *testing.T) {
	type TestSet struct {
		n Note
		r Note
		i Interval
	}

	sets := []TestSet{
		{n: "D", r: "F", i: 3},
		{n: "E", r: "G", i: 3},
		{n: "F#", r: "G#", i: 2},
		{n: "G", r: "C", i: 5},
	}

	for _, set := range sets {
		result, err := TransposeByInterval(set.n, set.i)
		if err != nil {
			t.Errorf("incorrect err, expecting nil, got '%s'", err)
		}

		if result != set.r {
			t.Errorf("incorrect result, expecting '%s', got '%s'", set.r, result)
		}
	}
}
