package splash

import (
	"gioui.org/gpu/headless"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/gesellix/gioui-splash/assets"
	"github.com/gesellix/gioui-splash/f32color"
	"image"
	"image/color"
	"testing"
)

func TestRender(t *testing.T) {
	logo, _ := assets.GetLogo()
	size := image.Point{X: 640, Y: 360}
	splash := NewSplash(
		logo,
		layout.Inset{
			Top:    5,
			Bottom: 10,
			Left:   10,
			Right:  10,
		},
		color.NRGBA{R: 100, G: 200, B: 100, A: 127},
	)
	splash.SetProgress(0.5)
	run(t,
		size,
		func(ops *op.Ops) {
			splash.Layout(layout.NewContext(ops, system.FrameEvent{
				Size:   size,
				Metric: unit.Metric{PxPerDp: 1},
			}))
		}, func(r result) {
		},
	)
}

func TestHeadless(t *testing.T) {
	size := image.Point{X: 640, Y: 360}
	w, release := newTestWindow(t, size)
	defer release()

	var ops op.Ops

	colOrigin := color.NRGBA{R: 15, G: 26, B: 32, A: 255}
	colProgress := color.NRGBA{R: 73, G: 148, B: 76, A: 255}

	logo, _ := assets.GetLogo()
	splash := NewSplash(
		logo,
		layout.Inset{
			Top:    5,
			Bottom: 10,
			Left:   10,
			Right:  10,
		},
		color.NRGBA{R: 100, G: 200, B: 100, A: 127},
	)
	splash.SetProgress(0.5)
	splash.Layout(layout.NewContext(&ops, system.FrameEvent{
		Size:   size,
		Metric: unit.Metric{PxPerDp: 1},
	}))

	if err := w.Frame(&ops); err != nil {
		t.Fatal(err)
	}

	img := image.NewRGBA(image.Rectangle{Max: w.Size()})
	err := w.Screenshot(img)
	if err != nil {
		t.Fatal(err)
	}
	if isz := img.Bounds().Size(); isz != size {
		t.Errorf("got %v screenshot, expected %v", isz, size)
	}
	if got := img.RGBAAt(0, 0); got != f32color.NRGBAToRGBA(colOrigin) {
		t.Errorf("got color %v, expected %v", got, f32color.NRGBAToRGBA(colOrigin))
	}
	if got := img.RGBAAt(10, 345); got != f32color.NRGBAToRGBA(colProgress) {
		t.Errorf("got color %v, expected %v", got, f32color.NRGBAToRGBA(colProgress))
	}
}

func newTestWindow(t *testing.T, sz image.Point) (*headless.Window, func()) {
	t.Helper()
	w, err := headless.NewWindow(sz.X, sz.Y)
	if err != nil {
		t.Skipf("headless windows not supported: %v", err)
	}
	return w, func() {
		w.Release()
	}
}
