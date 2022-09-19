package graph

import (
	"math/rand"
	"sort"
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

func TestFollow_SuggestByUser(t *testing.T) {
	f := NewFollow()
	f.Add(1, 2)
	f.Add(2, 3)
	f.Add(3, 2)
	f.Add(4, 2)
	f.Add(5, 2)
	f.Add(3, 6)
	f.Add(4, 6)
	f.Add(5, 8)

	r := f.SuggestByUser(2)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })

	assert.Equal(t, []int64{6, 8}, r)
}
