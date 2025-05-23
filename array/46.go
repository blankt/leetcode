package array

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	track := make([]int, 0, len(nums))

	backTrack(nums, track, &res, used)
	return res
}

func backTrack(nums []int, track []int, res *[][]int, used []bool) {
	if len(nums) == len(track) {
		*res = append(*res, append([]int{}, track...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		track = append(track, nums[i])
		used[i] = true
		backTrack(nums, track, res, used)

		used[i] = false
		track = track[:len(track)-1]
	}
}
