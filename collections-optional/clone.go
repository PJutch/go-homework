package collections

func CloneSliceNaive(to_clone []int) []int {
	cloned := make([]int, 0)
	for _, elem := range to_clone {
		cloned = append(cloned, elem)
	}
	return cloned
}

func CloneSliceReserve(to_clone []int) []int {
	cloned := make([]int, len(to_clone))
	for i, elem := range to_clone {
		cloned[i] = elem
	}
	return cloned
}

func CloneMapNaive(to_clone map[int]int) map[int]int {
	cloned := make(map[int]int)
	for key, value := range to_clone {
		cloned[key] = value
	}
	return cloned
}

func CloneMapReserve(to_clone map[int]int) map[int]int {
	cloned := make(map[int]int, len(to_clone))
	for key, value := range to_clone {
		cloned[key] = value
	}
	return cloned
}
