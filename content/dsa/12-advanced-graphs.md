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
tags: topological-sort, premium

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
assert Solution().findMinHeightTrees(4, [[1, 0], [1, 2], [1, 3]]) == [1]
assert Solution().findMinHeightTrees(6, [[3, 0], [3, 1], [3, 2], [3, 4], [5, 4]]) == [3, 4]
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
