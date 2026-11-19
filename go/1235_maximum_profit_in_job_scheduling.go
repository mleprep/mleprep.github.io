//go:build ignore

// 1235. Maximum Profit in Job Scheduling (Hard)
// https://leetcode.com/problems/maximum-profit-in-job-scheduling/
// Run: go run 1235_maximum_profit_in_job_scheduling.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
)

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(a, b int) bool { return endTime[idx[a]] < endTime[idx[b]] })
	ends := make([]int, n)
	for i, j := range idx {
		ends[i] = endTime[j]
	}
	dp := make([]int, n+1) // dp[i] = best profit using the first i jobs by end time
	for i := 1; i <= n; i++ {
		job := idx[i-1]
		// jobs among the first i-1 that finish by this job's start
		j := sort.Search(i-1, func(k int) bool { return ends[k] > startTime[job] })
		dp[i] = max(dp[i-1], dp[j]+profit[job])
	}
	return dp[n]
}

func main() {
	assertEq(jobScheduling([]int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}), 120)
	assertEq(jobScheduling([]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60}), 150)
	assertEq(jobScheduling([]int{1, 1, 1}, []int{2, 3, 4}, []int{5, 6, 4}), 6)
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
