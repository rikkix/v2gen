package mean

import (
	"sort"
)

type Value interface {
	Sum(value Value) Value
	DividedBy(i int) Value
}

type Mean interface {
	Value(index int) Value
	sort.Interface
}

func Median(m Mean) Value {
	l := m.Len()
	sort.Sort(m)
	if l%2 == 0 {
		return m.Value(l / 2).Sum(m.Value(l/2 - 1)).DividedBy(2)
	} else {
		return m.Value(l/2 - 1)
	}
}

func ArithmeticMean(m Mean) Value {
	l := m.Len()
	var sum Value

	if l == 0 {
		return nil
	}

	for i := 0; i < l; i++ {
		if i == 0 {
			sum = m.Value(i)
		}
		sum = m.Value(i).Sum(sum)
	}
	return sum.DividedBy(l)
}
