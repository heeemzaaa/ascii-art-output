package main

import (
	"fmt"
	"os"
	"strings"

	fs "fs/ascii"
)

func main() {
	if len(os.Args) > 4 || len(os.Args) < 3 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	name := ""
	if len(os.Args) == 3 {
		name = "standard"
	} else {
		if os.Args[3] == "thinkertoy" || os.Args[3] == "standard" || os.Args[3] == "shadow" {
			name = os.Args[3]
		} else {
			fmt.Println("incorrect banner")
			return
		}
	}
	file := fs.Read_file(name)
	if file == nil {
		return
	}
	line := os.Args[2]
	if !fs.Is_ascii(line) {
		fmt.Println("Non Ascii character found")
		return
	}
	if len(line) < 1 {
		return
	}
	finalResult := ""
	lines_count := fs.Count_next_line(line)
	splitted_line := strings.Split(line, "\\n")
	splitted_line, lines_count = fs.Cleaned_split(splitted_line, lines_count)
	finalResult = fs.Print_art(file[1:], splitted_line, lines_count)

	if len(os.Args) > 1 && (strings.HasPrefix(os.Args[1], "--output=")) && (strings.HasSuffix(os.Args[1], ".txt")) {
		ExtractName := strings.Split(os.Args[1], "=")
		filename := ExtractName[1]
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
