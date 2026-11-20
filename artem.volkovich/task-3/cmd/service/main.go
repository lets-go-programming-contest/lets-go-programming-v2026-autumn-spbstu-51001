package main

import (
	"flag"

	"github.com/electricbayan/task-3/internal/parser"
)

func main() {
	configPath := flag.String("config", "", "")
	flag.Parse()
	parser.SortValutes(*configPath)
}
