package recorder

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/caddyserver/caddy"
)

func TestSetup(t *testing.T) {
	err := setup(caddy.NewTestController("dns", ""))
	assert.NoError(t, err)
}
