//go:build ignore

// 297. Serialize and Deserialize Binary Tree (Hard)
// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/
// Run: go run 0297_serialize_and_deserialize_binary_tree.go lc.go
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// serialize writes the preorder with "#" for missing children.
func (this *Codec) serialize(root *TreeNode) string {
	out := []string{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			out = append(out, "#")
			return
		}
		out = append(out, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return strings.Join(out, ",")
}

func (this *Codec) deserialize(data string) *TreeNode {
	vals, i := strings.Split(data, ","), 0
	var build func() *TreeNode
	build = func() *TreeNode {
		v := vals[i]
		i++
		if v == "#" {
			return nil
		}
		n, _ := strconv.Atoi(v)
		node := &TreeNode{Val: n}
		node.Left = build()
		node.Right = build()
		return node
	}
	return build()
}

func main() {
	c := Constructor()
	t := makeTree(1, 2, 3, nil, nil, 4, 5)
	assertEq(c.serialize(t), "1,2,#,#,3,4,#,#,5,#,#")
	assertEq(c.deserialize(c.serialize(t)), []any{1, 2, 3, nil, nil, 4, 5})
	assertEq(c.serialize(c.deserialize(c.serialize(nil))), "#")
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
