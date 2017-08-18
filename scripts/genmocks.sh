#!/usr/bin/env bash

set -e

mocks="
vendor/github.com/aws/aws-sdk-go/service/s3/s3iface:internal/filestore/mocks
"

for input in $mocks
do
  IFS=":" read -r -a mock <<< "${input}"
  input=${mock[0]}
  output=${mock[1]:-$input/mocks}
  echo "$input:"
  mockery -dir $input -output $output -all
done
