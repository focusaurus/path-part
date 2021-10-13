package pathpart

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNormalizeAliases(t *testing.T) {
	type test struct {
		input    string
		expected string
	}
	tests := []test{
		{"base", base},
		{"exts", extensions},
		{"extensions", extensions},
		{"ext", extension},
		{"extension", extension},
		{"path", path},
		{"dir", path},
		{"directory", path},
		{"dirname", path},
	}
	for _, test := range tests {
		result, err := normalize(test.input)
		require.NoError(t, err)
		assert.Equal(t, test.expected, result)
	}
}

func TestParsing(t *testing.T) {
	type test struct {
		input    string
		partName string
		expected string
		message string
	}

	tests := []test{
		{"foo", base, "foo", "base: no path separators"},
		{"/usr/bin/foo.bar", base, "foo", "base: path with separators"},
		{"/usr/bin/foo.bar.baz", base, "foo", "base: multiple extensions"},
		{"", base, "", "base: empty"},
		{"/usr/bin/foo.bar.baz", extensions, "bar.baz", "extensions: 2 extensions"},
		{"/usr/bin/foo.bar.baz.buux", extensions, "bar.baz.buux", "extensions: 3 extensions"},
		{"", extensions, "", "extensions: empty"},
		{"design.pdf", extension, "pdf", "extension: single extension"},
		{"~/dir/design.pdf", extension, "pdf", "extension: single extension and path"},
		{"design.pdf.final.pdf", extension, "pdf", "extension: multiple extensions"},
		{"", extension, "", "extension: empty"},
		{"/usr/bin/bash", path, "/usr/bin", "path: dirs and file"},
		{"/usr/bin", path, "/usr", "path: dirs"},
		{"Some Directory - BLAH/FILE NAME DOT TXT", path, "Some Directory - BLAH", ""},
		{"", path, ".", "path: empty"},
		{"/", path, "/", "path: slash"},
		{"////", path, "/", "path: multiple adjacent slashes"},
	}

	for _, test := range tests {
		output := parse(test.input, test.partName)
		assert.Equal(t, test.expected, output, test.message)
	}
}
