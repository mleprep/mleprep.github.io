# Advanced Graphs

| Problem | Algorithm | Complexity |
|---|---|---|
| Single-source shortest path, weights ≥ 0 | **Dijkstra** (min-heap, lazy deletion) | O(E log V) |
| Shortest path with ≤ k edges / negative weights | **Bellman–Ford** (k+1 relaxation rounds) | O(k·E) |
| Minimum spanning tree | **Prim** (heap) or **Kruskal** (sort + union-find) | O(E log V) |
| Use every edge once | **Hierholzer** (Eulerian path) | O(E log E) |
| Bridges / articulation points | **Tarjan** (discovery time + low-link) | O(V + E) |
| Minimax path ("max edge on path is minimal") | Dijkstra on max, or binary search + BFS | O(E log V) |

```python
# Dijkstra
dist = {src: 0}
pq = [(0, src)]
while pq:
    d, u = heapq.heappop(pq)
    if d > dist.get(u, inf): continue          # stale entry
    for v, w in graph[u]:
        if d + w < dist.get(v, inf):
            dist[v] = d + w
            heapq.heappush(pq, (d + w, v))
```

## 743. Network Delay Time | Medium
related: 787, 1631
tags: dijkstra

Dijkstra from `k`; answer = max shortest distance, or −1 if some node is unreachable.

**Complexity:** O(E log V).

```python
class Solution:
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        g = defaultdict(list)
        for u, v, w in times:
            g[u].append((v, w))
        dist = {}
        pq = [(0, k)]
        while pq:
            d, u = heappop(pq)
            if u in dist:
                continue
            dist[u] = d
            for v, w in g[u]:
                if v not in dist:
                    heappush(pq, (d + w, v))
        return max(dist.values()) if len(dist) == n else -1
# test
assert Solution().networkDelayTime([[2, 1, 1], [2, 3, 1], [3, 4, 1]], 4, 2) == 2
assert Solution().networkDelayTime([[1, 2, 1]], 2, 2) == -1
```

```java
class Solution {
    public int networkDelayTime(int[][] times, int n, int k) {
        List<List<int[]>> g = new ArrayList<>();
        for (int i = 0; i <= n; i++) g.add(new ArrayList<>());
        for (int[] t : times) g.get(t[0]).add(new int[]{t[1], t[2]});
        int[] dist = new int[n + 1];
        Arrays.fill(dist, Integer.MAX_VALUE);
        dist[k] = 0;
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[0] - b[0]); // {dist, node}
        pq.add(new int[]{0, k});
        while (!pq.isEmpty()) {
            int[] cur = pq.poll();
            int d = cur[0], u = cur[1];
            if (d > dist[u]) continue;
            for (int[] e : g.get(u))
                if (d + e[1] < dist[e[0]]) { dist[e[0]] = d + e[1]; pq.add(new int[]{dist[e[0]], e[0]}); }
        }
        int best = 0;
        for (int i = 1; i <= n; i++) best = Math.max(best, dist[i]);
        return best == Integer.MAX_VALUE ? -1 : best;
    }
}
// test
check(new Solution().networkDelayTime(new int[][]{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2), 2);
check(new Solution().networkDelayTime(new int[][]{{1, 2, 1}}, 2, 2), -1);
```

## 787. Cheapest Flights Within K Stops | Medium
related: 743
tags: bellman-ford

At most `k` stops = at most `k+1` edges → run `k+1` rounds of Bellman–Ford relaxation, each round relaxing from a **copy** of the previous distances (so one round adds at most one edge).

```viz
round 0: dist {0:0}
round 1: dist {0:0, 1:100}
round 2: dist {0:0, 1:100, 2:200, 3:700}    k=1 -> answer dist[3] = 700
```

**Complexity:** O(k · E).

```python
class Solution:
    def findCheapestPrice(self, n: int, flights: List[List[int]], src: int, dst: int, k: int) -> int:
        INF = float("inf")
        dist = [INF] * n
        dist[src] = 0
        for _ in range(k + 1):
            nxt = dist[:]
            for u, v, w in flights:
                if dist[u] + w < nxt[v]:
                    nxt[v] = dist[u] + w
            dist = nxt
        return -1 if dist[dst] == INF else dist[dst]
# test
F = [[0, 1, 100], [1, 2, 100], [2, 0, 100], [1, 3, 600], [2, 3, 200]]
assert Solution().findCheapestPrice(4, F, 0, 3, 1) == 700
assert Solution().findCheapestPrice(3, [[0, 1, 100], [1, 2, 100], [0, 2, 500]], 0, 2, 0) == 500
```

```java
class Solution {
    public int findCheapestPrice(int n, int[][] flights, int src, int dst, int k) {
        int INF = Integer.MAX_VALUE;
        int[] dist = new int[n];
        Arrays.fill(dist, INF);
        dist[src] = 0;
        for (int i = 0; i <= k; i++) {
            int[] nxt = dist.clone();
            for (int[] f : flights)
                if (dist[f[0]] != INF && dist[f[0]] + f[2] < nxt[f[1]]) nxt[f[1]] = dist[f[0]] + f[2];
            dist = nxt;
        }
        return dist[dst] == INF ? -1 : dist[dst];
    }
}
// test
int[][] F = {{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}};
check(new Solution().findCheapestPrice(4, F, 0, 3, 1), 700);
```

## 1584. Min Cost to Connect All Points | Medium
related: 1135
tags: mst, prim

Complete graph with Manhattan weights → **Prim's** O(n²) array version (no heap needed for dense graphs): repeatedly add the closest outside point, then update distances.

**Complexity:** O(n²).

```python
class Solution:
    def minCostConnectPoints(self, points: List[List[int]]) -> int:
        n = len(points)
        INF = float("inf")
        dist = [INF] * n
        dist[0] = 0
        used = [False] * n
        total = 0
        for _ in range(n):
            u = min((d, i) for i, d in enumerate(dist) if not used[i])[1]
            used[u] = True
            total += dist[u]
            ux, uy = points[u]
            for v in range(n):
                if not used[v]:
                    dist[v] = min(dist[v], abs(points[v][0] - ux) + abs(points[v][1] - uy))
        return total
# test
assert Solution().minCostConnectPoints([[0, 0], [2, 2], [3, 10], [5, 2], [7, 0]]) == 20
assert Solution().minCostConnectPoints([[3, 12], [-2, 5], [-4, 1]]) == 18
```

```java
class Solution {
    public int minCostConnectPoints(int[][] points) {
        int n = points.length, total = 0;
        int[] dist = new int[n];
        Arrays.fill(dist, Integer.MAX_VALUE);
        dist[0] = 0;
        boolean[] used = new boolean[n];
        for (int it = 0; it < n; it++) {
            int u = -1;
            for (int i = 0; i < n; i++) if (!used[i] && (u == -1 || dist[i] < dist[u])) u = i;
            used[u] = true;
            total += dist[u];
            for (int v = 0; v < n; v++)
                if (!used[v])
                    dist[v] = Math.min(dist[v], Math.abs(points[u][0] - points[v][0]) + Math.abs(points[u][1] - points[v][1]));
        }
        return total;
    }
}
// test
check(new Solution().minCostConnectPoints(new int[][]{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}), 20);
```

## 332. Reconstruct Itinerary | Hard
related: 753
tags: eulerian-path, hierholzer

Use every ticket exactly once = **Eulerian path**. Hierholzer: DFS taking the smallest destination first, **append a node after all its edges are used**, then reverse. Dead ends get placed at the end automatically.

```viz
JFK -> KUL (dead end) and JFK -> NRT -> JFK
DFS JFK: go KUL (smallest) -> stuck -> post-order [KUL]
back at JFK: go NRT -> JFK -> stuck ... -> [KUL, JFK, NRT, JFK]
reverse -> JFK NRT JFK KUL
```

**Complexity:** O(E log E).

```python
class Solution:
    def findItinerary(self, tickets: List[List[str]]) -> List[str]:
        g = defaultdict(list)
        for a, b in sorted(tickets, reverse=True):
            g[a].append(b)               # reverse-sorted so pop() gives smallest
        route, st = [], ["JFK"]
        while st:
            while g[st[-1]]:
                st.append(g[st[-1]].pop())
            route.append(st.pop())
        return route[::-1]
# test
T = [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
assert Solution().findItinerary(T) == ["JFK", "MUC", "LHR", "SFO", "SJC"]
assert Solution().findItinerary([["JFK", "KUL"], ["JFK", "NRT"], ["NRT", "JFK"]]) == ["JFK", "NRT", "JFK", "KUL"]
```

```java
class Solution {
    public List<String> findItinerary(List<List<String>> tickets) {
        Map<String, PriorityQueue<String>> g = new HashMap<>();
        for (List<String> t : tickets) g.computeIfAbsent(t.get(0), k -> new PriorityQueue<>()).add(t.get(1));
        LinkedList<String> route = new LinkedList<>();
        Deque<String> st = new ArrayDeque<>(List.of("JFK"));
        while (!st.isEmpty()) {
            PriorityQueue<String> out = g.get(st.peek());
            if (out != null && !out.isEmpty()) st.push(out.poll());
            else route.addFirst(st.pop());
        }
        return route;
    }
}
// test
check(new Solution().findItinerary(List.of(List.of("JFK", "KUL"), List.of("JFK", "NRT"), List.of("NRT", "JFK"))), "[JFK, NRT, JFK, KUL]");
```

## 778. Swim in Rising Water | Hard
related: 1631, 743
tags: dijkstra, minimax

Cost of a path = the **max** cell on it. Dijkstra where the "distance" is `max(current, grid[next])`: always expand the lowest-elevation frontier cell.

**Complexity:** O(n² log n).

```python
class Solution:
    def swimInWater(self, grid: List[List[int]]) -> int:
        n = len(grid)
        pq, seen = [(grid[0][0], 0, 0)], {(0, 0)}
        while pq:
            t, r, c = heappop(pq)
            if r == c == n - 1:
                return t
            for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                if 0 <= nr < n and 0 <= nc < n and (nr, nc) not in seen:
                    seen.add((nr, nc))
                    heappush(pq, (max(t, grid[nr][nc]), nr, nc))
# test
assert Solution().swimInWater([[0, 2], [1, 3]]) == 3
G = [[0, 1, 2, 3, 4], [24, 23, 22, 21, 5], [12, 13, 14, 15, 16], [11, 17, 18, 19, 20], [10, 9, 8, 7, 6]]
assert Solution().swimInWater(G) == 16
```

```java
class Solution {
    public int swimInWater(int[][] grid) {
        int n = grid.length;
        boolean[][] seen = new boolean[n][n];
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[0] - b[0]); // {time, r, c}
        pq.add(new int[]{grid[0][0], 0, 0});
        seen[0][0] = true;
        int[][] dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        while (true) {
            int[] cur = pq.poll();
            if (cur[1] == n - 1 && cur[2] == n - 1) return cur[0];
            for (int[] d : dirs) {
                int nr = cur[1] + d[0], nc = cur[2] + d[1];
                if (nr >= 0 && nc >= 0 && nr < n && nc < n && !seen[nr][nc]) {
                    seen[nr][nc] = true;
                    pq.add(new int[]{Math.max(cur[0], grid[nr][nc]), nr, nc});
                }
            }
        }
    }
}
// test
check(new Solution().swimInWater(new int[][]{{0, 2}, {1, 3}}), 3);
```

## 269. Alien Dictionary | Hard
related: 207, 210, 953
tags: blind75, topological-sort, premium

Compare adjacent words: the first differing letters give an edge `a → b`. If a word is followed by its own proper prefix, the input is invalid. Topologically sort the letters (Kahn); a cycle → `""`.

```viz
wrt, wrf  -> t < f
wrf, er   -> w < e
er, ett   -> r < t
ett, rftt -> e < r          order: w e r t f
```

**Complexity:** O(total characters).

```python
class Solution:
    def alienOrder(self, words: List[str]) -> str:
        g = {ch: set() for w in words for ch in w}
        indeg = {ch: 0 for ch in g}
        for a, b in zip(words, words[1:]):
            for x, y in zip(a, b):
                if x != y:
                    if y not in g[x]:
                        g[x].add(y)
                        indeg[y] += 1
                    break
            else:
                if len(a) > len(b):
                    return ""
        q = deque(sorted(ch for ch in indeg if indeg[ch] == 0))
        out = []
        while q:
            u = q.popleft()
            out.append(u)
            for v in sorted(g[u]):
                indeg[v] -= 1
                if indeg[v] == 0:
                    q.append(v)
        return "".join(out) if len(out) == len(g) else ""
# test
assert Solution().alienOrder(["wrt", "wrf", "er", "ett", "rftt"]) == "wertf"
assert Solution().alienOrder(["z", "x", "z"]) == ""
assert Solution().alienOrder(["abc", "ab"]) == ""
```

```java
class Solution {
    public String alienOrder(String[] words) {
        Map<Character, Set<Character>> g = new TreeMap<>();
        Map<Character, Integer> indeg = new TreeMap<>();
        for (String w : words) for (char c : w.toCharArray()) { g.putIfAbsent(c, new TreeSet<>()); indeg.putIfAbsent(c, 0); }
        for (int i = 0; i + 1 < words.length; i++) {
            String a = words[i], b = words[i + 1];
            int j = 0;
            while (j < Math.min(a.length(), b.length()) && a.charAt(j) == b.charAt(j)) j++;
            if (j == Math.min(a.length(), b.length())) {
                if (a.length() > b.length()) return "";
                continue;
            }
            if (g.get(a.charAt(j)).add(b.charAt(j))) indeg.merge(b.charAt(j), 1, Integer::sum);
        }
        Deque<Character> q = new ArrayDeque<>();
        for (Map.Entry<Character, Integer> e : indeg.entrySet()) if (e.getValue() == 0) q.add(e.getKey());
        StringBuilder sb = new StringBuilder();
        while (!q.isEmpty()) {
            char u = q.poll();
            sb.append(u);
            for (char v : g.get(u)) if (indeg.merge(v, -1, Integer::sum) == 0) q.add(v);
        }
        return sb.length() == g.size() ? sb.toString() : "";
    }
}
// test
check(new Solution().alienOrder(new String[]{"wrt", "wrf", "er", "ett", "rftt"}), "wertf");
check(new Solution().alienOrder(new String[]{"z", "x", "z"}), "");
```

## 1631. Path With Minimum Effort | Medium
related: 778, 743
tags: dijkstra, minimax

Effort of a path = max |height difference| along it. Dijkstra with `cost = max(cost, |Δh|)`.

**Complexity:** O(R·C log(R·C)).

```python
class Solution:
    def minimumEffortPath(self, heights: List[List[int]]) -> int:
        R, C = len(heights), len(heights[0])
        best = [[float("inf")] * C for _ in range(R)]
        best[0][0] = 0
        pq = [(0, 0, 0)]
        while pq:
            e, r, c = heappop(pq)
            if (r, c) == (R - 1, C - 1):
                return e
            if e > best[r][c]:
                continue
            for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                if 0 <= nr < R and 0 <= nc < C:
                    ne = max(e, abs(heights[nr][nc] - heights[r][c]))
                    if ne < best[nr][nc]:
                        best[nr][nc] = ne
                        heappush(pq, (ne, nr, nc))
# test
assert Solution().minimumEffortPath([[1, 2, 2], [3, 8, 2], [5, 3, 5]]) == 2
assert Solution().minimumEffortPath([[1, 2, 1, 1, 1], [1, 2, 1, 2, 1], [1, 2, 1, 2, 1], [1, 2, 1, 2, 1], [1, 1, 1, 2, 1]]) == 0
```

```java
class Solution {
    public int minimumEffortPath(int[][] heights) {
        int R = heights.length, C = heights[0].length;
        int[][] best = new int[R][C];
        for (int[] row : best) Arrays.fill(row, Integer.MAX_VALUE);
        best[0][0] = 0;
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[0] - b[0]);
        pq.add(new int[]{0, 0, 0});
        int[][] dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        while (!pq.isEmpty()) {
            int[] cur = pq.poll();
            int e = cur[0], r = cur[1], c = cur[2];
            if (r == R - 1 && c == C - 1) return e;
            if (e > best[r][c]) continue;
            for (int[] d : dirs) {
                int nr = r + d[0], nc = c + d[1];
                if (nr < 0 || nc < 0 || nr >= R || nc >= C) continue;
                int ne = Math.max(e, Math.abs(heights[nr][nc] - heights[r][c]));
                if (ne < best[nr][nc]) { best[nr][nc] = ne; pq.add(new int[]{ne, nr, nc}); }
            }
        }
        return 0;
    }
}
// test
check(new Solution().minimumEffortPath(new int[][]{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}), 2);
```

## 310. Minimum Height Trees | Medium
related: 207
tags: topological-sort, leaf-trimming

The best roots are the tree's **centers** (1 or 2 nodes). Peel leaves layer by layer (like Kahn's algorithm on degree 1) until ≤ 2 nodes remain.

```viz
0 - 1 - 2 - 3 - 4      trim leaves {0,4} -> {1,3} -> remaining {2}  = center
```

**Complexity:** O(n).

```python
class Solution:
    def findMinHeightTrees(self, n: int, edges: List[List[int]]) -> List[int]:
        if n <= 2:
            return list(range(n))
        g = defaultdict(set)
        for a, b in edges:
            g[a].add(b)
            g[b].add(a)
        leaves = [i for i in range(n) if len(g[i]) == 1]
        remaining = n
        while remaining > 2:
            remaining -= len(leaves)
            nxt = []
            for leaf in leaves:
                nb = g[leaf].pop()
                g[nb].remove(leaf)
                if len(g[nb]) == 1:
                    nxt.append(nb)
            leaves = nxt
        return sorted(leaves)
# test
assert sorted(Solution().findMinHeightTrees(4, [[1, 0], [1, 2], [1, 3]])) == sorted([1])
assert sorted(Solution().findMinHeightTrees(6, [[3, 0], [3, 1], [3, 2], [3, 4], [5, 4]])) == sorted([3, 4])
```

```java
class Solution {
    public List<Integer> findMinHeightTrees(int n, int[][] edges) {
        if (n <= 2) { List<Integer> r = new ArrayList<>(); for (int i = 0; i < n; i++) r.add(i); return r; }
        List<Set<Integer>> g = new ArrayList<>();
        for (int i = 0; i < n; i++) g.add(new HashSet<>());
        for (int[] e : edges) { g.get(e[0]).add(e[1]); g.get(e[1]).add(e[0]); }
        List<Integer> leaves = new ArrayList<>();
        for (int i = 0; i < n; i++) if (g.get(i).size() == 1) leaves.add(i);
        int remaining = n;
        while (remaining > 2) {
            remaining -= leaves.size();
            List<Integer> nxt = new ArrayList<>();
            for (int leaf : leaves) {
                int nb = g.get(leaf).iterator().next();
                g.get(nb).remove(leaf);
                if (g.get(nb).size() == 1) nxt.add(nb);
            }
            leaves = nxt;
        }
        Collections.sort(leaves);
        return leaves;
    }
}
// test
check(new Solution().findMinHeightTrees(6, new int[][]{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}), "[3, 4]");
```

## 1192. Critical Connections in a Network | Hard
related: 684
tags: tarjan, bridges

**Tarjan's bridge-finding**: DFS assigns discovery time `tin[u]`; `low[u]` = smallest `tin` reachable from u's subtree using at most one back edge. Tree edge `(u, v)` is a bridge iff `low[v] > tin[u]` (v's subtree can't reach u or above without it).

```viz
0 - 1 - 3            tin: 0:0 1:1 2:2 3:3
 \ /                 low: 0:0 1:0 2:0 3:3
  2                  low[3]=3 > tin[1]=1  -> bridge (1,3)
```

**Complexity:** O(V + E).

```python
class Solution:
    def criticalConnections(self, n: int, connections: List[List[int]]) -> List[List[int]]:
        g = defaultdict(list)
        for a, b in connections:
            g[a].append(b)
            g[b].append(a)
        tin, low, res, timer = [-1] * n, [0] * n, [], [0]

        def dfs(u, parent):
            tin[u] = low[u] = timer[0]
            timer[0] += 1
            for v in g[u]:
                if v == parent:
                    continue
                if tin[v] == -1:
                    dfs(v, u)
                    low[u] = min(low[u], low[v])
                    if low[v] > tin[u]:
                        res.append([u, v])
                else:
                    low[u] = min(low[u], tin[v])

        dfs(0, -1)
        return res
# test
assert [sorted(e) for e in Solution().criticalConnections(4, [[0, 1], [1, 2], [2, 0], [1, 3]])] == [[1, 3]]
assert sorted(sorted(e) for e in Solution().criticalConnections(2, [[0, 1]])) == [[0, 1]]
assert sorted(sorted(e) for e in Solution().criticalConnections(6, [[0, 1], [1, 2], [2, 0], [1, 3], [3, 4], [4, 5], [5, 3]])) == [[1, 3]]
assert sorted(sorted(e) for e in Solution().criticalConnections(4, [[0, 1], [1, 2], [2, 3]])) == [[0, 1], [1, 2], [2, 3]]
```

```java
class Solution {
    private int timer = 0;
    private int[] tin, low;
    private List<List<Integer>> g;
    private final List<List<Integer>> res = new ArrayList<>();

    public List<List<Integer>> criticalConnections(int n, List<List<Integer>> connections) {
        g = new ArrayList<>();
        for (int i = 0; i < n; i++) g.add(new ArrayList<>());
        for (List<Integer> e : connections) { g.get(e.get(0)).add(e.get(1)); g.get(e.get(1)).add(e.get(0)); }
        tin = new int[n];
        low = new int[n];
        Arrays.fill(tin, -1);
        dfs(0, -1);
        return res;
    }

    private void dfs(int u, int parent) {
        tin[u] = low[u] = timer++;
        for (int v : g.get(u)) {
            if (v == parent) continue;
            if (tin[v] == -1) {
                dfs(v, u);
                low[u] = Math.min(low[u], low[v]);
                if (low[v] > tin[u]) res.add(List.of(u, v));
            } else low[u] = Math.min(low[u], tin[v]);
        }
    }
}
// test
check(new Solution().criticalConnections(4, List.of(List.of(0, 1), List.of(1, 2), List.of(2, 0), List.of(1, 3))), "[[1, 3]]");
```

## 802. Find Eventual Safe States | Medium
related: 207
tags: cycle-detection, three-color-dfs

A node is safe iff no path from it reaches a cycle. **Three-color DFS**: 0 = unvisited, 1 = on stack, 2 = safe. Reaching a node with color 1 means a cycle → unsafe. (Alternatively: Kahn's on the reversed graph.)

**Complexity:** O(V + E).

```python
class Solution:
    def eventualSafeNodes(self, graph: List[List[int]]) -> List[int]:
        color = [0] * len(graph)

        def safe(u):
            if color[u]:
                return color[u] == 2
            color[u] = 1
            if all(safe(v) for v in graph[u]):
                color[u] = 2
                return True
            return False          # stays 1 -> unsafe

        return [u for u in range(len(graph)) if safe(u)]
# test
assert Solution().eventualSafeNodes([[1, 2], [2, 3], [5], [0], [5], [], []]) == [2, 4, 5, 6]
assert Solution().eventualSafeNodes([[1, 2, 3, 4], [1, 2], [3, 4], [0, 4], []]) == [4]
```

```java
class Solution {
    public List<Integer> eventualSafeNodes(int[][] graph) {
        int[] color = new int[graph.length];
        List<Integer> res = new ArrayList<>();
        for (int u = 0; u < graph.length; u++) if (safe(graph, u, color)) res.add(u);
        return res;
    }

    private boolean safe(int[][] g, int u, int[] color) {
        if (color[u] != 0) return color[u] == 2;
        color[u] = 1;
        for (int v : g[u]) if (!safe(g, v, color)) return false;
        color[u] = 2;
        return true;
    }
}
// test
check(new Solution().eventualSafeNodes(new int[][]{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}), "[2, 4, 5, 6]");
```


## 1514. Path with Maximum Probability | Medium
related: 743
tags: dijkstra

Dijkstra with a **max-heap** on probability: probabilities only shrink along a path (multiplying by ≤ 1), so the greedy "best first" argument still holds.

**Complexity:** O(E log V).

```python
class Solution:
    def maxProbability(self, n: int, edges: List[List[int]], succProb: List[float], start_node: int, end_node: int) -> float:
        g = defaultdict(list)
        for (a, b), p in zip(edges, succProb):
            g[a].append((b, p))
            g[b].append((a, p))
        best = [0.0] * n
        best[start_node] = 1.0
        pq = [(-1.0, start_node)]
        while pq:
            p, u = heappop(pq)
            p = -p
            if u == end_node:
                return p
            if p < best[u]:
                continue
            for v, w in g[u]:
                if p * w > best[v]:
                    best[v] = p * w
                    heappush(pq, (-best[v], v))
        return 0.0
# test
assert abs(Solution().maxProbability(3, [[0, 1], [1, 2], [0, 2]], [0.5, 0.5, 0.2], 0, 2) - 0.25) < 1e-9
assert Solution().maxProbability(3, [[0, 1]], [0.5], 0, 2) == 0.0
```

```java
class Solution {
    public double maxProbability(int n, int[][] edges, double[] succProb, int start, int end) {
        List<List<double[]>> g = new ArrayList<>();
        for (int i = 0; i < n; i++) g.add(new ArrayList<>());
        for (int i = 0; i < edges.length; i++) {
            g.get(edges[i][0]).add(new double[]{edges[i][1], succProb[i]});
            g.get(edges[i][1]).add(new double[]{edges[i][0], succProb[i]});
        }
        double[] best = new double[n];
        best[start] = 1;
        PriorityQueue<double[]> pq = new PriorityQueue<>((a, b) -> Double.compare(b[0], a[0]));
        pq.add(new double[]{1, start});
        while (!pq.isEmpty()) {
            double[] cur = pq.poll();
            int u = (int) cur[1];
            if (u == end) return cur[0];
            if (cur[0] < best[u]) continue;
            for (double[] e : g.get(u)) {
                int v = (int) e[0];
                if (cur[0] * e[1] > best[v]) { best[v] = cur[0] * e[1]; pq.add(new double[]{best[v], v}); }
            }
        }
        return 0;
    }
}
// test
check(new Solution().maxProbability(3, new int[][]{{0, 1}, {1, 2}, {0, 2}}, new double[]{0.5, 0.5, 0.2}, 0, 2), 0.25);
```

## 1976. Number of Ways to Arrive at Destination | Medium
related: 743
tags: dijkstra, counting

Dijkstra while counting shortest paths: a strictly shorter distance to `v` resets `ways[v] = ways[u]`; an equal distance adds `ways[u]`.

**Complexity:** O(E log V).

```python
class Solution:
    def countPaths(self, n: int, roads: List[List[int]]) -> int:
        MOD = 10 ** 9 + 7
        g = defaultdict(list)
        for a, b, t in roads:
            g[a].append((b, t))
            g[b].append((a, t))
        dist, ways = [float("inf")] * n, [0] * n
        dist[0], ways[0] = 0, 1
        pq = [(0, 0)]
        while pq:
            d, u = heappop(pq)
            if d > dist[u]:
                continue
            for v, t in g[u]:
                if d + t < dist[v]:
                    dist[v], ways[v] = d + t, ways[u]
                    heappush(pq, (dist[v], v))
                elif d + t == dist[v]:
                    ways[v] = (ways[v] + ways[u]) % MOD
        return ways[n - 1]
# test
R = [[0, 6, 7], [0, 1, 2], [1, 2, 3], [1, 3, 3], [6, 3, 3], [3, 5, 1], [6, 5, 1], [2, 5, 1], [0, 4, 5], [4, 6, 2]]
assert Solution().countPaths(7, R) == 4
assert Solution().countPaths(2, [[1, 0, 10]]) == 1
```

```java
class Solution {
    public int countPaths(int n, int[][] roads) {
        final int MOD = 1_000_000_007;
        List<List<int[]>> g = new ArrayList<>();
        for (int i = 0; i < n; i++) g.add(new ArrayList<>());
        for (int[] r : roads) { g.get(r[0]).add(new int[]{r[1], r[2]}); g.get(r[1]).add(new int[]{r[0], r[2]}); }
        long[] dist = new long[n];
        Arrays.fill(dist, Long.MAX_VALUE);
        long[] ways = new long[n];
        dist[0] = 0; ways[0] = 1;
        PriorityQueue<long[]> pq = new PriorityQueue<>((a, b) -> Long.compare(a[0], b[0]));
        pq.add(new long[]{0, 0});
        while (!pq.isEmpty()) {
            long[] cur = pq.poll();
            int u = (int) cur[1];
            if (cur[0] > dist[u]) continue;
            for (int[] e : g.get(u)) {
                long nd = cur[0] + e[1];
                if (nd < dist[e[0]]) { dist[e[0]] = nd; ways[e[0]] = ways[u]; pq.add(new long[]{nd, e[0]}); }
                else if (nd == dist[e[0]]) ways[e[0]] = (ways[e[0]] + ways[u]) % MOD;
            }
        }
        return (int) ways[n - 1];
    }
}
// test
int[][] R = {{0, 6, 7}, {0, 1, 2}, {1, 2, 3}, {1, 3, 3}, {6, 3, 3}, {3, 5, 1}, {6, 5, 1}, {2, 5, 1}, {0, 4, 5}, {4, 6, 2}};
check(new Solution().countPaths(7, R), 4);
```

## 1334. Find the City With the Smallest Number of Neighbors at a Threshold Distance | Medium
related: 743
tags: floyd-warshall

All-pairs shortest paths with **Floyd–Warshall** (`dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])` for every intermediate k). Count reachable cities within the threshold; ties go to the larger index.

**Complexity:** O(n³).

```python
class Solution:
    def findTheCity(self, n: int, edges: List[List[int]], distanceThreshold: int) -> int:
        INF = float("inf")
        d = [[0 if i == j else INF for j in range(n)] for i in range(n)]
        for a, b, w in edges:
            d[a][b] = d[b][a] = w
        for k in range(n):
            for i in range(n):
                for j in range(n):
                    if d[i][k] + d[k][j] < d[i][j]:
                        d[i][j] = d[i][k] + d[k][j]
        best, city = INF, -1
        for i in range(n):
            cnt = sum(d[i][j] <= distanceThreshold for j in range(n)) - 1
            if cnt <= best:
                best, city = cnt, i
        return city
# test
assert Solution().findTheCity(4, [[0, 1, 3], [1, 2, 1], [1, 3, 4], [2, 3, 1]], 4) == 3
assert Solution().findTheCity(5, [[0, 1, 2], [0, 4, 8], [1, 2, 3], [1, 4, 2], [2, 3, 1], [3, 4, 1]], 2) == 0
```

```java
class Solution {
    public int findTheCity(int n, int[][] edges, int distanceThreshold) {
        int INF = 1_000_000_000;
        int[][] d = new int[n][n];
        for (int[] row : d) Arrays.fill(row, INF);
        for (int i = 0; i < n; i++) d[i][i] = 0;
        for (int[] e : edges) d[e[0]][e[1]] = d[e[1]][e[0]] = e[2];
        for (int k = 0; k < n; k++)
            for (int i = 0; i < n; i++)
                for (int j = 0; j < n; j++) d[i][j] = Math.min(d[i][j], d[i][k] + d[k][j]);
        int best = Integer.MAX_VALUE, city = -1;
        for (int i = 0; i < n; i++) {
            int cnt = 0;
            for (int j = 0; j < n; j++) if (j != i && d[i][j] <= distanceThreshold) cnt++;
            if (cnt <= best) { best = cnt; city = i; }
        }
        return city;
    }
}
// test
check(new Solution().findTheCity(4, new int[][]{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, 4), 3);
```

## 685. Redundant Connection II | Hard
related: 684
tags: union-find, directed-graph

A rooted tree plus one extra **directed** edge creates either (a) a node with two parents, (b) a cycle, or both. If some node has two parents, the answer is one of those two edges: skip the later one and check with Union-Find whether a cycle remains — if yes, the earlier edge is the culprit. Otherwise return the edge that closes the cycle.

**Complexity:** O(n · α(n)).

```python
class Solution:
    def findRedundantDirectedConnection(self, edges: List[List[int]]) -> List[int]:
        n = len(edges)
        parent_of, cand1, cand2 = {}, None, None
        for u, v in edges:
            if v in parent_of:
                cand1, cand2 = [parent_of[v], v], [u, v]
            else:
                parent_of[v] = u
        root = list(range(n + 1))

        def find(x):
            while root[x] != x:
                root[x] = root[root[x]]
                x = root[x]
            return x

        for u, v in edges:
            if [u, v] == cand2:
                continue
            ru, rv = find(u), find(v)
            if ru == rv:                       # cycle found
                return cand1 if cand1 else [u, v]
            root[rv] = ru
        return cand2
# test
assert Solution().findRedundantDirectedConnection([[1, 2], [1, 3], [2, 3]]) == [2, 3]
assert Solution().findRedundantDirectedConnection([[1, 2], [2, 3], [3, 4], [4, 1], [1, 5]]) == [4, 1]
assert Solution().findRedundantDirectedConnection([[2, 1], [3, 1], [4, 2], [1, 4]]) == [2, 1]
```

```java
class Solution {
    private int[] root;

    private int find(int x) {
        while (root[x] != x) { root[x] = root[root[x]]; x = root[x]; }
        return x;
    }

    public int[] findRedundantDirectedConnection(int[][] edges) {
        int n = edges.length;
        int[] parentOf = new int[n + 1];
        int[] cand1 = null, cand2 = null;
        for (int[] e : edges) {
            if (parentOf[e[1]] != 0) { cand1 = new int[]{parentOf[e[1]], e[1]}; cand2 = e; }
            else parentOf[e[1]] = e[0];
        }
        root = new int[n + 1];
        for (int i = 0; i <= n; i++) root[i] = i;
        for (int[] e : edges) {
            if (e == cand2) continue;
            int ru = find(e[0]), rv = find(e[1]);
            if (ru == rv) return cand1 != null ? cand1 : e;
            root[rv] = ru;
        }
        return cand2;
    }
}
// test
check(new Solution().findRedundantDirectedConnection(new int[][]{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}}), new int[]{4, 1});
check(new Solution().findRedundantDirectedConnection(new int[][]{{2, 1}, {3, 1}, {4, 2}, {1, 4}}), new int[]{2, 1});
```

## 753. Cracking the Safe | Hard
related: 332
tags: de-bruijn, eulerian-path

The shortest string containing every length-n password is a **De Bruijn sequence** of length `kⁿ + n − 1`. Nodes = length-(n−1) strings, each appended digit = an edge; an Eulerian circuit uses every edge (password) once. Hierholzer's DFS: append the edge digit after exploring it.

**Complexity:** O(kⁿ).

```python
class Solution:
    def crackSafe(self, n: int, k: int) -> str:
        seen, out = set(), []

        def dfs(node):
            for d in map(str, range(k)):
                edge = node + d
                if edge not in seen:
                    seen.add(edge)
                    dfs(edge[1:])
                    out.append(d)

        start = "0" * (n - 1)
        dfs(start)
        return "".join(out) + start
# test
def covers(s, n, k):
    return len({s[i:i + n] for i in range(len(s) - n + 1)}) == k ** n and len(s) == k ** n + n - 1
assert covers(Solution().crackSafe(1, 2), 1, 2)
assert covers(Solution().crackSafe(2, 2), 2, 2)
assert covers(Solution().crackSafe(3, 3), 3, 3)
```

```java
class Solution {
    private final Set<String> seen = new HashSet<>();
    private final StringBuilder out = new StringBuilder();

    public String crackSafe(int n, int k) {
        String start = "0".repeat(n - 1);
        dfs(start, k);
        return out.append(start).toString();
    }

    private void dfs(String node, int k) {
        for (int d = 0; d < k; d++) {
            String edge = node + d;
            if (seen.add(edge)) {
                dfs(edge.substring(1), k);
                out.append(d);
            }
        }
    }
}
// test
String s = new Solution().crackSafe(3, 3);
Set<String> all = new HashSet<>();
for (int i = 0; i + 3 <= s.length(); i++) all.add(s.substring(i, i + 3));
check(all.size() == 27 && s.length() == 29, true);
```

## 1489. Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree | Hard
related: 1584, 1192
tags: mst, union-find

Compute the MST weight once with Kruskal. Then for each edge: if the MST **without** it is heavier (or the graph disconnects), it's **critical**; otherwise, if an MST **forced to include** it still has the optimal weight, it's **pseudo-critical**. Each check is one Kruskal run.

**Complexity:** O(E² · α(V)) (after one sort).

```python
class Solution:
    def findCriticalAndPseudoCriticalEdges(self, n: int, edges: List[List[int]]) -> List[List[int]]:
        order = sorted(range(len(edges)), key=lambda i: edges[i][2])

        def mst(skip=-1, force=-1):
            parent = list(range(n))

            def find(x):
                while parent[x] != x:
                    parent[x] = parent[parent[x]]
                    x = parent[x]
                return x

            weight = used = 0
            if force >= 0:
                a, b, w = edges[force]
                parent[find(a)] = find(b)
                weight, used = w, 1
            for i in order:
                a, b, w = edges[i]
                if i != skip and find(a) != find(b):
                    parent[find(a)] = find(b)
                    weight, used = weight + w, used + 1
            return weight if used == n - 1 else float("inf")

        base = mst()
        critical, pseudo = [], []
        for i in range(len(edges)):
            if mst(skip=i) > base:
                critical.append(i)
            elif mst(force=i) == base:
                pseudo.append(i)
        return [critical, pseudo]
# test
r = Solution().findCriticalAndPseudoCriticalEdges(5, [[0, 1, 1], [1, 2, 1], [2, 3, 2], [0, 3, 2], [0, 4, 3], [3, 4, 3], [1, 4, 6]])
assert [sorted(r[0]), sorted(r[1])] == [[0, 1], [2, 3, 4, 5]]
r = Solution().findCriticalAndPseudoCriticalEdges(4, [[0, 1, 1], [1, 2, 1], [2, 3, 1], [0, 3, 1]])
assert [sorted(r[0]), sorted(r[1])] == [[], [0, 1, 2, 3]]
r = Solution().findCriticalAndPseudoCriticalEdges(3, [[0, 1, 1], [1, 2, 2], [0, 2, 3]])
assert [sorted(r[0]), sorted(r[1])] == [[0, 1], []]
```

```java
class Solution {
    private int n;
    private int[][] edges;
    private Integer[] order;

    public List<List<Integer>> findCriticalAndPseudoCriticalEdges(int n, int[][] edges) {
        this.n = n;
        this.edges = edges;
        order = new Integer[edges.length];
        for (int i = 0; i < edges.length; i++) order[i] = i;
        Arrays.sort(order, (x, y) -> edges[x][2] - edges[y][2]);
        int base = mst(-1, -1);
        List<Integer> critical = new ArrayList<>(), pseudo = new ArrayList<>();
        for (int i = 0; i < edges.length; i++) {
            if (mst(i, -1) > base) critical.add(i);
            else if (mst(-1, i) == base) pseudo.add(i);
        }
        return List.of(critical, pseudo);
    }

    private int mst(int skip, int force) {
        int[] parent = new int[n];
        for (int i = 0; i < n; i++) parent[i] = i;
        int weight = 0, used = 0;
        if (force >= 0) { union(parent, edges[force][0], edges[force][1]); weight = edges[force][2]; used = 1; }
        for (int i : order) if (i != skip && union(parent, edges[i][0], edges[i][1])) { weight += edges[i][2]; used++; }
        return used == n - 1 ? weight : Integer.MAX_VALUE;
    }

    private boolean union(int[] p, int a, int b) {
        int ra = find(p, a), rb = find(p, b);
        if (ra == rb) return false;
        p[ra] = rb;
        return true;
    }

    private int find(int[] p, int x) {
        while (p[x] != x) x = p[x] = p[p[x]];
        return x;
    }
}
// test
check(new Solution().findCriticalAndPseudoCriticalEdges(5, new int[][]{{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {0, 3, 2}, {0, 4, 3}, {3, 4, 3}, {1, 4, 6}}),
      List.of(List.of(0, 1), List.of(2, 3, 4, 5)));
```

## 2392. Build a Matrix With Conditions | Hard
related: 210, 269
tags: topological-sort

Rows and columns are independent. Topologically sort the numbers 1..k by `rowConditions` (gives each number's row) and separately by `colConditions` (its column); a cycle in either means no answer. Place number v at `(row[v], col[v])` — every row and column gets exactly one number. Any valid matrix is accepted.

**Complexity:** O(k² + conditions) (the k × k output dominates).

```python
class Solution:
    def buildMatrix(self, k: int, rowConditions: List[List[int]], colConditions: List[List[int]]) -> List[List[int]]:
        def topo(conds):
            adj, indeg = defaultdict(list), [0] * (k + 1)
            for a, b in conds:
                adj[a].append(b)
                indeg[b] += 1
            q = deque(v for v in range(1, k + 1) if indeg[v] == 0)
            order = []
            while q:
                v = q.popleft()
                order.append(v)
                for w in adj[v]:
                    indeg[w] -= 1
                    if indeg[w] == 0:
                        q.append(w)
            return {v: i for i, v in enumerate(order)} if len(order) == k else None

        row, col = topo(rowConditions), topo(colConditions)
        if row is None or col is None:
            return []
        m = [[0] * k for _ in range(k)]
        for v in range(1, k + 1):
            m[row[v]][col[v]] = v
        return m
# test
def valid(k, rc, cc, m):  # any matrix satisfying the conditions is accepted
    pos = {m[i][j]: (i, j) for i in range(k) for j in range(k) if m[i][j]}
    return (len(m) == k and all(len(r) == k for r in m) and sorted(pos) == list(range(1, k + 1))
            and all(pos[a][0] < pos[b][0] for a, b in rc) and all(pos[a][1] < pos[b][1] for a, b in cc))
assert valid(3, [[1, 2], [3, 2]], [[2, 1], [3, 2]], Solution().buildMatrix(3, [[1, 2], [3, 2]], [[2, 1], [3, 2]]))
assert Solution().buildMatrix(3, [[1, 2], [2, 3], [3, 1], [2, 3]], [[2, 1]]) == []
assert valid(4, [[4, 3], [3, 2], [2, 1]], [[1, 2], [2, 3]], Solution().buildMatrix(4, [[4, 3], [3, 2], [2, 1]], [[1, 2], [2, 3]]))
```

```java
class Solution {
    public int[][] buildMatrix(int k, int[][] rowConditions, int[][] colConditions) {
        int[] row = topo(k, rowConditions), col = topo(k, colConditions);
        if (row == null || col == null) return new int[0][0];
        int[][] m = new int[k][k];
        for (int v = 1; v <= k; v++) m[row[v]][col[v]] = v;
        return m;
    }

    private int[] topo(int k, int[][] conds) {                 // position of each value, or null on a cycle
        List<List<Integer>> adj = new ArrayList<>();
        for (int i = 0; i <= k; i++) adj.add(new ArrayList<>());
        int[] indeg = new int[k + 1], pos = new int[k + 1];
        for (int[] c : conds) { adj.get(c[0]).add(c[1]); indeg[c[1]]++; }
        Deque<Integer> q = new ArrayDeque<>();
        for (int v = 1; v <= k; v++) if (indeg[v] == 0) q.add(v);
        int idx = 0;
        while (!q.isEmpty()) {
            int v = q.poll();
            pos[v] = idx++;
            for (int w : adj.get(v)) if (--indeg[w] == 0) q.add(w);
        }
        return idx == k ? pos : null;
    }
}
// test
int[][] bm = new Solution().buildMatrix(3, new int[][]{{1, 2}, {3, 2}}, new int[][]{{2, 1}, {3, 2}});
int[][] at = new int[4][];
for (int i = 0; i < 3; i++) for (int j = 0; j < 3; j++) if (bm[i][j] > 0) at[bm[i][j]] = new int[]{i, j};
check(at[1][0] < at[2][0] && at[3][0] < at[2][0] && at[2][1] < at[1][1] && at[3][1] < at[2][1], true);
check(new Solution().buildMatrix(3, new int[][]{{1, 2}, {2, 3}, {3, 1}}, new int[][]{{2, 1}}).length, 0);
```

## 2709. Greatest Common Divisor Traversal | Hard
related: 952, 547
tags: union-find, sieve

Indices i and j are directly connected iff they share a prime factor, so connect each index to its primes: with a **smallest-prime-factor sieve**, factor every number and union the index with the first index seen for each of its primes. Traversal between all pairs ⇔ one connected component. A `1` (no prime factors) is isolated unless it's the only element.

**Complexity:** O(M log log M + n log M) for M = max(nums).

```python
class Solution:
    def canTraverseAllPairs(self, nums: List[int]) -> bool:
        n = len(nums)
        if n == 1:
            return True
        if 1 in nums:
            return False
        M = max(nums)
        spf = list(range(M + 1))                      # smallest prime factor
        for i in range(2, int(M ** 0.5) + 1):
            if spf[i] == i:
                for j in range(i * i, M + 1, i):
                    if spf[j] == j:
                        spf[j] = i
        parent = list(range(n))

        def find(x):
            while parent[x] != x:
                parent[x] = parent[parent[x]]
                x = parent[x]
            return x

        first = {}                                    # prime -> first index having it
        for i, x in enumerate(nums):
            while x > 1:
                p = spf[x]
                if p in first:
                    parent[find(i)] = find(first[p])
                else:
                    first[p] = i
                while x % p == 0:
                    x //= p
        return len({find(i) for i in range(n)}) == 1
# test
assert Solution().canTraverseAllPairs([2, 3, 6]) is True
assert Solution().canTraverseAllPairs([3, 9, 5]) is False
assert Solution().canTraverseAllPairs([4, 3, 12, 8]) is True
assert Solution().canTraverseAllPairs([1]) is True
assert Solution().canTraverseAllPairs([1, 1]) is False
assert Solution().canTraverseAllPairs([2, 2]) is True
assert Solution().canTraverseAllPairs([35, 77, 22, 6]) is True
```

```java
class Solution {
    public boolean canTraverseAllPairs(int[] nums) {
        int n = nums.length, M = 0;
        if (n == 1) return true;
        for (int x : nums) { if (x == 1) return false; M = Math.max(M, x); }
        int[] spf = new int[M + 1];
        for (int i = 2; i <= M; i++) if (spf[i] == 0) for (int j = i; j <= M; j += i) if (spf[j] == 0) spf[j] = i;
        int[] parent = new int[n];
        for (int i = 0; i < n; i++) parent[i] = i;
        Map<Integer, Integer> first = new HashMap<>();
        for (int i = 0; i < n; i++) {
            for (int x = nums[i]; x > 1; ) {
                int p = spf[x];
                Integer j = first.putIfAbsent(p, i);
                if (j != null) parent[find(parent, i)] = find(parent, j);
                while (x % p == 0) x /= p;
            }
        }
        int root = find(parent, 0);
        for (int i = 1; i < n; i++) if (find(parent, i) != root) return false;
        return true;
    }

    private int find(int[] p, int x) {
        while (p[x] != x) x = p[x] = p[p[x]];
        return x;
    }
}
// test
check(new Solution().canTraverseAllPairs(new int[]{2, 3, 6}), true);
check(new Solution().canTraverseAllPairs(new int[]{3, 9, 5}), false);
```
