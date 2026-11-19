//go:build ignore

// LeetCode's ListNode / TreeNode plus the small test helpers every solution file uses.
//
//	go run 0001_two_sum.go lc.go
//
// The in-browser judge interprets this same file (tools/gojudge), so a test means the same thing
// locally and on the site.
package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// ListNode is LeetCode's singly-linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode is LeetCode's binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// makeList(1, 2, 3) builds 1 -> 2 -> 3.
func makeList(vals ...int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}

// listVals is the inverse of makeList (it stops after 10,000 nodes in case of a cycle).
func listVals(head *ListNode) []int {
	out := []int{}
	for ; head != nil && len(out) < 10000; head = head.Next {
		out = append(out, head.Val)
	}
	return out
}

// makeTree builds a tree from LeetCode's level order, nil marking a missing child:
// makeTree(3, 9, 20, nil, nil, 15, 7).
func makeTree(vals ...any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	for i := 1; len(queue) > 0 && i < len(vals); i += 2 {
		node := queue[0]
		queue = queue[1:]
		if vals[i] != nil {
			node.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Left)
		}
		if i+1 < len(vals) && vals[i+1] != nil {
			node.Right = &TreeNode{Val: vals[i+1].(int)}
			queue = append(queue, node.Right)
		}
	}
	return root
}

// treeVals is the inverse of makeTree (level order, trailing nils trimmed).
func treeVals(root *TreeNode) []any {
	out := []any{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			out = append(out, nil)
			continue
		}
		out = append(out, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	for len(out) > 0 && out[len(out)-1] == nil {
		out = out[:len(out)-1]
	}
	return out
}

// findNode returns the node holding val (first in pre-order), or nil.
func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if node := findNode(root.Left, val); node != nil {
		return node
	}
	return findNode(root.Right, val)
}

// byteGrid("1.1", "..1") builds the [][]byte boards LeetCode uses for grids.
func byteGrid(rows ...string) [][]byte {
	out := make([][]byte, len(rows))
	for i, r := range rows {
		out[i] = []byte(r)
	}
	return out
}

// assertEq is one test case: got must equal want. Numbers compare by value (floats within 1e-5),
// slices element by element (nil equals empty), lists and trees by their values.
func assertEq(got, want any) { lcCheck(0, got, want) }

// assertUnordered: got holds the same elements as want, in any order ("return the answer in any order").
func assertUnordered(got, want any) { lcCheck(1, got, want) }

// assertUnorderedDeep is assertUnordered where each inner list may be in any order too (e.g. Group Anagrams).
func assertUnorderedDeep(got, want any) { lcCheck(2, got, want) }

// ---------------------------------------------------------------- internals

var lcFails int

// lcReport receives every comparison; the in-browser judge replaces it to collect results.
var lcReport = func(ok bool, got, want string) {
	if !ok {
		lcFails++
		fmt.Println("FAIL: got " + got + ", expected " + want)
	}
}

func lcCheck(mode int, got, want any) {
	g, w := lcCanon(got), lcCanon(want)
	ok := false
	if mode == 0 {
		ok = lcEq(reflect.ValueOf(g), reflect.ValueOf(w))
	} else {
		gs, gok := lcBag(reflect.ValueOf(g), mode == 2)
		ws, wok := lcBag(reflect.ValueOf(w), mode == 2)
		ok = gok && wok && strings.Join(gs, "\n") == strings.Join(ws, "\n")
	}
	lcReport(ok, lcFmt(reflect.ValueOf(g), 0), lcFmt(reflect.ValueOf(w), 0))
}

// lcCanon replaces lists and trees by their values so they compare and print like LeetCode's.
// (Type assertions rather than a type switch: the judge's interpreter mixes up recursive types in switches.)
func lcCanon(v any) any {
	if x, ok := v.(*ListNode); ok {
		return listVals(x)
	}
	if x, ok := v.(*TreeNode); ok {
		return treeVals(x)
	}
	if x, ok := v.([]*ListNode); ok {
		out := make([]any, len(x))
		for i, n := range x {
			out[i] = listVals(n)
		}
		return out
	}
	if x, ok := v.([]*TreeNode); ok {
		out := make([]any, len(x))
		for i, n := range x {
			out[i] = treeVals(n)
		}
		return out
	}
	return v
}

func lcDeref(v reflect.Value) reflect.Value {
	for v.IsValid() && (v.Kind() == reflect.Interface || v.Kind() == reflect.Pointer) {
		if v.IsNil() {
			return reflect.Value{}
		}
		v = v.Elem()
	}
	return v
}

func lcIsList(v reflect.Value) bool { return v.Kind() == reflect.Slice || v.Kind() == reflect.Array }

func lcEmpty(v reflect.Value) bool {
	return !v.IsValid() || ((lcIsList(v) || v.Kind() == reflect.Map) && v.Len() == 0)
}

func lcNum(v reflect.Value) (float64, bool) {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int()), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return float64(v.Uint()), true
	case reflect.Float32, reflect.Float64:
		return v.Float(), true
	}
	return 0, false
}

func lcIsFloat(v reflect.Value) bool { return v.Kind() == reflect.Float32 || v.Kind() == reflect.Float64 }

func lcEq(a, b reflect.Value) bool {
	a, b = lcDeref(a), lcDeref(b)
	if !a.IsValid() || !b.IsValid() {
		return lcEmpty(a) && lcEmpty(b)
	}
	if x, ok := lcNum(a); ok {
		y, ok := lcNum(b)
		if !ok {
			return false
		}
		if lcIsFloat(a) || lcIsFloat(b) {
			return math.Abs(x-y) <= 1e-5*math.Max(1, math.Abs(y)) || (math.IsInf(x, 0) && x == y)
		}
		return lcFmt(a, 0) == lcFmt(b, 0) || x == y && math.Abs(x) < 1<<53 // exact for integers
	}
	switch {
	case lcIsList(a) && lcIsList(b):
		if a.Len() != b.Len() {
			return false
		}
		for i := 0; i < a.Len(); i++ {
			if !lcEq(a.Index(i), b.Index(i)) {
				return false
			}
		}
		return true
	case lcIsList(a) || lcIsList(b):
		return lcEmpty(a) && lcEmpty(b)
	case a.Kind() == reflect.Map && b.Kind() == reflect.Map:
		if a.Len() != b.Len() {
			return false
		}
		ka, kb := lcSortedKeys(a), lcSortedKeys(b)
		for i := range ka {
			if !lcEq(ka[i], kb[i]) || !lcEq(a.MapIndex(ka[i]), b.MapIndex(kb[i])) {
				return false
			}
		}
		return true
	}
	return lcFmt(a, 0) == lcFmt(b, 0)
}

func lcSortedKeys(m reflect.Value) []reflect.Value {
	keys := m.MapKeys()
	sort.Slice(keys, func(i, j int) bool { return lcFmt(keys[i], 0) < lcFmt(keys[j], 0) })
	return keys
}

// lcBag formats the elements of a list (each inner list sorted too when deep), sorted.
func lcBag(v reflect.Value, deep bool) ([]string, bool) {
	v = lcDeref(v)
	if !v.IsValid() {
		return []string{}, true
	}
	if !lcIsList(v) {
		return nil, false
	}
	out := make([]string, v.Len())
	for i := range out {
		e := lcDeref(v.Index(i))
		if inner, ok := lcBag(e, deep); deep && ok && e.IsValid() {
			out[i] = "[" + strings.Join(inner, ",") + "]"
		} else {
			out[i] = lcFmt(e, 0)
		}
	}
	sort.Strings(out)
	return out, true
}

// lcFmt prints a value the way LeetCode shows it: [1,2,3], "abc", null, true, 2.5.
func lcFmt(v reflect.Value, depth int) string {
	v = lcDeref(v)
	if !v.IsValid() {
		return "null"
	}
	if depth > 40 {
		return "…"
	}
	switch v.Kind() {
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint8:
		if c := v.Uint(); c >= 32 && c < 127 {
			return strconv.QuoteRune(rune(c))
		}
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		if f := v.Float(); math.Abs(f) < 1e21 {
			return strconv.FormatFloat(f, 'f', -1, 64)
		}
		return strconv.FormatFloat(v.Float(), 'g', -1, 64)
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Slice, reflect.Array:
		parts := make([]string, v.Len())
		for i := range parts {
			parts[i] = lcFmt(v.Index(i), depth+1)
		}
		return "[" + strings.Join(parts, ",") + "]"
	case reflect.Map:
		keys := lcSortedKeys(v)
		parts := make([]string, len(keys))
		for i, k := range keys {
			parts[i] = lcFmt(k, depth+1) + ":" + lcFmt(v.MapIndex(k), depth+1)
		}
		return "{" + strings.Join(parts, ",") + "}"
	case reflect.Struct:
		parts := make([]string, v.NumField())
		for i := range parts {
			parts[i] = lcFmt(v.Field(i), depth+1)
		}
		return "{" + strings.Join(parts, ",") + "}"
	}
	return fmt.Sprint(v)
}
