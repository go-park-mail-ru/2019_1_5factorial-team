package stats

import (
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)

//var Hits = prometheus.NewCounterVec(prometheus.CounterOpts{
//	Name: "hits",
//}, []string{"status", "path"})
//
//func init() {
//	prometheus.MustRegister(Hits)
//}

type Statistic struct {
	Bad         *prometheus.CounterVec
	Hits        prometheus.Counter
	ActiveRooms prometheus.Gauge
}

var Stats = Statistic{}

func init() {
	Stats.Bad = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hits",
	}, []string{"status", "path"})

	Stats.Hits = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "hits_total",
		Help: "just dummy hits on service",
	})

	Stats.ActiveRooms = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "active_rooms",
		Help: "just dummy counter of room in game",
	})

	prometheus.MustRegister(Stats.Hits, Stats.Bad, Stats.ActiveRooms)
}

func (s *Statistic) AddBadResponse(status int, path string) {
	s.Bad.WithLabelValues(strconv.Itoa(status), path).Inc()
}

func (s *Statistic) AddHit() {
	s.Hits.Inc()
}

func (s *Statistic) AddActiveRoom() {
	s.ActiveRooms.Inc()
}

func (s *Statistic) RemoveActiveRoom() {
	s.ActiveRooms.Dec()
}
