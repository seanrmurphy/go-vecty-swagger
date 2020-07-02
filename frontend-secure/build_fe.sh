#! /usr/bin/env bash

env GOOS=js GOARCH=wasm go build -o build/out.wasm ./src

