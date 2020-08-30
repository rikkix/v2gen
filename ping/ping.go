package ping

import (
	"time"

	"iochen.com/v2gen/v2/common/mean"
)

type Duration time.Duration

func (d Duration) Precision(i int64) Duration {
	p := Duration(i)
	return (d / p) * p
}

func (d Duration) String() string {
	return time.Duration(d).String()
}

func (d Duration) Sum(value mean.Value) mean.Value {
	return d + value.(Duration)
}

func (d Duration) DividedBy(i int) mean.Value {
	return d / Duration(i)
}

type DurationList []Duration

func (dList *DurationList) Value(index int) mean.Value {
	return (*dList)[index]
}

func (dList *DurationList) Len() int {
	return len(*dList)
}

func (dList *DurationList) Less(i, j int) bool {
	return (*dList)[i] < (*dList)[j]
}

func (dList *DurationList) Swap(i, j int) {
	(*dList)[i], (*dList)[j] = (*dList)[j], (*dList)[i]
}

type Status struct {
	Durations *DurationList
	Errors    []*error
}

func (ps *Status) Value(index int) mean.Value {
	return (*ps.Durations)[index]
}
