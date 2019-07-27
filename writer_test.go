package recorder

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/coredns/coredns/plugin/test"
	"github.com/miekg/dns"
)

func TestRecorderWriter(t *testing.T) {
	sut := NewRecorderWriter(&test.ResponseWriter{})
	testMsg := new(dns.Msg)
	sut.WriteMsg(testMsg)
	assert.Equal(t, sut.msg, testMsg)
	assert.Equal(t, sut.LocalAddr().String(), "127.0.0.1:53")
	assert.Equal(t, sut.RemoteAddr().String(), "10.240.0.1:40212")
}
