package recorder

import (
	"net"

	"github.com/miekg/dns"
)

type RecorderWriter struct {
	rw  dns.ResponseWriter
	msg *dns.Msg
}

// NewRecorderWriter automagically creates a recorder instance with the underlying writer backing it for most of it's operations
func NewRecorderWriter(w dns.ResponseWriter) *RecorderWriter {
	return &RecorderWriter{
		rw: w,
	}
}

// WriteMsg both records the message, and facades the underlying writemsg call
func (recw *RecorderWriter) WriteMsg(res *dns.Msg) error {
	recw.msg = res
	return recw.rw.WriteMsg(res)
}

// LocalAddr just facades the underlying LocalAddr call
func (recw *RecorderWriter) LocalAddr() net.Addr {
	return recw.rw.LocalAddr()
}

// RemoteAddr just facades the underlying RemoteAddr call
func (recw *RecorderWriter) RemoteAddr() net.Addr {
	return recw.rw.RemoteAddr()
}

// Write just facades the underlying Write call
func (recw *RecorderWriter) Write(b []byte) (i int, err error) {
	return recw.rw.Write(b)
}

// Close just facades the underlying Close call
func (recw *RecorderWriter) Close() error {
	return recw.rw.Close()
}

// TsigStatus just facades the underlying TsigStatus call
func (recw *RecorderWriter) TsigStatus() error {
	return recw.rw.TsigStatus()
}

// TsigTimersOnly just facades the underlying TsigTimersOnly call
func (recw *RecorderWriter) TsigTimersOnly(b bool) {
	recw.rw.TsigTimersOnly(b)
}

// Hijack just facades the underlying hijack call
func (recw *RecorderWriter) Hijack() {
	recw.rw.Hijack()
}

// Msg returns the message that has already been sent
func (recw *RecorderWriter) Msg() *dns.Msg {
	return recw.msg
}
