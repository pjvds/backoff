package main

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
