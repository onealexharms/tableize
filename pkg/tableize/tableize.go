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

func header(widths map[string]int) string {
	header := ""
	for field, width := range widths {
		format := fmt.Sprintf("%%-%ds ", width)
		header += fmt.Sprintf(format, field)
	}
	return strings.TrimRight(header, " ")
}

func separator(widths map[string]int) string {
	separator := ""
	for _, width := range widths {
		separator += strings.Repeat("-", width) + " "
	}
	return strings.TrimRight(separator, " ")
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
	fmt.Fprintf(out, "%s\n", header(widths))
	fmt.Fprintf(out, "%s\n", separator(widths))

	for _, record := range data {
		started := false
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
