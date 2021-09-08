package tableize

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"strings"
)

type tableizer struct {
}

func (t *tableizer) parseInput(in io.Reader) ([]map[string]string, error) {
	rawData, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}

	var data []map[string]string
	err = yaml.Unmarshal(rawData, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (t *tableizer) fieldWidths(data []map[string]string) (fieldWidths map[string]int) {
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

func (t *tableizer) header(widths map[string]int) string {
	header := ""
	for field, width := range widths {
		format := fmt.Sprintf("%%-%ds ", width)
		header += fmt.Sprintf(format, field)
	}
	return strings.TrimRight(header, " ")
}

func (t *tableizer) separator(widths map[string]int) string {
	separator := ""
	for _, width := range widths {
		separator += strings.Repeat("-", width) + " "
	}
	return strings.TrimRight(separator, " ")
}

func (t *tableizer) row(record map[string]string, widths map[string]int) string {
	row := ""
	for field, width := range widths {
		value := record[field]
		format := fmt.Sprintf("%%-%ds", width)
		row += fmt.Sprintf(format, value)
	}
	return strings.TrimRight(row, " ")
}

func Tableize(in io.Reader, out io.Writer) error {
	t := tableizer{}
	data, err := t.parseInput(in)
	if err != nil {
		return err
	}

	widths := t.fieldWidths(data)
	fmt.Fprintf(out, "%s\n", t.header(widths))
	fmt.Fprintf(out, "%s\n", t.separator(widths))
	for _, record := range data {
		fmt.Fprintf(out, "%s\n", t.row(record, widths))
	}

	return nil
}
