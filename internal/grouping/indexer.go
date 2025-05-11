package grouping

// TODO: Recomment this
// This is a simple indexer that will index a string to a uint32.
// Basically like the Set object
type indexer struct {
	Indexes map[string]int
}

func newIndexer() *indexer {
	return &indexer{Indexes: make(map[string]int)}
}

func (i *indexer) AddIndexFor(val string) {
	if _, exists := i.Indexes[val]; !exists {
		i.Indexes[val] = len(i.Indexes)
	}
}
