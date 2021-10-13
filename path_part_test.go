package pathpart

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"strings"
)

func TestParsing(t *testing.T) {
	type test struct {
		input    string
		partName string
		expected string
		message  string
	}

	tests := []test{
		{"foo", "base", "foo", "base: no path separators"},
		{"/usr/bin/foo.bar", "base", "foo", "base: path with separators"},
		{"/usr/bin/foo.bar.baz", "base", "foo", "base: multiple extensions"},
		{"", "base", "", "base: empty"},
		{"/usr/bin/foo.bar.baz", "extensions", "bar.baz", "extensions: 2 extensions"},
		{"/usr/bin/foo.bar.baz.buux", "extensions", "bar.baz.buux", "extensions: 3 extensions"},
		{"", "extensions", "", "extensions: empty"},
		{"/usr/bin/foo.bar.baz", "exts", "bar.baz", "extensions: alias exts"},
		{"design.pdf", "extension", "pdf", "extension: single extension"},
		{"~/dir/design.pdf", "extension", "pdf", "extension: single extension and path"},
		{"design.pdf.final.pdf", "extension", "pdf", "extension: multiple extensions"},
		{"", "extension", "", "extension: empty"},
		{"design.pdf.final.pdf", "ext", "pdf", "extension: alias ext"},
		{"$HOME/projects/hi.md", "name", "hi.md", "name: dirs and file"},
		{"$HOME/projects/hi.md", "basename", "hi.md", "name: alias basename"},
		{"$HOME/projects/hi.md.zip", "last", "hi.md.zip", "name: alias last"},
		{"/usr/bin/bash", "path", "/usr/bin", "path: dirs and file"},
		{"/usr/bin", "path", "/usr", "path: dirs"},
		{"Some Directory - BLAH/FILE NAME DOT TXT", "path", "Some Directory - BLAH", ""},
		{"", "path", ".", "path: empty"},
		{"/", "path", "/", "path: slash"},
		{"////", "path", "/", "path: multiple adjacent slashes"},
	}

	for _, test := range tests {
		output, err := Parse(test.input, test.partName)
		require.NoError(t, err)
		assert.Equal(t, test.expected, output, test.message)
	}
}

func TestErrorTruncation(t *testing.T) {
	_, err := Parse("some input", "aaaaaaaaaabbbbbbbbbbcccccccccc")
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "aabb"), true)
	assert.True(t, strings.Contains(err.Error(), "c"), false, "Lengthy part name input should be truncated in error message.")
}
