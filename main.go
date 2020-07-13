package main

import (
	"bufio"
	"path/filepath"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		base := filepath.Base(scanner.Text())
		fmt.Println(base)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
