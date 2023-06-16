package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var deleteTests = []struct {
	ss       []int
	idx      []int
	expected []int
}{
	// idx out of bounds
	{
		[]int{1, 2},
		[]int{-1},
		[]int{1, 2},
	},
	{
		[]int{1, 2},
		[]int{2},
		[]int{1, 2},
	},
	// remove from empty slice
	{
		[]int{},
		[]int{0},
		[]int{},
	},
	{
		[]int{1},
		[]int{0},
		[]int{},
	},
	{
		[]int{1, 2, 3, 4, 5},
		[]int{2},
		[]int{1, 2, 4, 5},
	},
	{
		[]int{1, 2, 3, 4, 5},
		[]int{1, 3},
		[]int{1, 3, 5},
	},
	// mixed indices
	{
		[]int{1, 2, 3, 4, 5},
		[]int{1, -1, 5, 3},
		[]int{1, 3, 5},
	},
}

func TestDelete(t *testing.T) {
	for _, test := range deleteTests {

		t.Run("", func(t *testing.T) {
			// could be inplace! changes input slice
			ss := make([]int, len(test.ss))

			copy(ss, test.ss)
			assert.Equal(t, test.expected, pie.Delete(ss, test.idx...))

			copy(ss, test.ss)
			assert.Equal(t, test.expected, pie.DeleteSet(ss, test.idx...))
		})
	}
}

var big_arr []int

func init() {
	for i := 0; i < 1_000_000; i++ {
		big_arr = append(big_arr, i)
	}
}

func BenchmarkDelete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pie.Delete(big_arr,
			700_000,
			500_000,
			200_000,
			100_000,
			50_000,
			10_000,
			5000,
			1000,
		)
	}
}

func BenchmarkDeleteSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pie.DeleteSet(big_arr,
			700_000,
			500_000,
			200_000,
			100_000,
			50_000,
			10_000,
			5000,
			1000,
		)
	}
}
