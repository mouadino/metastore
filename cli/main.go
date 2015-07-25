package cli

import (
	"github.com/codegangsta/cli"
	"github.com/mouadino/metastore/server"
)

func Main(args []string) error {
	app := cli.NewApp()
	app.Name = "Metastore"
	app.Usage = "Metadata Storage"
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

		server.Main(host, port)
	}
	return app.Run(args)
}
