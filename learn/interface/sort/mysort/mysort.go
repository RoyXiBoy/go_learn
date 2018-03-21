package mysort

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter) {
	for i := 1 ; i < data.Len() ; i++{
		for j := 0; j < data.Len() - i ; j++ {
			if data.Less(j + 1, j){
				data.Swap(j,j + 1)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	for i := 0 ; i < data.Len() - 1 ; i++ {
		if data.Less(i + 1, i){
			return false
		}
	}
	return true
}

type IntArray []int

func (p IntArray)Len() int{
	return len(p)
}

func (p IntArray)Less(i, j int) bool{
	return p[i] < p[j]
}

func (p IntArray)Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

