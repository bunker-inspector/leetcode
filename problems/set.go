package problems

type IntSet map[int]bool

func (i *IntSet) Contains(v int) bool {
	if _, ok := (*i)[v]; ok {
		return true
	}
	return false
}

type IntArray4Set map[[4]int]bool

func NewIntArray4Set(as ...[4]int) *IntArray4Set {
	i := IntArray4Set{}
	i.Add(as...)
	return &i
}

func (i *IntArray4Set) Contains(as ...[4]int) bool {
	for _, a := range as {
		if _, ok := (*i)[a]; !ok {
			return false
		}
	}
	return true
}

func (i *IntArray4Set) Add(as ...[4]int) *IntArray4Set {
	for _, a := range as {
		(*i)[a] = true
	}
	return i
}
