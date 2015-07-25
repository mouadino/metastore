package cli

import (
	"github.com/codegangsta/cli"
	"github.com/mouadino/metastore/server"
	"log"
)

const Version = "0.1.0"

func Main(args []string) error {
	app := cli.NewApp()
	app.Name = "metastore"
	app.Usage = "metadata storage"
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "port, p",
			Usage:  "port to listen to",
			Value:  4080,
			EnvVar: "PORT",
		},
		cli.StringFlag{
			Name:   "host, a",
			Usage:  "host to listen to",
			EnvVar: "HOST",
		},
	}
	app.Action = func(c *cli.Context) {
		host := c.String("host")
		port := c.Int("port")
		log.Fatal(server.ListenAndServe(host, port))
	}

	return app.Run(args)
}
