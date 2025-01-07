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
		// {
		// 	name: "Missing Authorization Header",
		// 	headers:     http.Header{},
		// 	want:        "",
		// 	expectError: ErrNoAuthHeaderIncluded,
		// },
		// {
		// 	name: "Malformed Authorization Header",
		// 	headers: http.Header{
		// 		"Authorization": []string{"Bearer some_token"},
		// 	},
		// 	want:        "",
		// 	expectError: errors.New("malformed authorization header"),
		// },
		// {
		// 	name: "Empty Authorization Header",
		// 	headers: http.Header{
		// 		"Authorization": []string{""},
		// 	},
		// 	want:        "",
		// 	expectError: ErrNoAuthHeaderIncluded,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)

			// Check the expected error
			if !errors.Is(err, tt.expectError) {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.expectError)
			}

			// Check the expected result
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
