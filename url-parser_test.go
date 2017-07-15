package main

import (
	"testing"
)

type TestParseData struct {
	urlString string
	part      string
	index     int
	field     string
	expected  string
}

func TestParse(t *testing.T) {
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#f"

	testDatas := []TestParseData{
		TestParseData{
			urlString: urlString,
			part:      "all",
			expected:  urlString,
		},
		TestParseData{
			urlString: urlString,
			part:      "scheme",
			expected:  "postgres",
		},

		TestParseData{
			urlString: urlString,
			part:      "user",
			expected:  "user",
		},
		TestParseData{
			urlString: "postgres://host.com",
			part:      "user",
			expected:  "",
		},

		TestParseData{
			urlString: urlString,
			part:      "password",
			expected:  "pass",
		},
		TestParseData{
			urlString: "postgres://host.com",
			part:      "password",
			expected:  "",
		},
		TestParseData{
			urlString: "postgres://user@host.com",
			part:      "password",
			expected:  "",
		},
		TestParseData{
			urlString: "postgres://user:my password@host.com",
			part:      "password",
			expected:  "my password",
		},

		TestParseData{
			urlString: urlString,
			part:      "hostname",
			expected:  "host.com",
		},

		TestParseData{
			urlString: urlString,
			part:      "port",
			expected:  "5432",
		},
		TestParseData{
			urlString: "postgres://host.com",
			part:      "port",
			expected:  "",
		},

		TestParseData{
			urlString: urlString,
			part:      "path",
			index:     -1,
			expected:  "/path/to",
		},
		TestParseData{
			urlString: urlString,
			part:      "path",
			index:     0,
			expected:  "path",
		},
		TestParseData{
			urlString: urlString,
			part:      "path",
			index:     1,
			expected:  "to",
		},
		TestParseData{
			urlString: "postgres://host.com",
			part:      "path",
			index:     -1,
			expected:  "",
		},

		TestParseData{
			urlString: urlString,
			part:      "query",
			expected:  "key=value&other=other%20value",
		},
		TestParseData{
			urlString: urlString,
			part:      "query",
			field:     "key",
			expected:  "value",
		},
		TestParseData{
			urlString: urlString,
			part:      "query",
			field:     "other",
			expected:  "other value",
		},
		TestParseData{
			urlString: "postgres://host.com",
			part:      "query",
			expected:  "",
		},
		TestParseData{
			urlString: "postgres://host.com",
			part:      "query",
			field:     "key",
			expected:  "",
		},
	}

	var (
		result string
		err    error
	)
	for _, testData := range testDatas {
		result, err = parse(testData.urlString, testData.part, testData.index, testData.field)
		if err != nil {
			t.Fatal(err)
		}
		if result != testData.expected {
			t.Fatalf("The part `%v`, should be `%v`", result, testData.expected)
		}
	}
}
