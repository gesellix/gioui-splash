package main

import (
	"fmt"
	"image/color"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	splash "github.com/gesellix/gioui-splash"
	"github.com/gesellix/gioui-splash/assets"
)

func main() {
	go func() {
		logo, err := assets.GetLogo()
		if err != nil {
			fmt.Println(fmt.Errorf("failed to read image data: %w", err))
			os.Exit(0)
		}
		size := logo.Bounds().Size()
		sizeXDp := unit.Dp(size.X)
		sizeYDp := unit.Dp(size.Y)

		options := []app.Option{
			app.Size(sizeXDp, sizeYDp),
			app.MinSize(sizeXDp, sizeYDp),
			app.MaxSize(sizeXDp, sizeYDp),
			app.Title("Splash Example"),
			app.Decorated(false),
		}
		w := app.NewWindow(options...)
		w.Perform(system.ActionCenter)

		splashWidget := splash.NewSplash(
			logo,
			5,  // Progress bar height is 5 Dp.
			10, // Progress bar is inset 10 Dp from window edge.
			color.NRGBA{R: 100, G: 200, B: 100, A: 42},
		)
		progress := 0.0
		go func() {
			tick := time.NewTicker(50 * time.Millisecond)
			defer tick.Stop()

			for {
				select {
				case <-tick.C:
					progress += 0.01
					splashWidget.SetProgress(progress)
					w.Invalidate()
					if progress >= 1 {
						return
					}
				}
			}
		}()

		go func() {
			duration := 5 * time.Second
			fmt.Println("closing in", duration)
			time.Sleep(duration)
			w.Perform(system.ActionClose)
			os.Exit(0)
		}()

		var ops op.Ops
		for e := range w.Events() {
			switch e := e.(type) {

			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				splashWidget.Layout(gtx)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
