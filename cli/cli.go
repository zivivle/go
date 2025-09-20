package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/zivivle/go/explorer"
	"github.com/zivivle/go/rest"
)

func usage() {
	fmt.Printf("Welcome to Z Coin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("rest:4000:		Set the PORT of the REST server\n")
	fmt.Printf("html:4000:		Set the PORT of the HTML server\n")
	// 프로그램을 종료시켜줌
	// 0은 에러가 없다는 뜻이고
	// 다른 숫자는 exit code를 보여줌
	os.Exit(0)
}

func run(mode string, port int) {
	switch mode {
	case "rest":
		rest.Start(port)
	case "html":
		explorer.Start(port)
	default:
		usage()
	}

}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	cmd := os.Args[1]

	switch cmd {
	case "rest":
		restCmd := flag.NewFlagSet("rest", flag.ExitOnError)
		port := restCmd.Int("port", 4000, "Set port of the server")
		restCmd.Parse(os.Args[2:])
		run("rest", *port)
	case "html":
		htmlCmd := flag.NewFlagSet("html", flag.ExitOnError)
		port := htmlCmd.Int("port", 4000, "Set port of the server")
		htmlCmd.Parse(os.Args[2:])
		run("html", *port)
	default:
		usage()
	}
}
