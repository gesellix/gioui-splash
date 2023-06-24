package splash

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"image"
	"image/color"
)

type Splash struct {
	img            image.Image
	progressHeight unit.Dp
	progressInset  layout.Inset
	progressCol    color.NRGBA
	progress       float64
}

func NewSplash(img image.Image, progressInset layout.Inset, progressCol color.NRGBA) *Splash {
	return &Splash{
		img,
		progressInset.Bottom - progressInset.Top,
		progressInset,
		progressCol,
		0,
	}
}

func (s *Splash) SetProgress(progress float64) {
	s.progress = progress
	if s.progress > 1 {
		s.progress = 1
	}
	if s.progress < 0 {
		s.progress = 0
	}
}

func (s *Splash) Layout(gtx layout.Context) layout.Dimensions {
	// Lay out the image and overlay the progress bar on top of it, aligning
	// to the south so that the progress bar is at the bottom.
	return layout.Stack{
		Alignment: layout.S,
	}.Layout(gtx,
		layout.Stacked(s.drawLogo),
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			// Lay out the progress bar within the inset.
			return s.progressInset.Layout(gtx, s.drawProgress)
		}),
	)
}

func (s *Splash) drawLogo(gtx layout.Context) layout.Dimensions {
	return widget.Image{
		Src: paint.NewImageOp(s.img),
		Fit: widget.Contain,
	}.Layout(gtx)
}

func (s *Splash) drawProgress(gtx layout.Context) layout.Dimensions {
	// Our gtx.Constraints.Min.X provide the width
	// that a full progress bar must occupy.
	spread := gtx.Constraints.Min.X
	rectangle := image.Rectangle{
		Max: image.Pt(
			// Since spread is already a unit of screen pixels,
			// no need to pass through gtx.Dp to convert to them.
			int(float64(spread)*s.progress),
			// Convert our progress bar height to screen pixels.
			gtx.Dp(s.progressHeight),
		),
	}
	paint.FillShape(
		gtx.Ops,
		s.progressCol,
		clip.Rect(rectangle).Op())
	// Return the logical dimensions occupied by the progress bar,
	// which is the width of a full bar.
	return layout.Dimensions{
		Size: image.Point{
			X: gtx.Constraints.Min.X,
			Y: gtx.Dp(s.progressHeight),
		},
	}
}
