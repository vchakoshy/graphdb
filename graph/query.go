package graph

const DefaultQueryLimit = 10

type QueryOptions struct {
	Limit int
	Skip  int
}

func QueryOptionsDefault() QueryOptions {
	return QueryOptions{Limit: DefaultQueryLimit}
}
