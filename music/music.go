package music

import (
	"errors"
)

type Interval int8
type Note string

var SharpNotes = []Note{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var FlatNotes = []Note{"Ab", "A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G"}

// TODO need direction
func TransposeByInterval(n Note, i Interval) (Note, error) {
	currNote := n
	for i > 0 {
		nextNote, err := Transpose(currNote)
		if err != nil {
			return "", err
		}
		currNote = nextNote
		i = i - 1
	}
	return currNote, nil
}

// TODO sharp or flat notes?
func Transpose(n Note) (Note, error) {
	l := len(SharpNotes)
	for i, currNote := range SharpNotes {
		if currNote == n {
			if i == l - 1 {
				return SharpNotes[0], nil
			}
			return SharpNotes[i + 1], nil
		}
	}
	return "", errors.New("invalid note")
}

// TODO tests
