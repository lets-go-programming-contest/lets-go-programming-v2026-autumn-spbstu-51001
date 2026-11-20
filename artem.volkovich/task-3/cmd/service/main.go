package main

import (
	"flag"

	"github.com/electricbayan/task-3/internal/parser"
)

func main() {
	config_path := flag.String("config", "", "")
	var input_path, output_path = parser.GetFilenames(*config_path)
	parser.Parse(input_path, output_path)
}