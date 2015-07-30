#!/bin/bash

go list ./...  | xargs -n1 go test
