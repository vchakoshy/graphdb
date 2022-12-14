package graph

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFollow_Fof(t *testing.T) {
	f := NewFollow()
	f.Add(1, 2)
	f.Add(2, 3)
	for i := 0; i < 100; i++ {
		f.Add(2, rand.Int63n(100))
	}

	fof := f.Fof(1, 0, 10)
	assert.Equal(t, 10, len(fof))
}

func TestFollow_Follow(t *testing.T) {
	f := NewFollow()
	f.Add(1, 2)
	f.Add(2, 1)
	f.Add(2, 3)

	assert.Equal(t, f.AdjMatrix[1][2], true)
	assert.Equal(t, f.AdjMatrix[2][1], true)
	assert.Equal(t, f.AdjMatrix[2][5], false)
}

func BenchmarkFollowAdd(b *testing.B) {
	f := NewFollow()
	b.ResetTimer()
	var from, to int64
	from = 1
	to = 100
	for i := 1; i < b.N; i++ {
		f.Add(from, to)
		from++
		to++
	}
}

func TestFollow_CountAll(t *testing.T) {
	f := NewFollow()
	f.Add(1, 2)
	f.Add(4, 6)
	f.Add(9, 100)
	assert.Equal(t, f.CountAll(), 3)
}

func TestFollow_Remove(t *testing.T) {
	f := NewFollow()
	f.Add(1, 2)
	f.Add(2, 1)
	f.Add(9, 100)

	f.Remove(1, 2)
	assert.Equal(t, f.AdjMatrix[1][2], true)
	assert.Equal(t, f.AdjMatrix[2][1], true)
	assert.Equal(t, f.AdjMatrix[2][5], false)
	f.Remove(2, 1)
	assert.Equal(t, f.AdjMatrix[1][2], false)
	assert.Equal(t, f.AdjMatrix[2][1], false)
}

func TestFollow_SuggestByUser(t *testing.T) {
	f := NewFollow()
	f.Add(1, 2)

	var i int64
	// users by id [3..50] follows 2
	for i = 3; i <= 50; i++ {
		f.Add(i, 2)
	}

	// users by id [3..50] also follows
	for i = 3; i <= 50; i++ {
		f.Add(i, 8)
	}
	for i = 4; i <= 50; i++ {
		f.Add(i, 9)
	}
	for i = 1; i <= 50; i++ {
		f.Add(i, 1)
	}
	for i = 10; i <= 50; i++ {
		f.Add(i, 12)
	}

	r := f.SuggestByUser(2, 2)
	// sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })

	assert.Equal(t, []int64{1, 8}, r)

	r = f.SuggestByUser(10000, 10)
	assert.Equal(t, []int64{}, r)
}
