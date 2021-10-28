package Analysis

type Set map[Variable]bool

func Union(a Set, b Set) Set {
	union := make(Set)
	for k := range a {
		union[k] = true
	}
	for k := range b {
		union[k] = true
	}
	return union
}

func Intersection(a Set, b Set) Set {
	intersection := make(Set)
	for k := range a {
		if b[k] {
			intersection[k] = true
		}
	}
	return intersection
}

func Subtract(a Set, b Set) Set {
	result := make(Set)
	for k, v := range a {
		result[k] = v
	}
	for k := range a {
		if b[k] {
			delete(result, k)
		}
	}
	return result
}

func Equals(a Set, b Set) bool {
	for k := range a {
		if !b[k] {
			return false
		}
	}
	for k := range b {
		if !a[k] {
			return false
		}
	}
	return true
}
