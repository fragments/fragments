package aws

import (
	"testing"

	"github.com/fragments/fragments/pkg/testutils"
)

func TestCompressPolicy(t *testing.T) {
	input := `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "arn:aws:logs:*:*:*"
    }
  ]
}`
	actual := mustCompress(input)

	testutils.AssertGolden(t, actual, "testdata/compressedPolicy.json")
}
