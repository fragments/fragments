#!/usr/bin/env bash

echo "" > coverage.txt

rm -rf /tmp/junit
mkdir -p /tmp/junit
mkdir -p testoutput/go-test

for pkg in $(go list ./...); do
    # Run tests, capture coverage and output to pkg.out
    # Capture exit code rather than exiting immediately so we can output junit results from it
    set +e
    go test -coverprofile=profile.out -race -covermode=atomic -tags=integration -v $pkg 2>&1 > pkg.out
    code=$(echo $?)
    set -e

    # Output to log
    cat pkg.out

    # Convert to junit xml
    filename=$(echo "$pkg.xml" | sed 's/\//-/g')
    cat pkg.out | go-junit-report -no-xml-header > "/tmp/junit/$filename"
    files=$(find /tmp/junit/*.xml | tr '\n' ' ')
    if [ "$files" ]; then
        xunitmerge $files testoutput/go-test/junit.xml
    fi

    # Exit if tests failed
    if [ $code -ne 0 ]; then
        exit $code
    fi

    # Merge coverage
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi

    # Cleanup
    rm pkg.out
done
