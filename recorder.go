package recorder

import (
	"context"

	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"
)

// Recorder is the 'logic' for enabling the recording of our message
type Recorder struct {
	Next plugin.Handler
}

// NewRecorder instantiates the recorder object, since we don't have a config this is easy
func NewRecorder() (*Recorder, error) {
	return &Recorder{}, nil
}

// Name is the name of our plugin
func (rec *Recorder) Name() string { return PluginName }

// ServeDNS just creates a new recorer writer and then chains that down to the next plugin
func (rec *Recorder) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (rc int, err error) {
	// All we are doing is replacing our writer with a new recorder. Defining in this plugin as a copy of the one that is found in the dns test utils
	rrw := NewRecorderWriter(w)
	return plugin.NextOrFailure(rec.Name(), rec.Next, ctx, rrw, r)
}
