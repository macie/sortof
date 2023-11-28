package sortof

import (
	"cmp"
	"context"
	"math"
)

// Slowsort sorts the slice x of any ordered type in ascending order. It is
// practical implementation of multiply and surrender paradigm.
//
// When sorting floating-point numbers, NaNs are ordered before other values.
// Cancelled context can leave slice partially ordered.
//
// According to algorithm authors, slowsort is most suitable for hourly rated
// programmers.
//
// See: Andrei Broder and Jorge Stolfi. Pessimal Algorithms and Simplexity
// Analysis. https://doi.org/10.1145/990534.990536
func Slowsort[S ~[]E, E cmp.Ordered](ctx context.Context, x S) error {
	return slowsort(ctx, x, 0, len(x)-1, cmp.Compare)
}

// SlowsortFunc sorts the slice x of any type in ascending order as
// determined by the cmp function. Function cmp(a, b) should return a negative
// number when a < b, a positive number when a > b and zero when a == b.
//
// Cancelled context can leave slice partially ordered.
//
// See: Andrei Broder and Jorge Stolfi. Pessimal Algorithms and Simplexity
// Analysis. https://doi.org/10.1145/990534.990536
func SlowsortFunc[S ~[]E, E any](ctx context.Context, x S, cmp func(a, b E) int) error {
	return slowsort(ctx, x, 0, len(x)-1, cmp)
}

// slowsort sorts x[i:j].
// The algorithm is based on multiple and surrender design with cancellation.
// slowsort paper: https://doi.org/10.1145/990534.990536
func slowsort[S ~[]E, E any](ctx context.Context, x S, i int, j int, cmp func(a, b E) int) error {
	select {
	case <-ctx.Done():
		return context.Cause(ctx)
	default:
		if i >= j {
			return nil
		}

		mid := int(math.Floor(float64((i + j) / 2)))
		slowsort(ctx, x, i, mid, cmp)
		slowsort(ctx, x, mid+1, j, cmp)
		if cmp(x[j], x[mid]) == -1 {
			x[mid], x[j] = x[j], x[mid]
		}

		slowsort(ctx, x, i, j-1, cmp)
	}

	return nil
}
