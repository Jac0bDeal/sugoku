#!/usr/bin/env bash

# STEP 1: Determinate the required values

VERSION="$(git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')"

# STEP 2: Build the ldflags

LDFLAGS=(
  "-X 'main.Version=${VERSION}'"
)

# STEP 3: Actual Go build process

go build -ldflags="${LDFLAGS[*]}" -o bin/sugoku ./cmd/sugoku
