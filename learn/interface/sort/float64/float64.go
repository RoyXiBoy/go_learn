package float64

type Float64Array []float64

func (p Float64Array)Len() int{
	return len(p)
}

func (p Float64Array)Less(i, j int) bool{
	return p[i] < p[j]
}

func (p Float64Array)Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Float64Array)NewFloat64Array()  [25]float64{
	var float64Array [25]float64
	return  float64Array
}

