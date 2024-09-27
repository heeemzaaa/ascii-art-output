package main

import (
	"fmt"
	"os"

	fs "fs/ascii"
)

func main() {
	if len(os.Args) > 4 || len(os.Args) < 2 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	if len(os.Args) == 4 {
		fs.Output(os.Args[1], os.Args[2], os.Args[3])
	} else if len(os.Args) == 3 {
		fs.FsArgument(os.Args[1], os.Args[2])
	} else if len(os.Args) == 2 {
		fs.SimpleArt(os.Args[1])
	}
}
