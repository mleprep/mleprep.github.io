package main

import (
	"strings"

	gen "github.com/traefik/yaegi/stdlib/generic"
)

// Yaegi interprets the generic packages (cmp, maps, slices) from source, but it cannot infer E from a
// constraint like S ~[]E, so most of the standard versions fail ("undefined type for E"), and it leaves
// out slices' sorting and searching. These replacements take []E and map[K]V instead (a named slice
// type has to be converted first) and cover the functions solutions use.
const slicesSource = `package slices

import (
	"cmp"
	"sort"
)

func Sort[E cmp.Ordered](x []E) {
	sort.Slice(x, func(i, j int) bool { return cmp.Less(x[i], x[j]) })
}

func SortFunc[E any](x []E, c func(a, b E) int) {
	sort.Slice(x, func(i, j int) bool { return c(x[i], x[j]) < 0 })
}

func SortStableFunc[E any](x []E, c func(a, b E) int) {
	sort.SliceStable(x, func(i, j int) bool { return c(x[i], x[j]) < 0 })
}

func IsSorted[E cmp.Ordered](x []E) bool {
	for i := len(x) - 1; i > 0; i-- {
		if cmp.Less(x[i], x[i-1]) {
			return false
		}
	}
	return true
}

func IsSortedFunc[E any](x []E, c func(a, b E) int) bool {
	for i := len(x) - 1; i > 0; i-- {
		if c(x[i], x[i-1]) < 0 {
			return false
		}
	}
	return true
}

func Min[E cmp.Ordered](x []E) E {
	if len(x) < 1 {
		panic("slices.Min: empty list")
	}
	m := x[0]
	for i := 1; i < len(x); i++ {
		if x[i] < m {
			m = x[i]
		}
	}
	return m
}

func Max[E cmp.Ordered](x []E) E {
	if len(x) < 1 {
		panic("slices.Max: empty list")
	}
	m := x[0]
	for i := 1; i < len(x); i++ {
		if x[i] > m {
			m = x[i]
		}
	}
	return m
}

func MinFunc[E any](x []E, c func(a, b E) int) E {
	if len(x) < 1 {
		panic("slices.MinFunc: empty list")
	}
	m := x[0]
	for i := 1; i < len(x); i++ {
		if c(x[i], m) < 0 {
			m = x[i]
		}
	}
	return m
}

func MaxFunc[E any](x []E, c func(a, b E) int) E {
	if len(x) < 1 {
		panic("slices.MaxFunc: empty list")
	}
	m := x[0]
	for i := 1; i < len(x); i++ {
		if c(x[i], m) > 0 {
			m = x[i]
		}
	}
	return m
}

func BinarySearch[E cmp.Ordered](x []E, target E) (int, bool) {
	i := sort.Search(len(x), func(i int) bool { return !cmp.Less(x[i], target) })
	return i, i < len(x) && !cmp.Less(target, x[i])
}

func BinarySearchFunc[E, T any](x []E, target T, c func(E, T) int) (int, bool) {
	i := sort.Search(len(x), func(i int) bool { return c(x[i], target) >= 0 })
	return i, i < len(x) && c(x[i], target) == 0
}

func Index[E comparable](s []E, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

func IndexFunc[E any](s []E, f func(E) bool) int {
	for i := range s {
		if f(s[i]) {
			return i
		}
	}
	return -1
}

func Contains[E comparable](s []E, v E) bool { return Index(s, v) >= 0 }

func ContainsFunc[E any](s []E, f func(E) bool) bool { return IndexFunc(s, f) >= 0 }

func Equal[E comparable](s1, s2 []E) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func EqualFunc[E1, E2 any](s1 []E1, s2 []E2, eq func(E1, E2) bool) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if !eq(s1[i], s2[i]) {
			return false
		}
	}
	return true
}

func Compare[E cmp.Ordered](s1, s2 []E) int {
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if c := cmp.Compare(s1[i], s2[i]); c != 0 {
			return c
		}
	}
	return cmp.Compare(len(s1), len(s2))
}

func Reverse[E any](s []E) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Clone[E any](s []E) []E {
	if s == nil {
		return nil
	}
	return append([]E{}, s...)
}

func Compact[E comparable](s []E) []E {
	if len(s) < 2 {
		return s
	}
	k := 1
	for i := 1; i < len(s); i++ {
		if s[i] != s[k-1] {
			s[k] = s[i]
			k++
		}
	}
	return s[:k]
}

func CompactFunc[E any](s []E, eq func(E, E) bool) []E {
	if len(s) < 2 {
		return s
	}
	k := 1
	for i := 1; i < len(s); i++ {
		if !eq(s[k-1], s[i]) {
			s[k] = s[i]
			k++
		}
	}
	return s[:k]
}

func Insert[E any](s []E, i int, v ...E) []E {
	out := make([]E, 0, len(s)+len(v))
	out = append(out, s[:i]...)
	out = append(out, v...)
	return append(out, s[i:]...)
}

func Delete[E any](s []E, i, j int) []E {
	return append(s[:i], s[j:]...)
}

func DeleteFunc[E any](s []E, del func(E) bool) []E {
	k := 0
	for _, v := range s {
		if !del(v) {
			s[k] = v
			k++
		}
	}
	return s[:k]
}

func Replace[E any](s []E, i, j int, v ...E) []E {
	out := make([]E, 0, len(s)-(j-i)+len(v))
	out = append(out, s[:i]...)
	out = append(out, v...)
	return append(out, s[j:]...)
}

func Concat[E any](ss ...[]E) []E {
	var out []E
	for _, s := range ss {
		out = append(out, s...)
	}
	return out
}

func Repeat[E any](x []E, count int) []E {
	out := make([]E, 0, len(x)*count)
	for i := 0; i < count; i++ {
		out = append(out, x...)
	}
	return out
}

func Grow[E any](s []E, n int) []E {
	if cap(s)-len(s) < n {
		s = append(s[:cap(s)], make([]E, n)...)[:len(s)]
	}
	return s
}

func Clip[E any](s []E) []E { return s[:len(s):len(s)] }
`

const mapsSource = `package maps

func Equal[K, V comparable](m1, m2 map[K]V) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func EqualFunc[K comparable, V1, V2 any](m1 map[K]V1, m2 map[K]V2, eq func(V1, V2) bool) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || !eq(v1, v2) {
			return false
		}
	}
	return true
}

func Clone[K comparable, V any](m map[K]V) map[K]V {
	if m == nil {
		return nil
	}
	out := make(map[K]V, len(m))
	for k, v := range m {
		out[k] = v
	}
	return out
}

func Copy[K comparable, V any](dst, src map[K]V) {
	for k, v := range src {
		dst[k] = v
	}
}

func DeleteFunc[K comparable, V any](m map[K]V, del func(K, V) bool) {
	for k, v := range m {
		if del(k, v) {
			delete(m, k)
		}
	}
}
`

func init() {
	for i, s := range gen.Sources {
		switch {
		case strings.Contains(s, "\npackage slices\n"):
			gen.Sources[i] = slicesSource
		case strings.Contains(s, "\npackage maps\n"):
			gen.Sources[i] = mapsSource
		}
	}
}
