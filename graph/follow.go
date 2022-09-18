package graph

import "fmt"

type Follow struct {
	Data map[int64]map[int64]Node
}

func NewFollow() Follow {
	return Follow{
		Data: make(map[int64]map[int64]Node),
	}
}

func (f *Follow) CountAll() int {
	return len(f.Data)
}

// GetLeaders returns the list of users who
// followed by the current from user
func (f *Follow) GetLeaders(from int64) []int64 {
	var leaders []int64
	if _, ok := f.Data[from]; ok {
		// map key to value
		leaders = make([]int64, len(f.Data[from]))
		i := 0
		for k := range f.Data[from] {
			leaders[i] = k
			i++
		}
		return leaders
	}
	return leaders
}

func (f *Follow) List(from int64) ([]int64, error) {
	leaders := f.GetLeaders(from)
	if len(leaders) > 0 {
		return leaders, nil
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

	f.Data[from][to] = Node{}
}

func (f *Follow) Remove(from, to int64) {
	delete(f.Data[from], to)
}
