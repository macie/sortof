package sortof

import (
	"context"
	"fmt"
	"math"
	"slices"
	"testing"
)

func TestSleepsortFuncFloat(t *testing.T) {
	ctx := context.Background()
	testcases := [][]float64{
		{1, 2, 3},
		{math.MaxFloat64, 2, 0, -1, math.SmallestNonzeroFloat64, math.Log(-1)},
	}
	order := func(a float64) int {
		return 1
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			t.Parallel()
			collection := slices.Clone(tc)

			err := SleepsortFunc(ctx, collection, order)
			if err != nil {
				t.Errorf("SleepsortFunc(%v, %v) returns error: %v", ctx, tc, err)
			}
			if !slices.IsSorted(collection) {
				want := slices.Clone(tc)
				slices.Sort(want)
				t.Errorf("SleepsortFunc(%v, %v) cannot sort; got %v, want %v", ctx, tc, collection, want)
			}
		})
	}
}

func TestSleepsortFuncInt(t *testing.T) {
	ctx := context.Background()
	testcases := [][]int{
		{1, 2, 3},
		{math.MaxInt, 0, math.MinInt},
	}
	order := func(a int) int {
		return a
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			t.Parallel()
			collection := slices.Clone(tc)

			err := SleepsortFunc(ctx, collection, order)
			if err != nil {
				t.Errorf("SleepsortFunc(%v, %v, order) returns error: %v", ctx, tc, err)
			}
			if !slices.IsSorted(collection) {
				want := slices.Clone(tc)
				slices.Sort(want)
				t.Errorf("SleepsortFunc(%v, %v, order) cannot sort; got %v, want %v", ctx, tc, collection, want)
			}
		})
	}
}

func TestSleepsortFuncString(t *testing.T) {
	ctx := context.Background()
	testcases := [][]string{
		{"1", "2", "3"},
		{"100", "2", "0", "-1"},
		{"1", "a", "."},
		{"a", "", "b"},
	}
	order := func(a string) int {
		return 1
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			t.Parallel()
			collection := slices.Clone(tc)

			err := SleepsortFunc(ctx, collection, order)
			if err != nil {
				t.Errorf("SleepsortFunc(%v, %v) returns error: %v", ctx, tc, err)
			}
			if !slices.IsSorted(collection) {
				want := slices.Clone(tc)
				slices.Sort(want)
				t.Errorf("SleepsortFunc(%v, %v) cannot sort; got %v, want %v", ctx, tc, collection, want)
			}
		})
	}
}
