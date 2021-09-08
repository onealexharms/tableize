package tableize

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatsSingleColumnData(t *testing.T) {
	in := bytes.NewReader([]byte(`
        - foo: hello, there
        `))
	out := bytes.NewBuffer([]byte{})
	err := Tableize(in, out)
	assert.NoError(t, err)
	assert.Equal(t, `foo
------------
hello, there
`, out.String())
}

func TestDeterministicallySortsMultipleColumns(t *testing.T) {
	in := bytes.NewReader([]byte(`
        - b: 1
          e: 2
          a: 3
          d: 4
          c: 5
        `))
	out := bytes.NewBuffer([]byte{})
	err := Tableize(in, out)
	assert.NoError(t, err)
	assert.Equal(t, `a b c d e
- - - - -
3 1 5 4 2
`, out.String())
}
