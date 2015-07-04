// Package queries provides a nice wrapper to keep SQL *mostly* out of Go
package queries

var queries = make(map[string]string)

func init() {
	addUserQueries()
	addBoardQueries()
	addPostQueries()
}

// Get returns the named query
func Get(queryName string) string {
	return queries[queryName]
}

// All returns all queries currently loaded
func All() (q []string) {
	for _, name := range queries {
		q = append(q, name)
	}

	return q
}
