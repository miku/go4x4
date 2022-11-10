.PHONY: run
run: main
	GODEBUG=schedtrace=100 ./main

main: main.go
	go build -o main main.go
