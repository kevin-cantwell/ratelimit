package ratelimit_test

import (
	"log"
	"time"

	"github.com/kevin-cantwell/ratelimit"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RateLimiter", func() {
	Describe("#Do", func() {
		It("Should limit 5 executions per second", func() {
			rl := ratelimit.New(5, time.Second)
			start := time.Now()
			for i := 0; i < 15; i++ {
				rl.Do(func() {
					log.Println(i)
				})
				switch {
				case i >= 10:
					Expect(time.Now()).To(BeTemporally("~", start.Add(2*time.Second), 50*time.Millisecond))
				case i >= 5:
					Expect(time.Now()).To(BeTemporally("~", start.Add(1*time.Second), 50*time.Millisecond))
				case i >= 0:
					Expect(time.Now()).To(BeTemporally("~", start, 50*time.Millisecond))
				}
			}
		})
	})
})
