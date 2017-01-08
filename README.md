# go-testhttpd [![Build Status](http://ec2-54-194-144-141.eu-west-1.compute.amazonaws.com/api/badges/cp2017/go-testhttpd/status.svg)](http://ec2-54-194-144-141.eu-west-1.compute.amazonaws.com/cp2017/go-testhttpd)
http-server which returns sequences of error codes for testing purposes.

## Usage

The binary consumes a sequence of status codes, which will be returned in a round-robin fashion.

```
./resources/bin/go-testhttpd_darwin --help
NAME:
   go-testhttpd - A new cli application

USAGE:
   go-testhttpd_darwin [global options] command [command options] [arguments...]

VERSION:
   1.0.2

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --status-sequence value  Sequence of HTTP status codes to return (default: "200,201,202")
   --help, -h               show help
   --version, -v            print the version
```

## Example

Start the daemon...

```
$ ./resources/bin/go-testhttpd_darwin --status-sequence 200,201,400,500
```

Fetch the resulting header:

```
$ curl -sI http://localhost:8080/
HTTP/1.1 200 OK
Date: Sun, 08 Jan 2017 09:48:46 GMT
Content-Length: 8
Content-Type: text/plain; charset=utf-8
$ curl -sI http://localhost:8080/
HTTP/1.1 201 Created
Date: Sun, 08 Jan 2017 09:48:48 GMT
Content-Length: 13
Content-Type: text/plain; charset=utf-8
$ curl -sI http://localhost:8080/
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Sun, 08 Jan 2017 09:48:48 GMT
Content-Length: 29
$ curl -sI http://localhost:8080/
HTTP/1.1 500 Internal Server Error
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Sun, 08 Jan 2017 09:48:48 GMT
Content-Length: 49
$ curl -sI http://localhost:8080/
HTTP/1.1 200 OK
Date: Sun, 08 Jan 2017 09:48:46 GMT
Content-Length: 8
Content-Type: text/plain; charset=utf-8
```

The returned content is `<status_code>, StatusText(<status_code>)`, e.g.:

```
$ curl -s http://localhost:8080/
200, OK
$ curl -s http://localhost:8080/
201, Created
$ curl -s http://localhost:8080/
Bad Request
400, Bad Request
$ curl -s http://localhost:8080/
Internal Server Error
500, Internal Server Error
```
