package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/unit"
	syscall "golang.org/x/sys/windows"
	"unsafe"
)

func init() {
	//dpiAware = unit.Metric{PxPerDp: 1.25, PxPerSp: 1.25}
	//dpiAware = getDpiAwareMetric()
	SetProcessDPIAware()
	dpiAware = GetDPIAwareMetricForSystemDPI()
}

func getDpiAwareMetric() unit.Metric {
	m := unit.Metric{}
	w := app.NewWindow(
		app.Size(1, 1),
		app.Decorated(false),
	)
	for e := range w.Events() {
		switch et := e.(type) {

		case system.FrameEvent:
			m = et.Metric
			w.Perform(system.ActionClose)
		}
	}
	return m
}

const (
	MONITOR_DEFAULTTOPRIMARY = 1
	MDT_EFFECTIVE_DPI        = 0
	LOGPIXELSX               = 88
)

var (
	user32            = syscall.NewLazySystemDLL("user32.dll")
	_MonitorFromPoint = user32.NewProc("MonitorFromPoint")
	_GetDC            = user32.NewProc("GetDC")
	_ReleaseDC        = user32.NewProc("ReleaseDC")

	shcore            = syscall.NewLazySystemDLL("shcore")
	_GetDpiForMonitor = shcore.NewProc("GetDpiForMonitor")

	gdi32          = syscall.NewLazySystemDLL("gdi32")
	_GetDeviceCaps = gdi32.NewProc("GetDeviceCaps")
)

func SetProcessDPIAware() {
	user32 := syscall.NewLazySystemDLL("user32.dll")
	_SetProcessDPIAware := user32.NewProc("SetProcessDPIAware")
	_SetProcessDPIAware.Call()
}

// GetDPIAwareMetricForSystemDPI requires the process to be DPI aware. See SetProcessDPIAware.
func GetDPIAwareMetricForSystemDPI() unit.Metric {
	systemDpi := getSystemDPI()
	return getDPIAwareMetric(systemDpi)
}

// GetSystemDPI returns the effective DPI of the system.
func getSystemDPI() int {
	if _GetDpiForMonitor.Find() == nil {
		r, _, _ := _MonitorFromPoint.Call(uintptr(0), uintptr(0), uintptr(MONITOR_DEFAULTTOPRIMARY))
		hmon := syscall.Handle(r)
		var dpiX, dpiY uintptr
		_GetDpiForMonitor.Call(uintptr(hmon), uintptr(MDT_EFFECTIVE_DPI), uintptr(unsafe.Pointer(&dpiX)), uintptr(unsafe.Pointer(&dpiY)))
		return int(dpiX)
	} else {
		screenDC, err := getDC(0)
		if err != nil {
			return 96
		}
		defer releaseDC(screenDC)
		return getDeviceCaps(screenDC, LOGPIXELSX)
	}
}

func getDeviceCaps(hdc syscall.Handle, index int32) int {
	c, _, _ := _GetDeviceCaps.Call(uintptr(hdc), uintptr(index))
	return int(c)
}

func releaseDC(hdc syscall.Handle) {
	_ReleaseDC.Call(uintptr(hdc))
}

func getDC(hwnd syscall.Handle) (syscall.Handle, error) {
	hdc, _, err := _GetDC.Call(uintptr(hwnd))
	if hdc == 0 {
		return 0, fmt.Errorf("GetDC failed: %v", err)
	}
	return syscall.Handle(hdc), nil
}

func getDPIAwareMetric(dpi int) unit.Metric {
	const inchPrDp = 1.0 / 96.0
	ppdp := float32(dpi) * inchPrDp
	return unit.Metric{
		PxPerDp: ppdp,
		PxPerSp: ppdp,
	}
}
