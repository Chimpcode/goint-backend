#!/bin/sh

cd "$(dirname "$0")"

go build -o goint-backend
./goint-backend

