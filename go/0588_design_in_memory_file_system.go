//go:build ignore

// 588. Design In-Memory File System (Hard)
// https://leetcode.com/problems/design-in-memory-file-system/
// Run: go run 0588_design_in_memory_file_system.go lc.go
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type fsNode struct {
	kids    map[string]*fsNode
	isFile  bool
	content string
}

type FileSystem struct {
	root *fsNode
}

func Constructor() FileSystem {
	return FileSystem{root: &fsNode{kids: map[string]*fsNode{}}}
}

func (this *FileSystem) walk(path string) *fsNode { // creates missing directories on the way
	node := this.root
	for _, part := range strings.Split(path, "/") {
		if part == "" {
			continue
		}
		next, ok := node.kids[part]
		if !ok {
			next = &fsNode{kids: map[string]*fsNode{}}
			node.kids[part] = next
		}
		node = next
	}
	return node
}

func (this *FileSystem) Ls(path string) []string {
	node := this.walk(path)
	if node.isFile {
		return []string{path[strings.LastIndex(path, "/")+1:]}
	}
	names := []string{}
	for name := range node.kids {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (this *FileSystem) Mkdir(path string) {
	this.walk(path)
}

func (this *FileSystem) AddContentToFile(filePath string, content string) {
	node := this.walk(filePath)
	node.isFile = true
	node.content += content
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
	return this.walk(filePath).content
}

func main() {
	fs := Constructor()
	assertEq(fs.Ls("/"), []string{})
	fs.Mkdir("/a/b/c")
	fs.AddContentToFile("/a/b/c/d", "hello")
	assertEq(fs.Ls("/"), []string{"a"})
	assertEq(fs.ReadContentFromFile("/a/b/c/d"), "hello")
	fs.AddContentToFile("/a/b/c/d", " world")
	assertEq(fs.ReadContentFromFile("/a/b/c/d"), "hello world")
	assertEq(fs.Ls("/a/b/c/d"), []string{"d"})
	fs.Mkdir("/a/b/e")
	fs.AddContentToFile("/a/b/aa", "x")
	assertEq(fs.Ls("/a/b"), []string{"aa", "c", "e"})
	if lcFails > 0 {
		fmt.Println(lcFails, "test(s) FAILED")
		os.Exit(1)
	}
	fmt.Println("OK")
}
