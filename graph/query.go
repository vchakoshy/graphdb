package graph

const _defaultQueryLimit = 10

type QueryOptions struct {
	Limit int
	Skip  int
}

func QueryOptionsDefault() QueryOptions {
	return QueryOptions{Limit: _defaultQueryLimit}
}
