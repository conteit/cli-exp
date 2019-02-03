#! /bin/bash

GO111MODULE=on goreleaser release --skip-publish --rm-dist $*

