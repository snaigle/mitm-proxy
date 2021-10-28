#!/bin/bash
GO111MODULE=on
go-bindata -nomemcopy -pkg assets -tags release -debug=false -o assets/assets_release.go resources/...
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-extldflags '-static' -s -w" -tags=release -o build/goproxy2_linux_amd64
