package tableize

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"strings"
)

func fieldWidths(data []map[string]string) (fieldWidths map[string]int) {
	fieldWidths = make(map[string]int)
	for _, record := range data {
		for field, value := range record {
			if len(field) > fieldWidths[field] {
				fieldWidths[field] = len(field)
			}
			if len(value) > fieldWidths[field] {
				fieldWidths[field] = len(value)
			}
		}
	}
	return
}

func Tableize(in io.Reader, out io.Writer) error {
	rawData, err := io.ReadAll(in)
	if err != nil {
		return err
	}

	var data []map[string]string
	err = yaml.Unmarshal(rawData, &data)
	if err != nil {
		return err
	}

	widths := fieldWidths(data)

	header := ""
	started := false
	for field, width := range widths {
		if started {
			header += " "
		}
		started = true
		format := fmt.Sprintf("%%-%ds", width)
		header += fmt.Sprintf(format, field)
	}
	fmt.Fprintf(out, "%s\n", strings.TrimRight(header, " "))

	started = false
	for _, width := range widths {
		if started {
			fmt.Fprintf(out, " ")
		}
		started = true
		fmt.Fprintf(out, "%s", strings.Repeat("-", width))
	}
	fmt.Fprintf(out, "\n")

	for _, record := range data {
		started = false
		for field, width := range widths {
			if started {
				fmt.Fprintf(out, " ")
			}
			started = true
			value := record[field]
			format := fmt.Sprintf("%%-%ds", width)
			fmt.Fprintf(out, format, value)
		}
		fmt.Fprintf(out, "\n")
	}

	return nil
}
