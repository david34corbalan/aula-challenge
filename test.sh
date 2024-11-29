#!/bin/zsh
go test $(find . -type f -name '*.go' -not -path '*/example/*' -not -path '*/mocks/*' -not -path '*/migrations/*' -exec dirname {} \; | sort -u) -coverprofile=cover.out &&
go tool cover -html=cover.out &&
go tool cover -func=cover.out
