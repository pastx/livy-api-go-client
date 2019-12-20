package livy

import (
	"math"
	"regexp"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
)


func TestClient_GetSessionLogs(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock Rest endpoint
	// Endpoint specified https://livy.incubator.apache.org/docs/latest/rest-api.html
	httpmock.RegisterRegexpResponder("GET", regexp.MustCompile(`^/sessions/(\d+)/log\?from=(\d+)&size=(\d+)\z`),
		httpmock.NewStringResponder(200, `{}`))

	mockClient := NewClient("localhost", 2 * time.Second)

	_, err := mockClient.GetSessionLogs(0, 0, 100)
	if err != nil {
		t.Errorf("unexpecred error: %s", err)
	}

	_, err = mockClient.GetSessionLogs(5, 0, math.MaxInt64)
	if err != nil {
		t.Errorf("unexpecred error: %s", err)
	}

	_, err = mockClient.GetSessionLogs(10, 500, 0)
	if err != nil {
		t.Errorf("unexpecred error: %s", err)
	}

	// Get total amount of registered calls
	count := httpmock.GetTotalCallCount()
	if count != 3 {
		t.Errorf("got %d api calls, expected 3", count)
	}
}
