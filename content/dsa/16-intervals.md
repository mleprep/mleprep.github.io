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

## 57. Insert Interval | Medium
related: 56
tags: blind75, sweep-line

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

## 435. Non-overlapping Intervals | Medium
related: 452, 56
tags: blind75, sort-by-end

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

## 452. Minimum Number of Arrows to Burst Balloons | Medium
related: 435
tags: sort-by-end

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

## 252. Meeting Rooms | Easy
related: 253, 56
tags: blind75, premium

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

## 253. Meeting Rooms II | Medium
related: 252, 1851
tags: blind75, sweep-line, premium

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

## 986. Interval List Intersections | Medium
related: 56
tags: two-pointers

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

## 1851. Minimum Interval to Include Each Query | Hard
related: 253, 56
tags: offline-queries, heap

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


## 1288. Remove Covered Intervals | Medium
related: 56, 435
tags: sort-by-start

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

## 729. My Calendar I | Medium
related: 731, 56
tags: sorted-map, design

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

## 352. Data Stream as Disjoint Intervals | Hard
related: 57, 228
tags: sorted-map, merging

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
