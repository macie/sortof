package sortof

import (
	"context"

	"time"
)

// SleepsortFunc sorts the slice x of any type in ascending order as
// determined by the asInt function. Function asInt(el) should return
// an ordening number for each element of x.
//
// Cancelled context can leave slice in a unknown state.
func SleepsortFunc[S ~[]E, E any](ctx context.Context, x S, asInt func(el E) int) error {
	ordered := make(chan E)

	for _, element := range x {
		go func(el E) {
			aWhile := time.Duration(asInt(el)) * time.Millisecond
			time.Sleep(aWhile)
			ordered <- el
		}(element)
	}

	i := 0
	select {
	case <-ctx.Done():
		return context.Cause(ctx)
	case x[i] = <-ordered:
		i++
	default:
	}

	return nil
}
