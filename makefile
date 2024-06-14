.DEFAULT_GOAL := build
build:
	go-winres make
	go build -ldflags="-s -H=windowsgui" -o "FFVIPR_Save_Editor.exe"
	upx -9 -k "FFVIPR_Save_Editor.exe"
	rm "FFVIPR_Save_Editor.ex~"
