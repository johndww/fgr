package routes

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/stretchr/testify/assert"
	"html"
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
		Name        string
		Description string
	}

	// Create a mock JSON string
	mockJSON := `{
    "name": "name ",
    "description": "i'm testing"
}`

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
	expectedAge := "i'm testing"
	assert.Equal(t, expectedAge, output.Description, "Sanitization should not have affected Description field")
}

func TestSanitizePlainText(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"I'm testing the bluemonday library. <script>alert('XSS');</script>", "I'm testing the bluemonday library. "},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := sanitizePlainText(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func sanitizePlainText(input string) string {
	p := bluemonday.StrictPolicy()

	// Sanitize the input
	return html.UnescapeString(p.Sanitize(input))
}
