package eredis

import (
	"github.com/gomodule/redigo/redis"
	"time"
	"github.com/whf-sky/easygo"
)

var configs = map[string]map[string]DialConfig{}//环境 组 配置

var Config *config

type config struct {
	env string
}

type DialConfig struct {
	MaxIdle int
	IdleTimeout time.Duration
	Address string
	Network string
	Options []redis.DialOption
}

func (c *config) Env(name string) *config {
	c.env = name
	return c
}

func (c *config) Group(name string, cnf DialConfig)  {
	if c.env == "" {
		c.env = easygo.EASYGO_ENV
	}
	if _, ok := configs[c.env]; !ok {
		configs[c.env] = map[string]DialConfig{}
	}
	if _, ok := configs[c.env][name]; !ok {
		configs[c.env][name] = cnf
	}
}

func init()  {
	Config = &config{}
}
