package skiena

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// FindBestPartitioning finds the maximum sum of the largest of a split of nums into <= k partitions.
// Example: The best partitioning for nums=[1 2 3 4 5], k=3 is [1 2 3][4][5]
func FindBestPartitioning(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([][]int, k)
	for i := 0; i < k; i++ {
		dp[i] = make([]int, len(nums))
	}
	for i := 0; i < k; i++ {
		dp[i][0] = nums[0]
	}
	for j := 1; j < len(nums); j++ {
		dp[0][j] = dp[0][j-1] + nums[j]
	}
	cumsum := make([]int, len(nums))
	cumsum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		cumsum[i] = cumsum[i-1] + nums[i]
	}
	for i := 1; i < k; i++ {
		for j := 1; j < len(nums); j++ {
			// Best split of nums[:j] into (i+1) partitions is Min_{m=1..j}(Max(dp[i-1][j-m]), Sum(nums[j-m:j]))
			best := -1
			for m := 0; m <= j; m++ {
				sum := max(dp[i-1][j-m], cumsum[j]-cumsum[j-m])
				if best == -1 || sum < best {
					best = sum
				}
			}
			dp[i][j] = best
			fmt.Printf("dp@%d:%d is %v\n", i, j, dp)
		}
	}
	return dp[k-1][len(nums)-1]
}
