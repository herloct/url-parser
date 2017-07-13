package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

const (
	ALL_PART uint = iota
	SCHEME_PART
	USER_PART
	PASSWORD_PART
	HOSTNAME_PART
	PORT_PART
	PATH_PART
	QUERY_PART
)

func parse(urlString string, part uint, pathIndex uint, queryField string) string {
	result := ""

	url, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}

	switch part {
	case SCHEME_PART:
		result = url.Scheme

	case USER_PART:
		if url.User != nil {
			result = url.User.Username()
		}

	case PASSWORD_PART:
		if url.User == nil {
			break
		}

		pass, hasPassword := url.User.Password()
		if hasPassword {
			result = pass
		}

	case HOSTNAME_PART:
		result = url.Hostname()

	case PORT_PART:
		result = url.Port()

	case PATH_PART:
		result = url.Path
		if pathIndex > 0 {
			paths := strings.Split(result, "/")
			result = paths[pathIndex]
		}

	case QUERY_PART:
		result = url.RawQuery
		if queryField != "" && result != "" {
			result = url.Query().Get(queryField)
		}

	default:
		result = url.String()
	}

	return result
}

func usage() {
	appName := "url-parser"
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", appName)
	fmt.Fprintf(os.Stderr, "\t%s [flags] some_url\n", appName)
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage

	schemePtr := flag.Bool("scheme", false, "show the scheme part")
	userPtr := flag.Bool("user", false, "show the user part")
	passwordPtr := flag.Bool("password", false, "show the password part")
	hostnamePtr := flag.Bool("hostname", false, "show the hostname part")
	portPtr := flag.Bool("port", false, "show the port part")
	pathPtr := flag.Bool("path", false, "show the raw path part")
	queryPtr := flag.Bool("query", false, "show the raw query string part")

	pathIndexPtr := flag.Uint("path-index", 0, "filter parsed path by index")
	queryFieldPtr := flag.String("query-field", "", "filter parsed query by field name")

	flag.Parse()

	if len(flag.Args()) < 1 {
		err := errors.New("Please provide an URL to parse")
		panic(err)
	}
	urlString := flag.Args()[0]
	part := ALL_PART
	pathIndex := uint(0)
	queryField := ""

	switch {
	case *schemePtr:
		part = SCHEME_PART

	case *userPtr:
		part = USER_PART

	case *passwordPtr:
		part = PASSWORD_PART

	case *hostnamePtr:
		part = HOSTNAME_PART

	case *portPtr:
		part = PORT_PART

	case *pathPtr:
		part = PATH_PART
		fallthrough

	case *pathPtr && *pathIndexPtr > 0:
		pathIndex = *pathIndexPtr + 1

	case *queryPtr:
		part = QUERY_PART
		fallthrough

	case *queryPtr && *queryFieldPtr != "":
		queryField = *queryFieldPtr

	default:
		part = ALL_PART
	}

	result := parse(urlString, part, pathIndex, queryField)
	fmt.Println(result)
}
