package sorting_tests

import (
	"reflect"
	"testing"

	"github.com/SuvorovSergey/go-algorithms/sorting"
)

func TestSelectionSorting(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "test1",
			input: []int{6, 8, 2, 9, 10, 1, 3, 7, 5, 0, 4},
			want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:  "test2",
			input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:  "test3",
			input: []int{5},
			want:  []int{5},
		},
		{
			name:  "test4",
			input: []int{5, 5, 5, 5, 5, 5, 5, 5},
			want:  []int{5, 5, 5, 5, 5, 5, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorting.SelectionSort(tt.input)
			if !reflect.DeepEqual(tt.want, tt.input) {
				t.Errorf("selectionSort() = %v, want %v", tt.input, tt.want)
			}
		})
	}
}
