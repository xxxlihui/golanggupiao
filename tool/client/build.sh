#!/usr/bin/env bash

export GOOS=windows
go build -o nnclient.exe client.go
