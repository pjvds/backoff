# Backoff

Simple backoff algorithms support package for Go. Usually used in retries in error situations.

# Exponential backoff

Here is an example to for an exponential backoff in Go:

``` go
import (
    "fmt"
    "time"

    "github.com/pjvds/backoff"
)

func main() {
    delay := backoff.Exp(1*time.Millisecond, 10*time.Second)

    for {
        started := time.Now()
        delay.Delay()

        fmt.Printf("%v\n", time.Since(started))
    }
}
```

Outputs:

``` terminal
$ go run main.go
1.221325ms
2.097915ms
4.251798ms
8.360996ms
16.76387ms
32.906347ms
64.342088ms
128.824612ms
256.234941ms
513.376278ms
1.025531565s
2.048582267s
4.097193995s
8.192315835s
10.000477256s
```