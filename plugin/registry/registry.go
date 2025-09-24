// This file will be generated to include all customer specific middlewares
package registry

import (
	"github.com/ffhenkes/vulcand/plugin"
	"github.com/ffhenkes/vulcand/plugin/cbreaker"
	"github.com/ffhenkes/vulcand/plugin/connlimit"
	"github.com/ffhenkes/vulcand/plugin/ratelimit"
	"github.com/ffhenkes/vulcand/plugin/rewrite"
	"github.com/ffhenkes/vulcand/plugin/trace"
)

func GetRegistry() *plugin.Registry {
	r := plugin.NewRegistry()

	specs := []*plugin.MiddlewareSpec{
		ratelimit.GetSpec(),
		connlimit.GetSpec(),
		rewrite.GetSpec(),
		cbreaker.GetSpec(),
		trace.GetSpec(),
	}

	for _, spec := range specs {
		if err := r.AddSpec(spec); err != nil {
			panic(err)
		}
	}

	return r
}
