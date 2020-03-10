package ping

import (
	"iochen.com/v2gen/infra/mean"
	"time"
)

type Duration time.Duration

type PingStat struct {
	Durations  []Duration
	ErrCounter uint
}

func (d Duration) String() string {
	return time.Duration((d / 1000000) * 1000000).String()
}

func (d Duration) Sum(value mean.Value) mean.Value {
	return d + value.(Duration)
}

func (d Duration) DividedBy(i int) mean.Value {
	return d / Duration(i)
}

func (ps *PingStat) Value(index int) mean.Value {
	return ps.Durations[index]
}

func (ps *PingStat) Len() int {
	return len(ps.Durations)
}

func (ps *PingStat) Less(i, j int) bool {
	return ps.Durations[i] < ps.Durations[j]
}

func (ps *PingStat) Swap(i, j int) {
	ps.Durations[i], ps.Durations[j] = ps.Durations[j], ps.Durations[i]
}
