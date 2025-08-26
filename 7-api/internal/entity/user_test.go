package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Gustavo Freitas", "gufofreitas@gmail.com", "11234")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "Gustavo Freitas", user.Name)
	assert.Equal(t, "gufofreitas@gmail.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Gustavo Freitas", "gufofreitas@gmail.com", "11234")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("11234"))
	assert.False(t, user.ValidatePassword("12345"))
	assert.NotEqual(t, "11234", user.Password)
}
