package auth

import (
	"errors"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		headers http.Header
		want    string
		err     error
	}{
		"simple":      {headers: http.Header{"Authorization": {"ApiKey key"}}, want: "key", err: nil},
		"errorNoAuth": {headers: http.Header{"Auth": {"Key key"}}, want: "", err: ErrNoAuthHeaderIncluded},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.headers)
			diff := cmp.Diff(tc.want, got)
			errIs := errors.Is(tc.err, err)
			if diff != "" {
				t.Fatalf(diff)
			}
			if !errIs {
				t.Error(err)
			}
		})
	}
}
