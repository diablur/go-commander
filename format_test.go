package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func Test_Format_Description(t *testing.T) {
	assert.Equal(t,
		Format.Description("a", "desc..."),
		"a             desc...")
	assert.Equal(t,
		Format.Description("abcdefghijklmn", "desc..."),
		"abcdefghijklmn\n              desc...")
}

func TestFormat_DescriptionLine(t *testing.T) {
	assert.Equal(t,
		formatDescription("a", "desc...", 2, 5, false),
		"a    desc...")
	assert.Equal(t,
		formatDescription("abcdef", "desc...", 2, 5, false),
		"abcdef  desc...")

	assert.Equal(t,
		formatDescription("a", "desc...", 2, 5, true),
		"a    desc...")
	assert.Equal(t,
		formatDescription("abcde", "desc...", 2, 5, true),
		"abcde\n     desc...")
	assert.Equal(t,
		formatDescription("abcdef", "desc...", 2, 5, true),
		"abcdef\n     desc...")

	assert.Equal(t,
		formatDescription("a", "desc...", -1, 5, false),
		"a    desc...")
	assert.Equal(t,
		formatDescription("a", "desc...", 0, 5, false),
		"a    desc...")
	assert.Equal(t,
		formatDescription("a", "desc...", 5, 4, false),
		"a     desc...")
	assert.Equal(t,
		formatDescription("a", "desc...", 5, 5, false),
		"a     desc...")
	assert.Equal(t,
		formatDescription("a", "desc...", 5, 6, false),
		"a     desc...")
}
