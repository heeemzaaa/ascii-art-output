package fs

import (
	"fmt"
	"os"
	"strings"
)

// a function that read ascii file

func Read_file(s string) []string {
	file, err := os.ReadFile("art/" + s + ".txt")
	if err != nil {
		fmt.Println("Ascii file not found")
		return nil
	}
	ret := strings.Split(string(file), "\n")
	for i := 0; i < len(ret) && s == "thinkertoy"; i++ {
		ret[i] = strings.ReplaceAll(ret[i], "\r", "")
	}
	return ret
}

// counts how many new lines to print

func Count_next_line(line string) []int {
	var ret []int
	result := []int{}
	j := 0
	ret = append(ret, 0)
	for i := 0; i < len(line); i++ {
		if i+1 < len(line) && line[i] == '\\' {
			if line[i+1] == 'n' {
				ret[j]++
			} else {
				ret = append(ret, 0)
			}
			i++
			if i+1 < len(line) && line[i+1] != '\\' {
				ret = append(ret, 0)
				j++
			}
		}
	}
	if line[0] != '\\' && (1 < len(line) && line[1] != 'n') {
		result = append(result, 0)
	}
	for i := 0; i < len(ret); i++ {
		result = append(result, ret[i])
	}
	return result
}

// a function that prints ascii charachters

func Print_art(file []string, splitted_line []string, lines_count []int) string {
	holder := 0
	result := ""
	i := 0
	for ; i < len(splitted_line); i++ {
		for ; (i == 0 && len(lines_count) > 1) && lines_count[i] > 0; lines_count[i]-- {
			result += "\n"
		}
		for ; (i < len(lines_count) && len(lines_count) > 0) && lines_count[i] > 1; lines_count[i]-- {
			result += "\n"
		}
		for j := 0; j < 8; j++ {
			for k := 0; k < len(splitted_line[i]); k++ {
				holder = (int(splitted_line[i][k])-32)*9 + j
				result += file[holder]
			}
			result += "\n"
		}
	}
	i = len(lines_count) - 1
	if i >= 0 {
		for ; lines_count[i] > 0 && len(lines_count) > 0; lines_count[i]-- {
			result += "\n"
		}
	}
	return result
}

func Check_if_empty(s []string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			return true
		}
	}
	return false
}

// clean splitted line from empty strings
func Cleaned_split(s []string, lines_count []int) ([]string, []int) {
	var ret []string
	i := 0
	j := 0
	if s[0] == "" {
		if !Check_if_empty(s) {
			i++
		}
		for ; (i < len(s) && s[i] == "") && j > 0; i++ {
			j = lines_count[i]
			ret = append(ret, "")
			j--
		}
	}
	for ; i < len(s); i++ {
		if s[i] != "" {
			ret = append(ret, s[i])
		}
	}
	// fmt.Printf("ret : %q\n", ret)
	// fmt.Println(lines_count)
	return ret, lines_count
}

func Is_ascii(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 32 || s[i] > 126 {
			return false
		}
	}
	return true
}
