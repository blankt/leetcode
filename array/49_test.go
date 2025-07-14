package array

import "testing"

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs     []string
		expected [][]string
	}{
		{
			strs:     []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			strs:     []string{"bdddddddddd", "bbbbbbbbbbc"},
			expected: [][]string{{"bbbbbbbbbbc"}, {"bdddddddddd"}},
		},
	}

	for _, test := range tests {
		result := groupAnagrams2(test.strs)
		if len(result) != len(test.expected) {
			t.Errorf("expected %d groups, got %d", len(test.expected), len(result))
			continue
		}

		for _, expectedGroup := range test.expected {
			found := false
			for _, resultGroup := range result {
				if len(resultGroup) == len(expectedGroup) && containsAll(resultGroup, expectedGroup) {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("expected group %v not found in result", expectedGroup)
			}
		}
	}

}

func containsAll(strs1, strs2 []string) bool {
	for _, str1 := range strs1 {
		found := false
		for _, str2 := range strs2 {
			if str1 == str2 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
