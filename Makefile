default: build

build:
	go build -v -o ./bin/snake

run: build
	./bin/snake
