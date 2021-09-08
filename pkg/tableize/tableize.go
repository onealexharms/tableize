package tableize

import (
	"io"
	"gopkg.in/yaml.v2"
)

func Tableize(in io.Reader, out io.Writer) error {
    bytes, err := io.ReadAll(in)
    if err != nil {
        return err
    }

    var data []map[string]interface{}
    err = yaml.Unmarshal(bytes, &data)
    if err != nil {
        return err
    }
    return nil
}
