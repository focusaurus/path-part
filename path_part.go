package pathpart

import (
	"errors"
	"fmt"
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

type parser func(string) string

func parseBase(line string) string {
	return strings.SplitN(filepath.Base(line), ".", 2)[0]
}

func parseExtension(line string) string {
	return strings.TrimLeft(filepath.Ext(line), ".")
}

func parseExtensions(line string) string {
	return strings.SplitN(filepath.Base(line), ".", 2)[1]
}

func parseName(line string) string {
	return filepath.Base(line)
}

func parsePath(line string) string {
	return filepath.Dir(line)
}

var dispatch = map[string]parser{
	"base":       parseBase,
	"extensions": parseExtensions,
	"exts":       parseExtensions,
	"extension":  parseExtension,
	"ext":        parseExtension,
	"name":       parseName,
	"last":       parseName,
	"basename":   parseName,
	"path":       parsePath,
	"directory":  parsePath,
	"dir":        parsePath,
	"dirname":    parsePath,
}

func Parse(line string, partName string) (string, error) {
	parse, ok := dispatch[strings.ToLower(partName)]
	if !ok {
		return line, errors.New("unrecognized part name: " + fmt.Sprintf("%.25s", partName) + `.
Valid names are:
base
extension (alias ext)
extensions (alias exts)
name (alias last, basename)
path (alias directory, dir, dirname)`)
	}
	return parse(line), nil
}
