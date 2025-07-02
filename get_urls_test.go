package main

import (
	"testing"
	"reflect"

)

func TestGetURLs(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		inputBody	  string
		expected      []string
	}{
		{
			name:     "inside and outside links",
			inputURL: "https://blog.boot.dev",
			inputBody : `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "inside div",
			inputURL: "https://blog.boot.dev",
			inputBody : `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
		<div>
		<a href="/path/two">
		<span>Boot.dev</span>
		</a>
		</div>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one", "https://blog.boot.dev/path/two"},
		},
		{
			name:     "in a list",
			inputURL: "https://blog.boot.dev",
			inputBody : `
<html>
	<body>
	<ul>
		<li>
		  <a href="/path/one">
			<span>Boot.dev</span>
		  </a>
		</li>
		<li>
		  <a href="https://other.com/path/one">
			<span>Boot.dev</span>
		  </a>
	    </li>
	</ul>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},

	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
