# ---------- DEV BUILD -------------------

ARG ALPINE_VERSION=3.15.4
ARG GOLANG_VERSION=1.18

FROM golang:$GOLANG_VERSION as dev

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOPROXY=https://proxy.golang.org,direct

ARG ServiceDir=services/server

RUN mkdir -p /app
WORKDIR /app

COPY $ServiceDir $ServiceDir

WORKDIR $ServiceDir

RUN go mod vendor

RUN go build -mod=vendor -o /app/server ./cmd/main.go

# ---------- PRODUCTION BUILD -------------------

ARG ALPINE_VERSION=3.15.4
ARG GOLANG_VERSION=1.18

FROM golang:$GOLANG_VERSION as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOPROXY=https://proxy.golang.org,direct

ARG VERSION
ARG GITCOMMIT
ARG DATE

ARG ServiceDir=services/server

RUN mkdir -p /app
WORKDIR /app

COPY $ServiceDir $ServiceDir

WORKDIR $ServiceDir

RUN go mod vendor

RUN \
    VERSION=$VERSION \
    DATE=$DATE \
    GITCOMMIT=$GITCOMMIT \
    go build -mod=vendor -ldflags="-X 'server/pkg/version.BinVersion=${VERSION}' -X 'server/pkg/version.GitCommit=${GITCOMMIT}' -X 'server/pkg/version.BuildTime=${DATE}'" -o /app/server ./cmd/main.go

FROM alpine:$ALPINE_VERSION as prod

ARG ServiceDir=services/server

WORKDIR /app
COPY --from=builder /app/server .

CMD ["/app/server"]
