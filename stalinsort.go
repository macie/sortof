package sortof

import (
	"cmp"
	"context"
)

// Stalinsort returns slice created from x by deleting elements which are not
// in ascending order. For compatibility with other functions from package,
// context controls cancellation.
//
// For floating-point types, a NaN is considered less than any non-NaN,
// a NaN is considered equal to a NaN, and -0.0 is equal to 0.0.
//
// See https://mastodon.social/@mathew/100958177234287431.
func Stalinsort[S ~[]E, E cmp.Ordered](ctx context.Context, x S) (S, error) {
	return StalinsortFunc(ctx, x, cmp.Compare)
}

// StalinsortFunc returns slice created from slice x by deleting elements which
// are not in order determined by the cmp function. For compatibility with
// other functions from package, context controls cancellation. Function
// cmp(a, b) should return a negative number when a < b, a positive number
// when a > b and zero when a == b.
//
// See https://mastodon.social/@mathew/100958177234287431.
func StalinsortFunc[S ~[]E, E any](ctx context.Context, x S, cmp func(a, b E) int) (S, error) {
	sorted := make(S, 0)
	for i := range x {
		select {
		case <-ctx.Done():
			return nil, context.Cause(ctx)
		default:
			prev := len(sorted) - 1
			if (i == 0) || (cmp(x[i], sorted[prev]) != -1) {
				sorted = append(sorted, x[i])
			}
		}
	}

	return sorted, nil
}
