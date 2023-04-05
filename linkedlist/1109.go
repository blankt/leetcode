package leetcode

func CorpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n)

	for _, v := range bookings {
		diff[v[0]-1] += v[2]
		if v[1] < n {
			diff[v[1]] -= v[2]
		}
	}

	for i := 1; i < n; i++ {
		diff[i] = diff[i] + diff[i-1]
	}
	return diff
}
