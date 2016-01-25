#!/bin/bash

# Get build + test dependencies. -d also doesn't bother with installing the
# packages, it just downloads them
go get -t -d
go test
