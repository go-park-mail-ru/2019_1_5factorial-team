package stats

import (
	"github.com/prometheus/client_golang/prometheus"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
)

var Hits = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "hits",
}, []string{"status", "path"})
