package util

const DEFAULT_X0 = 0 // seconds
const DEFAULT_X = 30 // seconds

var (
	TimeStepper = NewTimeStepper(DEFAULT_X0, DEFAULT_X)
)

func NewTimeStepper(x0, step uint64) func(uint64) uint64 {
	return func(timestamp uint64) uint64 {
		return (timestamp - x0) / step
	}
}
