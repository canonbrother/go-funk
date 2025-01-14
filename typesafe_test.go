package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsBool(t *testing.T) {
	is := assert.New(t)

	is.True(ContainsBool([]bool{true, false}, true))
	is.False(ContainsBool([]bool{true}, false))
}

func TestContainsInt(t *testing.T) {
	is := assert.New(t)

	is.True(ContainsInt([]int{1, 2, 3, 4}, 4))
	is.False(ContainsInt([]int{1, 2, 3, 4}, 5))

	is.True(ContainsInt32([]int32{1, 2, 3, 4}, 4))
	is.False(ContainsInt32([]int32{1, 2, 3, 4}, 5))

	is.True(ContainsInt64([]int64{1, 2, 3, 4}, 4))
	is.False(ContainsInt64([]int64{1, 2, 3, 4}, 5))

	is.True(ContainsUInt([]uint{1, 2, 3, 4}, 4))
	is.False(ContainsUInt([]uint{1, 2, 3, 4}, 5))

	is.True(ContainsUInt32([]uint32{1, 2, 3, 4}, 4))
	is.False(ContainsUInt32([]uint32{1, 2, 3, 4}, 5))

	is.True(ContainsUInt64([]uint64{1, 2, 3, 4}, 4))
	is.False(ContainsUInt64([]uint64{1, 2, 3, 4}, 5))
}

func TestContainsString(t *testing.T) {
	is := assert.New(t)

	is.True(ContainsString([]string{"flo", "gilles"}, "flo"))
	is.False(ContainsString([]string{"flo", "gilles"}, "alex"))
}

func TestFilterBool(t *testing.T) {
	is := assert.New(t)

	r := FilterBool([]bool{true, true, false, true}, func(x bool) bool {
		return x == true
	})

	is.Equal(r, []bool{true, true, true})
}

func TestFilterString(t *testing.T) {
	is := assert.New(t)

	r := FilterString([]string{"a", "b", "c", "d"}, func(x string) bool {
		return x >= "c"
	})

	is.Equal(r, []string{"c", "d"})
}

func TestFilterInt(t *testing.T) {
	is := assert.New(t)

	r := FilterInt([]int{1, 2, 3, 4}, func(x int) bool {
		return x%2 == 0
	})

	is.Equal(r, []int{2, 4})
}

func TestFilterInt32(t *testing.T) {
	is := assert.New(t)

	r := FilterInt32([]int32{1, 2, 3, 4}, func(x int32) bool {
		return x%2 == 0
	})

	is.Equal(r, []int32{2, 4})
}

func TestFilterInt64(t *testing.T) {
	is := assert.New(t)

	r := FilterInt64([]int64{1, 2, 3, 4}, func(x int64) bool {
		return x%2 == 0
	})

	is.Equal(r, []int64{2, 4})
}

func TestFilterUInt(t *testing.T) {
	is := assert.New(t)

	r := FilterUInt([]uint{1, 2, 3, 4}, func(x uint) bool {
		return x%2 == 0
	})

	is.Equal(r, []uint{2, 4})
}

func TestFilterUInt32(t *testing.T) {
	is := assert.New(t)

	r := FilterUInt32([]uint32{1, 2, 3, 4}, func(x uint32) bool {
		return x%2 == 0
	})

	is.Equal(r, []uint32{2, 4})
}

func TestFilterUInt64(t *testing.T) {
	is := assert.New(t)

	r := FilterUInt64([]uint64{1, 2, 3, 4}, func(x uint64) bool {
		return x%2 == 0
	})

	is.Equal(r, []uint64{2, 4})
}

func TestFilterFloat64(t *testing.T) {
	is := assert.New(t)

	r := FilterFloat64([]float64{1.0, 2.0, 3.0, 4.0}, func(x float64) bool {
		return int(x)%2 == 0
	})

	is.Equal(r, []float64{2.0, 4.0})
}

func TestFilterFloat32(t *testing.T) {
	is := assert.New(t)

	r := FilterFloat32([]float32{1.0, 2.0, 3.0, 4.0}, func(x float32) bool {
		return int(x)%2 == 0
	})

	is.Equal(r, []float32{2.0, 4.0})
}

func TestContainsFloat(t *testing.T) {
	is := assert.New(t)

	is.True(ContainsFloat64([]float64{0.1, 0.2}, 0.1))
	is.False(ContainsFloat64([]float64{0.1, 0.2}, 0.3))

	is.True(ContainsFloat32([]float32{0.1, 0.2}, 0.1))
	is.False(ContainsFloat32([]float32{0.1, 0.2}, 0.3))
}

func TestSumNumeral(t *testing.T) {
	is := assert.New(t)

	is.Equal(SumInt([]int{1, 2, 3}), 6)
	is.Equal(SumInt64([]int64{1, 2, 3}), int64(6))

	is.Equal(SumUInt([]uint{1, 2, 3}), uint(6))
	is.Equal(SumUInt64([]uint64{1, 2, 3}), uint64(6))

	is.Equal(SumFloat32([]float32{0.1, 0.2, 0.1}), float32(0.4))
	is.Equal(SumFloat64([]float64{0.1, 0.2, 0.1}), float64(0.4))
}

func TestTypesafeReverse(t *testing.T) {
	is := assert.New(t)

	is.Equal(ReverseBools([]bool{true, false, false}), []bool{false, false, true})
	is.Equal(ReverseString("abcdefg"), "gfedcba")
	is.Equal(ReverseInt([]int{1, 2, 3, 4}), []int{4, 3, 2, 1})
	is.Equal(ReverseInt64([]int64{1, 2, 3, 4}), []int64{4, 3, 2, 1})
	is.Equal(ReverseUInt([]uint{1, 2, 3, 4}), []uint{4, 3, 2, 1})
	is.Equal(ReverseUInt64([]uint64{1, 2, 3, 4}), []uint64{4, 3, 2, 1})
	is.Equal(ReverseStrings([]string{"flo", "gilles"}), []string{"gilles", "flo"})
	is.Equal(ReverseFloat64([]float64{0.1, 0.2, 0.3}), []float64{0.3, 0.2, 0.1})
	is.Equal(ReverseFloat32([]float32{0.1, 0.2, 0.3}), []float32{0.3, 0.2, 0.1})
}

func TestTypesafeIndexOf(t *testing.T) {
	is := assert.New(t)

	is.Equal(IndexOfBool([]bool{true, false}, false), 1)
	is.Equal(IndexOfBool([]bool{true}, false), -1)

	is.Equal(IndexOfString([]string{"foo", "bar"}, "bar"), 1)
	is.Equal(IndexOfString([]string{"foo", "bar"}, "flo"), -1)

	is.Equal(IndexOfInt([]int{0, 1, 2}, 1), 1)
	is.Equal(IndexOfInt([]int{0, 1, 2}, 3), -1)

	is.Equal(IndexOfInt64([]int64{0, 1, 2}, 1), 1)
	is.Equal(IndexOfInt64([]int64{0, 1, 2}, 3), -1)

	is.Equal(IndexOfUInt64([]uint64{0, 1, 2}, 1), 1)
	is.Equal(IndexOfUInt64([]uint64{0, 1, 2}, 3), -1)

	is.Equal(IndexOfFloat64([]float64{0.1, 0.2, 0.3}, 0.2), 1)
	is.Equal(IndexOfFloat64([]float64{0.1, 0.2, 0.3}, 0.4), -1)
}

func TestTypesafeLastIndexOf(t *testing.T) {
	is := assert.New(t)

	is.Equal(LastIndexOfBool([]bool{true, true, false, true}, true), 3)
	is.Equal(LastIndexOfString([]string{"foo", "bar", "bar"}, "bar"), 2)
	is.Equal(LastIndexOfInt([]int{1, 2, 2, 3}, 2), 2)
	is.Equal(LastIndexOfInt64([]int64{1, 2, 2, 3}, 4), -1)
	is.Equal(LastIndexOfUInt([]uint{1, 2, 2, 3}, 2), 2)
	is.Equal(LastIndexOfUInt64([]uint64{1, 2, 2, 3}, 4), -1)
}

func TestTypesafeUniq(t *testing.T) {
	is := assert.New(t)

	is.Equal(UniqBool([]bool{true, false, false, true, false}), []bool{true, false})
	is.Equal(UniqInt64([]int64{0, 1, 1, 2, 3, 0, 0, 12}), []int64{0, 1, 2, 3, 12})
	is.Equal(UniqInt([]int{0, 1, 1, 2, 3, 0, 0, 12}), []int{0, 1, 2, 3, 12})
	is.Equal(UniqUInt([]uint{0, 1, 1, 2, 3, 0, 0, 12}), []uint{0, 1, 2, 3, 12})
	is.Equal(UniqUInt64([]uint64{0, 1, 1, 2, 3, 0, 0, 12}), []uint64{0, 1, 2, 3, 12})
	is.Equal(UniqFloat64([]float64{0.0, 0.1, 0.1, 0.2, 0.3, 0.0, 0.0, 0.12}), []float64{0.0, 0.1, 0.2, 0.3, 0.12})
	is.Equal(UniqString([]string{"foo", "bar", "foo", "bar"}), []string{"foo", "bar"})
}

func TestTypesafeShuffle(t *testing.T) {
	is := assert.New(t)

	initial := []int{1, 2, 3, 5}

	results := ShuffleInt(initial)

	is.Len(results, 4)

	for _, entry := range initial {
		is.True(ContainsInt(results, entry))
	}
}

func TestDropBool(t *testing.T) {
	results := DropBool([]bool{true, false, false, true, true}, 3)

	is := assert.New(t)

	is.Len(results, 2)

	is.Equal([]bool{true, true}, results)
}

func TestDropString(t *testing.T) {
	results := DropString([]string{"the", "quick", "brown", "fox", "jumps", "..."}, 3)

	is := assert.New(t)

	is.Len(results, 3)

	is.Equal([]string{"fox", "jumps", "..."}, results)
}

func TestDropInt(t *testing.T) {
	results := DropInt([]int{0, 0, 0, 0}, 3)

	is := assert.New(t)

	is.Len(results, 1)

	is.Equal([]int{0}, results)
}

func TestDropInt32(t *testing.T) {
	results := DropInt32([]int32{1, 2, 3, 4}, 3)

	is := assert.New(t)

	is.Len(results, 1)

	is.Equal([]int32{4}, results)
}

func TestDropInt64(t *testing.T) {
	results := DropInt64([]int64{1, 2, 3, 4}, 3)

	is := assert.New(t)

	is.Len(results, 1)

	is.Equal([]int64{4}, results)
}

func TestDropUInt(t *testing.T) {
	results := DropUInt([]uint{0, 0, 0, 0}, 3)

	is := assert.New(t)

	is.Len(results, 1)

	is.Equal([]uint{0}, results)
}

func TestDropUInt32(t *testing.T) {
	results := DropUInt32([]uint32{1, 2, 3, 4}, 3)

	is := assert.New(t)

	is.Len(results, 1)

	is.Equal([]uint32{4}, results)
}

func TestDropUInt64(t *testing.T) {
	results := DropUInt64([]uint64{1, 2, 3, 4}, 3)

	is := assert.New(t)

	is.Len(results, 1)

	is.Equal([]uint64{4}, results)
}

func TestDropFloat32(t *testing.T) {
	results := DropFloat32([]float32{1.1, 2.2, 3.3, 4.4}, 3)

	is := assert.New(t)

	is.Len(results, 1)

	is.Equal([]float32{4.4}, results)
}

func TestDropFloat64(t *testing.T) {
	results := DropFloat64([]float64{1.1, 2.2, 3.3, 4.4}, 3)

	is := assert.New(t)

	is.Len(results, 1)

	is.Equal([]float64{4.4}, results)
}
