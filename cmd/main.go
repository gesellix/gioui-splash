package main

import (
	"fmt"
	"image/color"
	"os"
	"sync"
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
		w := app.Window{}
		w.Option(options...)
		w.Perform(system.ActionCenter)

		splashWidget := splash.NewSplash(
			logo,
			// (Bottom-Top) sets the height of the progress bar
			layout.Inset{
				Top:    5,
				Bottom: 10,
				Left:   10,
				Right:  10,
			},
			color.NRGBA{R: 100, G: 200, B: 100, A: 127},
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
					// The widget will not be updated until the next FrameEvent.
					// We're going to trigger that event now, so that
					// the changed progress will be visible.
					w.Invalidate()
					if progress >= 1 {
						return
					}
				}
			}
		}()

		// This is a placeholder for any background process.
		go func() {
			duration := 5 * time.Second
			fmt.Println("background processing will be finished in", duration)
			time.Sleep(duration)
			w.Perform(system.ActionClose)
			os.Exit(0)
		}()

		// TODO work around https://todo.sr.ht/~eliasnaur/gio/602
		// this should only be required shortly after creating the window w.
		performCenter := sync.OnceFunc(func() {
			w.Perform(system.ActionCenter)
		})
		var ops op.Ops
		for {
			switch e := w.Event().(type) {
			case app.FrameEvent:
				// TODO work around https://todo.sr.ht/~eliasnaur/gio/602
				// this should only be required shortly after creating the window w.
				performCenter()
				gtx := app.NewContext(&ops, e)
				splashWidget.Layout(gtx)
				e.Frame(gtx.Ops)
			case app.DestroyEvent:
				// Omitted
			}
		}
	}()
	app.Main()
}
