package graph

import "fmt"

type Follow struct {
	data map[int64]map[int64]Node
}

func NewFollow() Follow {
	return Follow{
		data: make(map[int64]map[int64]Node),
	}
}

func (f *Follow) CountAll() int {
	return len(f.data)
}

// getLeaders returns the list of users who
// followed by the current from user
func (f *Follow) getLeaders(from int64) []int64 {
	var leaders []int64
	if _, ok := f.data[from]; ok {
		leaders = make([]int64, len(f.data[from]))
		i := 0
		for k := range f.data[from] {
			leaders[i] = k
			i++
		}
		return leaders
	}
	return leaders
}

func (f *Follow) List(from int64) ([]int64, error) {
	leaders := f.getLeaders(from)
	if len(leaders) > 0 {
		return leaders, nil
	}
	return []int64{}, fmt.Errorf("follow not found")
}

func (f *Follow) Exists(from, to int64) bool {
	if _, ex := f.data[from][to]; ex {
		return true
	}
	return false
}

func (f *Follow) Add(from, to int64) {
	if _, ok := f.data[from]; !ok {
		f.data[from] = make(map[int64]Node)
	}

	f.data[from][to] = Node{}
}

func (f *Follow) Remove(from, to int64) {
	delete(f.data[from], to)
}
