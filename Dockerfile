FROM golang:1.12 as builder
WORKDIR /go/src/github.com/gerald1248/event-exporter
ADD . ./
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GO111MODULE on
RUN \
  go mod download && \
  go get && \
  go vet && \
  go test -v && \
  go build

FROM ubuntu:18.10
WORKDIR /app/
RUN groupadd app && useradd -g app app
COPY --from=builder /go/src/github.com/gerald1248/event-exporter /usr/local/bin/event-exporter
USER app
CMD ["event-exporter", "-types=Warning"]
