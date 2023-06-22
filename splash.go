package splash

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"image"
	"image/color"
)

type Splash struct {
	img          image.Image
	progressRect image.Rectangle
	progressCol  color.NRGBA
	progress     float64
}

func NewSplash(img image.Image, progressRect image.Rectangle, progressCol color.NRGBA) *Splash {
	return &Splash{
		img,
		progressRect,
		progressCol,
		0,
	}
}

func (s *Splash) SetProgress(progress float64) {
	s.progress = progress
}

func (s *Splash) Layout(gtx layout.Context) layout.Dimensions {
	dimensions := s.drawLogo(gtx)
	s.drawProgress(gtx, s.progress)
	return dimensions
}

func (s *Splash) drawLogo(gtx layout.Context) layout.Dimensions {
	return widget.Image{
		Src: paint.NewImageOp(s.img),
		Fit: widget.Contain,
	}.Layout(gtx)
}

func (s *Splash) drawProgress(gtx layout.Context, progress float64) {
	spread := s.progressRect.Max.X - s.progressRect.Min.X
	rectangle := image.Rectangle{
		Min: s.progressRect.Min,
		Max: image.Pt(
			s.progressRect.Min.X+int(progress*float64(spread)),
			s.progressRect.Max.Y,
		),
	}
	paint.FillShape(
		gtx.Ops,
		s.progressCol,
		clip.Rect(rectangle).Op())
}
