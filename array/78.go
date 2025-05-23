package array

// [1,2,3] dfs +回溯
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	track, index := make([]int, 0, len(nums)), make([]int, 0, len(nums))
	used := make([]bool, len(nums))
	singleSubsets(nums, track, index, &res, used)
	res = append(res, []int{})
	return res
}

func singleSubsets(nums []int, index, track []int, res *[][]int, used []bool) {
	if len(track) > 0 {
		*res = append(*res, append([]int{}, track...))
	}

	for i, v := range nums {
		if used[i] {
			continue
		}
		if len(index) >= 1 && index[len(index)-1] > i {
			continue
		}

		used[i] = true
		index = append(index, i)
		track = append(track, v)

		singleSubsets(nums, index, track, res, used)

		used[i] = false
		index = index[:len(index)-1]
		track = track[:len(track)-1]
	}
}

func subsets2(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0, len(nums))
	subsetsBackTrack(nums, track, &res, 0)
	return res
}

func subsetsBackTrack(nums []int, track []int, res *[][]int, start int) {
	*res = append(*res, append([]int{}, track...))

	for i := start; i < len(nums); i++ {

		track = append(track, nums[i])
		subsetsBackTrack(nums, track, res, i+1)

		track = track[:len(track)-1]
	}
}
