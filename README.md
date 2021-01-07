# Example mock rest server

This implements a fantasy REST api to handle quotas.

## How to build

Get a git checkout of the code.

Run:
```
$ go build -o ./qrest ./cmd/mgrd
$ ./qrest
$ curl http://localhost:8080/
```

## How to test

```
$ go test ./...
```

## Working with the code

Please use gofmt/go fmt on your code.

## External libraries used

This code is using the Gorilla web toolkit:
https://www.gorillatoolkit.org/

Test are using check.v1:
http://labix.org/gocheck

