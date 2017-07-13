# url-parser [![Build Status](https://travis-ci.org/herloct/url-parser.svg?branch=master)](https://travis-ci.org/herloct/url-parser)

Inspired by [urlp](https://github.com/clayallsopp/urlp), a simple command-line utility for parsing URLs.

Implemented in Go, using standard library.

```bash
$ url-parser -host https://somedomain.com
somedomain.com

$ url-parser -user https://herloct@somedomain.com
herloct

$ url-parser -path https://somedomain.com/path/to
/path/to

$ url-parser -path -path-index=1 https://somedomain.com/path/to
to

$ url-parser -query https://somedomain.com/?some-key=somevalue
some-key=somevalue

$ url-parser -query -query-field=some-key https://somedomain.com/?some-key=somevalue
somevalue
```

## Instalation

url-parser is available for Linux and OS X, 64-bit only for now:

```curl
curl -L https://github.com/herloct/url-parser/releases/download/1.0.0-beta2/url-parser-`uname -s`-x86_64 > /usr/local/bin/url-parser; chmod +x /usr/local/bin/url-parser

```

For Windows, you could download them here:

```curl
https://github.com/herloct/url-parser/releases/download/1.0.0-beta2/url-parser-Windows-x86_64.exe
```

## Usage

```bash
$ url-parser --help
Usage of url-parser:
        url-parser [flags] some_url
Flags:
  -hostname
        show the hostname part
  -password
        show the password part
  -path
        show the raw path part
  -path-index uint
        filter parsed path by index
  -port
        show the port part
  -query
        show the raw query string part
  -query-field string
        filter parsed query by field name
  -scheme
        show the scheme part
  -user
        show the user part
```
