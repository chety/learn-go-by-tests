package select_concurrency

import (
	"errors"
	"net/http"
	"time"
)

/*
//first edition
func Race(url1, url2 string) string {
	url1Duration := measureDuration(url1)
	url2Duration := measureDuration(url2)

	if url1Duration < url2Duration {
		return url1
	}
	return url2
}

func measureDuration(url string) time.Duration {
	now := time.Now()
	_, _ = http.Get(url)
	return time.Since(now)
}
*/

var (
	defaultTimeOut time.Duration = 10 * time.Second
)

func Race(url1, url2 string) (string, error) {
	return ConfigurableRace(url1, url2, defaultTimeOut)
}

func ConfigurableRace(url1, url2 string, timeOut time.Duration) (string, error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil

	case <-time.After(timeOut):
		return "", errors.New("timeout")
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		_, _ = http.Get(url)
		close(ch)
	}()
	return ch
}
