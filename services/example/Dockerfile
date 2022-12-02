ARG ALPINE_VERSION=3.15.4
ARG GOLANG_VERSION=1.18

FROM golang:$GOLANG_VERSION as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOPROXY=https://proxy.golang.org,direct

ARG VERSION
ARG ServiceDir=services/example

RUN mkdir -p /app
WORKDIR /app

COPY $ServiceDir $ServiceDir

WORKDIR $ServiceDir

RUN go mod vendor

RUN go build -ldflags="-X example/version.Version=$VERSION" -o /app/example ./cmd/main.go

FROM alpine:$ALPINE_VERSION as prod

ARG ServiceDir=services/example

WORKDIR /app
COPY --from=builder /app/example .

CMD ["/app/example"]
