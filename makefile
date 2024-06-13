build:
	go build -ldflags="-s -H=windowsgui" -o ffvi_editor.exe
	upx -9 -k ffvi_editor.exe
	rm ffvi_editor.ex~
