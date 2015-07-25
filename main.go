package main

import (
	"github.com/mouadino/metastore/cli"
	"log"
	"os"
)

func main() {
	err := cli.Main(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
