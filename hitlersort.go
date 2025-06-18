package sortof

import (
	"cmp"
	"context"
	"slices"
)

// Hitlersort sorts the slice x of any ordered type in ascending order by
// repeatedly removing odd elements until the slice is sorted.
//
// When sorting floating-point numbers, NaNs are ordered before other values.
// Cancelled context can leave slice partially ordered.
//
// See https://devrant.com/rants/2066915/hitler-sort-delete-every-odd-element
func Hitlersort[S ~[]E, E cmp.Ordered](ctx context.Context, x S) (S, error) {
	return HitlersortFunc(ctx, x, cmp.Compare)
}

// HitlersortFunc sorts the slice x of any type as determined by cmp,
// by repeatedly removing odd elements until sorted.
//
// Function cmp(a, b) returns negative if a<b, positive if a>b, zero if equal.
//
// Cancelled context can leave slice partially ordered.
//
// See https://devrant.com/rants/2066915/hitler-sort-delete-every-odd-element
func HitlersortFunc[S ~[]E, E any](ctx context.Context, x S, cmp func(a, b E) int) (S, error) {
	for !slices.IsSortedFunc(x, cmp) {
		select {
		case <-ctx.Done():
			return nil, context.Cause(ctx)
		default:
			filtered := x[:0] // without new slice allocation
			for i, v := range x {
				if i%2 == 0 {
					filtered = append(filtered, v)
				}
			}
			x = filtered
		}
	}
	return x, nil
}
