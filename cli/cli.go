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
	fmt.Printf("-port=4000:		Set the PORT of the server\n")
	fmt.Printf("-mode=rest:		Choose between 'html' and 'rest'\n\n")
	// 프로그램을 종료시켜줌
	// 0은 에러가 없다는 뜻이고
	// 다른 숫자는 exit code를 보여줌
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

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
