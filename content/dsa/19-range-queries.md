# Prefix Sums & Range Queries

**When to use:** many queries/updates over ranges of an array — "sum of a[l..r]", "add v to a[l..r]", "how many smaller elements to the right".

| Need | Structure | Query | Update |
|---|---|---|---|
| Static range sums | prefix sums `pre[i+1] = pre[i] + a[i]` | O(1) | — |
| Many range **adds**, read at the end | **difference array** `d[l] += v, d[r+1] -= v`, prefix-sum once | O(n) total | O(1) |
| Range sums + point updates | **Fenwick tree (BIT)** | O(log n) | O(log n) |
| Range min/max/sum + range updates | **segment tree** (lazy propagation) | O(log n) | O(log n) |
| Counting inversions / pairs | merge sort with counting, or BIT over ranks | O(n log n) | — |

```python
class BIT:                         # Fenwick tree, 1-indexed
    def __init__(self, n): self.t = [0] * (n + 1)
    def add(self, i, v):           # a[i] += v
        while i < len(self.t):
            self.t[i] += v
            i += i & -i            # jump to the next node covering i
    def sum(self, i):              # a[1] + ... + a[i]
        s = 0
        while i:
            s += self.t[i]
            i -= i & -i            # drop the lowest set bit
        return s
```

```viz
difference array: add 10 to [1,3] and 5 to [2,4] on n=6
d:      0  10   5   0 -10  -5      (d[l] += v, d[r+1] -= v)
prefix: 0  10  15  15   5   0      <- final values
```

## 303. Range Sum Query - Immutable | Easy
related: 304, 307, 560
tags: prefix-sum

Precompute `pre[i] = a[0] + … + a[i−1]`; then `sum(l, r) = pre[r+1] − pre[l]`.

**Complexity:** O(n) build, O(1) query.

```python
class NumArray:
    def __init__(self, nums: List[int]):
        self.pre = [0]
        for x in nums:
            self.pre.append(self.pre[-1] + x)

    def sumRange(self, left: int, right: int) -> int:
        return self.pre[right + 1] - self.pre[left]
# test
na = NumArray([-2, 0, 3, -5, 2, -1])
assert [na.sumRange(0, 2), na.sumRange(2, 5), na.sumRange(0, 5)] == [1, -1, -3]
assert NumArray([5]).sumRange(0, 0) == 5
```

```java
class NumArray {
    private final int[] pre;

    public NumArray(int[] nums) {
        pre = new int[nums.length + 1];
        for (int i = 0; i < nums.length; i++) pre[i + 1] = pre[i] + nums[i];
    }

    public int sumRange(int left, int right) { return pre[right + 1] - pre[left]; }
}
// test
NumArray na = new NumArray(new int[]{-2, 0, 3, -5, 2, -1});
check(na.sumRange(2, 5), -1);
```

## 304. Range Sum Query 2D - Immutable | Medium
related: 303, 221
tags: prefix-sum, 2d

`P[r][c]` = sum of the rectangle above-left of `(r, c)` (exclusive). Inclusion–exclusion: `sum = P[r2+1][c2+1] − P[r1][c2+1] − P[r2+1][c1] + P[r1][c1]`.

```viz
+-----+---------+
|  A  |    B    |      region = (A+B+C+D) - (A+B) - (A+C) + A
+-----+---------+
|  C  |  [D]    |
+-----+---------+
```

**Complexity:** O(m·n) build, O(1) query.

```python
class NumMatrix:
    def __init__(self, matrix: List[List[int]]):
        R, C = len(matrix), len(matrix[0])
        self.P = [[0] * (C + 1) for _ in range(R + 1)]
        for r in range(R):
            for c in range(C):
                self.P[r + 1][c + 1] = matrix[r][c] + self.P[r][c + 1] + self.P[r + 1][c] - self.P[r][c]

    def sumRegion(self, row1: int, col1: int, row2: int, col2: int) -> int:
        P = self.P
        return P[row2 + 1][col2 + 1] - P[row1][col2 + 1] - P[row2 + 1][col1] + P[row1][col1]
# test
nm = NumMatrix([[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]])
assert [nm.sumRegion(2, 1, 4, 3), nm.sumRegion(1, 1, 2, 2), nm.sumRegion(1, 2, 2, 4)] == [8, 11, 12]
assert NumMatrix([[1, 2], [3, 4]]).sumRegion(0, 0, 1, 1) == 10
assert NumMatrix([[1, 2], [3, 4]]).sumRegion(1, 1, 1, 1) == 4
```

```java
class NumMatrix {
    private final int[][] P;

    public NumMatrix(int[][] matrix) {
        int R = matrix.length, C = matrix[0].length;
        P = new int[R + 1][C + 1];
        for (int r = 0; r < R; r++)
            for (int c = 0; c < C; c++) P[r + 1][c + 1] = matrix[r][c] + P[r][c + 1] + P[r + 1][c] - P[r][c];
    }

    public int sumRegion(int r1, int c1, int r2, int c2) {
        return P[r2 + 1][c2 + 1] - P[r1][c2 + 1] - P[r2 + 1][c1] + P[r1][c1];
    }
}
// test
NumMatrix nm = new NumMatrix(new int[][]{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}});
check(nm.sumRegion(2, 1, 4, 3), 8);
check(nm.sumRegion(1, 2, 2, 4), 12);
```

## 307. Range Sum Query - Mutable | Medium
related: 303, 315
tags: fenwick-tree, segment-tree

Point updates break plain prefix sums (O(n) per update). A **Fenwick tree** stores partial sums over power-of-two ranges: both update and prefix query walk O(log n) nodes. `update` adds the delta `val − nums[i]`.

**Complexity:** O(log n) update and query.

```python
class NumArray:
    def __init__(self, nums: List[int]):
        self.n, self.nums = len(nums), nums[:]
        self.t = [0] * (self.n + 1)
        for i, x in enumerate(nums):
            self._add(i + 1, x)

    def _add(self, i, v):
        while i <= self.n:
            self.t[i] += v
            i += i & -i

    def _pre(self, i):
        s = 0
        while i:
            s += self.t[i]
            i -= i & -i
        return s

    def update(self, index: int, val: int) -> None:
        self._add(index + 1, val - self.nums[index])
        self.nums[index] = val

    def sumRange(self, left: int, right: int) -> int:
        return self._pre(right + 1) - self._pre(left)
# test
na = NumArray([1, 3, 5])
assert na.sumRange(0, 2) == 9
na.update(1, 2)
assert na.sumRange(0, 2) == 8 and na.sumRange(1, 1) == 2
```

```java
class NumArray {
    private final int[] t, nums;
    private final int n;

    public NumArray(int[] nums) {
        n = nums.length;
        this.nums = nums.clone();
        t = new int[n + 1];
        for (int i = 0; i < n; i++) add(i + 1, nums[i]);
    }

    private void add(int i, int v) { for (; i <= n; i += i & -i) t[i] += v; }

    private int pre(int i) { int s = 0; for (; i > 0; i -= i & -i) s += t[i]; return s; }

    public void update(int index, int val) { add(index + 1, val - nums[index]); nums[index] = val; }

    public int sumRange(int left, int right) { return pre(right + 1) - pre(left); }
}
// test
NumArray na = new NumArray(new int[]{1, 3, 5});
check(na.sumRange(0, 2), 9);
na.update(1, 2);
check(na.sumRange(0, 2), 8);
```

## 1109. Corporate Flight Bookings | Medium
related: 1094, 370
tags: difference-array

Each booking adds `seats` to a range of flights: `d[first] += seats`, `d[last + 1] -= seats`. One prefix-sum pass produces all totals.

**Complexity:** O(n + bookings).

```python
class Solution:
    def corpFlightBookings(self, bookings: List[List[int]], n: int) -> List[int]:
        d = [0] * (n + 1)
        for first, last, seats in bookings:
            d[first - 1] += seats
            d[last] -= seats
        return list(itertools.accumulate(d[:n]))
# test
assert Solution().corpFlightBookings([[1, 2, 10], [2, 3, 20], [2, 5, 25]], 5) == [10, 55, 45, 25, 25]
assert Solution().corpFlightBookings([[1, 2, 10], [2, 2, 15]], 2) == [10, 25]
```

```java
class Solution {
    public int[] corpFlightBookings(int[][] bookings, int n) {
        int[] d = new int[n + 1];
        for (int[] b : bookings) { d[b[0] - 1] += b[2]; d[b[1]] -= b[2]; }
        int[] res = new int[n];
        for (int i = 0, s = 0; i < n; i++) res[i] = s += d[i];
        return res;
    }
}
// test
check(new Solution().corpFlightBookings(new int[][]{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5), new int[]{10, 55, 45, 25, 25});
```

## 1094. Car Pooling | Medium
related: 1109, 253
tags: difference-array, sweep-line

Passengers board at `from` and leave at `to`: `d[from] += n`, `d[to] −= n`. Sweep the locations; if the running load ever exceeds capacity → false.

**Complexity:** O(trips + max location).

```python
class Solution:
    def carPooling(self, trips: List[List[int]], capacity: int) -> bool:
        d = [0] * 1001
        for n, a, b in trips:
            d[a] += n
            d[b] -= n
        load = 0
        for x in d:
            load += x
            if load > capacity:
                return False
        return True
# test
assert Solution().carPooling([[2, 1, 5], [3, 3, 7]], 4) is False
assert Solution().carPooling([[2, 1, 5], [3, 3, 7]], 5) is True
assert Solution().carPooling([[2, 1, 5], [3, 5, 7]], 3) is True
```

```java
class Solution {
    public boolean carPooling(int[][] trips, int capacity) {
        int[] d = new int[1001];
        for (int[] t : trips) { d[t[1]] += t[0]; d[t[2]] -= t[0]; }
        int load = 0;
        for (int x : d) if ((load += x) > capacity) return false;
        return true;
    }
}
// test
check(new Solution().carPooling(new int[][]{{2, 1, 5}, {3, 3, 7}}, 4), false);
check(new Solution().carPooling(new int[][]{{2, 1, 5}, {3, 5, 7}}, 3), true);
```

## 315. Count of Smaller Numbers After Self | Hard
related: 493, 307
tags: fenwick-tree, coordinate-compression

Scan from the right. Compress values to ranks `1..m`; a Fenwick tree counts how many already-seen (i.e. to the right) values have a smaller rank: `query(rank − 1)`, then `add(rank, 1)`.

```viz
nums 5 2 6 1        ranks 3 2 4 1
i=3 (1): seen {}        smaller 0   add rank1
i=2 (6): seen {1}       smaller 1   add rank4
i=1 (2): seen {1,6}     smaller 1   add rank2
i=0 (5): seen {1,6,2}   smaller 2   -> [2, 1, 1, 0]
```

**Complexity:** O(n log n).

```python
class Solution:
    def countSmaller(self, nums: List[int]) -> List[int]:
        rank = {v: i + 1 for i, v in enumerate(sorted(set(nums)))}
        t = [0] * (len(rank) + 1)
        res = []
        for x in reversed(nums):
            i, s = rank[x] - 1, 0
            while i:
                s += t[i]
                i -= i & -i
            res.append(s)
            i = rank[x]
            while i < len(t):
                t[i] += 1
                i += i & -i
        return res[::-1]
# test
assert Solution().countSmaller([5, 2, 6, 1]) == [2, 1, 1, 0]
assert Solution().countSmaller([-1, -1]) == [0, 0]
```

```java
class Solution {
    public List<Integer> countSmaller(int[] nums) {
        int[] sorted = Arrays.stream(nums).distinct().sorted().toArray();
        int m = sorted.length;
        int[] t = new int[m + 1];
        Integer[] res = new Integer[nums.length];
        for (int k = nums.length - 1; k >= 0; k--) {
            int r = Arrays.binarySearch(sorted, nums[k]) + 1, s = 0;
            for (int i = r - 1; i > 0; i -= i & -i) s += t[i];
            res[k] = s;
            for (int i = r; i <= m; i += i & -i) t[i]++;
        }
        return Arrays.asList(res);
    }
}
// test
check(new Solution().countSmaller(new int[]{5, 2, 6, 1}), "[2, 1, 1, 0]");
```

## 493. Reverse Pairs | Hard
related: 315
tags: merge-sort, divide-and-conquer

Merge sort: after sorting both halves, count pairs `i` in the left, `j` in the right with `a[i] > 2·a[j]` using a single moving pointer (both halves sorted), then merge normally.

**Complexity:** O(n log n).

```python
class Solution:
    def reversePairs(self, nums: List[int]) -> int:
        def sort(a):
            if len(a) <= 1:
                return a, 0
            mid = len(a) // 2
            left, cl = sort(a[:mid])
            right, cr = sort(a[mid:])
            count, j = cl + cr, 0
            for x in left:
                while j < len(right) and x > 2 * right[j]:
                    j += 1
                count += j
            merged, i, k = [], 0, 0
            while i < len(left) and k < len(right):
                if left[i] <= right[k]:
                    merged.append(left[i]); i += 1
                else:
                    merged.append(right[k]); k += 1
            return merged + left[i:] + right[k:], count

        return sort(nums)[1]
# test
assert Solution().reversePairs([1, 3, 2, 3, 1]) == 2
assert Solution().reversePairs([2, 4, 3, 5, 1]) == 3
assert Solution().reversePairs([2147483647, 2147483647, 2147483647]) == 0
```

```java
class Solution {
    public int reversePairs(int[] nums) { return sort(nums, 0, nums.length - 1, new int[nums.length]); }

    private int sort(int[] a, int lo, int hi, int[] tmp) {
        if (lo >= hi) return 0;
        int mid = (lo + hi) / 2;
        int count = sort(a, lo, mid, tmp) + sort(a, mid + 1, hi, tmp);
        for (int i = lo, j = mid + 1; i <= mid; i++) {
            while (j <= hi && a[i] > 2L * a[j]) j++;
            count += j - (mid + 1);
        }
        int i = lo, j = mid + 1, k = lo;
        while (i <= mid && j <= hi) tmp[k++] = a[i] <= a[j] ? a[i++] : a[j++];
        while (i <= mid) tmp[k++] = a[i++];
        while (j <= hi) tmp[k++] = a[j++];
        System.arraycopy(tmp, lo, a, lo, hi - lo + 1);
        return count;
    }
}
// test
check(new Solution().reversePairs(new int[]{2, 4, 3, 5, 1}), 3);
check(new Solution().reversePairs(new int[]{2147483647, 2147483647, 2147483647}), 0);
```

## 218. The Skyline Problem | Hard
related: 253, 850
tags: sweep-line, max-heap

Sweep x-coordinates of all building edges. At each x, add buildings starting here to a max-heap of `(height, right)`, lazily pop buildings whose `right ≤ x`. The skyline changes whenever the current max height differs from the last output height.

```viz
buildings [2,9,10] [3,7,15] [5,12,12] [15,20,10] [19,24,8]
x=2  heap {10}          max 10  -> (2,10)
x=3  heap {15,10}       max 15  -> (3,15)
x=7  15 expires         max 12  -> (7,12)
x=12 all expire         max 0   -> (12,0)   ... (15,10) (20,8) (24,0)
```

**Complexity:** O(n log n).

```python
class Solution:
    def getSkyline(self, buildings: List[List[int]]) -> List[List[int]]:
        xs = sorted({x for l, r, _ in buildings for x in (l, r)})
        res, h, i = [], [(0, float("inf"))], 0        # heap of (-height, right); ground never ends
        for x in xs:
            while i < len(buildings) and buildings[i][0] <= x:
                l, r, ht = buildings[i]
                heappush(h, (-ht, r))
                i += 1
            while h[0][1] <= x:
                heappop(h)
            top = -h[0][0]
            if not res or res[-1][1] != top:
                res.append([x, top])
        return res
# test
B = [[2, 9, 10], [3, 7, 15], [5, 12, 12], [15, 20, 10], [19, 24, 8]]
assert Solution().getSkyline(B) == [[2, 10], [3, 15], [7, 12], [12, 0], [15, 10], [20, 8], [24, 0]]
assert Solution().getSkyline([[0, 2, 3], [2, 5, 3]]) == [[0, 3], [5, 0]]
```

```java
class Solution {
    public List<List<Integer>> getSkyline(int[][] buildings) {
        TreeSet<Integer> xs = new TreeSet<>();
        for (int[] b : buildings) { xs.add(b[0]); xs.add(b[1]); }
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> b[0] - a[0]);   // {height, right}
        pq.add(new int[]{0, Integer.MAX_VALUE});
        List<List<Integer>> res = new ArrayList<>();
        int i = 0;
        for (int x : xs) {
            while (i < buildings.length && buildings[i][0] <= x) { pq.add(new int[]{buildings[i][2], buildings[i][1]}); i++; }
            while (pq.peek()[1] <= x) pq.poll();
            int top = pq.peek()[0];
            if (res.isEmpty() || res.get(res.size() - 1).get(1) != top) res.add(List.of(x, top));
        }
        return res;
    }
}
// test
int[][] B = {{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}};
check(new Solution().getSkyline(B), "[[2, 10], [3, 15], [7, 12], [12, 0], [15, 10], [20, 8], [24, 0]]");
```

## 1732. Find the Highest Altitude | Easy
related: 724, 303
tags: prefix-sum

The altitudes are the prefix sums of `gain`, starting from 0; return the largest one (including the starting 0).

**Complexity:** O(n).

```python
class Solution:
    def largestAltitude(self, gain: List[int]) -> int:
        return max(0, max(itertools.accumulate(gain)))
# test
assert Solution().largestAltitude([-5, 1, 5, 0, -7]) == 1
assert Solution().largestAltitude([-4, -3, -2, -1, 4, 3, 2]) == 0
assert Solution().largestAltitude([3]) == 3
```

```java
class Solution {
    public int largestAltitude(int[] gain) {
        int alt = 0, best = 0;
        for (int g : gain) best = Math.max(best, alt += g);
        return best;
    }
}
// test
check(new Solution().largestAltitude(new int[]{-5, 1, 5, 0, -7}), 1);
check(new Solution().largestAltitude(new int[]{-4, -3, -2, -1, 4, 3, 2}), 0);
```

## 724. Find Pivot Index | Easy
related: 1732, 238, 560
tags: prefix-sum

With `total` known, the right sum at `i` is `total − left − nums[i]`. Scan once, maintaining `left`; return the first `i` where both sides match.

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def pivotIndex(self, nums: List[int]) -> int:
        total, left = sum(nums), 0
        for i, x in enumerate(nums):
            if left == total - left - x:
                return i
            left += x
        return -1
# test
assert Solution().pivotIndex([1, 7, 3, 6, 5, 6]) == 3
assert Solution().pivotIndex([1, 2, 3]) == -1
assert Solution().pivotIndex([2, 1, -1]) == 0
assert Solution().pivotIndex([0]) == 0
assert Solution().pivotIndex([-1, -1, 0, 1, 1, 0]) == 5
```

```java
class Solution {
    public int pivotIndex(int[] nums) {
        int total = 0, left = 0;
        for (int x : nums) total += x;
        for (int i = 0; i < nums.length; i++) {
            if (left == total - left - nums[i]) return i;
            left += nums[i];
        }
        return -1;
    }
}
// test
check(new Solution().pivotIndex(new int[]{1, 7, 3, 6, 5, 6}), 3);
check(new Solution().pivotIndex(new int[]{2, 1, -1}), 0);
```
