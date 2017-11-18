package aws

import (
	"encoding/json"
	"log"
)

const defaultAssumeLambdaExecPolicy = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}`

func mustCompress(input string) string {
	var tmp interface{}
	if err := json.Unmarshal([]byte(input), &tmp); err != nil {
		log.Panic(err)
	}
	out, err := json.Marshal(tmp)
	if err != nil {
		log.Panic(err)
	}
	return string(out)
}
