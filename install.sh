#!/bin/bash
# 
# Build & Install Golang CLI
#

CGO_ENABLED=0 go build -ldflags="-s -w" -buildvcs=false 
mv crypt ~/go/bin/