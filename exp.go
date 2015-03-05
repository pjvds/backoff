package backoff

import (
	"math"
	"time"
)

type ExpDelay struct {
	delay float64
	count float64
	max   time.Duration
}

// Creates an exponential delay. Call Delay to delay the current
// context, or Reset to reset the delay time to its initial state.
func Exp(delay time.Duration, max time.Duration) *ExpDelay {
	return &ExpDelay{
		delay: float64(delay),
		max:   max,
	}
}

func (this *ExpDelay) Delay() {
	delay := time.Duration(math.Pow(2, this.count) * this.delay)

	if delay > this.max {
		delay = this.max
	}
	time.Sleep(delay)
	this.count++
}

func (this *ExpDelay) Reset() {
	this.count = 0
}
