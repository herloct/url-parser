package main

import (
	"testing"
)

type TestParseData struct {
	urlString  string
	part       uint
	pathIndex  uint
	queryField string
	expected   string
}

func (d *TestParseData) init() {
	d.pathIndex = 0
	d.queryField = ""
}

func TestParse(t *testing.T) {
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#f"

	testDatas := []TestParseData{
		TestParseData{
			urlString: urlString,
			part:      ALL_PART,
			expected:  urlString,
		},
		TestParseData{
			urlString: urlString,
			part:      SCHEME_PART,
			expected:  "postgres",
		},

		TestParseData{
			urlString: urlString,
			part:      USER_PART,
			expected:  "user",
		},
		TestParseData{
			urlString: "postgres://host.com",
			part:      USER_PART,
			expected:  "",
		},

		TestParseData{
			urlString: urlString,
			part:      PASSWORD_PART,
			expected:  "pass",
		},
		TestParseData{
			urlString: "postgres://host.com",
			part:      PASSWORD_PART,
			expected:  "",
		},
		TestParseData{
			urlString: "postgres://user@host.com",
			part:      PASSWORD_PART,
			expected:  "",
		},
		TestParseData{
			urlString: "postgres://user:my password@host.com",
			part:      PASSWORD_PART,
			expected:  "my password",
		},

		TestParseData{
			urlString: urlString,
			part:      HOSTNAME_PART,
			expected:  "host.com",
		},

		TestParseData{
			urlString: urlString,
			part:      PORT_PART,
			expected:  "5432",
		},
		TestParseData{
			urlString: "postgres://host.com",
			part:      PORT_PART,
			expected:  "",
		},

		TestParseData{
			urlString: urlString,
			part:      PATH_PART,
			expected:  "/path/to",
		},
		TestParseData{
			urlString: urlString,
			part:      PATH_PART,
			pathIndex: 1,
			expected:  "path",
		},
		TestParseData{
			urlString: urlString,
			part:      PATH_PART,
			pathIndex: 2,
			expected:  "to",
		},
		TestParseData{
			urlString: "postgres://host.com",
			part:      PATH_PART,
			expected:  "",
		},

		TestParseData{
			urlString: urlString,
			part:      QUERY_PART,
			expected:  "key=value&other=other%20value",
		},
		TestParseData{
			urlString:  urlString,
			part:       QUERY_PART,
			queryField: "key",
			expected:   "value",
		},
		TestParseData{
			urlString:  urlString,
			part:       QUERY_PART,
			queryField: "other",
			expected:   "other value",
		},
		TestParseData{
			urlString: "postgres://host.com",
			part:      QUERY_PART,
			expected:  "",
		},
		TestParseData{
			urlString:  "postgres://host.com",
			part:       QUERY_PART,
			queryField: "key",
			expected:   "",
		},
	}

	result := ""
	for _, testData := range testDatas {
		result = parse(testData.urlString, testData.part, testData.pathIndex, testData.queryField)
		if result != testData.expected {
			t.Fatalf("The part `%v`, should be `%v`", result, testData.expected)
		}
	}
}
