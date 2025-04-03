package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testCases := []struct {
		headers   http.Header
		expected  string
		expectErr bool
	}{
		{
			headers:   http.Header{},
			expected:  "",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		got, error := GetAPIKey(tc.headers)
		if !(got == tc.expected) {
			t.Errorf("Didn't get expected output")
		}
		if error != nil && tc.expectErr == false {
			t.Errorf("Expected error")
		}
	}
}
