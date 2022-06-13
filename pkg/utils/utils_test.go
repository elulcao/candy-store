package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeapString(t *testing.T) {
	tests := []struct {
		name         string
		data         map[string]int64
		expectedData string
	}{
		{
			name:         "Test one element",
			data:         map[string]int64{"Annika": 100},
			expectedData: "Annika",
		},
		{
			name:         "Test two elements",
			data:         map[string]int64{"Annika": 100, "Jonas": 200},
			expectedData: "Jonas",
		},
		{
			name:         "Test three elements",
			data:         map[string]int64{"Annika": 100, "Jane": 200, "Jonas": 1},
			expectedData: "Jane",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedData, MaxHeapString(test.data))
	}
}
