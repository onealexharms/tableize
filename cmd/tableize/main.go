package main

import (
	"github.com/onealexharms/tableize/pkg/tableize"
	"os"
)

func main() {
	tableize.Tableize(os.Stdin, os.Stdout)
}
