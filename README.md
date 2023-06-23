# gioui-splash

A cross-platform splash screen for Go, based on Gio UI

## About the Gio UI Splash Widget

Sometimes you'll need a splash screen, or loading indicator, optionally with a progress bar.
_gioui-splash_ is a small widget for Gio UI, supporting:

- any image to be displayed in a non-decorated window
- an optional progress bar to display asynchronous task progress
- tweaking of the position, size, and color of the progress bar

[Gio UI](https://gioui.org/) is a library for writing cross-platform immediate mode GUI-s in Go.

A minimal example can be run like this:

```shell
go run github.com/gesellix/gioui-splash/cmd@latest
```

The example should work on any platform supported by Gio UI, including WebAssembly. See below for details.

## Build and package with gogio

The gogio tool simplifies building and packaging your binaries.

```shell
go install gioui.org/cmd/gogio@latest
```

### MacOS

```shell
gogio -x -target macos -ldflags "-s -w" -appid splash -icon appicon.png -o splash.app ./cmd
open splash_arm64.app 
#open splash_amd64.app 
```

### Windows

```shell
gogio -x -target windows -ldflags "-s -w" -appid splash -icon appicon.png -o splash.exe ./cmd
.\splash.exe
```

### WebAssembly

```shell
gogio -x -target js -ldflags "-s -w" -appid splash -icon appicon.png -o wasm ./cmd
go run ./wasm
```

## Contributing/Developing

A modern [Golang](https://go.dev/dl/) installation is required. Some additional dependencies might be necessary, please consult the documentation at https://gioui.org/doc/install.

A fresh clone can be run like this:

```shell
git clone git@github.com:/gesellix/gioui-splash.git
cd guiui-splash
go run ./cmd
```

## Roadmap/ideas

- Refactor parts of the example code from cmd/main.go, so that custom Go code (the application code) won't be mixed with UI code.
- Reduce the packaged binary size (can we exclude Gio's themes and text modules?)
- Add tests, maybe using the `gioui.org/gpu/headless` module?
- Consider using Gio's widget/material/progressbar - without forgetting about the binary size ;)

## License

MIT License

## Notes

Thanks to the Gui UI developers for the great project! Please head over to https://gioui.org/doc/contribute if you want to support their work.
