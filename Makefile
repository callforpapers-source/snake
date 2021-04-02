default: build

build:
	go build -v -o ./snake

run: build
	./snake
