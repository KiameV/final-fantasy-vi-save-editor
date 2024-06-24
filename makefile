.DEFAULT_GOAL := build
build:
	go build -ldflags="-s -w -H=windowsgui" -o "FFPR_Save_Editor.exe"
	upx -9 -k "FFPR_Save_Editor.exe"
	rm "FFPR_Save_Editor.ex~"

setup:
	go install github.com/tc-hib/go-winres@latest
	go-winres simply --icon icon.png
