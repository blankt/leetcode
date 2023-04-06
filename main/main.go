package main

import (
	"fmt"
	"leetcode/array"
)

// [[1,4,5],[1,3,4],[2,6]]
func main() {

	//x1 := &leetcode.ListNode{Val: 5, Next: nil}
	//x2 := &leetcode.ListNode{Val: 4, Next: x1}
	//x3 := &leetcode.ListNode{Val: 1, Next: x2}

	//y1 := &leetcode.ListNode{Val: 4, Next: nil}
	//y2 := &leetcode.ListNode{Val: 3, Next: y1}
	//y3 := &leetcode.ListNode{Val: 1, Next: y2}
	//
	//z1 := &leetcode.ListNode{Val: 6, Next: nil}
	//z2 := &leetcode.ListNode{Val: 2, Next: z1}
	//
	//lists := []*leetcode.ListNode{x3, y3, z2}
	//
	//result := leetcode.MergeKLists(lists)
	//
	//for result != nil {
	//	fmt.Println(result.Val)
	//	result = result.Next
	//}

	//result := leetcode.Reverse(x3)
	//
	//for result != nil {
	//	fmt.Println(result.Val)
	//	result = result.Next
	//}

	//s1 := "hello"
	//s2 := "ooolleoooleh"
	//result := leetcode.CheckInclusion(s1, s2)
	//fmt.Println(result)

	//s := "aab"
	//fmt.Println(leetcode.LengthOfLongestSubstring(s))

	//var nums = [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
	//result := leetcode.CorpFlightBookings(nums, 5)
	//fmt.Println(result)

	//var nums = [][]int{{2, 1, 5}, {3, 5, 7}}
	//result := leetcode.CarPooling(nums, 3)
	//fmt.Println(result)

	//var matrix = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//fmt.Println(leetcode.Rotate(matrix))

	//var matrix = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//fmt.Println(leetcode.SpiralOrder(matrix))
	var s = []int{2, 3, 1, 2, 4, 3}
	fmt.Println(array.MinSubArrayLen(7, s))
}
