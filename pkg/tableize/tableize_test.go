package tableize

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkCase(t *testing.T, in string, out string) {
	reader := bytes.NewReader([]byte(in[1:]))
	writer := bytes.NewBuffer([]byte{})
	err := Tableize(reader, writer)
	assert.NoError(t, err)
	assert.Equal(t, out[1:], writer.String())
}

func TestFormatsSingleColumnData(t *testing.T) {
	checkCase(t, `
- foo: hello, there
`, `
foo
------------
hello, there
`)
}

func TestDeterministicallySortsMultipleColumns(t *testing.T) {
	checkCase(t, `
- b: 1
  e: 2
  a: 3
  d: 4
  c: 5
`, `
a b c d e
- - - - -
3 1 5 4 2
`)
}
