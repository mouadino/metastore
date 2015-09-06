package config

import (
	"strings"

	"github.com/mouadino/cfg"
	"github.com/spf13/viper"

	"github.com/mouadino/metastore/api"
	"github.com/mouadino/metastore/storage"
)

type GlobalOptions struct {
	API     *api.Options     `name:"api"`
	Storage *storage.Options `name:"storage"`
}

func (opts *GlobalOptions) Validate() error {
	err := opts.API.Validate()
	if err != nil {
		return err
	}
	err = opts.Storage.Validate()
	if err != nil {
		return err
	}
	return nil
}

type Config struct {
	Options  *GlobalOptions
	FilePath string
}

func Load() (*Config, error) {
	viper := viper.New()

	viper.SetConfigName("metastore")
	viper.AddConfigPath("/etc/metastore/")
	viper.AddConfigPath("$HOME/.metastore/")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	opts := &GlobalOptions{
		API:     &api.Options{},
		Storage: &storage.Options{},
	}

	getter := func(keys ...string) interface{} {
		return viper.Get(strings.Join(keys, "."))
	}
	err = cfg.Parse(getter, opts)
	if err != nil {
		return nil, err
	}

	err = opts.Validate()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Options:  opts,
		FilePath: "", // FIXME
	}
	return cfg, nil
}
