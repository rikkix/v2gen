package ping

import (
	"iochen.com/v2gen/infra/mean"
	"time"
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
	Durations  *DurationList
	Result     Duration
	ErrCounter uint
}

func (ps *Status) Value(index int) mean.Value {
	return (*ps.Durations)[index]
}

type StatusList []Status

func (sList *StatusList) Len() int {
	return len(*sList)
}

func (sList *StatusList) Less(i, j int) bool {
	if (*sList)[i].ErrCounter != (*sList)[j].ErrCounter {
		return (*sList)[i].ErrCounter < (*sList)[j].ErrCounter
	}

	return (*sList)[i].Result < (*sList)[j].Result
}

func (sList *StatusList) Swap(i, j int) {
	(*sList)[i], (*sList)[j] = (*sList)[j], (*sList)[i]
}
