package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserGetHandler(t *testing.T) {

	under_test := UserServiceImpl{}

	assert.Equal(t, under_test.GetUser().Age, int32(40))
	assert.Equal(t, under_test.GetUser().Name, "Marco")
}
