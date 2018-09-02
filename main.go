package main

import "fmt"

func main() {

	x := []int{2, 7, 11, 15}
	y := 9

	fmt.Println(twoSum(x, y))
}

func twoSum(nums []int, target int) []int {

	temp := make(map[int]int)

	result := []int{}

	for k, v := range nums {
		temp[v] = k
	}

	for k, v := range nums {
		complement := target - v
		if val, ok := temp[complement]; ok != false && val != k {
			result = append(result, k, temp[complement])
			break
		}
	}
	return result
}
