package graph

import (
	"fmt"
)

type Follow struct {
	Data      map[int64]map[int64]Node
	AdjMatrix map[int64]map[int64]bool
}

func NewFollow() Follow {
	return Follow{
		Data:      make(map[int64]map[int64]Node),
		AdjMatrix: make(map[int64]map[int64]bool),
	}
}

func (f *Follow) CountAll() int {
	return len(f.Data)
}

// GetLeaders returns the list of users who
// followed by the current from user
func (f *Follow) GetLeaders(from int64) []int64 {
	var out []int64
	if _, ok := f.Data[from]; ok {
		// map key to value
		out = make([]int64, len(f.Data[from]))
		i := 0
		for k := range f.Data[from] {
			out[i] = k
			i++
		}
		return out
	}
	return out
}

func (f *Follow) Fof(from int64, skip, limit int) map[int64]Node {
	out := make(map[int64]Node)
	f1 := f.Data[from]
	for k := range f1 {
		for k2, v := range f.Data[k] {
			// from use remove from list
			if k2 == from {
				continue
			}
			// currently followed by user remove from list
			if _, ex := f1[k2]; ex {
				continue
			}
			out[k2] = v
		}
	}

	l := DefaultQueryLimit
	if limit > 0 {
		l = limit
	}

	return f.getLimited(out, l)

}

func (f *Follow) getLimited(d map[int64]Node, limit int) map[int64]Node {
	out := make(map[int64]Node)
	i := 0
	for k, v := range d {
		if i == limit {
			return out
		}
		out[k] = v
		i++
	}
	return out
}

func (f *Follow) List(from int64) ([]int64, error) {
	out := f.GetLeaders(from)
	if len(out) > 0 {
		return out, nil
	}
	return []int64{}, fmt.Errorf("follow not found")
}

func (f *Follow) Exists(from, to int64) bool {
	if _, ex := f.Data[from][to]; ex {
		return true
	}
	return false
}

func (f *Follow) Add(from, to int64) {
	if _, ok := f.Data[from]; !ok {
		f.Data[from] = make(map[int64]Node)
	}

	if _, ok := f.AdjMatrix[from]; !ok {
		f.AdjMatrix[from] = make(map[int64]bool)
	}

	if _, ok := f.AdjMatrix[to]; !ok {
		f.AdjMatrix[to] = make(map[int64]bool)
	}

	f.Data[from][to] = Node{}
	f.AdjMatrix[from][to] = true
	f.AdjMatrix[to][from] = true
}

func (f *Follow) SuggestByUser(user int64) []int64 {
	if _, ok := f.AdjMatrix[user]; !ok {
		return []int64{}
	}
	var followers []int64
	for u := range f.AdjMatrix[user] {
		followers = append(followers, u)
	}

	// followers also follow
	allF := make(map[int64]int64)
	for _, f1 := range followers {
		for _, i := range f.GetLeaders(f1) {
			if i == user {
				continue
			}
			if v, ok := allF[i]; ok {
				allF[i] = v + 1
			} else {
				allF[i] = 1
			}
		}
	}

	var out []int64
	for k := range allF {
		out = append(out, k)
	}

	return out
}

func (f *Follow) Remove(from, to int64) {
	delete(f.Data[from], to)
}
