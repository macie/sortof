package sortof

import (
	"context"
	"fmt"
	"math"
	"slices"
	"testing"
)

func TestSlowsortFloat(t *testing.T) {
	ctx := context.Background()
	testcases := [][]float64{
		{1, 2, 3},
		{math.MaxFloat64, 2, 0, -1, math.SmallestNonzeroFloat64, math.Log(-1)},
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			t.Parallel()

			collection := slices.Clone(tc)

			err := Slowsort(ctx, collection)
			if err != nil {
				t.Errorf("Slowsort(%v, %v) returns error: %v", ctx, tc, err)
			}
			if !slices.IsSorted(collection) {
				want := slices.Clone(tc)
				slices.Sort(want)
				t.Errorf("Slowsort(%v, %v) cannot sort; got %v, want %v", ctx, tc, collection, want)
			}
		})
	}
}

func TestSlowsortInt(t *testing.T) {
	ctx := context.Background()
	testcases := [][]int{
		{1, 2, 3},
		{math.MaxInt, 0, math.MinInt},
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			t.Parallel()

			collection := slices.Clone(tc)

			err := Slowsort(ctx, collection)
			if err != nil {
				t.Errorf("Slowsort(%v, %v) returns error: %v", ctx, tc, err)
			}
			if !slices.IsSorted(collection) {
				want := slices.Clone(tc)
				slices.Sort(want)
				t.Errorf("Slowsort(%v, %v) cannot sort; got %v, want %v", ctx, tc, collection, want)
			}
		})
	}
}

func TestSlowsortString(t *testing.T) {
	ctx := context.Background()
	testcases := [][]string{
		{"1", "2", "3"},
		{"100", "2", "0", "-1"},
		{"1", "a", "."},
		{"a", "", "b"},
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			t.Parallel()

			collection := slices.Clone(tc)

			err := Slowsort(ctx, collection)
			if err != nil {
				t.Errorf("Slowsort(%v, %v) returns error: %v", ctx, tc, err)
			}
			if !slices.IsSorted(collection) {
				want := slices.Clone(tc)
				slices.Sort(want)
				t.Errorf("Slowsort(%v, %v) cannot sort; got %v, want %v", ctx, tc, collection, want)
			}
		})
	}
}
