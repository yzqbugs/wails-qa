package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Config struct {
	ctx context.Context
}

func NewConf() *Config {
	return &Config{}
}

func (c *Config) SetOutDir() string {

	//dialogOpt := runtime.OpenDialogOptions{}
	fmt.Println(c.ctx)
	runtime.BrowserOpenURL(c.ctx, "e:/")
	//if dir != "" {
	//
	//	runtime.LogInfof(c.ctx, "set output directory: %s", dir)
	//
	//}
	return "success"
}

func (c *Config) SetConf() string {
	c.ctx = context.Background()
	runtime.LogInfo(c.ctx, "log info")
	return "setconf"
}
func (c *Config) WailsInit() error {

	return nil
}

func (c Config) Name() error {
	return nil
}
