CGO_ENABLED = 0

.PHONY: all
all:
	@echo "To create an image: make image"
	@echo "To remove image   : make clean"

fetch:
	# https://stackoverflow.com/questions/52969195/docker-container-running-golang-http-client-getting-error-certificate-signed-by
	CGO_ENABLED=$(CGO_ENABLED) go build -ldflags="-w -s" -o $@ main.go

.PHONY: clean
clean:
	rm -rf fetch
	docker rmi -f gosp/fetch

.PHONY: image
image: fetch
	# $ docker images -f 'reference=gosp/fetch'
	# REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
	# gosp/fetch          latest              210b92811047        2 minutes ago       7.08MB
	docker build -t gosp/fetch .

.PHONY: imagecgo
imagecgo: fetch
	# http://kefblog.com/2017-07-04/Golang-ang-docker
	docker build -t gosp/fetchcgo -f Dockerfile.cgo .

.PHONY: clean-dangling
clean-dangling:
	docker rmi -f $$(docker images -f dangling=true -q)

