# Intervals

**Almost always: sort first.**
- Sort by **start** → merging / inserting overlapping intervals.
- Sort by **end** → maximum set of non-overlapping intervals (activity selection), arrows.
- Overlap test for `[a, b]` and `[c, d]`: `a <= d and c <= b` (inclusive ends).
- "How many at the same time?" → **sweep line** over sorted start/end events, or a **min-heap of end times**.

```viz
sorted by start: [1,3] [2,6] [8,10] [15,18]
  1--3
    2-------6            2 <= 3  -> merge into [1,6]
              8---10     8 >  6  -> new block
result: [1,6] [8,10] [15,18]
```

## 56. Merge Intervals | Medium
related: 57, 435, 986
tags: blind75, sort-by-start

```problem
Merge all overlapping intervals (intervals that merely touch, like [1,4] and [4,5], overlap too) and return the non-overlapping intervals that cover the same ranges.

**Example:** `[[1,3],[2,6],[8,10],[15,18]]` → `[[1,6],[8,10],[15,18]]`

**Constraints:** up to 10⁴ intervals.
```

Sort by start; if an interval starts before the last merged one ends, extend; otherwise start a new block.

**Complexity:** O(n log n).

```python
class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        res = []
        for s, e in sorted(intervals):
            if res and s <= res[-1][1]:
                res[-1][1] = max(res[-1][1], e)
            else:
                res.append([s, e])
        return res
# test
assert Solution().merge([[1, 3], [2, 6], [8, 10], [15, 18]]) == [[1, 6], [8, 10], [15, 18]]
assert Solution().merge([[1, 4], [4, 5]]) == [[1, 5]]
```

```java
class Solution {
    public int[][] merge(int[][] intervals) {
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));
        List<int[]> res = new ArrayList<>();
        for (int[] iv : intervals) {
            if (!res.isEmpty() && iv[0] <= res.get(res.size() - 1)[1])
                res.get(res.size() - 1)[1] = Math.max(res.get(res.size() - 1)[1], iv[1]);
            else res.add(new int[]{iv[0], iv[1]});
        }
        return res.toArray(new int[0][]);
    }
}
// test
check(new Solution().merge(new int[][]{{1, 3}, {2, 6}, {8, 10}, {15, 18}}), new int[][]{{1, 6}, {8, 10}, {15, 18}});
```

```go
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	res := [][]int{}
	for _, iv := range intervals {
		if n := len(res); n > 0 && iv[0] <= res[n-1][1] { // overlaps the last merged interval
			res[n-1][1] = max(res[n-1][1], iv[1])
		} else {
			res = append(res, []int{iv[0], iv[1]})
		}
	}
	return res
}
// test
assertEq(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}), [][]int{{1, 6}, {8, 10}, {15, 18}})
assertEq(merge([][]int{{1, 4}, {4, 5}}), [][]int{{1, 5}})
```

## 57. Insert Interval | Medium
related: 56
tags: blind75, sweep-line

```problem
`intervals` are non-overlapping and sorted by start. Insert `newInterval`, merging wherever necessary, and return the list, still sorted and non-overlapping.

**Example:** `intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]` → `[[1,2],[3,10],[12,16]]`
```

Three phases on the already-sorted list: copy intervals ending before the new one; merge every overlapping interval into it; copy the rest.

**Complexity:** O(n).

```python
class Solution:
    def insert(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        res, i, n = [], 0, len(intervals)
        s, e = newInterval
        while i < n and intervals[i][1] < s:
            res.append(intervals[i]); i += 1
        while i < n and intervals[i][0] <= e:
            s, e = min(s, intervals[i][0]), max(e, intervals[i][1]); i += 1
        res.append([s, e])
        return res + intervals[i:]
# test
assert Solution().insert([[1, 3], [6, 9]], [2, 5]) == [[1, 5], [6, 9]]
assert Solution().insert([[1, 2], [3, 5], [6, 7], [8, 10], [12, 16]], [4, 8]) == [[1, 2], [3, 10], [12, 16]]
assert Solution().insert([], [5, 7]) == [[5, 7]]
```

```java
class Solution {
    public int[][] insert(int[][] intervals, int[] newInterval) {
        List<int[]> res = new ArrayList<>();
        int i = 0, n = intervals.length, s = newInterval[0], e = newInterval[1];
        while (i < n && intervals[i][1] < s) res.add(intervals[i++]);
        while (i < n && intervals[i][0] <= e) {
            s = Math.min(s, intervals[i][0]);
            e = Math.max(e, intervals[i][1]);
            i++;
        }
        res.add(new int[]{s, e});
        while (i < n) res.add(intervals[i++]);
        return res.toArray(new int[0][]);
    }
}
// test
check(new Solution().insert(new int[][]{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, new int[]{4, 8}), new int[][]{{1, 2}, {3, 10}, {12, 16}});
```

```go
func insert(intervals [][]int, newInterval []int) [][]int {
	res, i, n := [][]int{}, 0, len(intervals)
	s, e := newInterval[0], newInterval[1]
	for ; i < n && intervals[i][1] < s; i++ { // entirely before
		res = append(res, intervals[i])
	}
	for ; i < n && intervals[i][0] <= e; i++ { // overlapping: absorb
		s, e = min(s, intervals[i][0]), max(e, intervals[i][1])
	}
	res = append(res, []int{s, e})
	return append(res, intervals[i:]...) // entirely after
}
// test
assertEq(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}), [][]int{{1, 5}, {6, 9}})
assertEq(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}), [][]int{{1, 2}, {3, 10}, {12, 16}})
assertEq(insert([][]int{}, []int{5, 7}), [][]int{{5, 7}})
```

## 435. Non-overlapping Intervals | Medium
related: 452, 56
tags: blind75, sort-by-end

```problem
Return the minimum number of intervals to remove so that the rest don't overlap. Intervals that only touch (like [1,2] and [2,3]) don't overlap.

**Example:** `[[1,2],[2,3],[3,4],[1,3]]` → `1`; `[[1,2],[1,2],[1,2]]` → `2`

**Constraints:** up to 10⁵ intervals.
```

Min removals = n − (max non-overlapping set). Greedy activity selection: sort by **end**, keep an interval if it starts at/after the last kept end.

**Complexity:** O(n log n).

```python
class Solution:
    def eraseOverlapIntervals(self, intervals: List[List[int]]) -> int:
        keep, end = 0, float("-inf")
        for s, e in sorted(intervals, key=lambda x: x[1]):
            if s >= end:
                keep += 1
                end = e
        return len(intervals) - keep
# test
assert Solution().eraseOverlapIntervals([[1, 2], [2, 3], [3, 4], [1, 3]]) == 1
assert Solution().eraseOverlapIntervals([[1, 2], [1, 2], [1, 2]]) == 2
assert Solution().eraseOverlapIntervals([[1, 2], [2, 3]]) == 0
```

```java
class Solution {
    public int eraseOverlapIntervals(int[][] intervals) {
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[1], b[1]));
        int keep = 0;
        long end = Long.MIN_VALUE;
        for (int[] iv : intervals) if (iv[0] >= end) { keep++; end = iv[1]; }
        return intervals.length - keep;
    }
}
// test
check(new Solution().eraseOverlapIntervals(new int[][]{{1, 2}, {2, 3}, {3, 4}, {1, 3}}), 1);
```

```go
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	keep, end := 0, math.MinInt
	for _, iv := range intervals { // keep the interval that ends first, as often as possible
		if iv[0] >= end {
			keep++
			end = iv[1]
		}
	}
	return len(intervals) - keep
}
// test
assertEq(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}), 1)
assertEq(eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}), 2)
assertEq(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}}), 0)
```

## 452. Minimum Number of Arrows to Burst Balloons | Medium
related: 435
tags: sort-by-end

```problem
Balloons are horizontal ranges `[xstart, xend]`. An arrow shot straight up at x bursts every balloon with xstart ≤ x ≤ xend. Return the minimum number of arrows needed to burst them all.

**Example:** `[[10,16],[2,8],[1,6],[7,12]]` → `2` (shoot at x = 6 and x = 11)

**Constraints:** up to 10⁵ balloons; coordinates span the full 32-bit range.
```

Sort by end; shoot at the first end. Every balloon starting ≤ that point is burst; the next balloon starting after it needs a new arrow at its end.

**Complexity:** O(n log n).

```python
class Solution:
    def findMinArrowShots(self, points: List[List[int]]) -> int:
        arrows, pos = 0, float("-inf")
        for s, e in sorted(points, key=lambda p: p[1]):
            if s > pos:
                arrows += 1
                pos = e
        return arrows
# test
assert Solution().findMinArrowShots([[10, 16], [2, 8], [1, 6], [7, 12]]) == 2
assert Solution().findMinArrowShots([[1, 2], [3, 4], [5, 6], [7, 8]]) == 4
```

```java
class Solution {
    public int findMinArrowShots(int[][] points) {
        Arrays.sort(points, (a, b) -> Integer.compare(a[1], b[1]));  // compare() avoids overflow
        int arrows = 0;
        long pos = Long.MIN_VALUE;
        for (int[] p : points) if (p[0] > pos) { arrows++; pos = p[1]; }
        return arrows;
    }
}
// test
check(new Solution().findMinArrowShots(new int[][]{{10, 16}, {2, 8}, {1, 6}, {7, 12}}), 2);
```

```go
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	arrows, pos := 0, math.MinInt
	for _, p := range points { // shoot at the end of the first balloon not yet burst
		if p[0] > pos {
			arrows++
			pos = p[1]
		}
	}
	return arrows
}
// test
assertEq(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}), 2)
assertEq(findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}), 4)
```

## 252. Meeting Rooms | Easy
related: 253, 56
tags: blind75, premium

```problem
Given meeting time `intervals` `[start, end]`, decide whether one person could attend all of them. A meeting may start exactly when another ends.

**Example:** `[[0,30],[5,10],[15,20]]` → `false`; `[[7,10],[2,4]]` → `true`
```

Sort by start; any meeting starting before the previous one ends is a conflict.

**Complexity:** O(n log n).

```python
class Solution:
    def canAttendMeetings(self, intervals: List[List[int]]) -> bool:
        intervals.sort()
        return all(intervals[i][0] >= intervals[i - 1][1] for i in range(1, len(intervals)))
# test
assert Solution().canAttendMeetings([[0, 30], [5, 10], [15, 20]]) is False
assert Solution().canAttendMeetings([[7, 10], [2, 4]]) is True
```

```java
class Solution {
    public boolean canAttendMeetings(int[][] intervals) {
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));
        for (int i = 1; i < intervals.length; i++) if (intervals[i][0] < intervals[i - 1][1]) return false;
        return true;
    }
}
// test
check(new Solution().canAttendMeetings(new int[][]{{0, 30}, {5, 10}, {15, 20}}), false);
check(new Solution().canAttendMeetings(new int[][]{{7, 10}, {2, 4}}), true);
```

```go
func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] { // starts before the previous one ends
			return false
		}
	}
	return true
}
// test
assertEq(canAttendMeetings([][]int{{0, 30}, {5, 10}, {15, 20}}), false)
assertEq(canAttendMeetings([][]int{{7, 10}, {2, 4}}), true)
```

## 253. Meeting Rooms II | Medium
related: 252, 1851
tags: blind75, sweep-line, premium

```problem
Given meeting time `intervals` `[start, end]`, return the minimum number of conference rooms needed to hold them all. A room freed at time t can host a meeting starting at t.

**Example:** `[[0,30],[5,10],[15,20]]` → `2`; `[[7,10],[2,4]]` → `1`

**Constraints:** up to 10⁴ meetings.
```

Max overlap at any time. **Sweep**: sort starts and ends separately; walk starts, freeing a room whenever the earliest end ≤ current start. (Equivalent: min-heap of end times.)

```viz
time   0    5   10   15   20   30
       [0----------------------30)
            [5---10)  [15--20)
rooms  1    2    1    2    1        max 2
```

**Complexity:** O(n log n).

```python
class Solution:
    def minMeetingRooms(self, intervals: List[List[int]]) -> int:
        starts = sorted(s for s, _ in intervals)
        ends = sorted(e for _, e in intervals)
        rooms = best = j = 0
        for s in starts:
            while ends[j] <= s:
                rooms -= 1
                j += 1
            rooms += 1
            best = max(best, rooms)
        return best
# test
assert Solution().minMeetingRooms([[0, 30], [5, 10], [15, 20]]) == 2
assert Solution().minMeetingRooms([[7, 10], [2, 4]]) == 1
```

```java
class Solution {
    public int minMeetingRooms(int[][] intervals) {
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));
        PriorityQueue<Integer> ends = new PriorityQueue<>();   // end times of rooms in use
        for (int[] iv : intervals) {
            if (!ends.isEmpty() && ends.peek() <= iv[0]) ends.poll();   // reuse the freed room
            ends.add(iv[1]);
        }
        return ends.size();
    }
}
// test
check(new Solution().minMeetingRooms(new int[][]{{0, 30}, {5, 10}, {15, 20}}), 2);
```

```go
func minMeetingRooms(intervals [][]int) int {
	n := len(intervals)
	starts, ends := make([]int, n), make([]int, n)
	for i, iv := range intervals {
		starts[i], ends[i] = iv[0], iv[1]
	}
	sort.Ints(starts)
	sort.Ints(ends)
	rooms, best, j := 0, 0, 0
	for _, s := range starts { // sweep: free the rooms whose meetings ended by s
		for ends[j] <= s {
			rooms--
			j++
		}
		rooms++
		best = max(best, rooms)
	}
	return best
}
// test
assertEq(minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}), 2)
assertEq(minMeetingRooms([][]int{{7, 10}, {2, 4}}), 1)
```

## 986. Interval List Intersections | Medium
related: 56
tags: two-pointers

```problem
Two lists of closed intervals are each sorted and pairwise disjoint. Return their intersection: every range covered by an interval from both lists.

**Example:** `[[0,2],[5,10],[13,23],[24,25]]` and `[[1,5],[8,12],[15,24],[25,26]]` → `[[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]`
```

Two pointers over the two sorted lists. Intersection = `[max(starts), min(ends)]` if non-empty. Advance the list whose interval ends first.

**Complexity:** O(m + n).

```python
class Solution:
    def intervalIntersection(self, firstList: List[List[int]], secondList: List[List[int]]) -> List[List[int]]:
        i = j = 0
        res = []
        while i < len(firstList) and j < len(secondList):
            lo = max(firstList[i][0], secondList[j][0])
            hi = min(firstList[i][1], secondList[j][1])
            if lo <= hi:
                res.append([lo, hi])
            if firstList[i][1] < secondList[j][1]:
                i += 1
            else:
                j += 1
        return res
# test
A = [[0, 2], [5, 10], [13, 23], [24, 25]]
B = [[1, 5], [8, 12], [15, 24], [25, 26]]
assert Solution().intervalIntersection(A, B) == [[1, 2], [5, 5], [8, 10], [15, 23], [24, 24], [25, 25]]
assert Solution().intervalIntersection([[1, 3], [5, 9]], []) == []
assert Solution().intervalIntersection([[1, 7]], [[3, 10]]) == [[3, 7]]
```

```java
class Solution {
    public int[][] intervalIntersection(int[][] a, int[][] b) {
        List<int[]> res = new ArrayList<>();
        int i = 0, j = 0;
        while (i < a.length && j < b.length) {
            int lo = Math.max(a[i][0], b[j][0]), hi = Math.min(a[i][1], b[j][1]);
            if (lo <= hi) res.add(new int[]{lo, hi});
            if (a[i][1] < b[j][1]) i++; else j++;
        }
        return res.toArray(new int[0][]);
    }
}
// test
int[][] A = {{0, 2}, {5, 10}, {13, 23}, {24, 25}}, B = {{1, 5}, {8, 12}, {15, 24}, {25, 26}};
check(new Solution().intervalIntersection(A, B), new int[][]{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}});
```

```go
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := [][]int{}
	for i, j := 0, 0; i < len(firstList) && j < len(secondList); {
		lo := max(firstList[i][0], secondList[j][0])
		hi := min(firstList[i][1], secondList[j][1])
		if lo <= hi {
			res = append(res, []int{lo, hi})
		}
		if firstList[i][1] < secondList[j][1] { // the one ending first cannot meet anything else
			i++
		} else {
			j++
		}
	}
	return res
}
// test
A := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
B := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
assertEq(intervalIntersection(A, B), [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}})
assertEq(intervalIntersection([][]int{{1, 3}, {5, 9}}, [][]int{}), [][]int{})
assertEq(intervalIntersection([][]int{{1, 7}}, [][]int{{3, 10}}), [][]int{{3, 7}})
```

## 1851. Minimum Interval to Include Each Query | Hard
related: 253, 56
tags: offline-queries, heap

```problem
For each query point q, return the size (right − left + 1) of the smallest interval in `intervals` that contains q (left ≤ q ≤ right), or −1 if no interval does.

**Example:** `intervals = [[1,4],[2,4],[3,6],[4,4]], queries = [2,3,4,5]` → `[3,3,1,4]`

**Constraints:** up to 10⁵ intervals and 10⁵ queries.
```

Process queries **offline** in sorted order. Push every interval with `start ≤ q` into a min-heap keyed by size; pop intervals whose `end < q` (they're dead for all later queries too). Heap top is the answer.

**Complexity:** O((n + q) log n).

```python
class Solution:
    def minInterval(self, intervals: List[List[int]], queries: List[int]) -> List[int]:
        intervals.sort()
        h, i, ans = [], 0, {}
        for q in sorted(set(queries)):
            while i < len(intervals) and intervals[i][0] <= q:
                l, r = intervals[i]
                heappush(h, (r - l + 1, r))
                i += 1
            while h and h[0][1] < q:
                heappop(h)
            ans[q] = h[0][0] if h else -1
        return [ans[q] for q in queries]
# test
assert Solution().minInterval([[1, 4], [2, 4], [3, 6], [4, 4]], [2, 3, 4, 5]) == [3, 3, 1, 4]
assert Solution().minInterval([[2, 3], [2, 5], [1, 8], [20, 25]], [2, 19, 5, 22]) == [2, -1, 4, 6]
```

```java
class Solution {
    public int[] minInterval(int[][] intervals, int[] queries) {
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));
        Integer[] qi = new Integer[queries.length];
        for (int k = 0; k < qi.length; k++) qi[k] = k;
        Arrays.sort(qi, (a, b) -> Integer.compare(queries[a], queries[b]));
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[0] - b[0]); // {size, end}
        int[] res = new int[queries.length];
        int i = 0;
        for (int k : qi) {
            int q = queries[k];
            while (i < intervals.length && intervals[i][0] <= q) {
                pq.add(new int[]{intervals[i][1] - intervals[i][0] + 1, intervals[i][1]});
                i++;
            }
            while (!pq.isEmpty() && pq.peek()[1] < q) pq.poll();
            res[k] = pq.isEmpty() ? -1 : pq.peek()[0];
        }
        return res;
    }
}
// test
check(new Solution().minInterval(new int[][]{{1, 4}, {2, 4}, {3, 6}, {4, 4}}, new int[]{2, 3, 4, 5}), new int[]{3, 3, 1, 4});
```

```go
type sizeHeap [][2]int // (size, right end), smallest size first

func (h sizeHeap) Len() int           { return len(h) }
func (h sizeHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h sizeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *sizeHeap) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *sizeHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func minInterval(intervals [][]int, queries []int) []int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	order := make([]int, len(queries)) // answer queries in increasing order
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(a, b int) bool { return queries[order[a]] < queries[order[b]] })
	res := make([]int, len(queries))
	h, i := &sizeHeap{}, 0
	for _, qi := range order {
		q := queries[qi]
		for ; i < len(intervals) && intervals[i][0] <= q; i++ { // intervals starting by q
			heap.Push(h, [2]int{intervals[i][1] - intervals[i][0] + 1, intervals[i][1]})
		}
		for h.Len() > 0 && (*h)[0][1] < q { // drop those that ended before q
			heap.Pop(h)
		}
		res[qi] = -1
		if h.Len() > 0 {
			res[qi] = (*h)[0][0]
		}
	}
	return res
}
// test
assertEq(minInterval([][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}}, []int{2, 3, 4, 5}), []int{3, 3, 1, 4})
assertEq(minInterval([][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}}, []int{2, 19, 5, 22}), []int{2, -1, 4, 6})
```


## 1288. Remove Covered Intervals | Medium
related: 56, 435
tags: sort-by-start

```problem
Interval [a, b) is covered by [c, d) if c ≤ a and b ≤ d. Remove every interval covered by another one and return how many remain.

**Example:** `[[1,4],[3,6],[2,8]]` → `2` ([3,6] is covered by [2,8])

**Constraints:** up to 1000 intervals, all distinct.
```

Sort by start ascending and end **descending** (so a longer interval with the same start comes first). Sweep tracking the max end; an interval whose end ≤ max end is covered.

**Complexity:** O(n log n).

```python
class Solution:
    def removeCoveredIntervals(self, intervals: List[List[int]]) -> int:
        keep, max_end = 0, -1
        for s, e in sorted(intervals, key=lambda x: (x[0], -x[1])):
            if e > max_end:
                keep += 1
                max_end = e
        return keep
# test
assert Solution().removeCoveredIntervals([[1, 4], [3, 6], [2, 8]]) == 2
assert Solution().removeCoveredIntervals([[1, 4], [2, 3]]) == 1
```

```java
class Solution {
    public int removeCoveredIntervals(int[][] intervals) {
        Arrays.sort(intervals, (a, b) -> a[0] != b[0] ? a[0] - b[0] : b[1] - a[1]);
        int keep = 0, maxEnd = -1;
        for (int[] iv : intervals) if (iv[1] > maxEnd) { keep++; maxEnd = iv[1]; }
        return keep;
    }
}
// test
check(new Solution().removeCoveredIntervals(new int[][]{{1, 4}, {3, 6}, {2, 8}}), 2);
```

```go
func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { // start up; for equal starts the longest first
		a, b := intervals[i], intervals[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	keep, maxEnd := 0, -1
	for _, iv := range intervals { // covered exactly when an earlier interval reaches as far
		if iv[1] > maxEnd {
			keep++
			maxEnd = iv[1]
		}
	}
	return keep
}
// test
assertEq(removeCoveredIntervals([][]int{{1, 4}, {3, 6}, {2, 8}}), 2)
assertEq(removeCoveredIntervals([][]int{{1, 4}, {2, 3}}), 1)
```

## 729. My Calendar I | Medium
related: 731, 56
tags: sorted-map, design

```problem
Design `MyCalendar` whose `book(start, end)` adds the event [start, end) and returns `true` if it doesn't overlap any booked event; otherwise it books nothing and returns `false`.

**Example:** `book(10,20)` → true, `book(15,25)` → false, `book(20,30)` → true

**Constraints:** up to 1000 calls.
```

Keep booked intervals sorted by start. A new `[s, e)` is fine iff it ends before the next booking starts and starts after the previous booking ends — one binary search (Java `TreeMap.floorKey/ceilingKey`).

**Complexity:** O(log n) per booking (O(n) insert for a Python list).

```python
class MyCalendar:
    def __init__(self):
        self.starts, self.ends = [], []

    def book(self, start: int, end: int) -> bool:
        i = bisect.bisect_right(self.starts, start)
        if i > 0 and self.ends[i - 1] > start:          # previous overlaps
            return False
        if i < len(self.starts) and self.starts[i] < end:   # next overlaps
            return False
        self.starts.insert(i, start)
        self.ends.insert(i, end)
        return True
# test
c = MyCalendar()
assert [c.book(10, 20), c.book(15, 25), c.book(20, 30)] == [True, False, True]
c2 = MyCalendar()
assert [c2.book(5, 10), c2.book(10, 15), c2.book(9, 11), c2.book(0, 5), c2.book(4, 6)] == [True, True, False, True, False]
```

```java
class MyCalendar {
    private final TreeMap<Integer, Integer> cal = new TreeMap<>();   // start -> end

    public boolean book(int start, int end) {
        Integer prev = cal.floorKey(start), next = cal.ceilingKey(start);
        if (prev != null && cal.get(prev) > start) return false;
        if (next != null && next < end) return false;
        cal.put(start, end);
        return true;
    }
}
// test
MyCalendar c = new MyCalendar();
check(c.book(10, 20), true);
check(c.book(15, 25), false);
check(c.book(20, 30), true);
```

```go
type MyCalendar struct {
	starts, ends []int // bookings sorted by start, never overlapping
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	i := sort.SearchInts(this.starts, start+1) // first booking starting after `start`
	if i > 0 && this.ends[i-1] > start {       // the previous one runs into us
		return false
	}
	if i < len(this.starts) && this.starts[i] < end { // we run into the next one
		return false
	}
	this.starts = insertAt(this.starts, i, start)
	this.ends = insertAt(this.ends, i, end)
	return true
}

func insertAt(s []int, i, v int) []int {
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = v
	return s
}
// test
c := Constructor()
assertEq([]bool{c.Book(10, 20), c.Book(15, 25), c.Book(20, 30)}, []bool{true, false, true})
c2 := Constructor()
assertEq([]bool{c2.Book(5, 10), c2.Book(10, 15), c2.Book(9, 11), c2.Book(0, 5), c2.Book(4, 6)}, []bool{true, true, false, true, false})
```

## 352. Data Stream as Disjoint Intervals | Hard
related: 57, 228
tags: sorted-map, merging

```problem
Design `SummaryRanges` for a stream of non-negative integers: `addNum(value)` adds a number and `getIntervals()` returns the numbers seen so far as a sorted list of disjoint intervals `[start, end]`.

**Example:** add 1, 3, 7, 2, 6 → `getIntervals()` = `[[1,3],[6,7]]`

**Constraints:** follow-up: many merges and few distinct intervals compared with the stream length.
```

Store disjoint intervals keyed by start in a sorted structure. On `addNum(v)`: find the interval ending at or after `v−1` from the left and the one starting at `v+1`; merge with either, both, or create `[v, v]`.

**Complexity:** O(log n) per add with a balanced tree; `getIntervals` O(n).

```python
class SummaryRanges:
    def __init__(self):
        self.iv = []                                   # sorted, disjoint [start, end]

    def addNum(self, value: int) -> None:
        iv = self.iv
        i = bisect.bisect_left(iv, [value + 1])        # first interval starting > value
        if i > 0 and iv[i - 1][1] >= value:            # already covered
            return
        left = i > 0 and iv[i - 1][1] == value - 1
        right = i < len(iv) and iv[i][0] == value + 1
        if left and right:
            iv[i - 1][1] = iv[i][1]
            iv.pop(i)
        elif left:
            iv[i - 1][1] = value
        elif right:
            iv[i][0] = value
        else:
            iv.insert(i, [value, value])

    def getIntervals(self) -> List[List[int]]:
        return [x[:] for x in self.iv]
# test
sr = SummaryRanges()
out = []
for v in (1, 3, 7, 2, 6):
    sr.addNum(v)
    out.append(sr.getIntervals())
assert out == [[[1, 1]], [[1, 1], [3, 3]], [[1, 1], [3, 3], [7, 7]], [[1, 3], [7, 7]], [[1, 3], [6, 7]]]
```

```java
class SummaryRanges {
    private final TreeMap<Integer, Integer> iv = new TreeMap<>();   // start -> end

    public void addNum(int value) {
        Integer lo = iv.floorKey(value), hi = iv.ceilingKey(value + 1);
        if (lo != null && iv.get(lo) >= value) return;
        boolean left = lo != null && iv.get(lo) == value - 1, right = hi != null && hi == value + 1;
        if (left && right) { iv.put(lo, iv.get(hi)); iv.remove(hi); }
        else if (left) iv.put(lo, value);
        else if (right) { iv.put(value, iv.get(hi)); iv.remove(hi); }
        else iv.put(value, value);
    }

    public int[][] getIntervals() {
        int[][] res = new int[iv.size()][];
        int i = 0;
        for (Map.Entry<Integer, Integer> e : iv.entrySet()) res[i++] = new int[]{e.getKey(), e.getValue()};
        return res;
    }
}
// test
SummaryRanges sr = new SummaryRanges();
for (int v : new int[]{1, 3, 7, 2, 6}) sr.addNum(v);
check(sr.getIntervals(), new int[][]{{1, 3}, {6, 7}});
```

```go
type SummaryRanges struct {
	iv [][]int // sorted, disjoint [start, end]
}

func Constructor() SummaryRanges {
	return SummaryRanges{}
}

func (this *SummaryRanges) AddNum(value int) {
	iv := this.iv
	i := sort.Search(len(iv), func(k int) bool { return iv[k][0] > value }) // first interval starting after value
	if i > 0 && iv[i-1][1] >= value {
		return // already covered
	}
	left := i > 0 && iv[i-1][1] == value-1
	right := i < len(iv) && iv[i][0] == value+1
	switch {
	case left && right: // value bridges two intervals
		iv[i-1][1] = iv[i][1]
		this.iv = append(iv[:i], iv[i+1:]...)
	case left:
		iv[i-1][1] = value
	case right:
		iv[i][0] = value
	default:
		iv = append(iv, []int{value, value})
		copy(iv[i+1:], iv[i:])
		iv[i] = []int{value, value}
		this.iv = iv
	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	out := make([][]int, len(this.iv))
	for i, x := range this.iv {
		out[i] = []int{x[0], x[1]}
	}
	return out
}
// test
sr := Constructor()
out := [][][]int{}
for _, v := range []int{1, 3, 7, 2, 6} {
	sr.AddNum(v)
	out = append(out, sr.GetIntervals())
}
assertEq(out, [][][]int{{{1, 1}}, {{1, 1}, {3, 3}}, {{1, 1}, {3, 3}, {7, 7}}, {{1, 3}, {7, 7}}, {{1, 3}, {6, 7}}})
```

## 759. Employee Free Time | Hard
related: 56, 986, 253
tags: sweep-line, sort-by-start, premium

```problem
`schedule[i]` lists employee i's working intervals, sorted and non-overlapping. Return the finite intervals of positive length during which **all** employees are free, sorted.

**Example:** `[[[1,2],[5,6]],[[1,3]],[[4,10]]]` → `[[3,4]]`
```

**Problem (Premium):** each employee has a sorted list of non-overlapping working `Interval`s. Return the finite intervals (positive length) when **everyone** is free, sorted.

Nobody's identity matters: flatten all intervals, sort by start and sweep while tracking the furthest end so far. Any start beyond that end leaves a common gap `(end, start)`. (A min-heap k-way merge avoids the full sort: O(N log k).)

**Complexity:** O(N log N) for N intervals.

```python
class Interval:  # provided by LeetCode
    def __init__(self, start: int = None, end: int = None):
        self.start = start
        self.end = end


class Solution:
    def employeeFreeTime(self, schedule: '[[Interval]]') -> '[Interval]':
        ivs = sorted((iv for emp in schedule for iv in emp), key=lambda iv: iv.start)
        res, end = [], ivs[0].end
        for iv in ivs[1:]:
            if iv.start > end:
                res.append(Interval(end, iv.start))
            end = max(end, iv.end)
        return res
# test
def free(s):
    return [[iv.start, iv.end] for iv in Solution().employeeFreeTime([[Interval(a, b) for a, b in emp] for emp in s])]
assert free([[[1, 2], [5, 6]], [[1, 3]], [[4, 10]]]) == [[3, 4]]
assert free([[[1, 3], [6, 7]], [[2, 4]], [[2, 5], [9, 12]]]) == [[5, 6], [7, 9]]
assert free([[[1, 2]], [[2, 3]]]) == []
assert free([[[1, 2], [8, 9]], [[3, 4]]]) == [[2, 3], [4, 8]]
```

```java
class Interval { // provided by LeetCode
    public int start, end;
    public Interval() {}
    public Interval(int _start, int _end) { start = _start; end = _end; }
}

class Solution {
    public List<Interval> employeeFreeTime(List<List<Interval>> schedule) {
        List<Interval> all = new ArrayList<>();
        for (List<Interval> emp : schedule) all.addAll(emp);
        all.sort((a, b) -> Integer.compare(a.start, b.start));
        List<Interval> res = new ArrayList<>();
        int end = all.get(0).end;
        for (Interval iv : all) {
            if (iv.start > end) res.add(new Interval(end, iv.start));
            end = Math.max(end, iv.end);
        }
        return res;
    }
}
// test
List<Interval> fr = new Solution().employeeFreeTime(List.of(
    List.of(new Interval(1, 3), new Interval(6, 7)), List.of(new Interval(2, 4)), List.of(new Interval(2, 5), new Interval(9, 12))));
check(fr.stream().map(iv -> List.of(iv.start, iv.end)).collect(Collectors.toList()), List.of(List.of(5, 6), List.of(7, 9)));
```

```go
type Interval struct { // provided by LeetCode
	Start int
	End   int
}

func employeeFreeTime(schedule [][]*Interval) []*Interval {
	ivs := []*Interval{}
	for _, emp := range schedule { // whose interval it is does not matter
		ivs = append(ivs, emp...)
	}
	sort.Slice(ivs, func(i, j int) bool { return ivs[i].Start < ivs[j].Start })
	res := []*Interval{}
	end := ivs[0].End // furthest busy time so far
	for _, iv := range ivs[1:] {
		if iv.Start > end {
			res = append(res, &Interval{Start: end, End: iv.Start})
		}
		end = max(end, iv.End)
	}
	return res
}
// test
free := func(s [][][]int) [][]int {
	sched := [][]*Interval{}
	for _, emp := range s {
		row := []*Interval{}
		for _, iv := range emp {
			row = append(row, &Interval{Start: iv[0], End: iv[1]})
		}
		sched = append(sched, row)
	}
	out := [][]int{}
	for _, iv := range employeeFreeTime(sched) {
		out = append(out, []int{iv.Start, iv.End})
	}
	return out
}
assertEq(free([][][]int{{{1, 2}, {5, 6}}, {{1, 3}}, {{4, 10}}}), [][]int{{3, 4}})
assertEq(free([][][]int{{{1, 3}, {6, 7}}, {{2, 4}}, {{2, 5}, {9, 12}}}), [][]int{{5, 6}, {7, 9}})
assertEq(free([][][]int{{{1, 2}}, {{2, 3}}}), [][]int{})
assertEq(free([][][]int{{{1, 2}, {8, 9}}, {{3, 4}}}), [][]int{{2, 3}, {4, 8}})
```

## 2402. Meeting Rooms III | Hard
related: 253, 1834
tags: two-heaps, simulation

```problem
Rooms 0…n − 1 host `meetings[i] = [start, end)` (all starts distinct). Each meeting takes the free room with the lowest number; if none is free, it waits for the earliest-freed room and keeps its duration, and waiting meetings are served in order of original start. Return the room that hosted the most meetings (ties: lowest number).

**Example:** `n = 2, meetings = [[0,10],[1,5],[2,7],[3,4]]` → `0`

**Constraints:** n ≤ 100; up to 10⁵ meetings.
```

Process meetings by start time with two min-heaps: **free rooms** by number and **busy rooms** by `(end time, room)`. Before each meeting, release rooms that have ended. If a room is free, use the lowest-numbered one; otherwise wait for the earliest-ending room (lowest number on ties) and **delay** the meeting, keeping its duration. Count uses per room.

**Complexity:** O(m log m + m log n).

```python
class Solution:
    def mostBooked(self, n: int, meetings: List[List[int]]) -> int:
        free, busy, count = list(range(n)), [], [0] * n
        for s, e in sorted(meetings):
            while busy and busy[0][0] <= s:
                heappush(free, heappop(busy)[1])
            if free:
                room = heappop(free)
                heappush(busy, (e, room))
            else:
                end, room = heappop(busy)
                heappush(busy, (end + e - s, room))
            count[room] += 1
        return count.index(max(count))
# test
assert Solution().mostBooked(2, [[0, 10], [1, 5], [2, 7], [3, 4]]) == 0
assert Solution().mostBooked(3, [[1, 20], [2, 10], [3, 5], [4, 9], [6, 8]]) == 1
assert Solution().mostBooked(1, [[0, 1]]) == 0
assert Solution().mostBooked(4, [[18, 19], [3, 12], [17, 19], [2, 13], [7, 10]]) == 0
```

```java
class Solution {
    public int mostBooked(int n, int[][] meetings) {
        Arrays.sort(meetings, (a, b) -> Integer.compare(a[0], b[0]));
        PriorityQueue<Integer> free = new PriorityQueue<>();
        for (int i = 0; i < n; i++) free.add(i);
        PriorityQueue<long[]> busy = new PriorityQueue<>((a, b) -> a[0] != b[0] ? Long.compare(a[0], b[0]) : Long.compare(a[1], b[1]));
        int[] count = new int[n];
        for (int[] m : meetings) {
            while (!busy.isEmpty() && busy.peek()[0] <= m[0]) free.add((int) busy.poll()[1]);
            int room;
            if (!free.isEmpty()) { room = free.poll(); busy.add(new long[]{m[1], room}); }
            else { long[] b = busy.poll(); room = (int) b[1]; busy.add(new long[]{b[0] + m[1] - m[0], room}); }
            count[room]++;
        }
        int best = 0;
        for (int i = 1; i < n; i++) if (count[i] > count[best]) best = i;
        return best;
    }
}
// test
check(new Solution().mostBooked(2, new int[][]{{0, 10}, {1, 5}, {2, 7}, {3, 4}}), 0);
check(new Solution().mostBooked(3, new int[][]{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}), 1);
```

```go
type IntHeap []int // free rooms, lowest number first

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

type busyHeap [][2]int // (free-at time, room): earliest first, then the lowest room

func (h busyHeap) Len() int { return len(h) }
func (h busyHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0] || h[i][0] == h[j][0] && h[i][1] < h[j][1]
}
func (h busyHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *busyHeap) Push(x any)   { *h = append(*h, x.([2]int)) }
func (h *busyHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][0] < meetings[j][0] })
	free, busy, count := &IntHeap{}, &busyHeap{}, make([]int, n)
	for r := 0; r < n; r++ {
		heap.Push(free, r)
	}
	for _, m := range meetings {
		s, e := m[0], m[1]
		for busy.Len() > 0 && (*busy)[0][0] <= s { // release rooms free by s
			heap.Push(free, heap.Pop(busy).([2]int)[1])
		}
		var room int
		if free.Len() > 0 {
			room = heap.Pop(free).(int)
			heap.Push(busy, [2]int{e, room})
		} else { // wait for the earliest room; the meeting keeps its duration
			top := heap.Pop(busy).([2]int)
			room = top[1]
			heap.Push(busy, [2]int{top[0] + e - s, room})
		}
		count[room]++
	}
	best := 0
	for r := range count {
		if count[r] > count[best] {
			best = r
		}
	}
	return best
}
// test
assertEq(mostBooked(2, [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}), 0)
assertEq(mostBooked(3, [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}), 1)
assertEq(mostBooked(1, [][]int{{0, 1}}), 0)
assertEq(mostBooked(4, [][]int{{18, 19}, {3, 12}, {17, 19}, {2, 13}, {7, 10}}), 0)
```
