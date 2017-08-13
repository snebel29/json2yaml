# json2yaml

json2yaml is a command line tool that convert json to yaml either from stdin or from a file then prints the result to stdout 

This was created in order to learn [goland](https://golang.org/) and [spf13/cobra](https://github.com/spf13/cobra) providing something useful even if small and non innovate.

## Installation

Ensure your [_$GOPATH_ is configured correctly](https://golang.org/doc/code.html#GOPATH), then simply 

```
$ go get github.com/snebel29/json2yaml
```

## Examples

```
$ json2yaml --file=foo.json
$ json2yaml < foo.json
$ echo '{"foo": "bar"}' | json2yaml
```