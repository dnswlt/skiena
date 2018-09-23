package skiena

import "testing"

func TestFindBestPartitioningAscending(t *testing.T) {
	best := FindBestPartitioning([]int{1, 2, 3, 4, 5, 6}, 3)
	if best != 9 {
		t.Error("Baad result", best)
	}
}

func TestFindBestPartitioningOnes(t *testing.T) {
	best := FindBestPartitioning([]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 3)
	if best != 3 {
		t.Error("Baad result", best)
	}
}

func TestFindBestPartitioningDescending(t *testing.T) {
	best := FindBestPartitioning([]int{6, 5, 4, 3, 2, 1}, 3)
	if best != 9 {
		t.Error("Baad result", best)
	}
}
