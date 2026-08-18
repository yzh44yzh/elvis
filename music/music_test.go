package music

import (
	"slices"
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
		t.Errorf("incorrect error, got '%s'", err.Error())
	}
}

func Test_TransposeByInterval(t *testing.T) {
	type TestSet struct {
		n Note
		r Note
		i int
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
		{n: "D", r: "D", i: 0},
		{n: "D", r: "D", i: 12},
		{n: "D", r: "D#", i: 13},
		{n: "D", r: "E", i: 14},
		{n: "D", r: "F", i: 15},
		{n: "D", r: "F#", i: 16},
		{n: "D", r: "D", i: 24},
		{n: "D", r: "D#", i: 25},
		{n: "D", r: "C#", i: -1},
		{n: "D", r: "C", i: -2},
		{n: "D", r: "B", i: -3},
		{n: "D", r: "A#", i: -4},
		{n: "G", r: "D", i: -5},
		{n: "G", r: "C#", i: -6},
		{n: "G", r: "C", i: -7},
		{n: "G", r: "G", i: -12},
		{n: "G", r: "F#", i: -13},
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

	result, err := TransposeByInterval("H", 1)
	if result != "" {
		t.Errorf("incorrect result, expecting empty string, got %s", result)
	}
	if err == nil {
		t.Errorf("expecting error, got nil")
	}
	if err.Error() != "invalid note 'H'" {
		t.Errorf("incorrect error, got '%s'", err.Error())
	}
}

func Test_TransposeNotes(t *testing.T) {
	notes := []Note{"C", "C#", "D", "E", "F", "F#", "G", "A#", "B"}

	type TestSet struct {
		expect   []Note
		interval int
	}

	sets := []TestSet{
		{
			expect:   []Note{"C#", "D", "D#", "F", "F#", "G", "G#", "B", "C"},
			interval: 1,
		},
		{
			expect:   []Note{"D", "D#", "E", "F#", "G", "G#", "A", "C", "C#"},
			interval: 2,
		},
		{
			expect:   []Note{"G", "G#", "A", "B", "C", "C#", "D", "F", "F#"},
			interval: 7,
		},
		{
			expect:   []Note{"A#", "B", "C", "D", "D#", "E", "F", "G#", "A"},
			interval: -2,
		},
		{
			expect:   []Note{"F", "F#", "G", "A", "A#", "B", "C", "D#", "E"},
			interval: -7,
		},
	}

	for _, set := range sets {
		res, err := TransposeNotes(notes, set.interval)
		if err != nil {
			t.Errorf("incorrect err, expecting nil, got '%s'", err)
		}
		if !slices.Equal(res, set.expect) {
			t.Errorf("incorrect res %v, expecting %v, when interval %d", res, set.expect, set.interval)
		}
	}
}
