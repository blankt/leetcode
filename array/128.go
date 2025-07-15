package array

func longestConsecutive(nums []int) int {
	numsMap := make(map[int]struct{})
	for _, num := range nums {
		numsMap[num] = struct{}{}
	}

	consecutive := 0
	maxConsecutive := 0
	for num := range numsMap {
		if _, ok := numsMap[num-1]; ok {
			continue
		}
		consecutive++
		for i := num + 1; ; i++ {
			if _, ok := numsMap[i]; !ok {
				break
			}
			consecutive++
		}
		if consecutive > maxConsecutive {
			maxConsecutive = consecutive
		}
		consecutive = 0
	}
	return maxConsecutive
}
