#!/bin/bash


env GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o ./dist/server_win64.exe ./
env GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -o ./dist/server_win32.exe ./
env GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o ./dist/server_macos ./
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./dist/server_linux ./