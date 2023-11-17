FROM golang:1.21.3-alpine as compiler

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN mkdir -p /usr/local/bin/app && \
	go build -v -o /usr/local/bin/app ./...

FROM alpine

# installing libc6-compat cf. https://stackoverflow.com/a/66974607
RUN set -ex && \
	apk update && \
	apk add libc6-compat

COPY --from=compiler /usr/local/bin/app /app
COPY assets /assets
COPY templates /templates

ENTRYPOINT ["/app/gophercise-03-cyoa"]
