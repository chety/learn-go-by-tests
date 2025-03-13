package main

import (
	"github.com/chety/learn-go-with-tests/mocks"
	"os"
	"time"
)

func main() {
	configurableSleeper := &mocks.ConfigurableSleeper{Duration: 3 * time.Second, Sleeper: time.Sleep}
	mocks.CountDown(os.Stdout, configurableSleeper)
}
