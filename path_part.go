package pathpart

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
Given a typical path:

some/directories/info.dat.zip

here's our terminology:

"info" is the "base"

"zip" is the "extension"

"dat.zip" is the "extensions"

"info.dat.zip" is the "name"

"some/directories" is the "path"
*/
const (
	base       = "base"
	extension  = "extension"
	extensions = "extensions"
	name       = "name"
	path       = "path"
)

func Parse(line string, partName string) string {
	switch partName {
	case base:
		return strings.SplitN(filepath.Base(line), ".", 2)[0]
	case extension:
		return strings.TrimLeft(filepath.Ext(line), ".")
	case extensions:
		return strings.SplitN(filepath.Base(line), ".", 2)[1]
	case name:
		return filepath.Base(line)
	case path:
		return filepath.Dir(line)
	}
	return line
}

func Normalize(partName string) (string, error) {
	switch partName {
	case "base":
		return base, nil
	case "exts":
		fallthrough
	case "extensions":
		return extensions, nil
	case "ext":
		fallthrough
	case "extension":
		return extension, nil
	case "name":
		fallthrough
	case "last":
		return name, nil
	case "directory":
		fallthrough
	case "dirname":
		fallthrough
	case "dir":
		fallthrough
	case "path":
		return path, nil
	}
	return partName, errors.New("unrecognized part name: " + partName + ". valid names are extension, extensions, base, and path.")
}

func main() {
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for _, partName := range flag.Args() {
			normalName, err := Normalize(partName)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(Parse(line, normalName))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
