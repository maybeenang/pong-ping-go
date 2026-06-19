package response_test

import (
	"net/http"
	"testing"

	"github.com/maybeenang/pong-ping-v2/pkg/response"
)

func TestJSON(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		w      http.ResponseWriter
		status int
		data   any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response.JSON(tt.w, tt.status, tt.data)
		})
	}
}
