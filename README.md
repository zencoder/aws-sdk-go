# AWS SDK for Go

[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/awslabs/aws-sdk-go)
[![Build Status](https://img.shields.io/travis/awslabs/aws-sdk-go.svg)](https://travis-ci.org/awslabs/aws-sdk-go)
[![Apache V2 License](http://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/awslabs/aws-sdk-go/blob/master/LICENSE)

aws-sdk-go is a set of clients for all Amazon Web Services APIs,
automatically generated from the JSON schemas shipped with
[botocore](http://github.com/boto/botocore).

It supports all known AWS services, and maps exactly to the documented
APIs, with some allowances for Go-specific idioms (e.g. `ID` vs. `Id`).

## Caution

It is currently **highly untested**, so please be patient and report any
bugs or problems you experience. The APIs may change radically without
much warning, so please vendor your dependencies w/ Godep or similar.

Please do not confuse this for a stable, feature-complete library.

Note that many services are currently not fully implemented. Some operations
might work, but not all, especially for XML-based services. We are currently
working to build out better service support, please bear with us!

## Installing

Let's say you want to use EC2:

    $ go get github.com/awslabs/aws-sdk-go/service/ec2
    
**NOTE**: If you are trying to use the development branch, after performing the command above, you must additionally check out the development branch:
 
    $ cd $GOPATH/src/github.com/awslabs/aws-sdk-go; git checkout develop

## Using

```go
import "github.com/awslabs/aws-sdk-go/aws"
import "github.com/awslabs/aws-sdk-go/service/ec2"

creds := aws.Creds(accessKey, secretKey, "")
cli := ec2.New(creds, "us-west-2", nil)
resp, err := cli.DescribeInstances(nil)
if err != nil {
    panic(err)
}
fmt.Println(resp.Reservations)
```
