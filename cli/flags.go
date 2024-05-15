package cli

import (
	"flag"
)

func ParseFlags() (bool) {
	concurrencyFlag := flag.Bool("concurrency", false, "Test concurrency")
	flag.Parse()

	return *concurrencyFlag
}
