package main

import (
	"regexp"
	"strings"
	"testing"
	"time"
)

func TestNewAppConfig(t *testing.T) {
	testcases := []struct {
		args []string
		want AppConfig
	}{
		{[]string{"-h"}, AppConfig{ExitMessage: helpMsg}},
		{[]string{"-v"}, AppConfig{ExitMessage: "sortof local-dev"}},
		{[]string{"bogo"}, AppConfig{SortFunc: BogosortFile}},
		{[]string{"bogo", "some_file"}, AppConfig{SortFunc: BogosortFile, Files: []string{"some_file"}}},
		{[]string{"bogo", "-t", "1s"}, AppConfig{SortFunc: BogosortFile, Timeout: time.Second}},
		{[]string{"bogo", "-t", "11s", "first_file", "second_file"}, AppConfig{
			SortFunc: BogosortFile, Timeout: 11 * time.Second, Files: []string{"first_file", "second_file"},
		}},
		{[]string{"slow"}, AppConfig{SortFunc: SlowsortFile}},
		{[]string{"slow", "-t", "5ns"}, AppConfig{SortFunc: SlowsortFile, Timeout: 5 * time.Nanosecond}},
		{[]string{"slow", "-t", "5ns", "-"}, AppConfig{
			SortFunc: SlowsortFile, Timeout: 5 * time.Nanosecond, Files: []string{"-"},
		}},
		{[]string{"stalin"}, AppConfig{SortFunc: StalinsortFile}},
		{[]string{"stalin", "-t", "2h"}, AppConfig{SortFunc: StalinsortFile, Timeout: 2 * time.Hour}},
		{[]string{"stalin", "-t", "2h", "-", "some_file"}, AppConfig{
			SortFunc: StalinsortFile, Timeout: 2 * time.Hour, Files: []string{"-", "some_file"},
		}},
		{[]string{"stalin", "-t", "2h", "some_file", "-"}, AppConfig{
			SortFunc: StalinsortFile, Timeout: 2 * time.Hour, Files: []string{"some_file", "-"},
		}},
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(strings.Join(tc.args, "_"), func(t *testing.T) {
			t.Parallel()
			got, err := NewAppConfig(tc.args)
			if err != nil {
				t.Errorf("NewAppConfig(%v) returns error: %v", tc.args, err)
			}
			if !got.Equal(tc.want) {
				t.Errorf("NewAppConfig(%v) = %v, want %v", tc.args, got, tc.want)
			}
		})
	}
}

func FuzzNewAppConfig(f *testing.F) {
	validArgs := []*regexp.Regexp{
		regexp.MustCompile(`\-h`), regexp.MustCompile(`\-v`),
		regexp.MustCompile(`bogo[ \-t \d+]?[ \w+]*`),
	}

	f.Add("-Z")
	f.Add("bogo -t 1z first_file second_file")
	f.Fuzz(func(t *testing.T, args string) {
		for _, re := range validArgs {
			if re.MatchString(args) {
				return
			}
		}

		got, err := NewAppConfig(strings.Split(args, " "))
		if err == nil || got.ExitMessage != "" {
			t.Errorf("NewAppConfig(%v) = %v, want error", args, got)
		}
	})
}
