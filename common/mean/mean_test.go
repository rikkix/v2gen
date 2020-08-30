package mean_test

import (
	"testing"

	"iochen.com/v2gen/v2/common/mean"
)

type T int

type S struct {
	v []T
}

func (v T) Sum(value mean.Value) mean.Value {
	return v + value.(T)
}

func (v T) DividedBy(i int) mean.Value {
	return v / T(i)
}

func (s *S) Value(index int) mean.Value {
	return s.v[index]
}

func (s *S) Len() int {
	return len(s.v)
}

func (s *S) Less(i, j int) bool {
	return s.v[i] < s.v[j]
}

func (s *S) Swap(i, j int) {
	s.v[i], s.v[j] = s.v[j], s.v[i]
}

func TestArithmeticMean(t *testing.T) {
	s := &S{v: []T{
		65, 75, 30, 10, 22,
		77, 13, 99, 60, 45,
		46, 40, 28, 34, 44,
		63, 90, 81, 74, 85,
		11, 93, 80, 41, 2,
		80, 20, 90, 71, 30,
		58, 84, 83, 79, 5,
		98, 69, 71, 21, 45,
		80, 61, 41, 75, 27,
		98, 21, 63, 71, 92,
	}}

	t.Log(mean.Median(s))
	t.Log(mean.ArithmeticMean(s))
}
