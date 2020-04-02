import (
	"sort"
	"qmabt/pb"
)

type Uint64Sorter []uint64

func (u Uint64Sorter) Len() int {
	return len(u)
}
func (u Uint64Sorter) Less(i, j int) bool {
	return u[i] < u[j]
}
func (u Uint64Sorter) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

sort.Sort(Uint64Sorter(orderedStrategy))
