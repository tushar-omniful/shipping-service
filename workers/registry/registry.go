package registry

import (
	"context"
	"github.com/omniful/go_commons/util"
	"github.com/omniful/go_commons/worker/listener"
)

const (
	AllGroups = "*"
)

type InitListener func(ctx context.Context) listener.ListenerServer

type Registry map[string][]InitListener

var _registry = make(Registry)

func Instance() Registry {
	return _registry
}

func RegisterListener(group string, il InitListener) {
	_registry[group] = append(_registry[group], il)
}

func (r Registry) GetWorkersFromGroups(ctx context.Context, gs ...string) (res []listener.ListenerServer) {
	if util.Contains(gs, AllGroups) {
		return r.GetAllWorkers(ctx)
	}

	for _, g := range gs {
		if lis, ok := r[g]; ok {
			for _, li := range lis {
				res = append(res, li(ctx))
			}
		}
	}

	return res
}

func (r Registry) GetAllWorkers(ctx context.Context) (res []listener.ListenerServer) {
	for _, lis := range r {
		for _, li := range lis {
			res = append(res, li(ctx))
		}
	}

	return res
}
