package fs

import (
	"fmt"
	"os"
	"strings"
)

// this function in the case of 2 arguments after the program's name
func FsArgument(text string, banner string) {
	name := ""
	if os.Args[2] == "thinkertoy" || os.Args[2] == "standard" || os.Args[2] == "shadow" {
		name = os.Args[2]
	} else {
		fmt.Println("incorrect banner")
		return
	}
	file := Read_file(name)
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
