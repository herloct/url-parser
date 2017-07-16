# url-parser [![Build Status](https://travis-ci.org/herloct/url-parser.svg?branch=master)](https://travis-ci.org/herloct/url-parser)

Inspired by [urlp](https://github.com/clayallsopp/urlp), a simple command-line utility for parsing URLs.

Implemented in Go, using standard library.

```bash
$ url-parser --part=host https://somedomain.com
somedomain.com

$ url-parser --part=user https://herloct@somedomain.com
herloct

$ url-parser --part=path https://somedomain.com/path/to
/path/to

$ url-parser --part=path --path-index=1 https://somedomain.com/path/to
to

$ url-parser --part=query https://somedomain.com/?some-key=somevalue
some-key=somevalue

$ url-parser --part=query --query-field=some-key https://somedomain.com/?some-key=somevalue
somevalue
```

## Instalation

url-parser is available for Linux and OS X, 64-bit only for now:

```curl
curl -L https://github.com/herloct/url-parser/releases/download/1.0.0-beta3/url-parser-`uname -s`-x86_64 > /usr/local/bin/url-parser; chmod +x /usr/local/bin/url-parser

```

For Windows, you could download them here:

```curl
https://github.com/herloct/url-parser/releases/download/1.0.0-beta3/url-parser-Windows-x86_64.exe
```

## Usage

```bash
$ url-parser --help
url-parser
  Parse URL and shows the part of it.

Usage:
  url-parser --part=PART <url>
  url-parser --part=path [--path-index=INDEX] <url>
  url-parser --part=query [--query-field=FIELD] <url>

Options:
  --part=PART          Part of URL to show [default: all].
                       Valid values: all, scheme, user, password,
                       hostname, port, path, query, or fragment.
  --path-index=INDEX   Filter parsed path by index.
  --query-field=FIELD  Filter parsed query string by field name.
```
