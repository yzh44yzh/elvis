package music

import (
	"fmt"
)

type Note string
type Interval int8

var SharpNotes = []Note{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var FlatNotes = []Note{"Ab", "A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G"}

func TransposeSemitone(n Note) (Note, error) {
	l := len(SharpNotes)
	for i, currNote := range SharpNotes {
		if currNote == n {
			if i == l-1 {
				return SharpNotes[0], nil
			}
			return SharpNotes[i+1], nil
		}
	}
	return "", fmt.Errorf("invalid note '%s'", n)
}

func TransposeByInterval(currNote Note, i Interval) (Note, error) {
	for i > 0 {
		nextNote, err := TransposeSemitone(currNote)
		if err != nil {
			return "", err
		}
		currNote = nextNote
		i -= 1
	}
	return currNote, nil
}
