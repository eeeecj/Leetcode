package main

import "fmt"

func main() {
	fmt.Println(beibaofull([]int{1, 3, 4}, []int{15, 20, 35}, 4))
}

//二维数组
func beibao01(weight, value []int, bagweight int) int {
	m, n := len(weight), bagweight
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := weight[0]; i <= bagweight; i++ {
		dp[0][i] = value[0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j <= bagweight; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= weight[i] {
				if dp[i-1][j] < dp[i-1][j-weight[i]]+value[i] {
					dp[i][j] = dp[i-1][j-weight[i]] + value[i]
				}
			}
		}
	}
	return dp[m-1][n]
}

//一维数组
func beibao01_2(weight, value []int, bagweight int) int {
	dp := make([]int, bagweight+1)
	for i := 0; i < len(weight); i++ {
		for j := bagweight; j >= weight[i]; j-- {
			if dp[j] < dp[j-weight[i]]+value[i] {
				dp[j] = dp[j-weight[i]] + value[i]
			}
		}
	}
	return dp[bagweight]
}

//完全背包问题，物品数量无限
func beibaofull(weight, value []int, bagweight int) int {
	dp := make([]int, bagweight+1)
	for i := 0; i < len(weight); i++ {
		//注意为前序遍历，且内外层循环可交换顺序
		for j := weight[i]; j <= bagweight; j++ {
			if dp[j] < dp[j-weight[i]]+value[i] {
				dp[j] = dp[j-weight[i]] + value[i]
			}
		}
	}
	fmt.Println(dp)
	return dp[bagweight]
}
