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
