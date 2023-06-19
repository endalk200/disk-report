package main

import (
	"testing"
)

func TestFormatSize(t *testing.T) {
	type ExpectedOutPut []struct {
		size     int64
		expected string
	}

	tests := ExpectedOutPut{
		{0, "0 B"},
		{512, "512 B"},
		{1024, "1.00 KB"},
		{1536, "1.50 KB"},
		{1048576, "1.00 MB"},
		{1572864, "1.50 MB"},
		{1073741824, "1.00 GB"},
		{1610612736, "1.50 GB"},
	}

	for _, test := range tests {
		result := formatSize(test.size)
		if result != test.expected {
			t.Errorf("formatSize(%d) = %s; expected %s", test.size, result, test.expected)
		}
	}
}
