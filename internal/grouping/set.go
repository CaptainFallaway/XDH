package grouping

type set struct {
	vals map[string]struct{}
}

func newSet() *set {
	return &set{vals: make(map[string]struct{})}
}

func (s *set) Add(val string) {
	s.vals[val] = struct{}{}
}

func (s *set) ToSlice() []string {
	ret := make([]string, 0, len(s.vals))

	for k := range s.vals {
		ret = append(ret, k)
	}

	return ret
}
