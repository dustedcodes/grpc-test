FROM golang:alpine AS build-env

WORKDIR /app
COPY . .
RUN go build -o /go/bin ./...

FROM alpine:latest
WORKDIR /app
COPY --from=build-env /go/bin/server .
ENTRYPOINT ["./server"]