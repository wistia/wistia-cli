#!/bin/bash
set -e

go mod tidy
go build -o wistia ./cmd/wistia/
