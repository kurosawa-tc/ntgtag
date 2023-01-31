package main

import (
	"testing"
	"time"

	"github.com/Songmu/flextime"
)

var cases = []struct {
	tag    string
	branch string
	date   time.Time
	want   string
}{
	{"", "", time.Date(2023, 1, 21, 12, 0, 0, 0, time.UTC), "20230121-01.000"},
	{"", "main", time.Date(2023, 1, 21, 12, 0, 0, 0, time.UTC), "20230121-01.000"},
	{"", "feature/000", time.Date(2023, 1, 21, 12, 0, 0, 0, time.UTC), "20230121-01.000"},
	{"", "feature/100", time.Date(2023, 1, 21, 12, 0, 0, 0, time.UTC), "20230121-01.100"},
	{"", "feature/100-test", time.Date(2023, 1, 21, 12, 0, 0, 0, time.UTC), "20230121-01.100"},
	{"20230121-01.000", "feature/000", time.Date(2023, 1, 21, 12, 0, 0, 0, time.UTC), "20230121-02.000"},
	{"20230121-09.000", "feature/000", time.Date(2023, 1, 21, 12, 0, 0, 0, time.UTC), "20230121-10.000"},
	{"20230121-99.000", "feature/000", time.Date(2023, 1, 21, 12, 0, 0, 0, time.UTC), "20230121-100.000"},
	{"20230121-01.000", "feature/000", time.Date(2023, 1, 22, 12, 0, 0, 0, time.UTC), "20230122-01.000"},
}

func TestToNtCase(t *testing.T) {
	for _, tt := range cases {
		t.Run("GenGitTag: "+tt.tag+tt.branch, func(t *testing.T) {
			flextime.Fix(tt.date)
			if got := GenGitTag([]string{tt.tag}, tt.branch); got != tt.want {
				t.Errorf("GenGitTag(%#q, %#q) = %#q, want %#q", tt.tag, tt.branch, got, tt.want)
			}
		})
	}
}
