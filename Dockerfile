FROM golang:1.17.8-bullseye

RUN apt-get update && apt-get install -y --no-install-recommends clang

ENV CGO_ENABLED 1
ENV CXX clang++

COPY . /go-zetasql

WORKDIR /go-zetasql

RUN --mount=type=cache,id=bazel,target=/root/.cache/bazel \
        --mount=type=cache,id=gomod,target=/go/pkg/mod \
        --mount=type=cache,id=gobuild,target=/root/.cache/go-build \
    go build -v . && go test -v .