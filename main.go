package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Given a typical path, here's our terminology
// some/directories/info.dat.zip
// some/directories is the "path"
// "info.dat.zip" is the "name"
// "zip" is the "extension"
// "dat.zip" is the "extensions"
// "info" is the "base"

func main() {
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for _, arg := range flag.Args() {
			switch arg {
			case "exts":
				fallthrough
			case "extensions":
				line = filepath.Base(line)
				line = strings.SplitN(line, ".", 2)[1]
			case "ext":
				fallthrough
			case "extension":
				line = strings.TrimLeft(filepath.Ext(line), ".")
			case "name":
				fallthrough
			case "last":
				line = filepath.Base(line)
			case "directory":
				fallthrough
			case "dirname":
				fallthrough
			case "path":
				line = filepath.Dir(line)
			case "base":
				line = filepath.Base(line)
				line = strings.SplitN(line, ".", 2)[0]
			}
		}
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
