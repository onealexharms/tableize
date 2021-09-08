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
