SHELL := /bin/bash

.PHONY: stats
stats:
	@grep -c '^* ' README.md && grep -c '^    * ' README.md

embed:
	embedmd -w *.md

