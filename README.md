goa-getting-started
===================

a customized goa-celler example for goa https://github.com/goadesign/goa/

## How go generate swagger and HTML documents

```
$ go get -u github.com/hnakamur/goa-getting-started
$ cd $GOPATH/src/github.com/hnakamur/goa-getting-started
$ goagen gen --pkg-path github.com/hnakmaur/localeoverlayswagger --design github.com/hnakamur/goa-getting-started/design
$ npm i
$ ./node_modules/.bin/bootprint swagger swagger/swagger.json api-doc
$ ./node_modules/.bin/bootprint swagger swagger/swagger.ja.json api-doc.ja
```
