package graph

import (
	"fmt"
	"sync"
)

const AppVersion = "0.0.1"

// Graph
type Graph struct {
	// store user following
	follow map[int64][]int64
	lock   sync.RWMutex
}

func NewGraph() *Graph {
	return &Graph{
		follow: make(map[int64][]int64),
	}
}

func (g *Graph) GetFollows(from int64) ([]int64, error) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if v, ok := g.follow[from]; ok {
		return v, nil
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
	if contains(g.follow[from], to) {
		return g
	}
	g.follow[from] = append(g.follow[from], to)
	return g
}

func (g *Graph) RemoveFollow(from, to int64) *Graph {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.follow[from] = removeFromSlice(g.follow[from], to)

	return g
}
