package cli

import (
	"log"

	"github.com/codegangsta/cli" // TODO: Get rid of this.

	"github.com/mouadino/metastore/api"
	"github.com/mouadino/metastore/config"
	"github.com/mouadino/metastore/storage"
)

func Main(args []string) error {
	app := cli.NewApp()
	app.Name = "Metastore"
	app.Usage = "Metadata Storage"
	app.Version = Version

	app.Action = run
	return app.Run(args)
}

func run(ctxt *cli.Context) {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.Init(cfg.Options.Storage)
	if err != nil {
		log.Fatal(err)
	}

	api.Init(cfg.Options.API, &db)
}
