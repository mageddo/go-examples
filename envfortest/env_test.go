package envfortest

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetMode(t *testing.T){
	assert.Equal(t, "TEST", GetMode())
}
