SHELL := /bin/bash

data.js: symtab.txt
	python3 simplify.py <(python3 tab2pydic.py symtab.txt) > data.js

main: main.go
	CGO_ENABLED=1 go build -o main main.go

symtab.txt: main
	go tool nm -size main | c++filt > symtab.txt


.PHONY: clean
clean:
	rm -f data.js
	rm -f symtab.txt
	rm -f main
