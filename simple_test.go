package iamissue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	clt, err := simple()
	assert.Error(t, err)
	assert.Nil(t, clt)
}
