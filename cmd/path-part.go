package main

import (
	pathpart "github.com/focusaurus/path-part"
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for _, partName := range flag.Args() {
			normalName, err := pathpart.Normalize(partName)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(pathpart.Parse(line, normalName))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
