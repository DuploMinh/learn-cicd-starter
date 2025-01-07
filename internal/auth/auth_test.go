package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		want        string
		expectError error
	}{
		{
			name: "Valid ApiKey Header",
			headers: http.Header{
				"Authorization": []string{"ApiKey some_valid_key"},
			},
			want:        "some_valid_key",
			expectError: nil,
		},
		{
			name: "Missing Authorization Header",
			headers:     http.Header{},
			want:        "",
			expectError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Malformed Authorization Header (Invalid Scheme)",
			headers: http.Header{
				"Authorization": []string{"Bearer some_token"},
			},
			want:        "",
			expectError: errors.New("malformed authorization header"),
		},
		{
			name: "Malformed Authorization Header (No Key Provided)",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			want:        "",
			expectError: errors.New("malformed authorization header"),
		},
		{
			name: "Empty Authorization Header",
			headers: http.Header{
				"Authorization": []string{""},
			},
			want:        "",
			expectError: ErrNoAuthHeaderIncluded,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)

			// Check for expected error
			if (err != nil && tt.expectError == nil) || (err == nil && tt.expectError != nil) {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.expectError)
			} else if err != nil && tt.expectError != nil && err.Error() != tt.expectError.Error() {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.expectError)
			}

			// Check for expected result
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
