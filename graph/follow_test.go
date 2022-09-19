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
