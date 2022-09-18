package graph

import (
	"fmt"
	"sync"
)

const AppVersion = "0.0.1"

// Graph
type Graph struct {
	// store user following
	follow       Follow
	lock         sync.RWMutex
	metrics      Metrics
	queryOptions QueryOptions
}

func NewGraph() *Graph {
	return &Graph{
		follow:       NewFollow(),
		metrics:      Metrics{},
		queryOptions: QueryOptionsDefault(),
	}
}

func (g *Graph) SetQueryOptions(opts QueryOptions) *Graph {
	g.queryOptions = opts
	return g
}

func (g *Graph) GetMetrics() Metrics {
	m := g.metrics
	m.FollowCount = g.follow.CountAll()
	return m
}

func (g *Graph) GetFollows(from int64) ([]int64, error) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if e, err := g.follow.List(from); err == nil {
		return e, nil
	}
	return []int64{}, fmt.Errorf("user not found: %d", from)
}

func (g *Graph) GetFriendsOfFriends(from int64) ([]int64, error) {
	f, err := g.GetFollows(from)
	if err != nil {
		return []int64{}, err
	}

	// get all followings of my followings
	var allFollow []int64
	for _, i := range f {
		f2, err := g.GetFollows(i)
		if err != nil {
			return []int64{}, err
		}
		allFollow = append(allFollow, f2...)
	}

	// final result, remove from and remove users already followed by from
	var fr []int64

	for _, f2 := range allFollow {
		if f2 == from {
			continue
		}
		if contains(f, f2) {
			continue
		}
		fr = append(fr, f2)
	}

	return fr, nil
}

func (g *Graph) AddFollow(from, to int64) *Graph {
	if from == to {
		// we should raise error
		return g
	}
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.follow.Exists(from, to) {
		return g
	}
	g.follow.Add(from, to)
	return g
}

func (g *Graph) RemoveFollow(from, to int64) *Graph {
	g.lock.Lock()
	g.follow.Remove(from, to)
	g.lock.Unlock()

	return g
}
