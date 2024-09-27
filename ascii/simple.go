package fs

import (
	"fmt"
	"os"
	"strings"
)

// this function works in case that there's only one string
func SimpleArt(text string) {
	file := Read_file("standard")
	if file == nil {
		return
	}
	line := os.Args[1]
	if !Is_ascii(line) {
		fmt.Println("Non Ascii character found")
		return
	}
	if len(line) < 1 {
		return
	}
	finalResult := ""
	lines_count := Count_next_line(line)
	splitted_line := strings.Split(line, "\\n")
	splitted_line, lines_count = Cleaned_split(splitted_line, lines_count)
	finalResult = Print_art(file[1:], splitted_line, lines_count)
	fmt.Print(finalResult)
}
