package main

import (
	"embed"
	"runtime"
	"time"

	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

func init() {
	// application.RegisterEvent[string]("time")
}

type RAMUsageStats struct {
	Alloc      uint64
	TotalAlloc uint64
	Sys        uint64
	NumGC      uint32
}

func GetRAMUsage() RAMUsageStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return RAMUsageStats{
		Alloc:      m.Alloc,
		TotalAlloc: m.TotalAlloc,
		Sys:        m.Sys,
		NumGC:      m.NumGC,
	}
}

func main() {
	app := application.New(application.Options{
		Name:        "harbour",
		Description: "A demo of using raw HTML & CSS",
		Services: []application.Service{
			application.NewService(new(NewHarbourService())),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:  "Window 1",
		Width:  500,
		Height: 300,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropLiquidGlass,
			TitleBar:                application.MacTitleBarHiddenInset,
			Appearance:              application.DefaultAppearance,
		},
		BackgroundColour: application.NewRGB(6, 7, 15),
		BackgroundType:   application.BackgroundTypeTransparent,
		URL:              "/",
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:  "Window 1",
		Width:  500,
		Height: 300,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropLiquidGlass,
			TitleBar:                application.MacTitleBarHiddenInset,
			Appearance:              application.DefaultAppearance,
		},
		BackgroundColour: application.NewRGB(6, 7, 15),
		BackgroundType:   application.BackgroundTypeTransparent,
		URL:              "/",
	})

	ticker := time.NewTicker(time.Second * 5)
	go func(t *time.Ticker) {
		for {
			<-ticker.C
			log.Printf("ram usage %+v", GetRAMUsage())
		}

	}(ticker)

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
