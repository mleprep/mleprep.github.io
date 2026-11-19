# Graphs

**Representations:** adjacency list `graph[u] = [v, …]`, or an implicit grid (4 neighbours).

| Need | Tool |
|---|---|
| Connected components / flood fill | DFS or BFS, or Union-Find |
| Shortest path, unweighted | **BFS** (multi-source: seed the queue with all sources) |
| Ordering with dependencies / cycle in DAG | **Topological sort** (Kahn's BFS on in-degrees) |
| Dynamic connectivity, "redundant edge" | **Union-Find** (path compression + union by size) |

```python
# grid BFS
q, seen = deque([(sr, sc)]), {(sr, sc)}
while q:
    r, c = q.popleft()
    for nr, nc in ((r+1,c), (r-1,c), (r,c+1), (r,c-1)):
        if 0 <= nr < R and 0 <= nc < C and (nr, nc) not in seen and grid[nr][nc] == OPEN:
            seen.add((nr, nc)); q.append((nr, nc))

# union-find
parent = list(range(n))
def find(x):
    while parent[x] != x:
        parent[x] = parent[parent[x]]      # path halving
        x = parent[x]
    return x
def union(a, b):
    ra, rb = find(a), find(b)
    if ra == rb: return False
    parent[ra] = rb; return True
```

## 200. Number of Islands | Medium
related: 695, 130, 323, 547
tags: blind75, grid-dfs

Scan the grid; each unvisited `1` starts a new island — flood-fill it (DFS/BFS) to mark the whole island visited.

```viz
1 1 0 0 0
1 1 0 0 0      -> 3 islands
0 0 1 0 0
0 0 0 1 1
```

**Complexity:** O(R·C).

```python
class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        R, C = len(grid), len(grid[0])

        def sink(r, c):
            st = [(r, c)]
            grid[r][c] = "0"
            while st:
                r, c = st.pop()
                for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                    if 0 <= nr < R and 0 <= nc < C and grid[nr][nc] == "1":
                        grid[nr][nc] = "0"
                        st.append((nr, nc))

        count = 0
        for r in range(R):
            for c in range(C):
                if grid[r][c] == "1":
                    count += 1
                    sink(r, c)
        return count
# test
g = [list(r) for r in ["11000", "11000", "00100", "00011"]]
assert Solution().numIslands(g) == 3
assert Solution().numIslands([["0"]]) == 0
assert Solution().numIslands([list(r) for r in ["11110", "11010", "11000", "00000"]]) == 1
assert Solution().numIslands([list(r) for r in ["101", "010", "101"]]) == 5
```

```java
class Solution {
    public int numIslands(char[][] grid) {
        int count = 0;
        for (int r = 0; r < grid.length; r++)
            for (int c = 0; c < grid[0].length; c++)
                if (grid[r][c] == '1') { count++; sink(grid, r, c); }
        return count;
    }

    private void sink(char[][] g, int r, int c) {
        if (r < 0 || c < 0 || r >= g.length || c >= g[0].length || g[r][c] != '1') return;
        g[r][c] = '0';
        sink(g, r + 1, c); sink(g, r - 1, c); sink(g, r, c + 1); sink(g, r, c - 1);
    }
}
// test
char[][] g = Stream.of("11000", "11000", "00100", "00011").map(String::toCharArray).toArray(char[][]::new);
check(new Solution().numIslands(g), 3);
```

```go
func numIslands(grid [][]byte) int {
	R, C := len(grid), len(grid[0])
	sink := func(r, c int) { // flood-fill one island with an explicit stack
		st := [][2]int{{r, c}}
		grid[r][c] = '0'
		for len(st) > 0 {
			p := st[len(st)-1]
			st = st[:len(st)-1]
			for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				nr, nc := p[0]+d[0], p[1]+d[1]
				if nr >= 0 && nr < R && nc >= 0 && nc < C && grid[nr][nc] == '1' {
					grid[nr][nc] = '0'
					st = append(st, [2]int{nr, nc})
				}
			}
		}
	}
	count := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if grid[r][c] == '1' {
				count++
				sink(r, c)
			}
		}
	}
	return count
}
// test
assertEq(numIslands(byteGrid("11000", "11000", "00100", "00011")), 3)
assertEq(numIslands(byteGrid("0")), 0)
assertEq(numIslands(byteGrid("11110", "11010", "11000", "00000")), 1)
assertEq(numIslands(byteGrid("101", "010", "101")), 5)
```

## 695. Max Area of Island | Medium
related: 200
tags: grid-dfs

Same flood fill, but each fill returns the island size; keep the max.

**Complexity:** O(R·C).

```python
class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        R, C = len(grid), len(grid[0])

        def area(r, c):
            if not (0 <= r < R and 0 <= c < C) or grid[r][c] != 1:
                return 0
            grid[r][c] = 0
            return 1 + area(r + 1, c) + area(r - 1, c) + area(r, c + 1) + area(r, c - 1)

        return max(area(r, c) for r in range(R) for c in range(C))
# test
g = [[0, 0, 1, 0, 0], [1, 1, 1, 0, 0], [0, 0, 0, 1, 1], [0, 0, 0, 1, 0]]
assert Solution().maxAreaOfIsland(g) == 4
assert Solution().maxAreaOfIsland([[0, 0, 0, 0, 0]]) == 0
assert Solution().maxAreaOfIsland([[1]]) == 1
```

```java
class Solution {
    public int maxAreaOfIsland(int[][] grid) {
        int best = 0;
        for (int r = 0; r < grid.length; r++)
            for (int c = 0; c < grid[0].length; c++) best = Math.max(best, area(grid, r, c));
        return best;
    }

    private int area(int[][] g, int r, int c) {
        if (r < 0 || c < 0 || r >= g.length || c >= g[0].length || g[r][c] != 1) return 0;
        g[r][c] = 0;
        return 1 + area(g, r + 1, c) + area(g, r - 1, c) + area(g, r, c + 1) + area(g, r, c - 1);
    }
}
// test
check(new Solution().maxAreaOfIsland(new int[][]{{0, 0, 1, 0, 0}, {1, 1, 1, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 0}}), 4);
```

```go
func maxAreaOfIsland(grid [][]int) int {
	R, C := len(grid), len(grid[0])
	var area func(r, c int) int
	area = func(r, c int) int {
		if r < 0 || r >= R || c < 0 || c >= C || grid[r][c] != 1 {
			return 0
		}
		grid[r][c] = 0 // visited
		return 1 + area(r+1, c) + area(r-1, c) + area(r, c+1) + area(r, c-1)
	}
	best := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			best = max(best, area(r, c))
		}
	}
	return best
}
// test
assertEq(maxAreaOfIsland([][]int{{0, 0, 1, 0, 0}, {1, 1, 1, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 0}}), 4)
assertEq(maxAreaOfIsland([][]int{{0, 0, 0, 0, 0}}), 0)
assertEq(maxAreaOfIsland([][]int{{1}}), 1)
```

## 133. Clone Graph | Medium
related: 138
tags: blind75, dfs, hash-map

DFS with a map `original → clone`. Create the clone before recursing into neighbours so cycles resolve through the map.

**Complexity:** O(V + E).

```python
class Node:  # provided by LeetCode
    def __init__(self, val=0, neighbors=None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []

class Solution:
    def cloneGraph(self, node: Optional["Node"]) -> Optional["Node"]:
        clones = {}

        def dfs(n):
            if n in clones:
                return clones[n]
            c = clones[n] = Node(n.val)
            c.neighbors = [dfs(nb) for nb in n.neighbors]
            return c

        return dfs(node) if node else None
# test
ns = [Node(i) for i in range(1, 5)]
for i, nbs in enumerate([[2, 4], [1, 3], [2, 4], [1, 3]]):
    ns[i].neighbors = [ns[j - 1] for j in nbs]
c = Solution().cloneGraph(ns[0])
assert c is not ns[0] and [n.val for n in c.neighbors] == [2, 4]
assert c.neighbors[0].neighbors[0] is c        # cycle preserved
```

```java
class Node { // provided by LeetCode
    public int val;
    public List<Node> neighbors = new ArrayList<>();
    public Node(int val) { this.val = val; }
}

class Solution {
    private final Map<Node, Node> clones = new HashMap<>();

    public Node cloneGraph(Node node) {
        if (node == null) return null;
        if (clones.containsKey(node)) return clones.get(node);
        Node c = new Node(node.val);
        clones.put(node, c);
        for (Node nb : node.neighbors) c.neighbors.add(cloneGraph(nb));
        return c;
    }
}
// test
Node a = new Node(1), b = new Node(2);
a.neighbors.add(b); b.neighbors.add(a);
Node c = new Solution().cloneGraph(a);
check(c != a && c.neighbors.get(0).val == 2 && c.neighbors.get(0).neighbors.get(0) == c, true);
```

```go
type Node struct { // provided by LeetCode
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	clones := map[*Node]*Node{}
	var dfs func(*Node) *Node
	dfs = func(n *Node) *Node {
		if c, ok := clones[n]; ok {
			return c
		}
		c := &Node{Val: n.Val}
		clones[n] = c // register before recursing: the graph has cycles
		for _, nb := range n.Neighbors {
			c.Neighbors = append(c.Neighbors, dfs(nb))
		}
		return c
	}
	if node == nil {
		return nil
	}
	return dfs(node)
}
// test
ns := make([]*Node, 4)
for i := range ns {
	ns[i] = &Node{Val: i + 1}
}
for i, nbs := range [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}} {
	for _, j := range nbs {
		ns[i].Neighbors = append(ns[i].Neighbors, ns[j-1])
	}
}
c := cloneGraph(ns[0])
assertEq(c != ns[0], true)
assertEq([]int{c.Neighbors[0].Val, c.Neighbors[1].Val}, []int{2, 4})
assertEq(c.Neighbors[0].Neighbors[0] == c, true) // cycle preserved
```

## 994. Rotting Oranges | Medium
related: 542, 286
tags: multi-source-bfs

Multi-source BFS from all rotten oranges at once; each BFS layer = 1 minute. At the end, any fresh orange left → −1.

```viz
minute 0     1     2     3     4
2 1 1      2 2 1  2 2 2  2 2 2  2 2 2
1 1 0  ->  2 1 0  2 2 0  2 2 0  2 2 0
0 1 1      0 1 1  0 1 1  0 2 1  0 2 2
```

**Complexity:** O(R·C).

```python
class Solution:
    def orangesRotting(self, grid: List[List[int]]) -> int:
        R, C = len(grid), len(grid[0])
        q = deque((r, c) for r in range(R) for c in range(C) if grid[r][c] == 2)
        fresh = sum(row.count(1) for row in grid)
        minutes = 0
        while q and fresh:
            minutes += 1
            for _ in range(len(q)):
                r, c = q.popleft()
                for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                    if 0 <= nr < R and 0 <= nc < C and grid[nr][nc] == 1:
                        grid[nr][nc] = 2
                        fresh -= 1
                        q.append((nr, nc))
        return -1 if fresh else minutes
# test
assert Solution().orangesRotting([[2, 1, 1], [1, 1, 0], [0, 1, 1]]) == 4
assert Solution().orangesRotting([[2, 1, 1], [0, 1, 1], [1, 0, 1]]) == -1
assert Solution().orangesRotting([[0, 2]]) == 0
```

```java
class Solution {
    public int orangesRotting(int[][] grid) {
        int R = grid.length, C = grid[0].length, fresh = 0, minutes = 0;
        Deque<int[]> q = new ArrayDeque<>();
        for (int r = 0; r < R; r++)
            for (int c = 0; c < C; c++) {
                if (grid[r][c] == 2) q.add(new int[]{r, c});
                else if (grid[r][c] == 1) fresh++;
            }
        int[][] dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        while (!q.isEmpty() && fresh > 0) {
            minutes++;
            for (int i = q.size(); i > 0; i--) {
                int[] cur = q.poll();
                for (int[] d : dirs) {
                    int nr = cur[0] + d[0], nc = cur[1] + d[1];
                    if (nr >= 0 && nc >= 0 && nr < R && nc < C && grid[nr][nc] == 1) {
                        grid[nr][nc] = 2; fresh--; q.add(new int[]{nr, nc});
                    }
                }
            }
        }
        return fresh > 0 ? -1 : minutes;
    }
}
// test
check(new Solution().orangesRotting(new int[][]{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}), 4);
check(new Solution().orangesRotting(new int[][]{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}), -1);
```

```go
func orangesRotting(grid [][]int) int {
	R, C := len(grid), len(grid[0])
	q, fresh := [][2]int{}, 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if grid[r][c] == 2 {
				q = append(q, [2]int{r, c})
			} else if grid[r][c] == 1 {
				fresh++
			}
		}
	}
	minutes := 0
	for len(q) > 0 && fresh > 0 { // multi-source BFS, one level per minute
		minutes++
		next := [][2]int{}
		for _, p := range q {
			for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				nr, nc := p[0]+d[0], p[1]+d[1]
				if nr >= 0 && nr < R && nc >= 0 && nc < C && grid[nr][nc] == 1 {
					grid[nr][nc] = 2
					fresh--
					next = append(next, [2]int{nr, nc})
				}
			}
		}
		q = next
	}
	if fresh > 0 {
		return -1
	}
	return minutes
}
// test
assertEq(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}), 4)
assertEq(orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}), -1)
assertEq(orangesRotting([][]int{{0, 2}}), 0)
```

## 417. Pacific Atlantic Water Flow | Medium
related: 130, 200
tags: blind75, reverse-search

Reverse the flow: from each ocean's border, walk **uphill** (to neighbours ≥ current). Cells reachable from both oceans are the answer.

**Complexity:** O(R·C).

```python
class Solution:
    def pacificAtlantic(self, heights: List[List[int]]) -> List[List[int]]:
        R, C = len(heights), len(heights[0])

        def reach(starts):
            seen, st = set(starts), list(starts)
            while st:
                r, c = st.pop()
                for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                    if (0 <= nr < R and 0 <= nc < C and (nr, nc) not in seen
                            and heights[nr][nc] >= heights[r][c]):
                        seen.add((nr, nc))
                        st.append((nr, nc))
            return seen

        pac = reach([(0, c) for c in range(C)] + [(r, 0) for r in range(R)])
        atl = reach([(R - 1, c) for c in range(C)] + [(r, C - 1) for r in range(R)])
        return sorted([r, c] for r, c in pac & atl)
# test
H = [[1, 2, 2, 3, 5], [3, 2, 3, 4, 4], [2, 4, 5, 3, 1], [6, 7, 1, 4, 5], [5, 1, 1, 2, 4]]
assert sorted(Solution().pacificAtlantic(H)) == sorted([[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]])
assert Solution().pacificAtlantic([[1]]) == [[0, 0]]
assert sorted(Solution().pacificAtlantic([[2, 1], [1, 2]])) == [[0, 0], [0, 1], [1, 0], [1, 1]]
```

```java
class Solution {
    public List<List<Integer>> pacificAtlantic(int[][] heights) {
        int R = heights.length, C = heights[0].length;
        boolean[][] pac = new boolean[R][C], atl = new boolean[R][C];
        for (int r = 0; r < R; r++) { dfs(heights, r, 0, pac); dfs(heights, r, C - 1, atl); }
        for (int c = 0; c < C; c++) { dfs(heights, 0, c, pac); dfs(heights, R - 1, c, atl); }
        List<List<Integer>> res = new ArrayList<>();
        for (int r = 0; r < R; r++)
            for (int c = 0; c < C; c++)
                if (pac[r][c] && atl[r][c]) res.add(List.of(r, c));
        return res;
    }

    private void dfs(int[][] h, int r, int c, boolean[][] seen) {
        if (seen[r][c]) return;
        seen[r][c] = true;
        int[][] dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        for (int[] d : dirs) {
            int nr = r + d[0], nc = c + d[1];
            if (nr >= 0 && nc >= 0 && nr < h.length && nc < h[0].length && h[nr][nc] >= h[r][c]) dfs(h, nr, nc, seen);
        }
    }
}
// test
int[][] H = {{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}};
check(new Solution().pacificAtlantic(H), "[[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]]");
```

```go
func pacificAtlantic(heights [][]int) [][]int {
	R, C := len(heights), len(heights[0])
	reach := func(starts [][2]int) [][]bool { // cells that can flow into this ocean (walk uphill from it)
		seen := make([][]bool, R)
		for i := range seen {
			seen[i] = make([]bool, C)
		}
		st := starts
		for _, p := range starts {
			seen[p[0]][p[1]] = true
		}
		for len(st) > 0 {
			p := st[len(st)-1]
			st = st[:len(st)-1]
			for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				nr, nc := p[0]+d[0], p[1]+d[1]
				if nr >= 0 && nr < R && nc >= 0 && nc < C && !seen[nr][nc] && heights[nr][nc] >= heights[p[0]][p[1]] {
					seen[nr][nc] = true
					st = append(st, [2]int{nr, nc})
				}
			}
		}
		return seen
	}
	var pacStart, atlStart [][2]int
	for c := 0; c < C; c++ {
		pacStart = append(pacStart, [2]int{0, c})
		atlStart = append(atlStart, [2]int{R - 1, c})
	}
	for r := 0; r < R; r++ {
		pacStart = append(pacStart, [2]int{r, 0})
		atlStart = append(atlStart, [2]int{r, C - 1})
	}
	pac, atl := reach(pacStart), reach(atlStart)
	res := [][]int{}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if pac[r][c] && atl[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}
	return res
}
// test
H := [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}
assertUnordered(pacificAtlantic(H), [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}})
assertEq(pacificAtlantic([][]int{{1}}), [][]int{{0, 0}})
assertUnordered(pacificAtlantic([][]int{{2, 1}, {1, 2}}), [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}})
```

## 130. Surrounded Regions | Medium
related: 200, 417
tags: border-dfs

An `O` survives only if connected to the border. Mark border-connected `O`s as safe (`S`) via DFS, then flip remaining `O → X` and `S → O`.

**Complexity:** O(R·C).

```python
class Solution:
    def solve(self, board: List[List[str]]) -> None:
        R, C = len(board), len(board[0])

        def mark(r, c):
            if 0 <= r < R and 0 <= c < C and board[r][c] == "O":
                board[r][c] = "S"
                mark(r + 1, c); mark(r - 1, c); mark(r, c + 1); mark(r, c - 1)

        for r in range(R):
            mark(r, 0); mark(r, C - 1)
        for c in range(C):
            mark(0, c); mark(R - 1, c)
        for r in range(R):
            for c in range(C):
                board[r][c] = "O" if board[r][c] == "S" else "X"
# test
b = [list(r) for r in ["XXXX", "XOOX", "XXOX", "XOXX"]]
Solution().solve(b)
assert ["".join(r) for r in b] == ["XXXX", "XXXX", "XXXX", "XOXX"]
b = [["X"]]; Solution().solve(b)
assert b == [["X"]]
b = [list(r) for r in ["XXX", "XOX", "XXX"]]; Solution().solve(b)
assert ["".join(r) for r in b] == ["XXX", "XXX", "XXX"]
b = [list(r) for r in ["OO", "OO"]]; Solution().solve(b)
assert ["".join(r) for r in b] == ["OO", "OO"]
```

```java
class Solution {
    public void solve(char[][] board) {
        int R = board.length, C = board[0].length;
        for (int r = 0; r < R; r++) { mark(board, r, 0); mark(board, r, C - 1); }
        for (int c = 0; c < C; c++) { mark(board, 0, c); mark(board, R - 1, c); }
        for (char[] row : board)
            for (int c = 0; c < C; c++) row[c] = row[c] == 'S' ? 'O' : 'X';
    }

    private void mark(char[][] b, int r, int c) {
        if (r < 0 || c < 0 || r >= b.length || c >= b[0].length || b[r][c] != 'O') return;
        b[r][c] = 'S';
        mark(b, r + 1, c); mark(b, r - 1, c); mark(b, r, c + 1); mark(b, r, c - 1);
    }
}
// test
char[][] b = Stream.of("XXXX", "XOOX", "XXOX", "XOXX").map(String::toCharArray).toArray(char[][]::new);
new Solution().solve(b);
check(new String(b[1]) + new String(b[3]), "XXXXXOXX");
```

```go
func solve(board [][]byte) {
	R, C := len(board), len(board[0])
	var mark func(r, c int) // "O"s connected to the border are safe
	mark = func(r, c int) {
		if r >= 0 && r < R && c >= 0 && c < C && board[r][c] == 'O' {
			board[r][c] = 'S'
			mark(r+1, c)
			mark(r-1, c)
			mark(r, c+1)
			mark(r, c-1)
		}
	}
	for r := 0; r < R; r++ {
		mark(r, 0)
		mark(r, C-1)
	}
	for c := 0; c < C; c++ {
		mark(0, c)
		mark(R-1, c)
	}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if board[r][c] == 'S' {
				board[r][c] = 'O'
			} else {
				board[r][c] = 'X'
			}
		}
	}
}
// test
b := byteGrid("XXXX", "XOOX", "XXOX", "XOXX")
solve(b)
assertEq(b, byteGrid("XXXX", "XXXX", "XXXX", "XOXX"))
b2 := byteGrid("X")
solve(b2)
assertEq(b2, byteGrid("X"))
b3 := byteGrid("XXX", "XOX", "XXX")
solve(b3)
assertEq(b3, byteGrid("XXX", "XXX", "XXX"))
b4 := byteGrid("OO", "OO")
solve(b4)
assertEq(b4, byteGrid("OO", "OO"))
```

## 286. Walls and Gates | Medium
related: 994, 542
tags: multi-source-bfs, premium

Multi-source BFS from every gate (0). The first time BFS reaches an empty room is its shortest distance.

**Complexity:** O(R·C).

```python
class Solution:
    def wallsAndGates(self, rooms: List[List[int]]) -> None:
        INF = 2147483647
        R, C = len(rooms), len(rooms[0])
        q = deque((r, c) for r in range(R) for c in range(C) if rooms[r][c] == 0)
        while q:
            r, c = q.popleft()
            for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                if 0 <= nr < R and 0 <= nc < C and rooms[nr][nc] == INF:
                    rooms[nr][nc] = rooms[r][c] + 1
                    q.append((nr, nc))
# test
I = 2147483647
g = [[I, -1, 0, I], [I, I, I, -1], [I, -1, I, -1], [0, -1, I, I]]
Solution().wallsAndGates(g)
assert g == [[3, -1, 0, 1], [2, 2, 1, -1], [1, -1, 2, -1], [0, -1, 3, 4]]
g = [[-1]]; Solution().wallsAndGates(g)
assert g == [[-1]]
g = [[I, 0]]; Solution().wallsAndGates(g)
assert g == [[1, 0]]
```

```java
class Solution {
    public void wallsAndGates(int[][] rooms) {
        int R = rooms.length, C = rooms[0].length, INF = Integer.MAX_VALUE;
        Deque<int[]> q = new ArrayDeque<>();
        for (int r = 0; r < R; r++)
            for (int c = 0; c < C; c++) if (rooms[r][c] == 0) q.add(new int[]{r, c});
        int[][] dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        while (!q.isEmpty()) {
            int[] cur = q.poll();
            for (int[] d : dirs) {
                int nr = cur[0] + d[0], nc = cur[1] + d[1];
                if (nr >= 0 && nc >= 0 && nr < R && nc < C && rooms[nr][nc] == INF) {
                    rooms[nr][nc] = rooms[cur[0]][cur[1]] + 1;
                    q.add(new int[]{nr, nc});
                }
            }
        }
    }
}
// test
int I = Integer.MAX_VALUE;
int[][] g = {{I, -1, 0, I}, {I, I, I, -1}, {I, -1, I, -1}, {0, -1, I, I}};
new Solution().wallsAndGates(g);
check(g, new int[][]{{3, -1, 0, 1}, {2, 2, 1, -1}, {1, -1, 2, -1}, {0, -1, 3, 4}});
```

```go
func wallsAndGates(rooms [][]int) {
	const INF = 2147483647
	R, C := len(rooms), len(rooms[0])
	q := [][2]int{}
	for r := 0; r < R; r++ { // BFS from every gate at once
		for c := 0; c < C; c++ {
			if rooms[r][c] == 0 {
				q = append(q, [2]int{r, c})
			}
		}
	}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			nr, nc := p[0]+d[0], p[1]+d[1]
			if nr >= 0 && nr < R && nc >= 0 && nc < C && rooms[nr][nc] == INF {
				rooms[nr][nc] = rooms[p[0]][p[1]] + 1
				q = append(q, [2]int{nr, nc})
			}
		}
	}
}
// test
const I = 2147483647
g := [][]int{{I, -1, 0, I}, {I, I, I, -1}, {I, -1, I, -1}, {0, -1, I, I}}
wallsAndGates(g)
assertEq(g, [][]int{{3, -1, 0, 1}, {2, 2, 1, -1}, {1, -1, 2, -1}, {0, -1, 3, 4}})
g2 := [][]int{{-1}}
wallsAndGates(g2)
assertEq(g2, [][]int{{-1}})
g3 := [][]int{{I, 0}}
wallsAndGates(g3)
assertEq(g3, [][]int{{1, 0}})
```

## 207. Course Schedule | Medium
related: 210, 269, 802
tags: blind75, topological-sort

Courses are a directed graph; finishable ⇔ no cycle. **Kahn's algorithm**: repeatedly take nodes with in-degree 0; if all n get taken, there's no cycle.

```viz
[[1,0]]        edge 0 -> 1     indeg 0:0 1:1
               queue [0] -> take 0 -> indeg[1]=0 -> take 1    taken 2 == n  ✓
[[1,0],[0,1]]  indeg 1,1 -> queue empty -> taken 0 < 2        ✗ cycle
```

**Complexity:** O(V + E).

```python
class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        g, indeg = defaultdict(list), [0] * numCourses
        for a, b in prerequisites:
            g[b].append(a)
            indeg[a] += 1
        q = deque(i for i in range(numCourses) if indeg[i] == 0)
        taken = 0
        while q:
            u = q.popleft()
            taken += 1
            for v in g[u]:
                indeg[v] -= 1
                if indeg[v] == 0:
                    q.append(v)
        return taken == numCourses
# test
assert Solution().canFinish(2, [[1, 0]]) is True
assert Solution().canFinish(2, [[1, 0], [0, 1]]) is False
```

```java
class Solution {
    public boolean canFinish(int numCourses, int[][] prerequisites) {
        List<List<Integer>> g = new ArrayList<>();
        for (int i = 0; i < numCourses; i++) g.add(new ArrayList<>());
        int[] indeg = new int[numCourses];
        for (int[] p : prerequisites) { g.get(p[1]).add(p[0]); indeg[p[0]]++; }
        Deque<Integer> q = new ArrayDeque<>();
        for (int i = 0; i < numCourses; i++) if (indeg[i] == 0) q.add(i);
        int taken = 0;
        while (!q.isEmpty()) {
            int u = q.poll();
            taken++;
            for (int v : g.get(u)) if (--indeg[v] == 0) q.add(v);
        }
        return taken == numCourses;
    }
}
// test
check(new Solution().canFinish(2, new int[][]{{1, 0}}), true);
check(new Solution().canFinish(2, new int[][]{{1, 0}, {0, 1}}), false);
```

```go
func canFinish(numCourses int, prerequisites [][]int) bool {
	g, indeg := make([][]int, numCourses), make([]int, numCourses)
	for _, p := range prerequisites {
		g[p[1]] = append(g[p[1]], p[0])
		indeg[p[0]]++
	}
	q := []int{}
	for i, d := range indeg {
		if d == 0 {
			q = append(q, i)
		}
	}
	taken := 0
	for len(q) > 0 { // Kahn's algorithm: a cycle leaves courses with indegree > 0
		u := q[0]
		q = q[1:]
		taken++
		for _, v := range g[u] {
			if indeg[v]--; indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return taken == numCourses
}
// test
assertEq(canFinish(2, [][]int{{1, 0}}), true)
assertEq(canFinish(2, [][]int{{1, 0}, {0, 1}}), false)
```

## 210. Course Schedule II | Medium
related: 207, 269
tags: topological-sort

Same Kahn's algorithm, but record the order nodes are taken. If fewer than n are taken → cycle → `[]`.

**Complexity:** O(V + E).

```python
class Solution:
    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        g, indeg = defaultdict(list), [0] * numCourses
        for a, b in prerequisites:
            g[b].append(a)
            indeg[a] += 1
        q = deque(i for i in range(numCourses) if indeg[i] == 0)
        order = []
        while q:
            u = q.popleft()
            order.append(u)
            for v in g[u]:
                indeg[v] -= 1
                if indeg[v] == 0:
                    q.append(v)
        return order if len(order) == numCourses else []
# test
# any valid topological order is accepted
def valid(n, pre, order):
    pos = {c: i for i, c in enumerate(order)}
    return len(order) == n == len(pos) and all(pos[b] < pos[a] for a, b in pre)
assert valid(4, [[1, 0], [2, 0], [3, 1], [3, 2]], Solution().findOrder(4, [[1, 0], [2, 0], [3, 1], [3, 2]]))
assert valid(2, [[1, 0]], Solution().findOrder(2, [[1, 0]]))
assert Solution().findOrder(2, [[0, 1], [1, 0]]) == []
```

```java
class Solution {
    public int[] findOrder(int numCourses, int[][] prerequisites) {
        List<List<Integer>> g = new ArrayList<>();
        for (int i = 0; i < numCourses; i++) g.add(new ArrayList<>());
        int[] indeg = new int[numCourses], order = new int[numCourses];
        for (int[] p : prerequisites) { g.get(p[1]).add(p[0]); indeg[p[0]]++; }
        Deque<Integer> q = new ArrayDeque<>();
        for (int i = 0; i < numCourses; i++) if (indeg[i] == 0) q.add(i);
        int k = 0;
        while (!q.isEmpty()) {
            int u = q.poll();
            order[k++] = u;
            for (int v : g.get(u)) if (--indeg[v] == 0) q.add(v);
        }
        return k == numCourses ? order : new int[0];
    }
}
// test
check(new Solution().findOrder(4, new int[][]{{1, 0}, {2, 0}, {3, 1}, {3, 2}}), new int[]{0, 1, 2, 3});
```

```go
func findOrder(numCourses int, prerequisites [][]int) []int {
	g, indeg := make([][]int, numCourses), make([]int, numCourses)
	for _, p := range prerequisites {
		g[p[1]] = append(g[p[1]], p[0])
		indeg[p[0]]++
	}
	q := []int{}
	for i, d := range indeg {
		if d == 0 {
			q = append(q, i)
		}
	}
	order := []int{}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		order = append(order, u)
		for _, v := range g[u] {
			if indeg[v]--; indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	if len(order) < numCourses {
		return []int{}
	}
	return order
}
// test
// any valid topological order is accepted
valid := func(n int, pre [][]int, order []int) bool {
	pos := map[int]int{}
	for i, c := range order {
		pos[c] = i
	}
	if len(order) != n || len(pos) != n {
		return false
	}
	for _, p := range pre {
		if pos[p[1]] >= pos[p[0]] {
			return false
		}
	}
	return true
}
pre := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
assertEq(valid(4, pre, findOrder(4, pre)), true)
assertEq(valid(2, [][]int{{1, 0}}, findOrder(2, [][]int{{1, 0}})), true)
assertEq(findOrder(2, [][]int{{0, 1}, {1, 0}}), []int{})
```

## 684. Redundant Connection | Medium
related: 323, 261, 685
tags: union-find

Process edges in order with Union-Find; the first edge whose endpoints are already connected closes a cycle → return it.

**Complexity:** O(E · α(V)).

```python
class Solution:
    def findRedundantConnection(self, edges: List[List[int]]) -> List[int]:
        parent = list(range(len(edges) + 1))

        def find(x):
            while parent[x] != x:
                parent[x] = parent[parent[x]]
                x = parent[x]
            return x

        for a, b in edges:
            ra, rb = find(a), find(b)
            if ra == rb:
                return [a, b]
            parent[ra] = rb
        return []
# test
assert Solution().findRedundantConnection([[1, 2], [1, 3], [2, 3]]) == [2, 3]
assert Solution().findRedundantConnection([[1, 2], [2, 3], [3, 4], [1, 4], [1, 5]]) == [1, 4]
```

```java
class Solution {
    private int[] parent;

    private int find(int x) {
        while (parent[x] != x) { parent[x] = parent[parent[x]]; x = parent[x]; }
        return x;
    }

    public int[] findRedundantConnection(int[][] edges) {
        parent = new int[edges.length + 1];
        for (int i = 0; i < parent.length; i++) parent[i] = i;
        for (int[] e : edges) {
            int ra = find(e[0]), rb = find(e[1]);
            if (ra == rb) return e;
            parent[ra] = rb;
        }
        return new int[0];
    }
}
// test
check(new Solution().findRedundantConnection(new int[][]{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}), new int[]{1, 4});
```

```go
func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	for i := range parent {
		parent[i] = i
	}
	find := func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]] // path halving
			x = parent[x]
		}
		return x
	}
	for _, e := range edges {
		ra, rb := find(e[0]), find(e[1])
		if ra == rb { // already connected: this edge closes the cycle
			return e
		}
		parent[ra] = rb
	}
	return nil
}
// test
assertEq(findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}), []int{2, 3})
assertEq(findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}), []int{1, 4})
```

## 323. Number of Connected Components in an Undirected Graph | Medium
related: 200, 547, 684
tags: blind75, union-find, premium

Start with n components; every successful union merges two → decrement.

**Complexity:** O(E · α(V)).

```python
class Solution:
    def countComponents(self, n: int, edges: List[List[int]]) -> int:
        parent = list(range(n))

        def find(x):
            while parent[x] != x:
                parent[x] = parent[parent[x]]
                x = parent[x]
            return x

        comps = n
        for a, b in edges:
            ra, rb = find(a), find(b)
            if ra != rb:
                parent[ra] = rb
                comps -= 1
        return comps
# test
assert Solution().countComponents(5, [[0, 1], [1, 2], [3, 4]]) == 2
assert Solution().countComponents(5, [[0, 1], [1, 2], [2, 3], [3, 4]]) == 1
```

```java
class Solution {
    public int countComponents(int n, int[][] edges) {
        int[] parent = new int[n];
        for (int i = 0; i < n; i++) parent[i] = i;
        int comps = n;
        for (int[] e : edges) {
            int ra = find(parent, e[0]), rb = find(parent, e[1]);
            if (ra != rb) { parent[ra] = rb; comps--; }
        }
        return comps;
    }

    private int find(int[] p, int x) {
        while (p[x] != x) { p[x] = p[p[x]]; x = p[x]; }
        return x;
    }
}
// test
check(new Solution().countComponents(5, new int[][]{{0, 1}, {1, 2}, {3, 4}}), 2);
```

```go
func countComponents(n int, edges [][]int) int {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	find := func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	comps := n
	for _, e := range edges {
		if ra, rb := find(e[0]), find(e[1]); ra != rb {
			parent[ra] = rb
			comps-- // each successful union merges two components
		}
	}
	return comps
}
// test
assertEq(countComponents(5, [][]int{{0, 1}, {1, 2}, {3, 4}}), 2)
assertEq(countComponents(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}), 1)
```

## 261. Graph Valid Tree | Medium
related: 684, 323
tags: blind75, union-find, premium

A tree on n nodes has exactly `n − 1` edges **and** no cycle (equivalently: connected). Union-Find: any union of already-connected nodes → cycle.

**Complexity:** O(E · α(V)).

```python
class Solution:
    def validTree(self, n: int, edges: List[List[int]]) -> bool:
        if len(edges) != n - 1:
            return False
        parent = list(range(n))

        def find(x):
            while parent[x] != x:
                parent[x] = parent[parent[x]]
                x = parent[x]
            return x

        for a, b in edges:
            ra, rb = find(a), find(b)
            if ra == rb:
                return False
            parent[ra] = rb
        return True
# test
assert Solution().validTree(5, [[0, 1], [0, 2], [0, 3], [1, 4]]) is True
assert Solution().validTree(5, [[0, 1], [1, 2], [2, 3], [1, 3], [1, 4]]) is False
```

```java
class Solution {
    public boolean validTree(int n, int[][] edges) {
        if (edges.length != n - 1) return false;
        int[] parent = new int[n];
        for (int i = 0; i < n; i++) parent[i] = i;
        for (int[] e : edges) {
            int ra = find(parent, e[0]), rb = find(parent, e[1]);
            if (ra == rb) return false;
            parent[ra] = rb;
        }
        return true;
    }

    private int find(int[] p, int x) {
        while (p[x] != x) { p[x] = p[p[x]]; x = p[x]; }
        return x;
    }
}
// test
check(new Solution().validTree(5, new int[][]{{0, 1}, {0, 2}, {0, 3}, {1, 4}}), true);
check(new Solution().validTree(5, new int[][]{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}), false);
```

```go
func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 { // a tree has exactly n-1 edges
		return false
	}
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	find := func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	for _, e := range edges {
		ra, rb := find(e[0]), find(e[1])
		if ra == rb { // ...and no cycle
			return false
		}
		parent[ra] = rb
	}
	return true
}
// test
assertEq(validTree(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}), true)
assertEq(validTree(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}), false)
```

## 547. Number of Provinces | Medium
related: 200, 323
tags: dfs, union-find

Adjacency-matrix connected components: DFS from each unvisited city, count starts.

**Complexity:** O(n²).

```python
class Solution:
    def findCircleNum(self, isConnected: List[List[int]]) -> int:
        n, seen = len(isConnected), set()

        def dfs(u):
            for v in range(n):
                if isConnected[u][v] and v not in seen:
                    seen.add(v)
                    dfs(v)

        count = 0
        for u in range(n):
            if u not in seen:
                seen.add(u)
                dfs(u)
                count += 1
        return count
# test
assert Solution().findCircleNum([[1, 1, 0], [1, 1, 0], [0, 0, 1]]) == 2
assert Solution().findCircleNum([[1, 0, 0], [0, 1, 0], [0, 0, 1]]) == 3
```

```java
class Solution {
    public int findCircleNum(int[][] isConnected) {
        int n = isConnected.length, count = 0;
        boolean[] seen = new boolean[n];
        for (int u = 0; u < n; u++)
            if (!seen[u]) { count++; dfs(isConnected, u, seen); }
        return count;
    }

    private void dfs(int[][] g, int u, boolean[] seen) {
        seen[u] = true;
        for (int v = 0; v < g.length; v++) if (g[u][v] == 1 && !seen[v]) dfs(g, v, seen);
    }
}
// test
check(new Solution().findCircleNum(new int[][]{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}), 2);
```

```go
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	seen := make([]bool, n)
	var dfs func(u int)
	dfs = func(u int) {
		for v := 0; v < n; v++ {
			if isConnected[u][v] == 1 && !seen[v] {
				seen[v] = true
				dfs(v)
			}
		}
	}
	count := 0
	for u := 0; u < n; u++ {
		if !seen[u] {
			seen[u] = true
			dfs(u)
			count++
		}
	}
	return count
}
// test
assertEq(findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}), 2)
assertEq(findCircleNum([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}), 3)
```

## 127. Word Ladder | Hard
related: 126, 433
tags: bfs, implicit-graph

Words are nodes; edges connect words differing by one letter. BFS from `beginWord`. Generate neighbours by trying all 26 letters at each position (O(26·L) per word) and only keep words in the dictionary set.

```viz
hit -> hot -> dot -> dog -> cog      length 5
             \-> lot -> log -/
```

**Complexity:** O(N · L · 26).

```python
class Solution:
    def ladderLength(self, beginWord: str, endWord: str, wordList: List[str]) -> int:
        words = set(wordList)
        if endWord not in words:
            return 0
        q, steps = deque([beginWord]), 1
        words.discard(beginWord)
        while q:
            for _ in range(len(q)):
                w = q.popleft()
                if w == endWord:
                    return steps
                for i in range(len(w)):
                    for ch in string.ascii_lowercase:
                        nxt = w[:i] + ch + w[i + 1:]
                        if nxt in words:
                            words.remove(nxt)
                            q.append(nxt)
            steps += 1
        return 0
# test
assert Solution().ladderLength("hit", "cog", ["hot", "dot", "dog", "lot", "log", "cog"]) == 5
assert Solution().ladderLength("hit", "cog", ["hot", "dot", "dog", "lot", "log"]) == 0
```

```java
class Solution {
    public int ladderLength(String beginWord, String endWord, List<String> wordList) {
        Set<String> words = new HashSet<>(wordList);
        if (!words.contains(endWord)) return 0;
        Deque<String> q = new ArrayDeque<>(List.of(beginWord));
        for (int steps = 1; !q.isEmpty(); steps++) {
            for (int k = q.size(); k > 0; k--) {
                String w = q.poll();
                if (w.equals(endWord)) return steps;
                char[] cs = w.toCharArray();
                for (int i = 0; i < cs.length; i++) {
                    char orig = cs[i];
                    for (char ch = 'a'; ch <= 'z'; ch++) {
                        cs[i] = ch;
                        String nxt = new String(cs);
                        if (words.remove(nxt)) q.add(nxt);
                    }
                    cs[i] = orig;
                }
            }
        }
        return 0;
    }
}
// test
check(new Solution().ladderLength("hit", "cog", Arrays.asList("hot", "dot", "dog", "lot", "log", "cog")), 5);
check(new Solution().ladderLength("hit", "cog", Arrays.asList("hot", "dot", "dog", "lot", "log")), 0);
```

```go
func ladderLength(beginWord string, endWord string, wordList []string) int {
	words := map[string]bool{}
	for _, w := range wordList {
		words[w] = true
	}
	if !words[endWord] {
		return 0
	}
	delete(words, beginWord)
	q := []string{beginWord}
	for steps := 1; len(q) > 0; steps++ { // BFS, one level per step
		next := []string{}
		for _, w := range q {
			if w == endWord {
				return steps
			}
			b := []byte(w)
			for i := range b {
				orig := b[i]
				for ch := byte('a'); ch <= 'z'; ch++ {
					b[i] = ch
					if s := string(b); words[s] {
						delete(words, s) // visited
						next = append(next, s)
					}
				}
				b[i] = orig
			}
		}
		q = next
	}
	return 0
}
// test
assertEq(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}), 5)
assertEq(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}), 0)
```

## 399. Evaluate Division | Medium
related: 990
tags: weighted-graph, dfs

`a / b = 2` ⇒ edges `a→b (2)` and `b→a (0.5)`. A query `x / y` = product of weights along any path from x to y (DFS/BFS). Unknown variable or no path → −1.

**Complexity:** O(Q · (V + E)).

```python
class Solution:
    def calcEquation(self, equations: List[List[str]], values: List[float], queries: List[List[str]]) -> List[float]:
        g = defaultdict(dict)
        for (a, b), v in zip(equations, values):
            g[a][b] = v
            g[b][a] = 1 / v

        def dfs(x, y, seen):
            if x == y:
                return 1.0
            seen.add(x)
            for nb, w in g[x].items():
                if nb not in seen:
                    r = dfs(nb, y, seen)
                    if r != -1:
                        return w * r
            return -1.0

        return [dfs(x, y, set()) if x in g and y in g else -1.0 for x, y in queries]
# test
eq, val = [["a", "b"], ["b", "c"]], [2.0, 3.0]
qs = [["a", "c"], ["b", "a"], ["a", "e"], ["a", "a"], ["x", "x"]]
assert Solution().calcEquation(eq, val, qs) == [6.0, 0.5, -1.0, 1.0, -1.0]
got = Solution().calcEquation([["a", "b"], ["b", "c"], ["bc", "cd"]], [1.5, 2.5, 5.0], [["a", "c"], ["c", "b"], ["bc", "cd"], ["cd", "bc"]])
assert all(abs(x - y) < 1e-5 for x, y in zip(got, [3.75, 0.4, 5.0, 0.2])) and len(got) == 4
assert Solution().calcEquation([["a", "b"]], [0.5], [["a", "b"], ["b", "a"], ["a", "c"], ["x", "y"]]) == [0.5, 2.0, -1.0, -1.0]
```

```java
class Solution {
    public double[] calcEquation(List<List<String>> equations, double[] values, List<List<String>> queries) {
        Map<String, Map<String, Double>> g = new HashMap<>();
        for (int i = 0; i < values.length; i++) {
            String a = equations.get(i).get(0), b = equations.get(i).get(1);
            g.computeIfAbsent(a, k -> new HashMap<>()).put(b, values[i]);
            g.computeIfAbsent(b, k -> new HashMap<>()).put(a, 1 / values[i]);
        }
        double[] res = new double[queries.size()];
        for (int i = 0; i < res.length; i++) {
            String x = queries.get(i).get(0), y = queries.get(i).get(1);
            res[i] = g.containsKey(x) && g.containsKey(y) ? dfs(g, x, y, new HashSet<>()) : -1.0;
        }
        return res;
    }

    private double dfs(Map<String, Map<String, Double>> g, String x, String y, Set<String> seen) {
        if (x.equals(y)) return 1.0;
        seen.add(x);
        for (Map.Entry<String, Double> e : g.get(x).entrySet()) {
            if (seen.contains(e.getKey())) continue;
            double r = dfs(g, e.getKey(), y, seen);
            if (r != -1.0) return e.getValue() * r;
        }
        return -1.0;
    }
}
// test
List<List<String>> eq = List.of(List.of("a", "b"), List.of("b", "c"));
List<List<String>> qs = List.of(List.of("a", "c"), List.of("b", "a"), List.of("a", "e"), List.of("a", "a"), List.of("x", "x"));
check(new Solution().calcEquation(eq, new double[]{2.0, 3.0}, qs), new double[]{6.0, 0.5, -1.0, 1.0, -1.0});
```

```go
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	g := map[string]map[string]float64{} // a -> b -> a/b
	for i, e := range equations {
		a, b := e[0], e[1]
		if g[a] == nil {
			g[a] = map[string]float64{}
		}
		if g[b] == nil {
			g[b] = map[string]float64{}
		}
		g[a][b] = values[i]
		g[b][a] = 1 / values[i]
	}
	var dfs func(x, y string, seen map[string]bool) float64
	dfs = func(x, y string, seen map[string]bool) float64 { // product of weights along a path, or -1
		if x == y {
			return 1
		}
		seen[x] = true
		for nb, w := range g[x] {
			if !seen[nb] {
				if r := dfs(nb, y, seen); r != -1 {
					return w * r
				}
			}
		}
		return -1
	}
	res := make([]float64, len(queries))
	for i, q := range queries {
		if g[q[0]] == nil || g[q[1]] == nil {
			res[i] = -1
		} else {
			res[i] = dfs(q[0], q[1], map[string]bool{})
		}
	}
	return res
}
// test
eq, val := [][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0}
qs := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
assertEq(calcEquation(eq, val, qs), []float64{6.0, 0.5, -1.0, 1.0, -1.0})
assertEq(calcEquation([][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}, []float64{1.5, 2.5, 5.0}, [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}), []float64{3.75, 0.4, 5.0, 0.2})
assertEq(calcEquation([][]string{{"a", "b"}}, []float64{0.5}, [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}}), []float64{0.5, 2.0, -1.0, -1.0})
```

## 785. Is Graph Bipartite? | Medium
related: 886
tags: bfs, coloring

Two-color each component with BFS; an edge between same colors → not bipartite (odd cycle).

**Complexity:** O(V + E).

```python
class Solution:
    def isBipartite(self, graph: List[List[int]]) -> bool:
        color = [-1] * len(graph)
        for s in range(len(graph)):
            if color[s] != -1:
                continue
            color[s] = 0
            q = deque([s])
            while q:
                u = q.popleft()
                for v in graph[u]:
                    if color[v] == -1:
                        color[v] = 1 - color[u]
                        q.append(v)
                    elif color[v] == color[u]:
                        return False
        return True
# test
assert Solution().isBipartite([[1, 2, 3], [0, 2], [0, 1, 3], [0, 2]]) is False
assert Solution().isBipartite([[1, 3], [0, 2], [1, 3], [0, 2]]) is True
```

```java
class Solution {
    public boolean isBipartite(int[][] graph) {
        int n = graph.length;
        int[] color = new int[n];
        Arrays.fill(color, -1);
        for (int s = 0; s < n; s++) {
            if (color[s] != -1) continue;
            color[s] = 0;
            Deque<Integer> q = new ArrayDeque<>(List.of(s));
            while (!q.isEmpty()) {
                int u = q.poll();
                for (int v : graph[u]) {
                    if (color[v] == -1) { color[v] = 1 - color[u]; q.add(v); }
                    else if (color[v] == color[u]) return false;
                }
            }
        }
        return true;
    }
}
// test
check(new Solution().isBipartite(new int[][]{{1, 3}, {0, 2}, {1, 3}, {0, 2}}), true);
```

```go
func isBipartite(graph [][]int) bool {
	color := make([]int, len(graph))
	for i := range color {
		color[i] = -1
	}
	for s := range graph { // the graph may be disconnected
		if color[s] != -1 {
			continue
		}
		color[s] = 0
		q := []int{s}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			for _, v := range graph[u] {
				if color[v] == -1 {
					color[v] = 1 - color[u]
					q = append(q, v)
				} else if color[v] == color[u] {
					return false
				}
			}
		}
	}
	return true
}
// test
assertEq(isBipartite([][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}), false)
assertEq(isBipartite([][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}), true)
```

## 542. 01 Matrix | Medium
related: 994, 286
tags: multi-source-bfs

Distance to the nearest 0 for each cell: multi-source BFS starting from **all zeros** (not from each 1!).

**Complexity:** O(R·C).

```python
class Solution:
    def updateMatrix(self, mat: List[List[int]]) -> List[List[int]]:
        R, C = len(mat), len(mat[0])
        dist = [[-1] * C for _ in range(R)]
        q = deque()
        for r in range(R):
            for c in range(C):
                if mat[r][c] == 0:
                    dist[r][c] = 0
                    q.append((r, c))
        while q:
            r, c = q.popleft()
            for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                if 0 <= nr < R and 0 <= nc < C and dist[nr][nc] == -1:
                    dist[nr][nc] = dist[r][c] + 1
                    q.append((nr, nc))
        return dist
# test
assert Solution().updateMatrix([[0, 0, 0], [0, 1, 0], [1, 1, 1]]) == [[0, 0, 0], [0, 1, 0], [1, 2, 1]]
assert Solution().updateMatrix([[0], [1], [1]]) == [[0], [1], [2]]
assert Solution().updateMatrix([[1, 1, 1], [1, 1, 1], [1, 1, 0]]) == [[4, 3, 2], [3, 2, 1], [2, 1, 0]]
```

```java
class Solution {
    public int[][] updateMatrix(int[][] mat) {
        int R = mat.length, C = mat[0].length;
        int[][] dist = new int[R][C];
        Deque<int[]> q = new ArrayDeque<>();
        for (int r = 0; r < R; r++)
            for (int c = 0; c < C; c++) {
                if (mat[r][c] == 0) q.add(new int[]{r, c});
                else dist[r][c] = -1;
            }
        int[][] dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        while (!q.isEmpty()) {
            int[] cur = q.poll();
            for (int[] d : dirs) {
                int nr = cur[0] + d[0], nc = cur[1] + d[1];
                if (nr >= 0 && nc >= 0 && nr < R && nc < C && dist[nr][nc] == -1) {
                    dist[nr][nc] = dist[cur[0]][cur[1]] + 1;
                    q.add(new int[]{nr, nc});
                }
            }
        }
        return dist;
    }
}
// test
check(new Solution().updateMatrix(new int[][]{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}), new int[][]{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}});
```

```go
func updateMatrix(mat [][]int) [][]int {
	R, C := len(mat), len(mat[0])
	dist := make([][]int, R)
	q := [][2]int{}
	for r := range dist {
		dist[r] = make([]int, C)
		for c := range dist[r] {
			if mat[r][c] == 0 {
				q = append(q, [2]int{r, c}) // BFS from every 0 at once
			} else {
				dist[r][c] = -1
			}
		}
	}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			nr, nc := p[0]+d[0], p[1]+d[1]
			if nr >= 0 && nr < R && nc >= 0 && nc < C && dist[nr][nc] == -1 {
				dist[nr][nc] = dist[p[0]][p[1]] + 1
				q = append(q, [2]int{nr, nc})
			}
		}
	}
	return dist
}
// test
assertEq(updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}), [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}})
assertEq(updateMatrix([][]int{{0}, {1}, {1}}), [][]int{{0}, {1}, {2}})
assertEq(updateMatrix([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 0}}), [][]int{{4, 3, 2}, {3, 2, 1}, {2, 1, 0}})
```

## 1091. Shortest Path in Binary Matrix | Medium
related: 542, 994
tags: bfs

BFS over 8 directions from `(0,0)` to `(n−1,n−1)`; path length counts cells.

**Complexity:** O(n²).

```python
class Solution:
    def shortestPathBinaryMatrix(self, grid: List[List[int]]) -> int:
        n = len(grid)
        if grid[0][0] or grid[n - 1][n - 1]:
            return -1
        q = deque([(0, 0, 1)])
        grid[0][0] = 1
        while q:
            r, c, d = q.popleft()
            if r == c == n - 1:
                return d
            for dr in (-1, 0, 1):
                for dc in (-1, 0, 1):
                    nr, nc = r + dr, c + dc
                    if 0 <= nr < n and 0 <= nc < n and grid[nr][nc] == 0:
                        grid[nr][nc] = 1
                        q.append((nr, nc, d + 1))
        return -1
# test
assert Solution().shortestPathBinaryMatrix([[0, 1], [1, 0]]) == 2
assert Solution().shortestPathBinaryMatrix([[0, 0, 0], [1, 1, 0], [1, 1, 0]]) == 4
assert Solution().shortestPathBinaryMatrix([[1, 0, 0], [1, 1, 0], [1, 1, 0]]) == -1
```

```java
class Solution {
    public int shortestPathBinaryMatrix(int[][] grid) {
        int n = grid.length;
        if (grid[0][0] == 1 || grid[n - 1][n - 1] == 1) return -1;
        Deque<int[]> q = new ArrayDeque<>();
        q.add(new int[]{0, 0, 1});
        grid[0][0] = 1;
        while (!q.isEmpty()) {
            int[] cur = q.poll();
            if (cur[0] == n - 1 && cur[1] == n - 1) return cur[2];
            for (int dr = -1; dr <= 1; dr++)
                for (int dc = -1; dc <= 1; dc++) {
                    int nr = cur[0] + dr, nc = cur[1] + dc;
                    if (nr >= 0 && nc >= 0 && nr < n && nc < n && grid[nr][nc] == 0) {
                        grid[nr][nc] = 1;
                        q.add(new int[]{nr, nc, cur[2] + 1});
                    }
                }
        }
        return -1;
    }
}
// test
check(new Solution().shortestPathBinaryMatrix(new int[][]{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}), 4);
```

```go
func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	q := [][3]int{{0, 0, 1}} // (r, c, path length)
	grid[0][0] = 1
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		if p[0] == n-1 && p[1] == n-1 {
			return p[2]
		}
		for dr := -1; dr <= 1; dr++ { // 8 directions
			for dc := -1; dc <= 1; dc++ {
				nr, nc := p[0]+dr, p[1]+dc
				if nr >= 0 && nr < n && nc >= 0 && nc < n && grid[nr][nc] == 0 {
					grid[nr][nc] = 1
					q = append(q, [3]int{nr, nc, p[2] + 1})
				}
			}
		}
	}
	return -1
}
// test
assertEq(shortestPathBinaryMatrix([][]int{{0, 1}, {1, 0}}), 2)
assertEq(shortestPathBinaryMatrix([][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}), 4)
assertEq(shortestPathBinaryMatrix([][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}}), -1)
```

## 721. Accounts Merge | Medium
related: 684, 547
tags: union-find

Union all emails within an account (to the account's first email). Then group emails by root, sort, and prefix the owner name.

**Complexity:** O(N log N) for N total emails (sorting).

```python
class Solution:
    def accountsMerge(self, accounts: List[List[str]]) -> List[List[str]]:
        parent, owner = {}, {}

        def find(x):
            parent.setdefault(x, x)
            while parent[x] != x:
                parent[x] = parent[parent[x]]
                x = parent[x]
            return x

        for name, *emails in accounts:
            for e in emails:
                owner[e] = name
                parent[find(e)] = find(emails[0])
        groups = defaultdict(list)
        for e in owner:
            groups[find(e)].append(e)
        return [[owner[root]] + sorted(es) for root, es in groups.items()]
# test
acc = [["John", "johnsmith@mail.com", "john_newyork@mail.com"], ["John", "johnsmith@mail.com", "john00@mail.com"],
       ["Mary", "mary@mail.com"], ["John", "johnnybravo@mail.com"]]
assert sorted(Solution().accountsMerge(acc)) == [
    ["John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"],
    ["John", "johnnybravo@mail.com"], ["Mary", "mary@mail.com"]]
assert sorted(Solution().accountsMerge([["Alex", "a@x.com"], ["Alex", "b@x.com"]])) == [["Alex", "a@x.com"], ["Alex", "b@x.com"]]
assert Solution().accountsMerge([["Bo", "b@x.com", "a@x.com", "b@x.com"]]) == [["Bo", "a@x.com", "b@x.com"]]
```

```java
class Solution {
    private final Map<String, String> parent = new HashMap<>();

    private String find(String x) {
        parent.putIfAbsent(x, x);
        while (!parent.get(x).equals(x)) {
            parent.put(x, parent.get(parent.get(x)));
            x = parent.get(x);
        }
        return x;
    }

    public List<List<String>> accountsMerge(List<List<String>> accounts) {
        Map<String, String> owner = new HashMap<>();
        for (List<String> acc : accounts)
            for (int i = 1; i < acc.size(); i++) {
                owner.put(acc.get(i), acc.get(0));
                parent.put(find(acc.get(i)), find(acc.get(1)));
            }
        Map<String, TreeSet<String>> groups = new HashMap<>();
        for (String e : owner.keySet()) groups.computeIfAbsent(find(e), k -> new TreeSet<>()).add(e);
        List<List<String>> res = new ArrayList<>();
        for (Map.Entry<String, TreeSet<String>> g : groups.entrySet()) {
            List<String> row = new ArrayList<>();
            row.add(owner.get(g.getKey()));
            row.addAll(g.getValue());
            res.add(row);
        }
        return res;
    }
}
// test
List<List<String>> acc = List.of(List.of("John", "a@x", "b@x"), List.of("John", "a@x", "c@x"), List.of("Mary", "m@x"));
List<List<String>> r = new Solution().accountsMerge(acc);
r.sort((p, q) -> p.get(1).compareTo(q.get(1)));
check(r, "[[John, a@x, b@x, c@x], [Mary, m@x]]");
```

```go
func accountsMerge(accounts [][]string) [][]string {
	parent, owner := map[string]string{}, map[string]string{}
	var find func(x string) string
	find = func(x string) string {
		if _, ok := parent[x]; !ok {
			parent[x] = x
		}
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	for _, acc := range accounts { // union every email of an account with its first
		for _, e := range acc[1:] {
			owner[e] = acc[0]
			parent[find(e)] = find(acc[1])
		}
	}
	groups := map[string][]string{}
	for e := range owner {
		root := find(e)
		groups[root] = append(groups[root], e)
	}
	res := [][]string{}
	for root, emails := range groups {
		sort.Strings(emails)
		res = append(res, append([]string{owner[root]}, emails...))
	}
	return res
}
// test
acc := [][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "johnsmith@mail.com", "john00@mail.com"},
	{"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}}
assertUnordered(accountsMerge(acc), [][]string{
	{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
	{"John", "johnnybravo@mail.com"}, {"Mary", "mary@mail.com"}})
assertUnordered(accountsMerge([][]string{{"Alex", "a@x.com"}, {"Alex", "b@x.com"}}), [][]string{{"Alex", "a@x.com"}, {"Alex", "b@x.com"}})
assertEq(accountsMerge([][]string{{"Bo", "b@x.com", "a@x.com", "b@x.com"}}), [][]string{{"Bo", "a@x.com", "b@x.com"}})
```

## 329. Longest Increasing Path in a Matrix | Hard
related: 300, 417
tags: dfs, memoization

Strictly increasing paths can't cycle, so the grid is a DAG: `best(r, c) = 1 + max(best(neighbor) for larger neighbours)`, memoized.

**Complexity:** O(R·C).

```python
class Solution:
    def longestIncreasingPath(self, matrix: List[List[int]]) -> int:
        R, C = len(matrix), len(matrix[0])

        @cache
        def best(r, c):
            v, res = matrix[r][c], 1
            for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                if 0 <= nr < R and 0 <= nc < C and matrix[nr][nc] > v:
                    res = max(res, 1 + best(nr, nc))
            return res

        return max(best(r, c) for r in range(R) for c in range(C))
# test
assert Solution().longestIncreasingPath([[9, 9, 4], [6, 6, 8], [2, 1, 1]]) == 4
assert Solution().longestIncreasingPath([[3, 4, 5], [3, 2, 6], [2, 2, 1]]) == 4
```

```java
class Solution {
    private int[][] memo;

    public int longestIncreasingPath(int[][] matrix) {
        int R = matrix.length, C = matrix[0].length, best = 0;
        memo = new int[R][C];
        for (int r = 0; r < R; r++)
            for (int c = 0; c < C; c++) best = Math.max(best, dfs(matrix, r, c));
        return best;
    }

    private int dfs(int[][] m, int r, int c) {
        if (memo[r][c] > 0) return memo[r][c];
        int res = 1;
        int[][] dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        for (int[] d : dirs) {
            int nr = r + d[0], nc = c + d[1];
            if (nr >= 0 && nc >= 0 && nr < m.length && nc < m[0].length && m[nr][nc] > m[r][c])
                res = Math.max(res, 1 + dfs(m, nr, nc));
        }
        return memo[r][c] = res;
    }
}
// test
check(new Solution().longestIncreasingPath(new int[][]{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}), 4);
```

```go
func longestIncreasingPath(matrix [][]int) int {
	R, C := len(matrix), len(matrix[0])
	memo := make([][]int, R) // longest path starting at (r, c); 0 = not computed
	for i := range memo {
		memo[i] = make([]int, C)
	}
	var best func(r, c int) int
	best = func(r, c int) int {
		if memo[r][c] > 0 {
			return memo[r][c]
		}
		res := 1
		for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < R && nc >= 0 && nc < C && matrix[nr][nc] > matrix[r][c] {
				res = max(res, 1+best(nr, nc))
			}
		}
		memo[r][c] = res
		return res
	}
	ans := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			ans = max(ans, best(r, c))
		}
	}
	return ans
}
// test
assertEq(longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}), 4)
assertEq(longestIncreasingPath([][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}), 4)
```


## 1971. Find if Path Exists in Graph | Easy
related: 323, 547
tags: union-find, bfs

Union every edge, then check whether `source` and `destination` share a root (or BFS from the source).

**Complexity:** O(E · α(V)).

```python
class Solution:
    def validPath(self, n: int, edges: List[List[int]], source: int, destination: int) -> bool:
        parent = list(range(n))

        def find(x):
            while parent[x] != x:
                parent[x] = parent[parent[x]]
                x = parent[x]
            return x

        for a, b in edges:
            parent[find(a)] = find(b)
        return find(source) == find(destination)
# test
assert Solution().validPath(3, [[0, 1], [1, 2], [2, 0]], 0, 2) is True
assert Solution().validPath(6, [[0, 1], [0, 2], [3, 5], [5, 4], [4, 3]], 0, 5) is False
```

```java
class Solution {
    public boolean validPath(int n, int[][] edges, int source, int destination) {
        int[] p = new int[n];
        for (int i = 0; i < n; i++) p[i] = i;
        for (int[] e : edges) p[find(p, e[0])] = find(p, e[1]);
        return find(p, source) == find(p, destination);
    }

    private int find(int[] p, int x) {
        while (p[x] != x) { p[x] = p[p[x]]; x = p[x]; }
        return x;
    }
}
// test
check(new Solution().validPath(6, new int[][]{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5), false);
```

```go
func validPath(n int, edges [][]int, source int, destination int) bool {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	find := func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	for _, e := range edges {
		parent[find(e[0])] = find(e[1])
	}
	return find(source) == find(destination)
}
// test
assertEq(validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2), true)
assertEq(validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5), false)
```

## 841. Keys and Rooms | Medium
related: 547, 1971
tags: dfs

Rooms are nodes, keys are edges. DFS from room 0; all rooms visited ⇔ true.

**Complexity:** O(V + E).

```python
class Solution:
    def canVisitAllRooms(self, rooms: List[List[int]]) -> bool:
        seen, st = {0}, [0]
        while st:
            for k in rooms[st.pop()]:
                if k not in seen:
                    seen.add(k)
                    st.append(k)
        return len(seen) == len(rooms)
# test
assert Solution().canVisitAllRooms([[1], [2], [3], []]) is True
assert Solution().canVisitAllRooms([[1, 3], [3, 0, 1], [2], [0]]) is False
```

```java
class Solution {
    public boolean canVisitAllRooms(List<List<Integer>> rooms) {
        boolean[] seen = new boolean[rooms.size()];
        seen[0] = true;
        Deque<Integer> st = new ArrayDeque<>(List.of(0));
        int count = 1;
        while (!st.isEmpty())
            for (int k : rooms.get(st.pop()))
                if (!seen[k]) { seen[k] = true; count++; st.push(k); }
        return count == rooms.size();
    }
}
// test
check(new Solution().canVisitAllRooms(List.of(List.of(1), List.of(2), List.of(3), List.of())), true);
```

```go
func canVisitAllRooms(rooms [][]int) bool {
	seen, st := map[int]bool{0: true}, []int{0}
	for len(st) > 0 {
		room := st[len(st)-1]
		st = st[:len(st)-1]
		for _, k := range rooms[room] {
			if !seen[k] {
				seen[k] = true
				st = append(st, k)
			}
		}
	}
	return len(seen) == len(rooms)
}
// test
assertEq(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}), true)
assertEq(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}), false)
```

## 886. Possible Bipartition | Medium
related: 785
tags: bfs, coloring

"Dislike" edges must cross groups → is the dislike graph bipartite? 2-color each component with BFS.

**Complexity:** O(V + E).

```python
class Solution:
    def possibleBipartition(self, n: int, dislikes: List[List[int]]) -> bool:
        g = defaultdict(list)
        for a, b in dislikes:
            g[a].append(b)
            g[b].append(a)
        color = {}
        for s in range(1, n + 1):
            if s in color:
                continue
            color[s] = 0
            q = deque([s])
            while q:
                u = q.popleft()
                for v in g[u]:
                    if v not in color:
                        color[v] = 1 - color[u]
                        q.append(v)
                    elif color[v] == color[u]:
                        return False
        return True
# test
assert Solution().possibleBipartition(4, [[1, 2], [1, 3], [2, 4]]) is True
assert Solution().possibleBipartition(3, [[1, 2], [1, 3], [2, 3]]) is False
```

```java
class Solution {
    public boolean possibleBipartition(int n, int[][] dislikes) {
        List<List<Integer>> g = new ArrayList<>();
        for (int i = 0; i <= n; i++) g.add(new ArrayList<>());
        for (int[] d : dislikes) { g.get(d[0]).add(d[1]); g.get(d[1]).add(d[0]); }
        int[] color = new int[n + 1];
        Arrays.fill(color, -1);
        for (int s = 1; s <= n; s++) {
            if (color[s] != -1) continue;
            color[s] = 0;
            Deque<Integer> q = new ArrayDeque<>(List.of(s));
            while (!q.isEmpty()) {
                int u = q.poll();
                for (int v : g.get(u)) {
                    if (color[v] == -1) { color[v] = 1 - color[u]; q.add(v); }
                    else if (color[v] == color[u]) return false;
                }
            }
        }
        return true;
    }
}
// test
check(new Solution().possibleBipartition(3, new int[][]{{1, 2}, {1, 3}, {2, 3}}), false);
```

```go
func possibleBipartition(n int, dislikes [][]int) bool {
	g := make([][]int, n+1)
	for _, d := range dislikes {
		g[d[0]] = append(g[d[0]], d[1])
		g[d[1]] = append(g[d[1]], d[0])
	}
	color := make([]int, n+1) // 0 = uncolored, 1 / 2 = the two groups
	for s := 1; s <= n; s++ {
		if color[s] != 0 {
			continue
		}
		color[s] = 1
		q := []int{s}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			for _, v := range g[u] {
				if color[v] == 0 {
					color[v] = 3 - color[u]
					q = append(q, v)
				} else if color[v] == color[u] {
					return false
				}
			}
		}
	}
	return true
}
// test
assertEq(possibleBipartition(4, [][]int{{1, 2}, {1, 3}, {2, 4}}), true)
assertEq(possibleBipartition(3, [][]int{{1, 2}, {1, 3}, {2, 3}}), false)
```

## 990. Satisfiability of Equality Equations | Medium
related: 399, 684
tags: union-find

First union all `a==b`; then every `a!=b` must have different roots.

**Complexity:** O(n · α(26)).

```python
class Solution:
    def equationsPossible(self, equations: List[str]) -> bool:
        parent = list(range(26))

        def find(x):
            while parent[x] != x:
                parent[x] = parent[parent[x]]
                x = parent[x]
            return x

        for e in equations:
            if e[1] == "=":
                parent[find(ord(e[0]) - 97)] = find(ord(e[3]) - 97)
        return all(find(ord(e[0]) - 97) != find(ord(e[3]) - 97) for e in equations if e[1] == "!")
# test
assert Solution().equationsPossible(["a==b", "b!=a"]) is False
assert Solution().equationsPossible(["b==a", "a==b"]) is True
assert Solution().equationsPossible(["a==b", "b==c", "a==c"]) is True
```

```java
class Solution {
    private final int[] p = new int[26];

    private int find(int x) {
        while (p[x] != x) { p[x] = p[p[x]]; x = p[x]; }
        return x;
    }

    public boolean equationsPossible(String[] equations) {
        for (int i = 0; i < 26; i++) p[i] = i;
        for (String e : equations) if (e.charAt(1) == '=') p[find(e.charAt(0) - 'a')] = find(e.charAt(3) - 'a');
        for (String e : equations) if (e.charAt(1) == '!' && find(e.charAt(0) - 'a') == find(e.charAt(3) - 'a')) return false;
        return true;
    }
}
// test
check(new Solution().equationsPossible(new String[]{"a==b", "b!=a"}), false);
```

```go
func equationsPossible(equations []string) bool {
	var parent [26]int
	for i := range parent {
		parent[i] = i
	}
	find := func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	for _, e := range equations { // union all "==" first
		if e[1] == '=' {
			parent[find(int(e[0]-'a'))] = find(int(e[3] - 'a'))
		}
	}
	for _, e := range equations { // then every "!=" must span two groups
		if e[1] == '!' && find(int(e[0]-'a')) == find(int(e[3]-'a')) {
			return false
		}
	}
	return true
}
// test
assertEq(equationsPossible([]string{"a==b", "b!=a"}), false)
assertEq(equationsPossible([]string{"b==a", "a==b"}), true)
assertEq(equationsPossible([]string{"a==b", "b==c", "a==c"}), true)
```

## 1466. Reorder Routes to Make All Paths Lead to the City Zero | Medium
related: 207
tags: dfs, directed-graph

Treat roads as undirected but remember direction: an original edge `a → b` costs 1 when traversed away from 0 (it must be reversed), its reverse costs 0. DFS from 0 summing costs.

**Complexity:** O(n).

```python
class Solution:
    def minReorder(self, n: int, connections: List[List[int]]) -> int:
        g = defaultdict(list)
        for a, b in connections:
            g[a].append((b, 1))     # original direction, points away from 0 -> needs flip
            g[b].append((a, 0))
        seen, st, flips = {0}, [0], 0
        while st:
            u = st.pop()
            for v, cost in g[u]:
                if v not in seen:
                    seen.add(v)
                    flips += cost
                    st.append(v)
        return flips
# test
assert Solution().minReorder(6, [[0, 1], [1, 3], [2, 3], [4, 0], [4, 5]]) == 3
assert Solution().minReorder(5, [[1, 0], [1, 2], [3, 2], [3, 4]]) == 2
```

```java
class Solution {
    public int minReorder(int n, int[][] connections) {
        List<List<int[]>> g = new ArrayList<>();
        for (int i = 0; i < n; i++) g.add(new ArrayList<>());
        for (int[] c : connections) { g.get(c[0]).add(new int[]{c[1], 1}); g.get(c[1]).add(new int[]{c[0], 0}); }
        boolean[] seen = new boolean[n];
        seen[0] = true;
        Deque<Integer> st = new ArrayDeque<>(List.of(0));
        int flips = 0;
        while (!st.isEmpty())
            for (int[] e : g.get(st.pop()))
                if (!seen[e[0]]) { seen[e[0]] = true; flips += e[1]; st.push(e[0]); }
        return flips;
    }
}
// test
check(new Solution().minReorder(6, new int[][]{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}), 3);
```

```go
func minReorder(n int, connections [][]int) int {
	g := make([][][2]int, n) // (neighbor, 1 if the road points away from city 0's side)
	for _, c := range connections {
		g[c[0]] = append(g[c[0]], [2]int{c[1], 1})
		g[c[1]] = append(g[c[1]], [2]int{c[0], 0})
	}
	seen, st, flips := make([]bool, n), []int{0}, 0
	seen[0] = true
	for len(st) > 0 {
		u := st[len(st)-1]
		st = st[:len(st)-1]
		for _, e := range g[u] {
			if !seen[e[0]] {
				seen[e[0]] = true
				flips += e[1]
				st = append(st, e[0])
			}
		}
	}
	return flips
}
// test
assertEq(minReorder(6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}), 3)
assertEq(minReorder(5, [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}), 2)
```

## 752. Open the Lock | Medium
related: 127, 433
tags: bfs, implicit-graph

States are 4-digit strings; each move turns one wheel ±1 → 8 neighbours. BFS from `"0000"`, skipping deadends.

**Complexity:** O(10⁴ · 8).

```python
class Solution:
    def openLock(self, deadends: List[str], target: str) -> int:
        dead = set(deadends)
        if "0000" in dead:
            return -1
        q, seen, steps = deque(["0000"]), {"0000"}, 0
        while q:
            for _ in range(len(q)):
                cur = q.popleft()
                if cur == target:
                    return steps
                for i in range(4):
                    d = int(cur[i])
                    for nd in ((d + 1) % 10, (d - 1) % 10):
                        nxt = cur[:i] + str(nd) + cur[i + 1:]
                        if nxt not in seen and nxt not in dead:
                            seen.add(nxt)
                            q.append(nxt)
            steps += 1
        return -1
# test
assert Solution().openLock(["0201", "0101", "0102", "1212", "2002"], "0202") == 6
assert Solution().openLock(["8888"], "0009") == 1
assert Solution().openLock(["8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"], "8888") == -1
```

```java
class Solution {
    public int openLock(String[] deadends, String target) {
        Set<String> seen = new HashSet<>(Arrays.asList(deadends));
        if (!seen.add("0000")) return -1;
        Deque<String> q = new ArrayDeque<>(List.of("0000"));
        for (int steps = 0; !q.isEmpty(); steps++) {
            for (int k = q.size(); k > 0; k--) {
                String cur = q.poll();
                if (cur.equals(target)) return steps;
                char[] c = cur.toCharArray();
                for (int i = 0; i < 4; i++) {
                    char orig = c[i];
                    for (int delta : new int[]{1, 9}) {
                        c[i] = (char) ('0' + (orig - '0' + delta) % 10);
                        String nxt = new String(c);
                        if (seen.add(nxt)) q.add(nxt);
                    }
                    c[i] = orig;
                }
            }
        }
        return -1;
    }
}
// test
check(new Solution().openLock(new String[]{"0201", "0101", "0102", "1212", "2002"}, "0202"), 6);
```

```go
func openLock(deadends []string, target string) int {
	dead := map[string]bool{}
	for _, d := range deadends {
		dead[d] = true
	}
	if dead["0000"] {
		return -1
	}
	q, seen := []string{"0000"}, map[string]bool{"0000": true}
	for steps := 0; len(q) > 0; steps++ {
		next := []string{}
		for _, cur := range q {
			if cur == target {
				return steps
			}
			b := []byte(cur)
			for i := 0; i < 4; i++ {
				orig := b[i]
				for _, delta := range []byte{1, 9} { // +1 or -1 (mod 10)
					b[i] = '0' + (orig-'0'+delta)%10
					if s := string(b); !seen[s] && !dead[s] {
						seen[s] = true
						next = append(next, s)
					}
				}
				b[i] = orig
			}
		}
		q = next
	}
	return -1
}
// test
assertEq(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"), 6)
assertEq(openLock([]string{"8888"}, "0009"), 1)
assertEq(openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"), -1)
```

## 433. Minimum Genetic Mutation | Medium
related: 127, 752
tags: bfs

Same as Word Ladder with a 4-letter alphabet: BFS over valid genes in the bank.

**Complexity:** O(B · 8 · 4).

```python
class Solution:
    def minMutation(self, startGene: str, endGene: str, bank: List[str]) -> int:
        valid = set(bank)
        q, seen, steps = deque([startGene]), {startGene}, 0
        while q:
            for _ in range(len(q)):
                g = q.popleft()
                if g == endGene:
                    return steps
                for i in range(len(g)):
                    for ch in "ACGT":
                        nxt = g[:i] + ch + g[i + 1:]
                        if nxt in valid and nxt not in seen:
                            seen.add(nxt)
                            q.append(nxt)
            steps += 1
        return -1
# test
assert Solution().minMutation("AACCGGTT", "AACCGGTA", ["AACCGGTA"]) == 1
assert Solution().minMutation("AACCGGTT", "AAACGGTA", ["AACCGGTA", "AACCGCTA", "AAACGGTA"]) == 2
assert Solution().minMutation("AAAAACCC", "AACCCCCC", ["AAAACCCC", "AAACCCCC", "AACCCCCC"]) == 3
```

```java
class Solution {
    public int minMutation(String startGene, String endGene, String[] bank) {
        Set<String> valid = new HashSet<>(Arrays.asList(bank)), seen = new HashSet<>(List.of(startGene));
        Deque<String> q = new ArrayDeque<>(List.of(startGene));
        for (int steps = 0; !q.isEmpty(); steps++)
            for (int k = q.size(); k > 0; k--) {
                String g = q.poll();
                if (g.equals(endGene)) return steps;
                char[] c = g.toCharArray();
                for (int i = 0; i < c.length; i++) {
                    char orig = c[i];
                    for (char ch : "ACGT".toCharArray()) {
                        c[i] = ch;
                        String nxt = new String(c);
                        if (valid.contains(nxt) && seen.add(nxt)) q.add(nxt);
                    }
                    c[i] = orig;
                }
            }
        return -1;
    }
}
// test
check(new Solution().minMutation("AACCGGTT", "AAACGGTA", new String[]{"AACCGGTA", "AACCGCTA", "AAACGGTA"}), 2);
```

```go
func minMutation(startGene string, endGene string, bank []string) int {
	valid := map[string]bool{}
	for _, g := range bank {
		valid[g] = true
	}
	q, seen := []string{startGene}, map[string]bool{startGene: true}
	for steps := 0; len(q) > 0; steps++ {
		next := []string{}
		for _, g := range q {
			if g == endGene {
				return steps
			}
			b := []byte(g)
			for i := range b {
				orig := b[i]
				for _, ch := range []byte("ACGT") {
					b[i] = ch
					if s := string(b); valid[s] && !seen[s] {
						seen[s] = true
						next = append(next, s)
					}
				}
				b[i] = orig
			}
		}
		q = next
	}
	return -1
}
// test
assertEq(minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}), 1)
assertEq(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}), 2)
assertEq(minMutation("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}), 3)
```

## 909. Snakes and Ladders | Medium
related: 752
tags: bfs, board-mapping

BFS over squares 1..n². The only tricky part is mapping square `s` to `(row, col)` on a boustrophedon board: `r = (s−1) // n`, `c = (s−1) % n`, reversed on odd rows, and rows counted from the bottom.

**Complexity:** O(n²).

```python
class Solution:
    def snakesAndLadders(self, board: List[List[int]]) -> int:
        n = len(board)

        def cell(s):
            r, c = divmod(s - 1, n)
            if r % 2:
                c = n - 1 - c
            return board[n - 1 - r][c]

        q, seen, moves = deque([1]), {1}, 0
        while q:
            for _ in range(len(q)):
                s = q.popleft()
                if s == n * n:
                    return moves
                for nxt in range(s + 1, min(s + 6, n * n) + 1):
                    dest = cell(nxt) if cell(nxt) != -1 else nxt
                    if dest not in seen:
                        seen.add(dest)
                        q.append(dest)
            moves += 1
        return -1
# test
B = [[-1, -1, -1, -1, -1, -1], [-1, -1, -1, -1, -1, -1], [-1, -1, -1, -1, -1, -1],
     [-1, 35, -1, -1, 13, -1], [-1, -1, -1, -1, -1, -1], [-1, 15, -1, -1, -1, -1]]
assert Solution().snakesAndLadders(B) == 4
assert Solution().snakesAndLadders([[-1, -1], [-1, 3]]) == 1
```

```java
class Solution {
    public int snakesAndLadders(int[][] board) {
        int n = board.length;
        boolean[] seen = new boolean[n * n + 1];
        Deque<Integer> q = new ArrayDeque<>(List.of(1));
        seen[1] = true;
        for (int moves = 0; !q.isEmpty(); moves++)
            for (int k = q.size(); k > 0; k--) {
                int s = q.poll();
                if (s == n * n) return moves;
                for (int nxt = s + 1; nxt <= Math.min(s + 6, n * n); nxt++) {
                    int r = (nxt - 1) / n, c = (nxt - 1) % n;
                    if (r % 2 == 1) c = n - 1 - c;
                    int v = board[n - 1 - r][c], dest = v == -1 ? nxt : v;
                    if (!seen[dest]) { seen[dest] = true; q.add(dest); }
                }
            }
        return -1;
    }
}
// test
int[][] B = {{-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1},
             {-1, 35, -1, -1, 13, -1}, {-1, -1, -1, -1, -1, -1}, {-1, 15, -1, -1, -1, -1}};
check(new Solution().snakesAndLadders(B), 4);
```

```go
func snakesAndLadders(board [][]int) int {
	n := len(board)
	cell := func(s int) int { // board value of square s (boustrophedon from the bottom-left)
		r, c := (s-1)/n, (s-1)%n
		if r%2 == 1 {
			c = n - 1 - c
		}
		return board[n-1-r][c]
	}
	q, seen := []int{1}, map[int]bool{1: true}
	for moves := 0; len(q) > 0; moves++ {
		next := []int{}
		for _, s := range q {
			if s == n*n {
				return moves
			}
			for nxt := s + 1; nxt <= min(s+6, n*n); nxt++ {
				dest := nxt
				if v := cell(nxt); v != -1 {
					dest = v
				}
				if !seen[dest] {
					seen[dest] = true
					next = append(next, dest)
				}
			}
		}
		q = next
	}
	return -1
}
// test
B := [][]int{{-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1},
	{-1, 35, -1, -1, 13, -1}, {-1, -1, -1, -1, -1, -1}, {-1, 15, -1, -1, -1, -1}}
assertEq(snakesAndLadders(B), 4)
assertEq(snakesAndLadders([][]int{{-1, -1}, {-1, 3}}), 1)
```

## 934. Shortest Bridge | Medium
related: 200, 542
tags: dfs, multi-source-bfs

Find one island with DFS and mark it; then multi-source BFS from all its cells across water. The BFS level at which you first touch the other island is the number of flips.

**Complexity:** O(n²).

```python
class Solution:
    def shortestBridge(self, grid: List[List[int]]) -> int:
        n = len(grid)
        start = next((r, c) for r in range(n) for c in range(n) if grid[r][c])
        q, st = deque(), [start]
        grid[start[0]][start[1]] = 2
        while st:                                   # mark island A as 2
            r, c = st.pop()
            q.append((r, c))
            for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                if 0 <= nr < n and 0 <= nc < n and grid[nr][nc] == 1:
                    grid[nr][nc] = 2
                    st.append((nr, nc))
        steps = 0
        while q:
            for _ in range(len(q)):
                r, c = q.popleft()
                for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                    if 0 <= nr < n and 0 <= nc < n:
                        if grid[nr][nc] == 1:
                            return steps
                        if grid[nr][nc] == 0:
                            grid[nr][nc] = 2
                            q.append((nr, nc))
            steps += 1
# test
assert Solution().shortestBridge([[0, 1], [1, 0]]) == 1
assert Solution().shortestBridge([[0, 1, 0], [0, 0, 0], [0, 0, 1]]) == 2
assert Solution().shortestBridge([[1, 1, 1, 1, 1], [1, 0, 0, 0, 1], [1, 0, 1, 0, 1], [1, 0, 0, 0, 1], [1, 1, 1, 1, 1]]) == 1
```

```java
class Solution {
    public int shortestBridge(int[][] grid) {
        int n = grid.length;
        int[][] dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        Deque<int[]> q = new ArrayDeque<>(), st = new ArrayDeque<>();
        outer:
        for (int r = 0; r < n; r++)
            for (int c = 0; c < n; c++)
                if (grid[r][c] == 1) { grid[r][c] = 2; st.push(new int[]{r, c}); break outer; }
        while (!st.isEmpty()) {
            int[] cur = st.pop();
            q.add(cur);
            for (int[] d : dirs) {
                int nr = cur[0] + d[0], nc = cur[1] + d[1];
                if (nr >= 0 && nc >= 0 && nr < n && nc < n && grid[nr][nc] == 1) { grid[nr][nc] = 2; st.push(new int[]{nr, nc}); }
            }
        }
        for (int steps = 0; !q.isEmpty(); steps++)
            for (int k = q.size(); k > 0; k--) {
                int[] cur = q.poll();
                for (int[] d : dirs) {
                    int nr = cur[0] + d[0], nc = cur[1] + d[1];
                    if (nr < 0 || nc < 0 || nr >= n || nc >= n) continue;
                    if (grid[nr][nc] == 1) return steps;
                    if (grid[nr][nc] == 0) { grid[nr][nc] = 2; q.add(new int[]{nr, nc}); }
                }
            }
        return -1;
    }
}
// test
check(new Solution().shortestBridge(new int[][]{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}), 2);
```

```go
func shortestBridge(grid [][]int) int {
	n := len(grid)
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var start [2]int
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 1 {
				start = [2]int{r, c}
			}
		}
	}
	q, st := [][2]int{}, [][2]int{start}
	grid[start[0]][start[1]] = 2
	for len(st) > 0 { // mark one island as 2; it seeds the BFS
		p := st[len(st)-1]
		st = st[:len(st)-1]
		q = append(q, p)
		for _, d := range dirs {
			nr, nc := p[0]+d[0], p[1]+d[1]
			if nr >= 0 && nr < n && nc >= 0 && nc < n && grid[nr][nc] == 1 {
				grid[nr][nc] = 2
				st = append(st, [2]int{nr, nc})
			}
		}
	}
	for steps := 0; len(q) > 0; steps++ { // expand over water until the other island
		next := [][2]int{}
		for _, p := range q {
			for _, d := range dirs {
				nr, nc := p[0]+d[0], p[1]+d[1]
				if nr < 0 || nr >= n || nc < 0 || nc >= n {
					continue
				}
				if grid[nr][nc] == 1 {
					return steps
				}
				if grid[nr][nc] == 0 {
					grid[nr][nc] = 2
					next = append(next, [2]int{nr, nc})
				}
			}
		}
		q = next
	}
	return -1
}
// test
assertEq(shortestBridge([][]int{{0, 1}, {1, 0}}), 1)
assertEq(shortestBridge([][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}), 2)
assertEq(shortestBridge([][]int{{1, 1, 1, 1, 1}, {1, 0, 0, 0, 1}, {1, 0, 1, 0, 1}, {1, 0, 0, 0, 1}, {1, 1, 1, 1, 1}}), 1)
```

## 733. Flood Fill | Easy
related: 200, 695, 463
tags: grid-dfs

DFS/BFS from the start cell through 4-neighbours that have the **original** color, recoloring as you go. If the new color equals the original, return immediately — otherwise the search never terminates (recolored cells still match).

**Complexity:** O(R·C).

```python
class Solution:
    def floodFill(self, image: List[List[int]], sr: int, sc: int, color: int) -> List[List[int]]:
        old = image[sr][sc]
        if old == color:
            return image
        R, C = len(image), len(image[0])
        image[sr][sc] = color
        stack = [(sr, sc)]
        while stack:
            r, c = stack.pop()
            for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                if 0 <= nr < R and 0 <= nc < C and image[nr][nc] == old:
                    image[nr][nc] = color
                    stack.append((nr, nc))
        return image
# test
assert Solution().floodFill([[1, 1, 1], [1, 1, 0], [1, 0, 1]], 1, 1, 2) == [[2, 2, 2], [2, 2, 0], [2, 0, 1]]
assert Solution().floodFill([[0, 0, 0], [0, 0, 0]], 0, 0, 0) == [[0, 0, 0], [0, 0, 0]]
assert Solution().floodFill([[0, 0, 0], [0, 1, 1]], 1, 1, 1) == [[0, 0, 0], [0, 1, 1]]
assert Solution().floodFill([[1]], 0, 0, 3) == [[3]]
```

```java
class Solution {
    public int[][] floodFill(int[][] image, int sr, int sc, int color) {
        int old = image[sr][sc];
        if (old != color) fill(image, sr, sc, old, color);
        return image;
    }

    private void fill(int[][] g, int r, int c, int old, int color) {
        if (r < 0 || c < 0 || r >= g.length || c >= g[0].length || g[r][c] != old) return;
        g[r][c] = color;
        fill(g, r + 1, c, old, color);
        fill(g, r - 1, c, old, color);
        fill(g, r, c + 1, old, color);
        fill(g, r, c - 1, old, color);
    }
}
// test
check(new Solution().floodFill(new int[][]{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2), new int[][]{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}});
check(new Solution().floodFill(new int[][]{{0, 0, 0}, {0, 0, 0}}, 0, 0, 0), new int[][]{{0, 0, 0}, {0, 0, 0}});
```

```go
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	old := image[sr][sc]
	if old == color {
		return image
	}
	R, C := len(image), len(image[0])
	image[sr][sc] = color
	st := [][2]int{{sr, sc}}
	for len(st) > 0 {
		p := st[len(st)-1]
		st = st[:len(st)-1]
		for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			nr, nc := p[0]+d[0], p[1]+d[1]
			if nr >= 0 && nr < R && nc >= 0 && nc < C && image[nr][nc] == old {
				image[nr][nc] = color
				st = append(st, [2]int{nr, nc})
			}
		}
	}
	return image
}
// test
assertEq(floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2), [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}})
assertEq(floodFill([][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 0), [][]int{{0, 0, 0}, {0, 0, 0}})
assertEq(floodFill([][]int{{0, 0, 0}, {0, 1, 1}}, 1, 1, 1), [][]int{{0, 0, 0}, {0, 1, 1}})
assertEq(floodFill([][]int{{1}}, 0, 0, 3), [][]int{{3}})
```

## 1730. Shortest Path to Get Food | Medium
related: 1926, 994, 1091
tags: bfs, grid, premium

**Problem (Premium):** a grid has exactly one `'*'` (you), any number of `'#'` (food), `'O'` (free) and `'X'` (wall). Moving up/down/left/right, return the length of the shortest path to **any** food cell, or −1.

Unweighted shortest path → BFS from `'*'`, layer by layer; the first food cell reached is the answer.

**Complexity:** O(R·C).

```python
class Solution:
    def getFood(self, grid: List[List[str]]) -> int:
        R, C = len(grid), len(grid[0])
        start = next((r, c) for r in range(R) for c in range(C) if grid[r][c] == "*")
        seen, q, steps = {start}, deque([start]), 0
        while q:
            steps += 1
            for _ in range(len(q)):
                r, c = q.popleft()
                for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                    if 0 <= nr < R and 0 <= nc < C and grid[nr][nc] != "X" and (nr, nc) not in seen:
                        if grid[nr][nc] == "#":
                            return steps
                        seen.add((nr, nc))
                        q.append((nr, nc))
        return -1
# test
G = lambda rows: [list(r) for r in rows]
assert Solution().getFood(G(["XXXXXX", "X*OOOX", "XOO#OX", "XXXXXX"])) == 3
assert Solution().getFood(G(["XXXXX", "X*XOX", "XOX#X", "XXXXX"])) == -1
assert Solution().getFood(G(["XXXXXXXX", "X*OXO#OX", "XOOXOOXX", "XOOOO#OX", "XXXXXXXX"])) == 6
assert Solution().getFood(G(["*#"])) == 1
```

```java
class Solution {
    public int getFood(char[][] grid) {
        int R = grid.length, C = grid[0].length;
        Deque<int[]> q = new ArrayDeque<>();
        boolean[][] seen = new boolean[R][C];
        for (int r = 0; r < R; r++) for (int c = 0; c < C; c++) if (grid[r][c] == '*') { q.add(new int[]{r, c}); seen[r][c] = true; }
        int[][] dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        for (int steps = 1; !q.isEmpty(); steps++) {
            for (int i = q.size(); i > 0; i--) {
                int[] cur = q.poll();
                for (int[] d : dirs) {
                    int r = cur[0] + d[0], c = cur[1] + d[1];
                    if (r < 0 || c < 0 || r >= R || c >= C || seen[r][c] || grid[r][c] == 'X') continue;
                    if (grid[r][c] == '#') return steps;
                    seen[r][c] = true;
                    q.add(new int[]{r, c});
                }
            }
        }
        return -1;
    }
}
// test
check(new Solution().getFood(new char[][]{"XXXXXX".toCharArray(), "X*OOOX".toCharArray(), "XOO#OX".toCharArray(), "XXXXXX".toCharArray()}), 3);
check(new Solution().getFood(new char[][]{"*#".toCharArray()}), 1);
```

```go
func getFood(grid [][]byte) int {
	R, C := len(grid), len(grid[0])
	var start [2]int
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if grid[r][c] == '*' {
				start = [2]int{r, c}
			}
		}
	}
	seen, q := map[[2]int]bool{start: true}, [][2]int{start}
	for steps := 1; len(q) > 0; steps++ {
		next := [][2]int{}
		for _, p := range q {
			for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				nb := [2]int{p[0] + d[0], p[1] + d[1]}
				if nb[0] < 0 || nb[0] >= R || nb[1] < 0 || nb[1] >= C || grid[nb[0]][nb[1]] == 'X' || seen[nb] {
					continue
				}
				if grid[nb[0]][nb[1]] == '#' {
					return steps
				}
				seen[nb] = true
				next = append(next, nb)
			}
		}
		q = next
	}
	return -1
}
// test
assertEq(getFood(byteGrid("XXXXXX", "X*OOOX", "XOO#OX", "XXXXXX")), 3)
assertEq(getFood(byteGrid("XXXXX", "X*XOX", "XOX#X", "XXXXX")), -1)
assertEq(getFood(byteGrid("XXXXXXXX", "X*OXO#OX", "XOOXOOXX", "XOOOO#OX", "XXXXXXXX")), 6)
assertEq(getFood(byteGrid("*#")), 1)
```

## 1197. Minimum Knight Moves | Medium
related: 1730, 752
tags: bfs, premium

**Problem (Premium):** on an infinite chessboard a knight starts at `(0, 0)`. Return the minimum number of moves to reach `(x, y)` (`|x| + |y| ≤ 300`).

BFS over squares. Two prunings make it fast: the board is **symmetric**, so solve for `(|x|, |y|)`; and an optimal path never needs to go more than 2 squares beyond the box `[0, x] × [0, y]`, so ignore squares outside `[−2, x+2] × [−2, y+2]`.

**Complexity:** O(x · y).

```python
class Solution:
    def minKnightMoves(self, x: int, y: int) -> int:
        x, y = abs(x), abs(y)
        seen, q = {(0, 0)}, deque([(0, 0, 0)])
        while q:
            a, b, d = q.popleft()
            if (a, b) == (x, y):
                return d
            for da, db in ((1, 2), (2, 1), (2, -1), (1, -2), (-1, -2), (-2, -1), (-2, 1), (-1, 2)):
                na, nb = a + da, b + db
                if -2 <= na <= x + 2 and -2 <= nb <= y + 2 and (na, nb) not in seen:
                    seen.add((na, nb))
                    q.append((na, nb, d + 1))
# test
assert Solution().minKnightMoves(2, 1) == 1
assert Solution().minKnightMoves(5, 5) == 4
assert Solution().minKnightMoves(0, 0) == 0
assert Solution().minKnightMoves(1, 1) == 2
assert Solution().minKnightMoves(1, 0) == 3
assert Solution().minKnightMoves(-1, -2) == 1
assert Solution().minKnightMoves(300, 0) == 150
```

```java
class Solution {
    public int minKnightMoves(int x, int y) {
        x = Math.abs(x); y = Math.abs(y);
        int off = 2;                                   // shift so coordinates >= -2 index from 0
        boolean[][] seen = new boolean[x + 5][y + 5];
        Deque<int[]> q = new ArrayDeque<>();
        q.add(new int[]{0, 0});
        seen[off][off] = true;
        int[][] moves = {{1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}};
        for (int d = 0; ; d++) {
            for (int i = q.size(); i > 0; i--) {
                int[] p = q.poll();
                if (p[0] == x && p[1] == y) return d;
                for (int[] m : moves) {
                    int a = p[0] + m[0], b = p[1] + m[1];
                    if (a < -2 || b < -2 || a > x + 2 || b > y + 2 || seen[a + off][b + off]) continue;
                    seen[a + off][b + off] = true;
                    q.add(new int[]{a, b});
                }
            }
        }
    }
}
// test
check(new Solution().minKnightMoves(5, 5), 4);
check(new Solution().minKnightMoves(1, 0), 3);
check(new Solution().minKnightMoves(300, 0), 150);
```

```go
func minKnightMoves(x int, y int) int {
	if x < 0 { // the board is symmetric
		x = -x
	}
	if y < 0 {
		y = -y
	}
	type state struct{ a, b int }
	seen := map[state]bool{{0, 0}: true}
	q := []state{{0, 0}}
	for d := 0; len(q) > 0; d++ {
		next := []state{}
		for _, s := range q {
			if s.a == x && s.b == y {
				return d
			}
			for _, m := range [][2]int{{1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}} {
				n := state{s.a + m[0], s.b + m[1]}
				if n.a >= -2 && n.a <= x+2 && n.b >= -2 && n.b <= y+2 && !seen[n] { // a small margin is enough
					seen[n] = true
					next = append(next, n)
				}
			}
		}
		q = next
	}
	return -1
}
// test
assertEq(minKnightMoves(2, 1), 1)
assertEq(minKnightMoves(5, 5), 4)
assertEq(minKnightMoves(0, 0), 0)
assertEq(minKnightMoves(1, 1), 2)
assertEq(minKnightMoves(1, 0), 3)
assertEq(minKnightMoves(-1, -2), 1)
assertEq(minKnightMoves(300, 0), 150)
```

## 815. Bus Routes | Hard
related: 127, 1730
tags: bfs, implicit-graph

BFS where one "step" is **riding one bus**: from the current stops, board every not-yet-used route through them and reach all of that route's stops. Index `stop → routes` first; mark routes as used so each route is expanded once.

```viz
routes: 0:[1 2 7]  1:[3 6 7]      source 1, target 6
bus 1: stop 1 -> route 0 -> stops {2, 7}
bus 2: stop 7 -> route 1 -> stops {3, 6}  -> target reached, answer 2
```

**Complexity:** O(Σ route lengths).

```python
class Solution:
    def numBusesToDestination(self, routes: List[List[int]], source: int, target: int) -> int:
        if source == target:
            return 0
        at = defaultdict(list)                        # stop -> routes through it
        for i, route in enumerate(routes):
            for s in route:
                at[s].append(i)
        used, seen, q, buses = set(), {source}, deque([source]), 0
        while q:
            buses += 1
            for _ in range(len(q)):
                for r in at[q.popleft()]:
                    if r in used:
                        continue
                    used.add(r)
                    for s in routes[r]:
                        if s == target:
                            return buses
                        if s not in seen:
                            seen.add(s)
                            q.append(s)
        return -1
# test
assert Solution().numBusesToDestination([[1, 2, 7], [3, 6, 7]], 1, 6) == 2
assert Solution().numBusesToDestination([[7, 12], [4, 5, 15], [6], [15, 19], [9, 12, 13]], 15, 12) == -1
assert Solution().numBusesToDestination([[1, 2]], 1, 1) == 0
assert Solution().numBusesToDestination([[1, 2]], 3, 1) == -1
assert Solution().numBusesToDestination([[1, 7], [3, 5]], 5, 5) == 0
```

```java
class Solution {
    public int numBusesToDestination(int[][] routes, int source, int target) {
        if (source == target) return 0;
        Map<Integer, List<Integer>> at = new HashMap<>();
        for (int i = 0; i < routes.length; i++) for (int s : routes[i]) at.computeIfAbsent(s, k -> new ArrayList<>()).add(i);
        boolean[] used = new boolean[routes.length];
        Set<Integer> seen = new HashSet<>(List.of(source));
        Deque<Integer> q = new ArrayDeque<>(List.of(source));
        for (int buses = 1; !q.isEmpty(); buses++) {
            for (int i = q.size(); i > 0; i--) {
                for (int r : at.getOrDefault(q.poll(), List.of())) {
                    if (used[r]) continue;
                    used[r] = true;
                    for (int s : routes[r]) {
                        if (s == target) return buses;
                        if (seen.add(s)) q.add(s);
                    }
                }
            }
        }
        return -1;
    }
}
// test
check(new Solution().numBusesToDestination(new int[][]{{1, 2, 7}, {3, 6, 7}}, 1, 6), 2);
check(new Solution().numBusesToDestination(new int[][]{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, 15, 12), -1);
```

```go
func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	at := map[int][]int{} // stop -> routes through it
	for i, route := range routes {
		for _, s := range route {
			at[s] = append(at[s], i)
		}
	}
	used, seen, q := map[int]bool{}, map[int]bool{source: true}, []int{source}
	for buses := 1; len(q) > 0; buses++ { // BFS over stops, one level per bus taken
		next := []int{}
		for _, stop := range q {
			for _, r := range at[stop] {
				if used[r] {
					continue
				}
				used[r] = true
				for _, s := range routes[r] {
					if s == target {
						return buses
					}
					if !seen[s] {
						seen[s] = true
						next = append(next, s)
					}
				}
			}
		}
		q = next
	}
	return -1
}
// test
assertEq(numBusesToDestination([][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6), 2)
assertEq(numBusesToDestination([][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, 15, 12), -1)
assertEq(numBusesToDestination([][]int{{1, 2}}, 1, 1), 0)
assertEq(numBusesToDestination([][]int{{1, 2}}, 3, 1), -1)
assertEq(numBusesToDestination([][]int{{1, 7}, {3, 5}}, 5, 5), 0)
```

## 1926. Nearest Exit from Entrance in Maze | Medium
related: 1730, 286
tags: bfs, grid

BFS from the entrance; the first **border** empty cell reached (other than the entrance itself) is the nearest exit.

**Complexity:** O(R·C).

```python
class Solution:
    def nearestExit(self, maze: List[List[str]], entrance: List[int]) -> int:
        R, C = len(maze), len(maze[0])
        start = tuple(entrance)
        seen, q, steps = {start}, deque([start]), 0
        while q:
            steps += 1
            for _ in range(len(q)):
                r, c = q.popleft()
                for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                    if 0 <= nr < R and 0 <= nc < C and maze[nr][nc] == "." and (nr, nc) not in seen:
                        if nr in (0, R - 1) or nc in (0, C - 1):
                            return steps
                        seen.add((nr, nc))
                        q.append((nr, nc))
        return -1
# test
G = lambda rows: [list(r) for r in rows]
assert Solution().nearestExit(G(["++.+", "...+", "+++."]), [1, 2]) == 1
assert Solution().nearestExit(G(["+++", "...", "+++"]), [1, 0]) == 2
assert Solution().nearestExit(G([".+"]), [0, 0]) == -1
```

```java
class Solution {
    public int nearestExit(char[][] maze, int[] entrance) {
        int R = maze.length, C = maze[0].length;
        boolean[][] seen = new boolean[R][C];
        seen[entrance[0]][entrance[1]] = true;
        Deque<int[]> q = new ArrayDeque<>(List.of(entrance));
        int[][] dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        for (int steps = 1; !q.isEmpty(); steps++) {
            for (int i = q.size(); i > 0; i--) {
                int[] cur = q.poll();
                for (int[] d : dirs) {
                    int r = cur[0] + d[0], c = cur[1] + d[1];
                    if (r < 0 || c < 0 || r >= R || c >= C || seen[r][c] || maze[r][c] != '.') continue;
                    if (r == 0 || c == 0 || r == R - 1 || c == C - 1) return steps;
                    seen[r][c] = true;
                    q.add(new int[]{r, c});
                }
            }
        }
        return -1;
    }
}
// test
check(new Solution().nearestExit(new char[][]{"++.+".toCharArray(), "...+".toCharArray(), "+++.".toCharArray()}, new int[]{1, 2}), 1);
check(new Solution().nearestExit(new char[][]{".+".toCharArray()}, new int[]{0, 0}), -1);
```

```go
func nearestExit(maze [][]byte, entrance []int) int {
	R, C := len(maze), len(maze[0])
	start := [2]int{entrance[0], entrance[1]}
	seen, q := map[[2]int]bool{start: true}, [][2]int{start}
	for steps := 1; len(q) > 0; steps++ {
		next := [][2]int{}
		for _, p := range q {
			for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				nr, nc := p[0]+d[0], p[1]+d[1]
				if nr < 0 || nr >= R || nc < 0 || nc >= C || maze[nr][nc] != '.' || seen[[2]int{nr, nc}] {
					continue
				}
				if nr == 0 || nr == R-1 || nc == 0 || nc == C-1 { // a border cell that is not the entrance
					return steps
				}
				seen[[2]int{nr, nc}] = true
				next = append(next, [2]int{nr, nc})
			}
		}
		q = next
	}
	return -1
}
// test
assertEq(nearestExit(byteGrid("++.+", "...+", "+++."), []int{1, 2}), 1)
assertEq(nearestExit(byteGrid("+++", "...", "+++"), []int{1, 0}), 2)
assertEq(nearestExit(byteGrid(".+"), []int{0, 0}), -1)
```

## 463. Island Perimeter | Easy
related: 200, 695
tags: grid, counting

Every land cell contributes 4 edges; every pair of adjacent land cells hides 2 of them. Count land cells and right/down land neighbours: `4·land − 2·adjacent`.

**Complexity:** O(R·C).

```python
class Solution:
    def islandPerimeter(self, grid: List[List[int]]) -> int:
        land = shared = 0
        for r, row in enumerate(grid):
            for c, v in enumerate(row):
                if v:
                    land += 1
                    shared += (r + 1 < len(grid) and grid[r + 1][c]) + (c + 1 < len(row) and row[c + 1])
        return 4 * land - 2 * shared
# test
assert Solution().islandPerimeter([[0, 1, 0, 0], [1, 1, 1, 0], [0, 1, 0, 0], [1, 1, 0, 0]]) == 16
assert Solution().islandPerimeter([[1]]) == 4
assert Solution().islandPerimeter([[1, 0]]) == 4
assert Solution().islandPerimeter([[1, 1], [1, 1]]) == 8
```

```java
class Solution {
    public int islandPerimeter(int[][] grid) {
        int land = 0, shared = 0;
        for (int r = 0; r < grid.length; r++)
            for (int c = 0; c < grid[0].length; c++)
                if (grid[r][c] == 1) {
                    land++;
                    if (r + 1 < grid.length && grid[r + 1][c] == 1) shared++;
                    if (c + 1 < grid[0].length && grid[r][c + 1] == 1) shared++;
                }
        return 4 * land - 2 * shared;
    }
}
// test
check(new Solution().islandPerimeter(new int[][]{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}), 16);
```

```go
func islandPerimeter(grid [][]int) int {
	land, shared := 0, 0
	for r, row := range grid {
		for c, v := range row {
			if v == 1 {
				land++
				if r+1 < len(grid) && grid[r+1][c] == 1 { // each shared edge hides two sides
					shared++
				}
				if c+1 < len(row) && row[c+1] == 1 {
					shared++
				}
			}
		}
	}
	return 4*land - 2*shared
}
// test
assertEq(islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}), 16)
assertEq(islandPerimeter([][]int{{1}}), 4)
assertEq(islandPerimeter([][]int{{1, 0}}), 4)
assertEq(islandPerimeter([][]int{{1, 1}, {1, 1}}), 8)
```

## 997. Find the Town Judge | Easy
related: 277, 1971
tags: graph, degree

The judge is trusted by everyone else (in-degree n − 1) and trusts nobody (out-degree 0). One score per person, `in − out`, equals n − 1 only for the judge.

**Complexity:** O(n + |trust|).

```python
class Solution:
    def findJudge(self, n: int, trust: List[List[int]]) -> int:
        score = [0] * (n + 1)
        for a, b in trust:
            score[a] -= 1
            score[b] += 1
        return next((p for p in range(1, n + 1) if score[p] == n - 1), -1)
# test
assert Solution().findJudge(2, [[1, 2]]) == 2
assert Solution().findJudge(3, [[1, 3], [2, 3]]) == 3
assert Solution().findJudge(3, [[1, 3], [2, 3], [3, 1]]) == -1
assert Solution().findJudge(1, []) == 1
```

```java
class Solution {
    public int findJudge(int n, int[][] trust) {
        int[] score = new int[n + 1];
        for (int[] t : trust) { score[t[0]]--; score[t[1]]++; }
        for (int p = 1; p <= n; p++) if (score[p] == n - 1) return p;
        return -1;
    }
}
// test
check(new Solution().findJudge(3, new int[][]{{1, 3}, {2, 3}}), 3);
check(new Solution().findJudge(3, new int[][]{{1, 3}, {2, 3}, {3, 1}}), -1);
```

```go
func findJudge(n int, trust [][]int) int {
	score := make([]int, n+1) // trusted-by count minus trusts count
	for _, t := range trust {
		score[t[0]]--
		score[t[1]]++
	}
	for p := 1; p <= n; p++ {
		if score[p] == n-1 {
			return p
		}
	}
	return -1
}
// test
assertEq(findJudge(2, [][]int{{1, 2}}), 2)
assertEq(findJudge(3, [][]int{{1, 3}, {2, 3}}), 3)
assertEq(findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}), -1)
assertEq(findJudge(1, [][]int{}), 1)
```

## 1462. Course Schedule IV | Medium
related: 207, 210, 399
tags: transitive-closure, floyd-warshall

Many reachability queries on a small graph (n ≤ 100) → precompute the **transitive closure** once: Floyd–Warshall on booleans (`reach[i][j] |= reach[i][k] and reach[k][j]`), then answer each query in O(1). (A BFS from every node or a topological pass with bitsets also works.)

**Complexity:** O(n³ + q).

```python
class Solution:
    def checkIfPrerequisite(self, numCourses: int, prerequisites: List[List[int]], queries: List[List[int]]) -> List[bool]:
        n = numCourses
        reach = [[False] * n for _ in range(n)]
        for a, b in prerequisites:
            reach[a][b] = True
        for k in range(n):
            for i in range(n):
                if reach[i][k]:
                    row_k, row_i = reach[k], reach[i]
                    for j in range(n):
                        if row_k[j]:
                            row_i[j] = True
        return [reach[u][v] for u, v in queries]
# test
assert Solution().checkIfPrerequisite(2, [[1, 0]], [[0, 1], [1, 0]]) == [False, True]
assert Solution().checkIfPrerequisite(2, [], [[1, 0], [0, 1]]) == [False, False]
assert Solution().checkIfPrerequisite(3, [[1, 2], [1, 0], [2, 0]], [[1, 0], [1, 2]]) == [True, True]
assert Solution().checkIfPrerequisite(4, [[0, 1], [1, 2], [2, 3]], [[0, 3], [3, 0], [1, 3]]) == [True, False, True]
```

```java
class Solution {
    public List<Boolean> checkIfPrerequisite(int numCourses, int[][] prerequisites, int[][] queries) {
        int n = numCourses;
        boolean[][] reach = new boolean[n][n];
        for (int[] p : prerequisites) reach[p[0]][p[1]] = true;
        for (int k = 0; k < n; k++)
            for (int i = 0; i < n; i++)
                if (reach[i][k]) for (int j = 0; j < n; j++) if (reach[k][j]) reach[i][j] = true;
        List<Boolean> res = new ArrayList<>();
        for (int[] q : queries) res.add(reach[q[0]][q[1]]);
        return res;
    }
}
// test
check(new Solution().checkIfPrerequisite(3, new int[][]{{1, 2}, {1, 0}, {2, 0}}, new int[][]{{1, 0}, {1, 2}}), List.of(true, true));
```

```go
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	n := numCourses
	reach := make([][]bool, n)
	for i := range reach {
		reach[i] = make([]bool, n)
	}
	for _, p := range prerequisites {
		reach[p[0]][p[1]] = true
	}
	for k := 0; k < n; k++ { // Floyd-Warshall transitive closure
		for i := 0; i < n; i++ {
			if reach[i][k] {
				for j := 0; j < n; j++ {
					if reach[k][j] {
						reach[i][j] = true
					}
				}
			}
		}
	}
	res := make([]bool, len(queries))
	for i, q := range queries {
		res[i] = reach[q[0]][q[1]]
	}
	return res
}
// test
assertEq(checkIfPrerequisite(2, [][]int{{1, 0}}, [][]int{{0, 1}, {1, 0}}), []bool{false, true})
assertEq(checkIfPrerequisite(2, [][]int{}, [][]int{{1, 0}, {0, 1}}), []bool{false, false})
assertEq(checkIfPrerequisite(3, [][]int{{1, 2}, {1, 0}, {2, 0}}, [][]int{{1, 0}, {1, 2}}), []bool{true, true})
assertEq(checkIfPrerequisite(4, [][]int{{0, 1}, {1, 2}, {2, 3}}, [][]int{{0, 3}, {3, 0}, {1, 3}}), []bool{true, false, true})
```
