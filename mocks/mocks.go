package mocks

import (
	"fmt"
	"io"
	"iter"
)

func countDownFromIter(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

const countdownStart = 3
const finalValue = "Go!"

func CountDown(out io.Writer, sleeper Sleeper) {
	for i := range countDownFromIter(3) {
		_, _ = fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	_, _ = fmt.Fprint(out, finalValue)
}
