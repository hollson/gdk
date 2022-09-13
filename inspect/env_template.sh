#!/bin/bash

go run  -ldflags \
   "-X 'main.version=v1.0.0'
   -X 'main.built=$(date "+%Y-%m-%d %H:%M:%S")'
   -X 'main.gitCommit=$(git rev-parse --short HEAD)'" \
  main.go