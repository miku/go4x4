SHELL := /bin/bash

SIZE = 1073741824
B = 1

# 1048576       |   1MB
# 10485760      |  10MB
# 104857600     | 100MB
# 1073741824    |   1GB
# 10737418240   |  10GB
# 107374182400  | 100GB
# 1099511627776 |   1TB

teebench: main.go
	go build -o $@ $<

.PHONY: run
run: teebench
	cat /dev/zero | head -c $(SIZE) | ./teebench > /dev/null

.PHONY: clean
clean:
	rm -f teebench
