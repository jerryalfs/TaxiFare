package faredata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	instance := New()
	assert.NotNil(t, instance)
}
