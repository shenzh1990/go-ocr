package main

import (
	"fmt"
	"time"
)

func main() {

	start1 := time.Now() // 获取当前时间
	//nums := []int{2,3,1}

	//1988832659
	//[63674,-34608,86424,-52056,-3992,93347,2084,-28546,-75702,-28400] 811768
	//5,-7,9,-6,8  57
	tr := numOfArrays(9, 1, 1)
	fmt.Println(tr)
	elapsed1 := time.Since(start1)
	fmt.Println("该函数执行完成耗时：", elapsed1)

}

/**
1420.生成数组
*/
func numOfArrays(n int, m int, k int) int {
	//mod
	var mod int64 = 1000000007
	if k == 0 {
		return 0
	}
	//设置三维动态规划数组  dp[i][j][w]  表示数组的第 i 位 最大值是 j 比较次数是w的方法数
	var dp [50][100][50]int64
	for i := 1; i <= m; i++ {
		//默认长度1的所有 m 默认比较是1
		dp[1][i][1] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for w := 1; w <= k; w++ {
				for l := 1; l < j; l++ {
					dp[i][j][w] += dp[i-1][l][w-1]
					dp[i][j][w] %= mod
				}
				dp[i][j][w] += dp[i-1][j][w] * (int64(j)) % mod
				dp[i][j][w] %= mod
			}
		}
	}
	var ans int64 = 0
	for i := 1; i <= m; i++ {
		ans += dp[n][i][k]
		ans %= mod
	}
	return (int)(ans)
}

func GetMaxArr(arr []int) int {
	mv := -1
	mi := -1
	sc := 0
	for key, value := range arr {
		if mv < value {
			mv = value
			mi = key
			sc++
		}

	}
	return mi
}
