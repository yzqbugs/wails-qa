package main

import (
	"context"
	"embed"
	"genshin/backend"
	"github.com/gookit/color"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	conf := backend.NewConf()
	// Create application with options
	err := wails.Run(&options.App{
		Title:            "genshin-wails",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			conf.StartUp(ctx)
		},
		Bind: []interface{}{
			app, conf,
		},
	})
	color.Redln("hello")
	if err != nil {
		println("Error:", err)
	}
}
