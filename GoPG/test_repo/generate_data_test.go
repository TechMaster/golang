package test_repo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generate_one_user(t *testing.T) {
	err := insert_user()
	assert.Nil(t, err)
}

func Test_generate_20_user(t *testing.T) {
	for i := 0; i < 20; i++ {
		err := insert_user()
		assert.Nil(t, err)
	}
}
