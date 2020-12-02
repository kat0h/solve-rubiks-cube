all: main

main: main.go
	mkdir target
	go build main.go

run:
	go run main.go

clean:
	rm target/*
