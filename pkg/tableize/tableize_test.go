package tableize

import (
        "bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItRuns(t *testing.T) {
        in := bytes.NewReader([]byte(`
        - foo: true
          bar: 42
        `))
        out := bytes.NewBuffer([]byte{})
	err := Tableize(in, out)
	assert.NoError(t, err)
}
