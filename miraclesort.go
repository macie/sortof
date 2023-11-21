package sortof

import (
	"cmp"
	"context"
	"slices"
	"time"
)

// Miraclesort sorts the slice x of any ordered type in ascending order.
// A context controls cancellation, because miracles are non-deterministic and
// there is no guarantees, that slice will be ever sorted. When sorting
// floating-point numbers, NaNs are ordered before other values.
//
// See https://en.wikipedia.org/wiki/Bogosort#Related_algorithms.
func Miraclesort[S ~[]E, E cmp.Ordered](ctx context.Context, x S) error {
	return MiraclesortFunc(ctx, x, cmp.Compare)
}

// MiraclesortFunc sorts the slice x of any type in ascending order as
// determined by the cmp function. A context controls cancellation, because
// miracles are non-deterministic and there is no guarantees, that slice
// will be ever sorted. Function cmp(a, b) should return a negative number
// when a < b, a positive number when a > b and zero when a == b.
//
// See https://en.wikipedia.org/wiki/Bogosort#Related_algorithms.
func MiraclesortFunc[S ~[]E, E any](ctx context.Context, x S, cmp func(a, b E) int) error {
	// This implementation is based on assumption, that miracles occures from
	// time to time.
	//
	// There are a couple of time units traditionally connected to miracles, but
	// some are not specified enough to be useful (eg. "every now and then").
	// The most appropriate units are:
	//  - Jubeljahre: https://en.wiktionary.org/wiki/alle_Jubeljahre
	//  - ruski rok: https://en.wiktionary.org/wiki/raz_na_ruski_rok
	//  - blue moon: https://en.wiktionary.org/wiki/once_in_a_blue_moon
	//  - zdrowaśka: https://en.wiktionary.org/wiki/zdrowa%C5%9Bka
	//
	// To provide smooth UI for user, interval should be short, and
	// millizdrowaśka fits the requirements.
	const millizdrowaska = 200 * time.Millisecond // 0.001 * 20 sec

	for !slices.IsSortedFunc(x, cmp) {
		select {
		case <-ctx.Done():
			return context.Cause(ctx)
		default:
			time.Sleep(millizdrowaska)
		}
	}

	return nil
}
