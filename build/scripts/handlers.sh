#!/usr/bin/env bash

for file in ../handler/*.go
do
    BINARY_NAME=$(basename $file .go)
    GOOS=linux GOARCH=amd64 go build -o handler/$BINARY_NAME $file
done