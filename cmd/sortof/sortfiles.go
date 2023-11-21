package main

import (
	"bufio"
	"context"
	"io"

	"github.com/macie/sortof"
)

// BogosortFile returns a sorted lines from the file in ascending order.
// A context controls cancellation.
func BogosortFile(ctx context.Context, file io.ReadCloser) ([]string, error) {
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return []string{}, context.Cause(ctx)
		default:
			lines = append(lines, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	if err := sortof.Bogosort(ctx, lines); err != nil {
		return []string{}, err
	}

	return lines, nil
}

// MiraclesortFile returns a sorted lines from the file in ascending order.
// A context controls cancellation.
func MiraclesortFile(ctx context.Context, file io.ReadCloser) ([]string, error) {
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return []string{}, context.Cause(ctx)
		default:
			lines = append(lines, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	if err := sortof.Miraclesort(ctx, lines); err != nil {
		return []string{}, err
	}

	return lines, nil
}

// SlowsortFile returns a sorted lines from the file in ascending order.
// A context controls cancellation.
func SlowsortFile(ctx context.Context, file io.ReadCloser) ([]string, error) {
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return []string{}, context.Cause(ctx)
		default:
			lines = append(lines, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	if err := sortof.Slowsort(ctx, lines); err != nil {
		return []string{}, err
	}

	return lines, nil
}

// StalinsortFile returns a sorted lines from the file in ascending order.
// A context controls cancellation.
func StalinsortFile(ctx context.Context, file io.ReadCloser) ([]string, error) {
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return []string{}, context.Cause(ctx)
		default:
			lines = append(lines, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	sorted, err := sortof.Stalinsort(ctx, lines)
	if err != nil {
		return []string{}, err
	}

	return sorted, nil
}
