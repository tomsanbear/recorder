package recorder

import (
	"github.com/caddyserver/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

// PluginName is the name of our plugin
const PluginName string = "recorder"

func init() {
	caddy.RegisterPlugin(PluginName, caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	// Normal Setup
	recorder, err := NewRecorder()
	if err != nil {
		return plugin.Error(PluginName, err)
	}

	// Pass plugin to our context
	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		recorder.Next = next
		return nil
	})

	// Done
	return nil
}
