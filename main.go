package main

import (
	"fmt"
	"os"

	"github.com/mouadino/metastore/cli"
)

func main() {
	err := cli.Main(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
