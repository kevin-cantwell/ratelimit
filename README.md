Package ratelimit demonstrates how to implement a rate limiter in Go.
Many apis restrict access by defining a limit expressed as a rate of
api calls per unit of time. To ensure that you never bust above those limits
you can funnel all calls through a rate limiter such as this.

Usage:

```go
// 100/hr
limiter := ratelimit.New(100, time.Hour)
for i := 0; i < 500; i++ {
  limiter.Do(func() {
    log.Println(i)
  })
}
```