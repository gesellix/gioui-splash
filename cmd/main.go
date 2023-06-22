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
		progress := 0.0
		logo, err := assets.GetLogo()
		if err != nil {
			fmt.Println(fmt.Errorf("failed to read image data: %w", err))
			os.Exit(0)
		}
		splashWidget := splash.NewSplash(
			logo,
			image.Rectangle{
				Min: image.Pt(100, 340),
				Max: image.Pt(540, 350),
			},
			color.NRGBA{R: 100, G: 200, B: 100, A: 42},
		)

		size := image.Point{X: 640, Y: 360}
		options := []app.Option{
			app.Size(unit.Dp(size.X), unit.Dp(size.Y)),
			app.Title("Splash Example"),
			app.Decorated(false),
		}
		w := app.NewWindow(options...)
		//w.Option(options...)
		w.Perform(system.ActionCenter)

		go func() {
			tick := time.NewTicker(50 * time.Millisecond)
			defer tick.Stop()

			for {
				select {
				case <-tick.C:
					progress += 0.01
					splashWidget.SetProgress(progress)
					w.Invalidate()
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
