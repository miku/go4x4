################################
# STEP 1 build executable binary
################################
FROM golang:alpine AS builder

# Install git, required for fetching the dependencies.
RUN apk update && apk add --no-cache git make

WORKDIR /app
COPY . .

# Fetch dependencies, using go get, download only, verbose.
RUN go get -d -v

# Build the binary.
RUN make fetch

############################
# STEP 2 build a small image
############################
FROM scratch

# Copy our static executable.
COPY --from=builder /app/fetch /app/fetch

# https://stackoverflow.com/questions/52969195/docker-container-running-golang-http-client-getting-error-certificate-signed-by
ADD ca-certificates.crt /etc/ssl/certs/

# Run the hello binary.
ENTRYPOINT ["/app/fetch"]

