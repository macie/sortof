package sortof

import (
	"context"
	"fmt"
	"math"
	"slices"
	"testing"
	"time"
)

func TestMiraclesortFloatSorted(t *testing.T) {
	testcases := [][]float64{
		{1, 2, 3},
		{math.Log(-1), -1, 0, math.SmallestNonzeroFloat64, 2, math.MaxFloat64},
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			t.Parallel()

			collection := slices.Clone(tc)
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			err := Miraclesort(ctx, collection)
			if err != nil {
				t.Errorf("Miraclesort(%v, %v) returns error: %v", ctx, tc, err)
			}
			if !slices.IsSorted(collection) {
				want := slices.Clone(tc)
				slices.Sort(want)
				t.Errorf("Miraclesort(%v, %v) cannot sort; got %v, want %v", ctx, tc, collection, want)
			}
		})
	}
}

func TestMiraclesortFloatUnsorted(t *testing.T) {
	testcases := [][]float64{
		{3, 1, 2},
		{math.MaxFloat64, 2, 0, -1, math.SmallestNonzeroFloat64, math.Log(-1)},
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			t.Parallel()

			collection := slices.Clone(tc)
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			err := Miraclesort(ctx, collection)
			if err != context.DeadlineExceeded {
				t.Errorf("Miraclesort(%v, %v) returns error: %v", ctx, tc, err)
			}
		})
	}
}

func TestMiraclesortIntSorted(t *testing.T) {
	testcases := [][]int{
		{1, 2, 3},
		{math.MinInt, 0, math.MaxInt},
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			t.Parallel()

			collection := slices.Clone(tc)
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			err := Miraclesort(ctx, collection)
			if err != nil {
				t.Errorf("Miraclesort(%v, %v) returns error: %v", ctx, tc, err)
			}
			if !slices.IsSorted(collection) {
				want := slices.Clone(tc)
				slices.Sort(want)
				t.Errorf("Miraclesort(%v, %v) cannot sort; got %v, want %v", ctx, tc, collection, want)
			}
		})
	}
}

func TestMiraclesortIntUnsorted(t *testing.T) {
	testcases := [][]int{
		{3, -2, math.MaxInt, 3},
		{math.MaxInt, math.MinInt},
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			t.Parallel()

			collection := slices.Clone(tc)
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			err := Miraclesort(ctx, collection)
			if err != context.DeadlineExceeded {
				t.Errorf("Miraclesort(%v, %v) returns error: %v", ctx, tc, err)
			}
		})
	}
}

func TestMiraclesortStringSorted(t *testing.T) {
	testcases := [][]string{
		{"-1", "0", "100", "2"},
		{".", "1", "a"},
		{"", "a", "b"},
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			t.Parallel()

			collection := slices.Clone(tc)
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			err := Miraclesort(ctx, collection)
			if err != nil {
				t.Errorf("Miraclesort(%v, %v) returns error: %v", ctx, tc, err)
			}
			if !slices.IsSorted(collection) {
				want := slices.Clone(tc)
				slices.Sort(want)
				t.Errorf("Miraclesort(%v, %v) cannot sort; got %v, want %v", ctx, tc, collection, want)
			}
		})
	}
}

func TestMiraclesortStringUnsorted(t *testing.T) {
	testcases := [][]string{
		{"100", "2", "0", "-1"},
		{"1", "a", "."},
		{"a", "", "b"},
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			t.Parallel()

			collection := slices.Clone(tc)
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			err := Miraclesort(ctx, collection)
			if err != context.DeadlineExceeded {
				t.Errorf("Miraclesort(%v, %v) returns error: %v", ctx, tc, err)
			}
		})
	}
}
