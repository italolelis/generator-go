# Generator Go

> A HelloFresh golang generator application

## Features

We will generate a new application that is completely compliant with hellofresh production ready services. It will come with:

 - Basic HTTP metrics reporting to prometheus
 - Basic Distributed Tracing metrics that snds tracing to Jaeger
 - Logging client using json log format
 - A resilient HTTP client that implements a [hystrix-like](https://github.com/afex/hystrix-go) circuit breaker and a retry mechanism
 - Simple Makefile that builds and run tests for your app
 - Simple Docker file that is ready to build your image
 - A health check library that you can add checks for your dependencies
 - A RAML API documentation

Your service will expose a few endpoints:

- `/` - A HelloWorld example
- `/github/repos` - An example of how to use the cirucit breaker and retry HTTP client
- `/docs` - Your API documentation
- `/hystrix` - Here you can have a stream of your circuits
- `/metrics` - This is a stream of your prometheus metrics

## Installation

### Using locally installed node.js

First, install [Yeoman](http://yeoman.io) and generator-go using [npm](https://www.npmjs.com/) (we assume you have pre-installed [node.js](https://nodejs.org/)).

```bash
npm install -g yo

npm install --global generator-go
```

Then generate your new project:

```bash
yo go
```

### Using pre-built docker container

```bash
docker pull italolelis/generator-go && docker run --rm -it -v $GOPATH:/home/yo/go italolelis/generator-go
```

## What do you get?

Scaffolds out a complete go directory structure for you:

```
.
├── Dockerfile
├── Gopkg.toml
├── Makefile
├── cmd
│   ├── root.go
│   ├── server_start.go
│   └── version.go
├── docs
│   └── specification
│       ├── api.raml
│       └── types
│           └── Status.raml
├── main.go
└── pkg
    ├── config
    │   ├── config.go
    │   └── context.go
    ├── handler
    │   ├── docs.go
    │   ├── github.go
    │   └── hello_world.go
    ├── log
    │   ├── context.go
    │   └── middleware.go
    ├── metrics
    │   ├── context.go
    │   └── middleware.go
    └── tracer
        ├── jaeger_log.go
        ├── middleware.go
        └── tracer.go
```

## Usage

After you've generated your application, you can also just build your application by running:

```sh
make
```

### Commands

A few commands are available for you out of the box

| Command                  | Description                          |
|--------------------------|--------------------------------------|
| `<your-service-name> start`    | Starts TCP server for your application|
| `<your-service-name> version`  | Prints the version information |
