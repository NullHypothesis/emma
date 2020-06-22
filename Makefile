output_dir = bin

build: *.go
	go build -o $(output_dir)/emma

osx: *.go
	# The command 'go tool dist list' shows you a list of all supported
	# platforms and architectures.
	GOOS=darwin go build -o $(output_dir)/emma.app

windows: *.go
	GOOS=windows go build -o $(output_dir)/emma.exe

linux: *.go
	GOOS=linux go build -o $(output_dir)/emma
