package sortof

import (
	"cmp"
	"context"
	"math/rand"
	"slices"
)

// Bogosort sorts the slice x of any ordered type in ascending order. A context
// controls cancellation, because the worst-case time complexity is O(infinity).
// When sorting floating-point numbers, NaNs are ordered before other values.
//
// See https://en.wikipedia.org/wiki/Bogosort.
func Bogosort[S ~[]E, E cmp.Ordered](ctx context.Context, x S) error {
	return BogosortFunc(ctx, x, cmp.Compare)
}

// BogosortFunc sorts the slice x of any type in ascending order as
// determined by the cmp function. A context controls cancellation, because
// the worst-case time complexity is O(infinity). Function cmp(a, b) should
// return a negative number when a < b, a positive number when a > b and zero
// when a == b.
//
// See https://en.wikipedia.org/wiki/Bogosort.
func BogosortFunc[S ~[]E, E any](ctx context.Context, x S, cmp func(a, b E) int) error {
	n := len(x)

	for !slices.IsSortedFunc(x, cmp) {
		select {
		case <-ctx.Done():
			return context.Cause(ctx)
		default:
			rand.Shuffle(n, func(i, j int) {
				x[i], x[j] = x[j], x[i]
			})
		}
	}

	return nil
}
