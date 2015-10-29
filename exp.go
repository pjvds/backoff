package backoff

import (
	"math"
	"time"
)

type ExpDelay struct {
	delay time.Duration
	count float64
	max   time.Duration
}

// Creates an exponential delay. Call Delay to delay the current
// context, or Reset to reset the delay time to its initial state.
func Exp(delay time.Duration, max time.Duration) *ExpDelay {
	return &ExpDelay{
		delay: delay,
		max:   max,
	}
}

// Next return the duration of the next delay.
func (this *ExpDelay) Next() time.Duration {
	return time.Duration(math.Pow(2, this.count) * float64(this.delay))
}

// Delays the current context with the initial delay times multiplied
// by the power of 2 of the delay count. The delay count is increased
// by 1 for every call to this method. Use Reset to reset this count to 0.
func (this *ExpDelay) Delay() {
	<-this.DelayC()
}

func (this *ExpDelay) DelayC() <-chan time.Time {
	delay := this.Next()

	if delay > this.max {
		delay = this.max
	} else {
		// only increase counter when the ceiling is not
		// reached yet, otherwise delay will eventually overflow
		// when reset is never called.
		this.count++
	}

	return time.After(delay)
}

// Reset the delay count to 0.
func (this *ExpDelay) Reset() {
	this.count = 0
}
