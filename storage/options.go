package storage

import (
	"fmt"
	"os"
	"path"

	"github.com/mouadino/metastore/storage/boltdb"
)

// TODO: Specific driver configuration.
type Options struct {
	Driver string `name:"driver" default:"boltdb"`
	Path   string `name:"path" default:"/tmp/"`
	Name   string `name:"name" default:"metastore.db"` // TODO: Rename me.
}

func (opts *Options) Validate() error {
	pathInfo, err := os.Stat(opts.Path)
	if err != nil {
		return err
	}
	if !pathInfo.IsDir() {
		return fmt.Errorf("%s is not a directory", opts.Path)
	}
	return nil
}

func (opts *Options) DBPath() string {
	return path.Join(opts.Path, opts.Name)
}

func (opts *Options) DB() (DB, error) {
	switch {
	case opts.Driver == "boltdb":
		return &boltdb.DB{}, nil
	default:
		return nil, fmt.Errorf("unknown storage driver %s", opts.Driver)
	}
}
