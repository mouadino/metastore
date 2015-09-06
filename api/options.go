package api

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	InvalidPort = errors.New("api: invalid port number 0")
	IPRegex     = regexp.MustCompile("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$")
)

type Options struct {
	IP   string `name:"ip" default:"127.0.0.1"`
	Port uint16 `name:"port" default:"4080"`
}

func (opts *Options) Address() string {
	return fmt.Sprintf("%s:%d", opts.IP, opts.Port)
}

func (opts *Options) Validate() error {
	if opts.Port == 0 {
		return InvalidPort
	}
	if match := IPRegex.MatchString(opts.IP); !match {
		return fmt.Errorf("api: invalid ip %s", opts.IP)
	}
	return nil
}
