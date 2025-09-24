package builder

import (
	"github.com/ffhenkes/vulcand/proxy"
	"github.com/ffhenkes/vulcand/proxy/mux"
	"github.com/ffhenkes/vulcand/stapler"
)

// NewProxy returns a new Proxy instance.
func NewProxy(id int, st stapler.Stapler, o proxy.Options) (proxy.Proxy, error) {
	return mux.New(id, st, o)
}
