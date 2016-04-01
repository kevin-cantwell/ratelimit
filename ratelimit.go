// Package ratelimit demonstrates how to implement a rate limiter in Go
// Many apis limit access by defining a limit expressed as a rate of
// api calls per unit of time. To ensure that you never bust those limits
// you can funnel all calls through a rate limiter such as this.
package ratelimit

import "time"

type RateLimiter struct {
	d  time.Duration
	ch chan time.Time
}

func New(n int, d time.Duration) *RateLimiter {
	return &RateLimiter{
		d:  d,
		ch: make(chan time.Time, n),
	}
}

func (rl *RateLimiter) Do(f func()) {
	rl.ch <- time.Now() // Will block until there is room
	// After duration d, release one item from the channel
	go func() {
		time.Sleep(rl.d)
		<-rl.ch
	}()
	f()
}
