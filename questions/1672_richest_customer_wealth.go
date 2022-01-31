package leetcode

import "fmt"

func maximumWealth(accounts [][]int) int {
	var maxWealth int = 0
	for x := 0; x < len(accounts); x++ {
		var wealth int = 0
		for y := 0; y < len(accounts[x]); y++ {
			wealth += accounts[x][y]
		}
		if wealth > maxWealth {
			maxWealth = wealth
		}
	}
	return maxWealth
}

func Q_1672() {
	accounts := [][]int{
		{1, 2, 3},
		{3, 2, 1},
	}
	res := maximumWealth(accounts)

	fmt.Println(res)
	// fmt.Println(maximumWealth(accounts))
}
