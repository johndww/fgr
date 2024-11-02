package routes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSanitizeFields_NoChange(t *testing.T) {
	name := "name"
	desc := "description"

	input := GiftRequestInput{
		Name:        name,
		Description: &desc,
	}

	SanitizeFields(&input)

	assert.Equal(t, name, input.Name)
	assert.Equal(t, desc, *input.Description)
}

func TestSanitizeFields_NormalDescription(t *testing.T) {
	name := "name"
	desc := "I'm not bringing a carryon to Ohio, so if I got it there, it would be easy to bring it back."

	input := GiftRequestInput{
		Name:        name,
		Description: &desc,
	}

	SanitizeFields(&input)

	assert.Equal(t, name, input.Name)
	assert.Equal(t, desc, *input.Description)
}

func TestSanitizeFields_XSS(t *testing.T) {
	name := "name"
	desc := "description <script>alert('XSS')</script> two"
	expectedDesc := "description  two"

	input := GiftRequestInput{
		Name:        name,
		Description: &desc,
	}

	SanitizeFields(&input)

	assert.Equal(t, name, input.Name)
	assert.Equal(t, expectedDesc, *input.Description)
}
