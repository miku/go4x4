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

noopfilter: main.go
	go build -o $@ $<

.PHONY: run
run: noopfilter
	cat /dev/zero | head -c $(SIZE) | ./noopfilter > /dev/null

.PHONY: runb
runb: noopfilter
	cat /dev/zero | head -c $(SIZE) | ./noopfilter -b $(B) > /dev/null

largefile:
	cat /dev/zero | head -c $(SIZE) > $@

.PHONY: clean
clean:
	rm -f noopfilter largefile

