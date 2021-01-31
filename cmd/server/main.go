package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/skakanigit/email-processor/pkg/parser"
)
func main()  {
	lambda.Start(parser.HandleRequest)
}
