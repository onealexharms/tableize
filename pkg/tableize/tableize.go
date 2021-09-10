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

func (t *tableizer) header() string {
	header := ""
	for i, field := range t.fields {
		if i > 0 {
			header += "\t"
		}
		header += fmt.Sprintf("%s", field)
	}
	return strings.TrimRight(header, " ")
}

func (t *tableizer) row(record map[string]string) string {
	row := ""
	for i, field := range t.fields {
		if i > 0 {
			row += "\t"
		}
		value := record[field]
		row += fmt.Sprintf("%s", value)
	}
	return strings.TrimRight(row, " ")
}

func Tableize(in io.Reader, out io.Writer) error {
	t := tableizer{}
	if err := t.parseInput(in); err != nil {
		return err
	}
	t.computeFieldList()
	fmt.Fprintf(out, "%s\n", t.header())
	for _, record := range t.records {
		fmt.Fprintf(out, "%s\n", t.row(record))
	}
	return nil
}
