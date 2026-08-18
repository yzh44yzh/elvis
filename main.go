package main

import (
	"com.github.yzh44yzh.elvis/music"
	"fmt"
)

func main() {
	var n music.Note = "E"
	n, _ = music.TransposeSemitone(n)
	fmt.Println(n)

	var i int = 7
	n, _ = music.TransposeByInterval(n, i)
	fmt.Println(n)
}
