package xiter_test

import (
	"iter"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/moeryomenko/xiter"
)

func Test_IterFunc(t *testing.T) {
	testcases := map[string]struct {
		seq  iter.Seq[int]
		fn   func(int) int
		want []int
	}{
		"identity": {
			seq:  slices.Values([]int{1, 2, 3}),
			fn:   func(e int) int { return e },
			want: []int{1, 2, 3},
		},
		"squaring": {
			seq:  slices.Values([]int{1, 2, 3}),
			fn:   func(e int) int { return e * e },
			want: []int{1, 4, 9},
		},
		"empty": {
			seq:  slices.Values([]int{}),
			fn:   func(e int) int { return e },
			want: nil,
		},
	}

	for caseName, tc := range testcases {
		tc := tc

		t.Run(caseName, func(t *testing.T) {
			t.Parallel()

			got := xiter.IterFunc(tc.seq, tc.fn)

			if diff := cmp.Diff(tc.want, slices.Collect(got)); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}

func Test_Filter(t *testing.T) {
	t.Parallel()

	testcases := map[string]struct {
		s         []int
		predicate func(int) bool
		want      []int
	}{
		"empty slice": {
			s:         []int{},
			predicate: func(int) bool { return true },
			want:      []int{},
		},
		"all elements pass": {
			s:         []int{1, 2, 3},
			predicate: func(v int) bool { return true },
			want:      []int{1, 2, 3},
		},
		"some elements pass": {
			s:         []int{1, 2, 3, 4},
			predicate: func(v int) bool { return v%2 == 0 },
			want:      []int{2, 4},
		},
		"no elements pass": {
			s:         []int{1, 3, 5},
			predicate: func(v int) bool { return v%2 == 0 },
			want:      []int{},
		},
		"nil slice": {
			s:         nil,
			predicate: func(int) bool { return true },
			want:      []int{},
		},
	}

	for caseName, tc := range testcases {
		tc := tc

		t.Run(caseName, func(t *testing.T) {
			t.Parallel()

			got := xiter.Filter(tc.s, tc.predicate)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}

func Test_FilterSeq(t *testing.T) {
	t.Parallel()

	testcases := map[string]struct {
		seq       iter.Seq[int]
		predicate func(int) bool
		want      []int
	}{
		"empty sequence": {
			seq:       slices.Values([]int{}),
			predicate: func(int) bool { return true },
			want:      nil,
		},
		"all elements pass": {
			seq:       slices.Values([]int{1, 2, 3}),
			predicate: func(int) bool { return true },
			want:      []int{1, 2, 3},
		},
		"some elements pass": {
			seq:       slices.Values([]int{1, 2, 3, 4}),
			predicate: func(v int) bool { return v%2 == 0 },
			want:      []int{2, 4},
		},
		"no elements pass": {
			seq:       slices.Values([]int{1, 3, 5}),
			predicate: func(v int) bool { return v%2 == 0 },
			want:      nil,
		},
	}

	for caseName, tc := range testcases {
		tc := tc

		t.Run(caseName, func(t *testing.T) {
			t.Parallel()

			got := xiter.FilterSeq(tc.seq, tc.predicate)

			if diff := cmp.Diff(tc.want, slices.Collect(got)); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}

func Test_AppendFunc(t *testing.T) {
	t.Parallel()

	testcases := map[string]struct {
		se   []int
		sv   []int
		fn   func(int) int
		want []int
	}{
		"empty slices": {
			se:   []int{},
			sv:   []int{},
			fn:   func(v int) int { return v },
			want: []int{},
		},
		"append elements": {
			se:   []int{1},
			sv:   []int{2},
			fn:   func(v int) int { return v + 1 },
			want: []int{1, 3},
		},
		"sv is nil": {
			se:   []int{1},
			sv:   nil,
			fn:   func(v int) int { return v },
			want: []int{1},
		},
		"append multiple elements": {
			se:   []int{1, 2},
			sv:   []int{3, 4},
			fn:   func(v int) int { return v },
			want: []int{1, 2, 3, 4},
		},
		"sv is empty": {
			se:   []int{1, 2},
			sv:   []int{},
			fn:   func(v int) int { return v * 2 },
			want: []int{1, 2},
		},
	}

	for caseName, tc := range testcases {
		tc := tc

		t.Run(caseName, func(t *testing.T) {
			t.Parallel()

			got := xiter.AppendFunc(tc.se, tc.sv, tc.fn)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}

func Test_AppendIf(t *testing.T) {
	testcases := map[string]struct {
		s1   []int
		s2   []int
		fn   func(int) bool
		want []int
	}{
		"append all elements": {
			s1:   []int{},
			s2:   []int{1, 2, 3},
			fn:   func(e int) bool { return true },
			want: []int{1, 2, 3},
		},
		"append no elements": {
			s1:   []int{1},
			s2:   []int{1, 2, 3},
			fn:   func(e int) bool { return false },
			want: []int{1},
		},
		"append some elements": {
			s1:   []int{0},
			s2:   []int{1, 2, 3},
			fn:   func(e int) bool { return e > 1 },
			want: []int{0, 2, 3},
		},
	}

	for caseName, tc := range testcases {
		tc := tc

		t.Run(caseName, func(t *testing.T) {
			t.Parallel()

			got := xiter.AppendIf(tc.s1, tc.s2, tc.fn)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}

func Test_Map(t *testing.T) {
	t.Parallel()

	testcases := map[string]struct {
		slice []int
		fn    func(int) string
		want  []string
	}{
		"empty slice": {
			slice: []int{},
			fn:    func(x int) string { return strconv.Itoa(x + 1) },
			want:  nil,
		},
		"non-empty slice": {
			slice: []int{1, 2, 3},
			fn:    func(x int) string { return strconv.Itoa(x + 1) },
			want:  []string{"2", "3", "4"},
		},
		"multiply by two": {
			slice: []int{1, 2, 3},
			fn:    func(x int) string { return strconv.Itoa(x * 2) },
			want:  []string{"2", "4", "6"},
		},
	}

	for caseName, tc := range testcases {
		tc := tc

		t.Run(caseName, func(t *testing.T) {
			t.Parallel()

			got := xiter.Map(tc.slice, tc.fn)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}

func Test_FoldLeft(t *testing.T) {
	t.Parallel()

	t.Run("sum", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}

		fn := func(acc, e int) int { return acc + e }

		want := 15

		got := xiter.FoldLeft(slice, 0, fn)

		if got != want {
			t.Errorf("FoldLeft() = %d, want %d", got, want)
		}
	})

	t.Run("build string", func(t *testing.T) {
		slice := []string{"Hello", " ", "World", "!"}

		fn := func(acc strings.Builder, e string) strings.Builder { acc.WriteString(e); return acc }

		want := "Hello World!"

		got := xiter.FoldLeft(slice, strings.Builder{}, fn)

		if diff := cmp.Diff(want, got.String()); diff != "" {
			t.Errorf("(-want +got):\n%s", diff)
		}
	})
}

func Test_FoldRight(t *testing.T) {
	t.Parallel()

	t.Run("minus", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}

		fn := func(acc, e int) int { return acc - e }

		want := 5

		got := xiter.FoldRight(slice, 20, fn)

		if got != want {
			t.Errorf("FoldLeft() = %d, want %d", got, want)
		}
	})

	t.Run("build string", func(t *testing.T) {
		slice := []string{"Hello", " ", "World", "!"}

		fn := func(acc strings.Builder, e string) strings.Builder { acc.WriteString(e); return acc }

		want := "!World Hello"

		got := xiter.FoldRight(slice, strings.Builder{}, fn)

		if diff := cmp.Diff(want, got.String()); diff != "" {
			t.Errorf("(-want +got):\n%s", diff)
		}
	})
}
