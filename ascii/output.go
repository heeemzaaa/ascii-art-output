package fs

import (
	"fmt"
	"os"
	"strings"
)

// this function works in the case of 3 arguments after the program's name
func Output(flag string, text string, banner string) {
	name := ""
	if os.Args[3] == "thinkertoy" || os.Args[3] == "standard" || os.Args[3] == "shadow" {
		name = os.Args[3]
	} else {
		fmt.Println("incorrect banner")
		return
	}
	file := Read_file(name)
	if file == nil {
		return
	}
	line := os.Args[2]
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

	if len(os.Args) > 1 && (strings.HasPrefix(os.Args[1], "--output=")) && (strings.HasSuffix(os.Args[1], ".txt")) {
		filename := MySplit(os.Args[1])
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("Error creating a file:", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(finalResult)
		if err != nil {
			fmt.Println("Error writing the result")
			return
		}
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
}

// this function gives me the part of the string after --output=
func MySplit(s string) string {
	slice := []rune(s)

	for i := 0; i < len(slice); i++ {
		if slice[i] == '=' {
			slice = slice[i+1:]
			break
		}
	}
	return string(slice)
}
