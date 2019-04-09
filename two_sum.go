package algorithms

// Give an array of integers, return indices of the two numbers
// such that they add up to a specified target

func twoSum(nums []int, target int) []int {
	mapData := make(map[int]int)
	for i := range nums {
		complement := target - nums[i]
		if val, found := mapData[complement]; !found {
			mapData[nums[i]] = i
		} else {
			return []int{val, i}
		}
	}
	return []int{}
}