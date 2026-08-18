package music

import (
	"fmt"
)

const octave int = 12

type Note string

var SharpNotes = []Note{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var FlatNotes = []Note{"Ab", "A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G"}

func TransposeSemitone(n Note) (Note, error) {
	return TransposeByInterval(n, 1)
}

func TransposeByInterval(note Note, interval int) (Note, error) {
	interval = interval % octave

	if interval < 0 {
		interval = octave + interval
	}

	for idx, currNote := range SharpNotes {
		if currNote == note {
			nextNoteIdx := idx + interval
			if nextNoteIdx >= octave {
				nextNoteIdx -= octave
			}
			return SharpNotes[nextNoteIdx], nil
		}
	}
	return "", fmt.Errorf("invalid note '%s'", note)
}

func TransposeNotes(notes []Note, interval int) ([]Note, error) {
	res := make([]Note, len(notes))
	for i, note := range notes {
		n, err := TransposeByInterval(note, interval)
		if err != nil {
			return nil, err
		}
		res[i] = n
	}
	return res, nil
}
