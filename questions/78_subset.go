package leetcode

import "fmt"

func subsetGenerator(nums [][]int, original []int, pos int) [][]int {
	res:= [][]int{}

	if(pos == len(original)){
		temp:=[]int{}
		for i:=0; i<len(nums); i++ {
				temp = append(temp, nums[i][0])			
		}
		return append(res, temp)
	}

	ans1:= subsetGenerator(append(nums, []int{original[pos]}), original, pos+1)
	ans2:= subsetGenerator(nums, original, pos+1)
	
	res = append(res, ans2...)
	res = append(res, ans1...)
	
	return res
}

func subsets(nums []int) [][]int {
	res := subsetGenerator([][]int{}, nums,0)

	return res
}

func Q_78(){

	fmt.Println(subsets([]int{1,2,3}))
}