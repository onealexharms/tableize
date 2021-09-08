package tableize

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"sort"
	"strings"
)

type tableizer struct {
	records []map[string]string
	fields  []string
	widths  map[string]int
}

func (t *tableizer) parseInput(in io.Reader) error {
	rawData, err := io.ReadAll(in)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(rawData, &t.records)
}

func (t *tableizer) computeFieldList() {
	for _, record := range t.records {
		for field := range record {
			t.fields = append(t.fields, field)
		}
	}
	sort.Sort(sort.StringSlice(t.fields))
}

func (t *tableizer) computeColumnWidths() {
	t.widths = make(map[string]int)
	for _, record := range t.records {
		for field, value := range record {
			if len(field) > t.widths[field] {
				t.widths[field] = len(field)
			}
			if len(value) > t.widths[field] {
				t.widths[field] = len(value)
			}
		}
	}
	return
}

func (t *tableizer) header() string {
	header := ""
	for _, field := range t.fields {
		width := t.widths[field]
		format := fmt.Sprintf("%%-%ds ", width)
		header += fmt.Sprintf(format, field)
	}
	return strings.TrimRight(header, " ")
}

func (t *tableizer) separator() string {
	separator := ""
	for _, field := range t.fields {
		width := t.widths[field]
		separator += strings.Repeat("-", width) + " "
	}
	return strings.TrimRight(separator, " ")
}

func (t *tableizer) row(record map[string]string) string {
	row := ""
	for _, field := range t.fields {
		width := t.widths[field]
		value := record[field]
		format := fmt.Sprintf("%%-%ds ", width)
		row += fmt.Sprintf(format, value)
	}
	return strings.TrimRight(row, " ")
}

func Tableize(in io.Reader, out io.Writer) error {
	t := tableizer{}
	if err := t.parseInput(in); err != nil {
		return err
	}
	t.computeFieldList()
	t.computeColumnWidths()
	fmt.Fprintf(out, "%s\n", t.header())
	fmt.Fprintf(out, "%s\n", t.separator())
	for _, record := range t.records {
		fmt.Fprintf(out, "%s\n", t.row(record))
	}
	return nil
}
