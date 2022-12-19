package parser

import (
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"testing"
)

func TestSubstring(t *testing.T) {
	tests := []struct {
		name, input, expected, start, length string
		err                                  error
	}{
		{
			name:     "correct",
			input:    "Test string",
			expected: "est",
			start:    "1",
			length:   "3",
			err:      nil,
		},
		{
			name:     "start > len must be empty ",
			input:    "Test",
			expected: "",
			start:    "5",
			length:   "10",
			err:      nil,
		},
		{
			name:     "start+length > len",
			input:    "Zimbabwe",
			expected: "bwe",
			start:    "5",
			length:   "10",
			err:      nil,
		},
		{
			name:     "start+length > len",
			input:    "Zimbabwe",
			expected: "bw",
			start:    "5",
			length:   "2",
			err:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res, err := substring(pcommon.NewValueString(tt.input), tt.start, tt.length)
			assert.Equal(t, err, tt.err)
			assert.Equal(t, tt.expected, res.AsString())

		})
	}
}
