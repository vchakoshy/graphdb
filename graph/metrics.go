package graph

import (
	"sync/atomic"
)

type Metrics struct {
	FollowCount int   `json:"follow_count"`
	WriteCount  int64 `json:"write_count"`
	ReadCount   int64 `json:"read_count"`
}

func (m *Metrics) IncWriteCount() *Metrics {
	atomic.AddInt64(&m.WriteCount, 1)
	return m
}
func (m *Metrics) IncReadCount() *Metrics {
	atomic.AddInt64(&m.ReadCount, 1)
	return m
}
