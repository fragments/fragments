#!/usr/bin/env bash

set -e

go generate $(go list ./...)
