package music

import (
	"fmt"
)

type Note string

const octave int = 12
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

	for i, currNote := range SharpNotes {
		if currNote == note {
			nextNoteIdx := i + interval
			if nextNoteIdx >= octave {
				nextNoteIdx -= octave
			}
			return SharpNotes[nextNoteIdx], nil
		}
	}
	return "", fmt.Errorf("invalid note '%s'", note)
}
