package dummyquerier

import (
	"context"

	"github.com/els0r/goProbe/v4/cmd/global-query/pkg/distributed"
	"github.com/els0r/goProbe/v4/cmd/global-query/pkg/hosts"
	"github.com/els0r/goProbe/v4/pkg/query"
	"github.com/els0r/goProbe/v4/pkg/results"
	"github.com/els0r/goProbe/v4/plugins"
)

const (
	// Name is the name of the Dummy querier
	Name = "dummy"
)

func init() {
	plugins.RegisterQuerier(Name, func(_ context.Context, cfgPath string) (distributed.Querier, error) {
		return New(cfgPath)
	})
}

type DummyQuerier struct{}

func New(_ string) (*DummyQuerier, error) {
	return &DummyQuerier{}, nil
}

func (d *DummyQuerier) Query(ctx context.Context, hosts hosts.Hosts, args *query.Args) (<-chan *results.Result, <-chan struct{}) {
	out := make(chan *results.Result, 1)
	go func() {
		out <- results.New()
	}()
	return out, nil
}
