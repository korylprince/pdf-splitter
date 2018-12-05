#!/usr/bin/bash

rm -Rf build
mkdir build

for GOOS in darwin linux windows; do
    #386 currently broken in UniDoc
    #for GOARCH in 386 amd64; do
    for GOARCH in amd64; do
        export GOOS
        export GOARCH
        if [ $GOOS = "windows" ]; then
            go build -v -o build/pdf-splitter-$GOOS-$GOARCH.exe
        else
            go build -v -o build/pdf-splitter-$GOOS-$GOARCH
        fi
    done
done
