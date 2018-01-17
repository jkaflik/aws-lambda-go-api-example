AWS Lambda Golang API example
===

AWS at 15 Jan [announced the Go support for Lambda](https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/). This repository contains simple approach over serverless API using [Serverless framework](https://serverless.com/).

Each handler consists of one `.go` file in `handler/` directory. Building those and packaging into AWS compliant package is managed by Makefile:

    make handlers package
    
Deploy using `serverless deploy` or:

    make deploy 

TODO
----
* unit tests
* e2e tests
* CLI command with HTTP server