# Example mock rest server

This implements a fantasy REST api to handle quotas.

## How to build

Run:
```
$ go build -o ./qrest github.com/mvo5/qrest-skeleton/cmd/mgrd
$ ./qrest
$ curl http://localhost:8080/
```

## How to test

```
$ go test github.com/mvo5/qrest-skeleton/...
```

## External libraries used

This code is using the Gorilla web toolkit:
https://www.gorillatoolkit.org/

Test are using check.v1:
http://labix.org/gocheck

