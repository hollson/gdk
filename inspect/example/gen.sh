#!/bin/bash


function build() {
    go build -ldflags \
     "-X 'main.name=example'
      -X 'main.version=V1.0.0'
      -X 'github.com/hollson/gdk/inspect.built=$(date "+%Y-%m-%d %H:%M:%S")'
      -X 'github.com/hollson/gdk/inspect.branch=$(git rev-parse --abbrev-ref @{u})'
      -X 'github.com/hollson/gdk/inspect.commit=$(git rev-parse --short HEAD)'
      -X 'github.com/hollson/gdk/inspect.author=$(git config user.name)'
      -X 'github.com/hollson/gdk/inspect.tag=$(git describe --tags --abbrev=0)'
      -X 'github.com/hollson/gdk/inspect.environment=develop'" \
      -ldflags="-s -w" \
      main.go
}

function run() {
    go run -ldflags \
     "-X 'main.name=example'
      -X 'main.version=V1.0.0'
      -X 'github.com/hollson/gdk/inspect.built=$(date "+%Y-%m-%d %H:%M:%S")'
      -X 'github.com/hollson/gdk/inspect.branch=$(git rev-parse --abbrev-ref @{u})'
      -X 'github.com/hollson/gdk/inspect.commit=$(git rev-parse --short HEAD)'
      -X 'github.com/hollson/gdk/inspect.author=$(git config user.name)'
      -X 'github.com/hollson/gdk/inspect.tag=$(git describe --tags --abbrev=0)'
      -X 'github.com/hollson/gdk/inspect.environment=develop'" \
      main.go
}

"$1"