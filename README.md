# json2yaml

json2yaml is a command line tool that convert json to yaml either from stdin or from a file then prints the result to stdout, this small project was created in order to learn [goland](https://golang.org/) and [spf13/cobra](https://github.com/spf13/cobra) providing something useful.

## Requirements

    - Linux

## Installation

```
$ curl -O https://github.com/snebel29/json2yaml/releases/download/v.1.0.0/json2yaml
$ chmod u+x json2yaml
$ sudo mv json2yaml /usr/local/bin/
```

Or get it using go command line, your [_$GOPATH_](https://golang.org/doc/code.html#GOPATH) must be configured correctly

```
$ go get github.com/snebel29/json2yaml
```

## Examples

```
$ json2yaml --file=foo.json
$ json2yaml < foo.json
$ echo '{"foo": "bar"}' | json2yaml
```