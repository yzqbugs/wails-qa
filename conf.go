package main

import (
	"context"
)

type Config struct {
	ctx context.Context
}

func (c *Config) StartUp(ctx context.Context) {
	c.ctx = ctx
}

func (c *Config) SayHello() string {
	return "hello"
}
func NewConf() *Config {
	return &Config{}
}


