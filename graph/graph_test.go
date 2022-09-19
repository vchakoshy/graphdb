package graph

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

const _testFileName = "data/graphdb_test.db"

func TestGraph_GetFollows(t *testing.T) {
	g := NewGraph()
	g.AddFollow(1, 5)
	g.AddFollow(1, 2)
	g.AddFollow(1, 3)
	g.AddFollow(1, 4)
	g.AddFollow(1, 4)
	g.AddFollow(1, 1)

	// follow of this user not exists
	_, err := g.GetFollows(100)
	assert.NotEqual(t, err, nil)

	r, err := g.GetFollows(1)
	if err != nil {
		t.Errorf("Graph.GetFollows() = %s", err.Error())
	}

	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })

	assert.EqualValues(t, r, []int64{2, 3, 4, 5})
}

func TestGraph_GetFriendsOfFriends(t *testing.T) {
	g := NewGraph()
	g.AddFollow(1, 2)
	g.AddFollow(2, 3)
	g.AddFollow(2, 5)
	g.AddFollow(3, 1)
	g.AddFollow(5, 6)
	r, err := g.GetFriendsOfFriends(1)
	if err != nil {
		t.Errorf("Graph.GetFriendsOfFriends() = %s", err.Error())
	}
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	assert.Equal(t, r, []int64{3, 5})

	// follow of this user not exists
	_, err = g.GetFriendsOfFriends(100)
	assert.NotEqual(t, err, nil)
}

func TestGraph_RemoveFollow(t *testing.T) {
	g := NewGraph()
	g.AddFollow(1, 2)
	g.AddFollow(1, 3)
	g.AddFollow(1, 4)
	g.AddFollow(1, 8)
	g.AddFollow(1, 5)
	g.AddFollow(1, 6)
	g.RemoveFollow(1, 3)
	r, err := g.GetFollows(1)
	if err != nil {
		t.Errorf("Graph.RemoveFollow() = %s", err.Error())
	}
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	assert.Equal(t, r, []int64{2, 4, 5, 6, 8})

}

// func TestGraph_save(t *testing.T) {
// 	g := NewGraph()
// 	g.AddFollow(1, 2)
// 	g.AddFollow(2, 3)
// 	err := g.save(_testFileName)
// 	assert.Equal(t, err, nil)
// }

// func TestGraph_load(t *testing.T) {
// 	TestGraph_save(t)
// 	g := NewGraph()
// 	err := g.Load(_testFileName)
// 	assert.Equal(t, err, nil)

// 	r := g.follow.Exists(1, 2)
// 	assert.Equal(t, r, true)

// 	r = g.follow.Exists(2, 3)
// 	assert.Equal(t, r, true)

// 	r = g.follow.Exists(2, 4)
// 	assert.Equal(t, r, false)

// }
