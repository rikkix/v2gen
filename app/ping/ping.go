package ping

import (
	"iochen.com/v2gen/infra/mean"
	"time"
)

type Duration time.Duration

type Status struct {
	Durations  []Duration
	ErrCounter uint
}

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

func (ps *Status) Value(index int) mean.Value {
	return ps.Durations[index]
}

func (ps *Status) Len() int {
	return len(ps.Durations)
}

func (ps *Status) Less(i, j int) bool {
	return ps.Durations[i] < ps.Durations[j]
}

func (ps *Status) Swap(i, j int) {
	ps.Durations[i], ps.Durations[j] = ps.Durations[j], ps.Durations[i]
}
