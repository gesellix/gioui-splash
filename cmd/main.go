package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	splash "github.com/gesellix/gioui-splash"
	"github.com/gesellix/gioui-splash/assets"
	"image"
	"image/color"
	"os"
	"time"
)

func main() {
	go func() {
		logo, err := assets.GetLogo()
		if err != nil {
			fmt.Println(fmt.Errorf("failed to read image data: %w", err))
			os.Exit(0)
		}
		size := logo.Bounds().Size()

		options := []app.Option{
			app.Size(unit.Dp(size.X), unit.Dp(size.Y)),
			app.Title("Splash Example"),
			app.Decorated(false),
		}
		w := app.NewWindow(options...)
		//w.Option(options...)
		w.Perform(system.ActionCenter)

		splashWidget := splash.NewSplash(
			logo,
			image.Rectangle{
				Min: image.Pt(10, size.Y-10),
				Max: image.Pt(size.X-10, size.Y-5),
			},
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

				//gtx.Constraints = layout.Exact(size)
				//w.Option(app.Size(unit.Dp(size.X), unit.Dp(size.Y)))

				splashDim := splashWidget.Layout(gtx)

				fmt.Printf("expected size: %v | splash dim: %v | actual gtx.Metric: %v gtx.Constraints %v e.size %v\n", size, splashDim, gtx.Metric, gtx.Constraints, e.Size)

				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
