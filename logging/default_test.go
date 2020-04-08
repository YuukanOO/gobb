package logging

import (
	"regexp"
	"testing"

	"github.com/yuukanoo/gobb/assert"
)

func TestDefaultLogger(t *testing.T) {
	tests := []struct {
		name        string
		debug       bool
		message     string
		data        []F
		outExpected *regexp.Regexp
		errExpected *regexp.Regexp
	}{
		{
			name:        "message without data",
			message:     "a message but with no data",
			outExpected: regexp.MustCompile("^INFO\\t.*a message but with no data\\n$"),
			errExpected: regexp.MustCompile("^ERROR\\t.*a message but with no data\\n$"),
		},
		{
			name:        "message without data and debug",
			message:     "a message but with no data",
			debug:       true,
			outExpected: regexp.MustCompile("^INFO\\t.*a message but with no data\\n$"),
			errExpected: regexp.MustCompile("^ERROR\\t.*a message but with no data\\n$"),
		},
		{
			name:    "message with one data",
			message: "this message contains one data dict",
			data: []F{
				{
					"a":       "value",
					"another": 1,
				},
			},
			outExpected: regexp.MustCompile("^INFO\\t.*this message contains one data dict\\n\\ta: value\\n\\tanother: 1\\n$"),
			errExpected: regexp.MustCompile("^ERROR\\t.*this message contains one data dict\\n\\ta: value\\n\\tanother: 1\\n$"),
		},
		{
			name:    "message with one data and debug",
			message: "this message contains one data dict",
			data: []F{
				{
					"a":       "value",
					"another": 1,
				},
			},
			debug:       true,
			outExpected: regexp.MustCompile("^INFO\\t.*this message contains one data dict\\n\\ta: value\\n\\tanother: 1\\n$"),
			errExpected: regexp.MustCompile("^ERROR\\t.*this message contains one data dict\\n\\ta: value\\n\\tanother: 1\\n$"),
		},
		{
			name:    "message with many data",
			message: "this message contains two data dicts",
			data: []F{
				{
					"a":       "value",
					"another": 1,
				},
				{
					"second": "dict",
				},
			},
			outExpected: regexp.MustCompile("^INFO\\t.*this message contains two data dicts\\n\\ta: value\\n\\tanother: 1\\n\\tsecond: dict\\n$"),
			errExpected: regexp.MustCompile("^ERROR\\t.*this message contains two data dicts\\n\\ta: value\\n\\tanother: 1\\n\\tsecond: dict\\n$"),
		},
		{
			name:    "message with many data and debug",
			message: "this message contains two data dicts",
			data: []F{
				{
					"a":       "value",
					"another": 1,
				},
				{
					"second": "dict",
				},
			},
			debug:       true,
			outExpected: regexp.MustCompile("^INFO\\t.*this message contains two data dicts\\n\\ta: value\\n\\tanother: 1\\n\\tsecond: dict\\n$"),
			errExpected: regexp.MustCompile("^ERROR\\t.*this message contains two data dicts\\n\\ta: value\\n\\tanother: 1\\n\\tsecond: dict\\n$"),
		},
	}

	for _, test := range tests {
		out := &fakeWriter{}
		err := &fakeWriter{}
		logger := New(test.debug, out, err)

		t.Run(test.name, func(t *testing.T) {
			// Test info output
			logger.Info(test.message, test.data...)

			assert.Match(t, test.outExpected, out.written, "out writer should have been called correctly")
			assert.Equals(t, "", err.written, "err writer should be empty")

			// Test error output
			logger.Error(test.message, test.data...)
			assert.Match(t, test.errExpected, err.written, "err writer should have been called correctly")

			// Reset out written to test the Debug method
			out.written = ""

			logger.Debug(test.message, test.data...)

			if test.debug {
				assert.Match(t, test.outExpected, out.written, "out writer should have been called correctly")
			} else {
				assert.Equals(t, "", out.written, "not in debug, out writer should be empty")
			}
		})
	}
}

// fakeWriter which implements the writer interface.
type fakeWriter struct {
	written string
}

func (w *fakeWriter) Write(data []byte) (int, error) {
	w.written = string(data)
	return len(data), nil
}
