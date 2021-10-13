package main

import (
	"bufio"
	"flag"
	"fmt"
	pathpart "github.com/focusaurus/path-part"
	"os"
)

func main() {
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for _, partName := range flag.Args() {
			part, err := pathpart.Parse(line, partName)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			fmt.Println(part)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
