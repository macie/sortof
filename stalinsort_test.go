package sortof

import (
	"context"
	"fmt"
	"math"
	"testing"
)

func TestStalinsortFloat(t *testing.T) {
	ctx := context.Background()
	testcases := map[string][]float64{
		"[1 2 3]":         {1, 2, 3},
		"[NaN NaN 0 0 0]": {math.Log(-1), math.Log(-1), 0, -0.0, 0, math.Log(-1)},
		fmt.Sprintf("[-1 2 %v]", math.MaxFloat64): {-1, 2, 0, math.Log(-1), math.MaxFloat64, math.Log(-1)},
	}
	for want, tc := range testcases {
		got, err := Stalinsort(ctx, tc)
		if err != nil {
			t.Errorf("Stalinsort(%v, %v) returns error: %v", ctx, tc, err)
		}
		if fmt.Sprint(got) != want {
			t.Errorf("Stalinsort(%v, %v) cannot sort; got %v, want %v", ctx, tc, got, want)
		}
	}
}

func TestStalinsortInt(t *testing.T) {
	ctx := context.Background()
	testcases := map[string][]int{
		"[1 2 3]": {1, 2, 3},
		fmt.Sprintf("[%v 2 %v]", math.MinInt, math.MaxInt): {math.MinInt, 2, 0, -1, math.MaxInt, 3},
	}
	for want, tc := range testcases {
		got, err := Stalinsort(ctx, tc)
		if err != nil {
			t.Errorf("Stalinsort(%v, %v) returns error: %v", ctx, tc, err)
		}
		if fmt.Sprint(got) != want {
			t.Errorf("Stalinsort(%v, %v) cannot sort; got %v, want %v", ctx, tc, got, want)
		}
	}
}

func TestStalinsortString(t *testing.T) {
	ctx := context.Background()
	testcases := map[string][]string{
		"[. 1 2 3 z]": {".", "1", "2", "3", "z", "-2"},
		"[100 2]":     {"100", "2", "0", "-1"},
		"[1 a]":       {"1", "a", "."},
		"[a b]":       {"a", "", "b"},
	}
	for want, tc := range testcases {
		got, err := Stalinsort(ctx, tc)
		if err != nil {
			t.Errorf("Stalinsort(%v, %v) returns error: %v", ctx, tc, err)
		}
		if fmt.Sprint(got) != want {
			t.Errorf("Stalinsort(%v, %v) cannot sort; got %v, want %v", ctx, tc, got, want)
		}
	}
}
