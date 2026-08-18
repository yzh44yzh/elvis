package music

import (
	"testing"
)

func Test_TransposeSemitone(t *testing.T) {
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
		result, err := TransposeSemitone(pair[0])

		if err != nil {
			t.Errorf("incorrect err, expecting nil, got '%s'", err)
		}

		if result != pair[1] {
			t.Errorf("incorrect result, expecting '%s', got '%s'", pair[1], result)
		}
	}

  result, err := TransposeSemitone("aaa")
	if result != "" {
		t.Errorf("incorrect result, expecting empty string, got %s", result)
	}
	if err == nil {
		t.Errorf("expecting error, got nil")
	}
	if err.Error() != "invalid note 'aaa'" {
		t.Errorf("invalid error value %s", err)
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
		{n: "C", r: "G", i: 7},
		{n: "C#", r: "G", i: 6},
		{n: "D", r: "F", i: 3},
		{n: "D", r: "D", i: 12},
		{n: "D", r: "C#", i: 11},
		{n: "D", r: "C", i: 10},
		{n: "C", r: "C", i: 12},
		{n: "C", r: "C", i: 24},
		{n: "C", r: "D", i: 14},
		{n: "C", r: "C#", i: 25},
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

	result, err := TransposeByInterval("bbb", 1)
	if result != "" {
		t.Errorf("incorrect result, expecting empty string, got %s", result)
	}
	if err == nil {
		t.Errorf("expecting error, got nil")
	}
	if err.Error() != "invalid note 'bbb'" {
		t.Errorf("invalid error value %s", err)
	}
}
