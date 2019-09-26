#!/bin/bash

PACKAGE=zip2png.go

printf "Building Linux...\n"
GOOS=linux go build -o "dist/linux/zip2png" $PACKAGE

printf "Building MacOS...\n"
GOOS=darwin go build -o "dist/macos/zip2png" $PACKAGE

printf "Building Windows...\n"
GOOS=windows go build -o "dist/windows/zip2png.exe" $PACKAGE

ls -alh dist/*

echo "Build finished"
