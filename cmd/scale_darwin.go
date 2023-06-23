package main

import (
	"gioui.org/unit"
)

func init() {
	scale := float32(2)
	dpiAware = unit.Metric{PxPerDp: scale, PxPerSp: scale}
	//scale = float32(C.getViewBackingScale(0))
}
