package sortof

import (
	"cmp"
	"context"
	"math"
	"slices"
	"testing"
)

func TestBogosortFloat(t *testing.T) {
	ctx := context.Background()
	testcases := [][]float64{
		{1, 2, 3},
		{math.MaxFloat64, 2, 0, -1, math.SmallestNonzeroFloat64, math.Log(-1)},
	}
	for _, tc := range testcases {
		collection := slices.Clone(tc)

		err := Bogosort(ctx, collection)
		if err != nil {
			t.Errorf("Bogosort(%v, %v) returns error: %v", ctx, tc, err)
		}
		if !slices.IsSorted(collection) {
			want := slices.Clone(tc)
			slices.Sort(want)
			t.Errorf("Bogosort(%v, %v) cannot sort; got %v, want %v", ctx, tc, collection, want)
		}
	}
}

func TestBogosortInt(t *testing.T) {
	ctx := context.Background()
	testcases := [][]int{
		{1, 2, 3},
		{math.MaxInt, 2, 0, -1, math.MinInt},
	}
	for _, tc := range testcases {
		collection := slices.Clone(tc)

		err := Bogosort(ctx, collection)
		if err != nil {
			t.Errorf("Bogosort(%v, %v) returns error: %v", ctx, tc, err)
		}
		if !slices.IsSorted(collection) {
			want := slices.Clone(tc)
			slices.Sort(want)
			t.Errorf("Bogosort(%v, %v) cannot sort; got %v, want %v", ctx, tc, collection, want)
		}
	}
}

func TestBogosortString(t *testing.T) {
	ctx := context.Background()
	testcases := [][]string{
		{"1", "2", "3"},
		{"100", "2", "0", "-1"},
		{"1", "a", "."},
		{"a", "", "b"},
	}
	for _, tc := range testcases {
		collection := slices.Clone(tc)

		err := Bogosort(ctx, collection)
		if err != nil {
			t.Errorf("Bogosort(%v, %v) returns error: %v", ctx, tc, err)
		}
		if !slices.IsSorted(collection) {
			want := slices.Clone(tc)
			slices.Sort(want)
			t.Errorf("Bogosort(%v, %v) cannot sort; got %v, want %v", ctx, tc, collection, want)
		}
	}
}

func TestBogosortFuncFloat(t *testing.T) {
	ctx := context.Background()
	cmpFloats := func(a, b float64) int { return cmp.Compare(a, b) }
	testcases := [][]float64{
		{1, 2, 3},
		{math.MaxFloat64, 2, 0, -1, math.SmallestNonzeroFloat64, math.Log(-1)},
	}
	for _, tc := range testcases {
		collection := slices.Clone(tc)

		err := BogosortFunc(ctx, collection, cmpFloats)
		if err != nil {
			t.Errorf("BogosortFunc(%v, %v, cmpFloats) returns error: %v", ctx, tc, err)
		}
		if !slices.IsSorted(collection) {
			want := slices.Clone(tc)
			slices.SortFunc(want, cmpFloats)
			t.Errorf("BogosortFunc(%v, %v, cmpFloats) cannot sort; got %v, want %v", ctx, tc, collection, want)
		}
	}
}

func TestBogosortFuncInt(t *testing.T) {
	ctx := context.Background()
	cmpInts := func(a, b int) int { return cmp.Compare(a, b) }
	testcases := [][]int{
		{1, 2, 3},
		{math.MaxInt, 2, 0, -1, math.MinInt},
	}
	for _, tc := range testcases {
		collection := slices.Clone(tc)

		err := BogosortFunc(ctx, collection, cmpInts)
		if err != nil {
			t.Errorf("BogosortFunc(%v, %v, cmpInts) returns error: %v", ctx, tc, err)
		}
		if !slices.IsSorted(collection) {
			want := slices.Clone(tc)
			slices.SortFunc(want, cmpInts)
			t.Errorf("BogosortFunc(%v, %v, cmpInts) cannot sort; got %v, want %v", ctx, tc, collection, want)
		}
	}
}

func TestBogosortFuncString(t *testing.T) {
	ctx := context.Background()
	cmpStrings := func(a, b string) int { return cmp.Compare(a, b) }
	testcases := [][]string{
		{"1", "2", "3"},
		{"100", "2", "0", "-1"},
		{"1", "a", "."},
		{"a", "", "b"},
	}
	for _, tc := range testcases {
		collection := slices.Clone(tc)

		err := BogosortFunc(ctx, collection, cmpStrings)
		if err != nil {
			t.Errorf("BogosortFunc(%v, %v, cmpStrings) returns error: %v", ctx, tc, err)
		}
		if !slices.IsSorted(collection) {
			want := slices.Clone(tc)
			slices.Sort(want)
			t.Errorf("BogosortFunc(%v, %v, cmpStrings) cannot sort; got %v, want %v", ctx, tc, collection, want)
		}
	}
}