package backoff

import "time"

type ExpBackoff struct {
	initial time.Duration
	delay   time.Duration
	max     time.Duration
}

func Exp(delay time.Duration, max time.Duration) *ExpBackoff {
	return &ExpBackoff{
		initial: delay,
		delay:   delay,
		max:     max,
	}
}

func (this *ExpBackoff) Backoff() {
	if this.delay += this.initial; this.delay > this.max {
		this.delay.max
	}

	time.Sleep(this.delay)
}

func (this *ExpBackoff) Reset() {
	this.delay = this.initial
}
