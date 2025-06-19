package greedy

import (
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []int
	}{
		{
			name: "test1",
			s:    "ababcbacadefegdehijhklij",
			want: []int{9, 7, 8},
		},
		{
			name: "test2",
			s:    "eccbbbbdec",
			want: []int{10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionLabels1(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
