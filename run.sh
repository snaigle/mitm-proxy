#!/bin/bash
GO111MODULE=on

go-bindata -nomemcopy -pkg assets -tags debug -debug=true -o assets/assets_debug.go resources/...

go run -tags=debug .