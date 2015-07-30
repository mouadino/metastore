package cli

import (
	"log"

	"github.com/codegangsta/cli"

	"github.com/mouadino/metastore/api"
	"github.com/mouadino/metastore/context"
	"github.com/mouadino/metastore/storage"
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
		cli.StringFlag{
			Name:   "driver, d",
			Usage:  "Storage driver to use e.g. boltdb",
			Value:  "boltdb",
			EnvVar: "DRIVER",
		},
		cli.StringFlag{
			Name:   "dbname, db",
			Usage:  "Storage database name",
			Value:  "metastore.db",
			EnvVar: "DBNAME",
		},
	}

	app.Action = run
	return app.Run(args)
}

func run(ctxt *cli.Context) {
	driverName := ctxt.String("driver")
	DBName := ctxt.String("dbname")

	db, err := storage.Init(driverName, DBName)
	if err != nil {
		// XXX Could we do better here ?
		log.Fatal(err)
	}

	host := ctxt.String("host")
	port := ctxt.Int("port")
	context := context.Create(host, port, db)

	api.Main(context)
}
