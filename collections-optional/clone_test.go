package collections

import "testing"

var slice_to_clone = func() []int {
	slice_to_clone := make([]int, 10000)
	for i := range slice_to_clone {
		slice_to_clone[i] = i + 1
	}
	return slice_to_clone
}()

func BenchmarkCloneSliceNaive(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		CloneSliceNaive(slice_to_clone)
	}
}

func BenchmarkCloneSliceReserve(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		CloneSliceReserve(slice_to_clone)
	}
}

var map_to_clone = func() map[int]int {
	map_size := 10000
	map_to_clone := make(map[int]int, map_size)
	for i := 0; i < map_size; i += 1 {
		map_to_clone[i] = i + 1
	}
	return map_to_clone
}()

func BenchmarkCloneMapNaive(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		CloneMapNaive(map_to_clone)
	}
}

func BenchmarkCloneMapReserve(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		CloneMapReserve(map_to_clone)
	}
}
