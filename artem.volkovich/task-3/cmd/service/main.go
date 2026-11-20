package main

import (
	"flag"

	"github.com/electricbayan/task-3/internal/parser"
)

func main() {
	config_path := flag.String("config", "", "")
	flag.Parse()
	parser.SortValutes(*config_path)
}
