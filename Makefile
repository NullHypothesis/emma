output_dir = bin

all: windows linux

windows: *.go
	GOOS=windows GOARCH=386 go build -o $(output_dir)/emma.exe

linux: *.go
	go build -o $(output_dir)/emma
