package workers

import (
	"context"
	"github.com/omniful/go_commons/log"
	"github.com/omniful/go_commons/shutdown"
	"github.com/omniful/go_commons/worker"
	"github.com/omniful/shipping-service/workers/registry"
	"strings"
)

func RunWorkers(ctx context.Context, gs string) {
	log.Info(gs)
	groups := ensureUnique(strings.Split(gs, ","))
	startListening(ctx, groups...)
}

func startListening(ctx context.Context, groups ...string) {
	listeners := worker.NewServer(registry.Instance().GetWorkersFromGroups(ctx, groups...))
	listeners.Run(ctx)

	shutdown.RegisterShutdownCallback("shipping service worker shutdown callback", listeners)
	<-shutdown.GetWaitChannel()
}

func ensureUnique(arr []string) (res []string) {
	m := make(map[string]bool)
	for _, e := range arr {
		m[e] = true
	}
	for k := range m {
		res = append(res, k)
	}
	return res
}
