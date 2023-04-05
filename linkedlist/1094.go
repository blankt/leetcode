package leetcode

func CarPooling(trips [][]int, capacity int) bool {
	diff := make([]int, 1001)

	for _, v := range trips {
		diff[v[1]] += v[0]
		if v[2] < 1000 {
			diff[v[2]] -= v[0]
		}
	}

	for i := 1; i < len(diff); i++ {
		diff[i] = diff[i] + diff[i-1]
	}
	for i := 1; i < len(diff); i++ {
		if diff[i] > capacity {
			return false
		}
	}
	return true
}
