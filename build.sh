#!/bin/sh

GOOS=linux GOARCH=amd64 docker run --rm -v "`pwd`:/xform" -v /tmp/xform-go-build:/root/.cache/go-build -w /xform cardboard/golang:1.17.6 go build -mod=vendor -v -o xformd ./cmd/xformd