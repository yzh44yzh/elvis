package main

import (
	"com.github.yzh44yzh.elvis/music"
	"fmt"
)

func main() {
	var n music.Note = "E"
	n, _ = music.Transpose(n)
	fmt.Println(n)

	var i music.Interval = 7
	n, _ = music.TransposeByInterval(n, i)
	fmt.Println(n)
}
