package problem0001

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if c, ok := m[target-v]; ok {
			return []int{c, i}
		}
		if _, ok := m[v]; !ok {
			m[v] = i
		}
	}
	return []int{}
}
