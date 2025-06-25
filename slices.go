package xiter

import (
	"iter"
	"slices"
)

// FoldRight returns the result of repeatedly applying fn to the elements of s, from right to left.
func FoldRight[Slice ~[]E, E, V any](s Slice, initial V, fn func(V, E) V) V {
	for i := len(s) - 1; i >= 0; i-- {
		initial = fn(initial, s[i])
	}

	return initial
}

// FoldLeft returns the result of repeatedly applying fn to the elements of s, from left to right.
func FoldLeft[Slice ~[]E, E, V any](s Slice, initial V, fn func(V, E) V) V {
	for _, v := range s {
		initial = fn(initial, v)
	}

	return initial
}

// Map returns a new slice containing the results of applying fn to each element of s.
func Map[E, V any](s []E, fn func(E) V) []V {
	if len(s) == 0 {
		return nil
	}

	return AppendFunc(make([]V, 0, len(s)), s, fn)
}

// MapIf returns a new slice containing the results of applying fn to each element of s, but only if fn returns true.
func MapIf[E, V any](s []E, fn func(E) (V, bool)) []V {
	if len(s) == 0 {
		return nil
	}

	ret := make([]V, 0, len(s))
	for _, v := range s {
		if val, ok := fn(v); ok {
			ret = append(ret, val)
		}
	}

	return ret
}

// IterFunc returns a sequence that yields the results of applying fn to each element of seq.
func IterFunc[E, V any](seq iter.Seq[E], fn func(E) V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if !yield(fn(v)) {
				return
			}
		}
	}
}

// Filter returns a new slice containing the elements of s that satisfy the given predicate.
func Filter[Slice ~[]E, E any](s Slice, predicate func(E) bool) Slice {
	ret := make(Slice, 0, len(s))

	return AppendIf(ret, s, predicate)
}

// FilterSeq returns a sequence that yields the elements of seq that satisfy the given predicate.
func FilterSeq[E any](seq iter.Seq[E], predicate func(E) bool) iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range seq {
			if !predicate(v) {
				continue
			}

			if !yield(v) {
				return
			}
		}
	}
}

// Append appends the results of applying fn to each element of s2 to s1 and returns the result.
func AppendFunc[SliceE ~[]E, SliceV ~[]V, E, V any](se SliceE, sv SliceV, fn func(V) E) SliceE {
	se = slices.Grow(se, len(sv))

	return AppendSeqFunc(se, slices.Values(sv), fn)
}

// AppendIf appends the elements of s2 that satisfy the given predicate to the slice s1.
func AppendIf[Slice ~[]E, E any](s1, s2 Slice, predicate func(E) bool) Slice {
	s1 = slices.Grow(s1, len(s2))

	return AppendSeqIf(s1, slices.Values(s2), predicate)
}

// AppendSeqIf appends the elements of seq that satisfy the given predicate to the slice s.
func AppendSeqIf[Slice ~[]E, E any](s Slice, seq iter.Seq[E], predicate func(E) bool) Slice {
	for v := range seq {
		if predicate(v) {
			s = append(s, v)
		}
	}

	return s
}

// AppendSeqFunc appends the results of applying fn to each element in seq to the slice s.
func AppendSeqFunc[Slice ~[]E, E, V any](s Slice, seq iter.Seq[V], fn func(V) E) Slice {
	for v := range seq {
		s = append(s, fn(v))
	}

	return s
}
