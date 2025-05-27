package array

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	trackbackCandidates(0, 0, target, []int{}, candidates, &res)
	return res
}

func trackbackCandidates(start, trace, target int, temp, candidates []int, res *[][]int) {
	if target == trace {
		*res = append(*res, append([]int{}, temp...))
		return
	}

	if trace > target {
		return
	}

	for i := start; i < len(candidates); i++ {
		v := candidates[i]
		trace += v
		temp = append(temp, v)

		trackbackCandidates(i, trace, target, temp, candidates, res)

		trace -= v
		temp = temp[:len(temp)-1]
	}
}
