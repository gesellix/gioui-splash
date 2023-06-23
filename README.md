# gioui-splash

Example of a splash widget for Gio UI

## Run

```shell
go run ./cmd
```

## Build w/ gogio

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
