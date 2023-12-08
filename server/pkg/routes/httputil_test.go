package routes

import (
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

// MockReadCloser is a mock implementation of io.ReadCloser for testing
type MockReadCloser struct {
	io.Reader
}

func (mrc *MockReadCloser) Close() error {
	return nil
}

// TestReadBody tests the ReadBody function
func TestReadBody(t *testing.T) {
	type TestData struct {
		Name string
		Age  int
	}

	// Create a mock JSON string
	mockJSON := `{"Name": "name <script>alert('XSS')</script>", "Age": 25}`

	// Create a mock ReadCloser using the JSON string
	mockBody := &MockReadCloser{strings.NewReader(mockJSON)}

	// Create a mock output structure
	var output TestData

	// Call ReadBody with the mock body and output
	err := ReadBody(mockBody, &output)
	assert.NoError(t, err, "ReadBody should not return an error")

	// Verify that sanitization has been applied to the output.Name field
	expectedName := "name "
	assert.Equal(t, expectedName, output.Name, "Sanitization failed for Name field")

	// Verify that other fields remain unchanged
	expectedAge := 25
	assert.Equal(t, expectedAge, output.Age, "Sanitization should not have affected Age field")
}
