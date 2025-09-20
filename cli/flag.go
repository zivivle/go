package cli

import (
	"flag"

	"github.com/zivivle/go/explorer"
	"github.com/zivivle/go/rest"
)

func Flags() {
	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
