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

```problem
Design `NumArray(nums)` whose `sumRange(left, right)` returns the sum of `nums[left..right]` (inclusive). The array never changes, and there are many queries.

**Example:** `NumArray([-2,0,3,-5,2,-1])` → `sumRange(0,2)` = 1, `sumRange(2,5)` = −1, `sumRange(0,5)` = −3

**Constraints:** up to 10⁴ queries.
```

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

```go
type NumArray struct {
	pre []int // pre[i] = sum of nums[:i]
}

func Constructor(nums []int) NumArray {
	pre := make([]int, len(nums)+1)
	for i, x := range nums {
		pre[i+1] = pre[i] + x
	}
	return NumArray{pre: pre}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.pre[right+1] - this.pre[left]
}
// test
na := Constructor([]int{-2, 0, 3, -5, 2, -1})
assertEq([]int{na.SumRange(0, 2), na.SumRange(2, 5), na.SumRange(0, 5)}, []int{1, -1, -3})
one := Constructor([]int{5})
assertEq(one.SumRange(0, 0), 5)
```

## 304. Range Sum Query 2D - Immutable | Medium
related: 303, 221
tags: prefix-sum, 2d

```problem
Design `NumMatrix(matrix)` whose `sumRegion(row1, col1, row2, col2)` returns the sum of the rectangle with those top-left and bottom-right corners (inclusive), in O(1) per query. The matrix never changes.

**Example:** for the 5×5 matrix `[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]`, `sumRegion(2,1,4,3)` → `8`

**Constraints:** up to 200 × 200; up to 10⁴ queries.
```

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

```go
type NumMatrix struct {
	P [][]int // P[r][c] = sum of the r x c top-left block
}

func Constructor(matrix [][]int) NumMatrix {
	R, C := len(matrix), len(matrix[0])
	P := make([][]int, R+1)
	for r := range P {
		P[r] = make([]int, C+1)
	}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			P[r+1][c+1] = matrix[r][c] + P[r][c+1] + P[r+1][c] - P[r][c]
		}
	}
	return NumMatrix{P: P}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	P := this.P // inclusion-exclusion on four prefix blocks
	return P[row2+1][col2+1] - P[row1][col2+1] - P[row2+1][col1] + P[row1][col1]
}
// test
nm := Constructor([][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}})
assertEq([]int{nm.SumRegion(2, 1, 4, 3), nm.SumRegion(1, 1, 2, 2), nm.SumRegion(1, 2, 2, 4)}, []int{8, 11, 12})
small := Constructor([][]int{{1, 2}, {3, 4}})
assertEq(small.SumRegion(0, 0, 1, 1), 10)
assertEq(small.SumRegion(1, 1, 1, 1), 4)
```

## 307. Range Sum Query - Mutable | Medium
related: 303, 315
tags: fenwick-tree, segment-tree

```problem
Design `NumArray(nums)` with `update(index, val)`, which sets `nums[index] = val`, and `sumRange(left, right)`, which returns the sum of `nums[left..right]`. Both kinds of calls are frequent.

**Example:** `NumArray([1,3,5])` → `sumRange(0,2)` = 9; `update(1,2)` → `sumRange(0,2)` = 8

**Constraints:** n ≤ 3·10⁴; up to 3·10⁴ calls.
```

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

```go
type NumArray struct {
	nums []int
	tree []int // Fenwick tree, 1-indexed
}

func Constructor(nums []int) NumArray {
	na := NumArray{nums: append([]int(nil), nums...), tree: make([]int, len(nums)+1)}
	for i, x := range nums {
		na.add(i+1, x)
	}
	return na
}

func (this *NumArray) add(i, v int) {
	for ; i < len(this.tree); i += i & -i {
		this.tree[i] += v
	}
}

func (this *NumArray) prefix(i int) int { // sum of nums[:i]
	s := 0
	for ; i > 0; i -= i & -i {
		s += this.tree[i]
	}
	return s
}

func (this *NumArray) Update(index int, val int) {
	this.add(index+1, val-this.nums[index])
	this.nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefix(right+1) - this.prefix(left)
}
// test
na := Constructor([]int{1, 3, 5})
assertEq(na.SumRange(0, 2), 9)
na.Update(1, 2)
assertEq([]int{na.SumRange(0, 2), na.SumRange(1, 1)}, []int{8, 2})
```

## 1109. Corporate Flight Bookings | Medium
related: 1094, 370
tags: difference-array

```problem
Flights are numbered 1…n. `bookings[i] = [first, last, seats]` reserves `seats` seats on every flight from `first` to `last`. Return the total number of seats reserved on each flight.

**Example:** `bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5` → `[10,55,45,25,25]`

**Constraints:** n and the number of bookings ≤ 2·10⁴.
```

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

```go
func corpFlightBookings(bookings [][]int, n int) []int {
	d := make([]int, n+1) // difference array: +seats at first, -seats after last
	for _, b := range bookings {
		d[b[0]-1] += b[2]
		d[b[1]] -= b[2]
	}
	res := make([]int, n)
	run := 0
	for i := 0; i < n; i++ {
		run += d[i]
		res[i] = run
	}
	return res
}
// test
assertEq(corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5), []int{10, 55, 45, 25, 25})
assertEq(corpFlightBookings([][]int{{1, 2, 10}, {2, 2, 15}}, 2), []int{10, 25})
```

## 1094. Car Pooling | Medium
related: 1109, 253
tags: difference-array, sweep-line

```problem
A car with `capacity` empty seats drives east only. `trips[i] = [passengers, from, to]` picks up `passengers` at kilometre `from` and drops them off at `to`. Decide whether every trip can be served without ever exceeding capacity.

**Example:** `trips = [[2,1,5],[3,3,7]]`: `capacity = 4` → `false`; `capacity = 5` → `true`

**Constraints:** up to 1000 trips; locations ≤ 1000.
```

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

```go
func carPooling(trips [][]int, capacity int) bool {
	d := make([]int, 1001) // passenger change at each location
	for _, t := range trips {
		d[t[1]] += t[0]
		d[t[2]] -= t[0] // they leave at t[2], before anyone boards there
	}
	load := 0
	for _, x := range d {
		load += x
		if load > capacity {
			return false
		}
	}
	return true
}
// test
assertEq(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 4), false)
assertEq(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 5), true)
assertEq(carPooling([][]int{{2, 1, 5}, {3, 5, 7}}, 3), true)
```

## 315. Count of Smaller Numbers After Self | Hard
related: 493, 307
tags: fenwick-tree, coordinate-compression

```problem
For each `nums[i]`, count how many elements to its right are strictly smaller, and return these counts.

**Example:** `[5,2,6,1]` → `[2,1,1,0]`

**Constraints:** n ≤ 10⁵.
```

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

```go
func countSmaller(nums []int) []int {
	vals := append([]int(nil), nums...)
	sort.Ints(vals)
	rank := map[int]int{} // value -> 1-based rank among distinct values
	for _, v := range vals {
		if _, ok := rank[v]; !ok {
			rank[v] = len(rank) + 1
		}
	}
	tree := make([]int, len(rank)+1) // Fenwick tree of counts by rank, filled right to left
	res := make([]int, len(nums))
	for k := len(nums) - 1; k >= 0; k-- {
		s := 0
		for i := rank[nums[k]] - 1; i > 0; i -= i & -i { // how many smaller values are to the right
			s += tree[i]
		}
		res[k] = s
		for i := rank[nums[k]]; i < len(tree); i += i & -i {
			tree[i]++
		}
	}
	return res
}
// test
assertEq(countSmaller([]int{5, 2, 6, 1}), []int{2, 1, 1, 0})
assertEq(countSmaller([]int{-1, -1}), []int{0, 0})
```

## 493. Reverse Pairs | Hard
related: 315
tags: merge-sort, divide-and-conquer

```problem
Count the pairs (i, j) with i < j and nums[i] > 2 × nums[j].

**Example:** `[1,3,2,3,1]` → `2`; `[2,4,3,5,1]` → `3`

**Constraints:** n ≤ 5·10⁴; values span the full 32-bit range.
```

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

```go
func reversePairs(nums []int) int {
	var sortCount func(a []int) ([]int, int) // merge sort that also counts the pairs
	sortCount = func(a []int) ([]int, int) {
		if len(a) <= 1 {
			return a, 0
		}
		mid := len(a) / 2
		left, cl := sortCount(append([]int(nil), a[:mid]...))
		right, cr := sortCount(append([]int(nil), a[mid:]...))
		count, j := cl+cr, 0
		for _, x := range left { // both halves sorted: j only moves forward
			for j < len(right) && x > 2*right[j] {
				j++
			}
			count += j
		}
		merged := make([]int, 0, len(a))
		i, k := 0, 0
		for i < len(left) && k < len(right) {
			if left[i] <= right[k] {
				merged = append(merged, left[i])
				i++
			} else {
				merged = append(merged, right[k])
				k++
			}
		}
		merged = append(append(merged, left[i:]...), right[k:]...)
		return merged, count
	}
	_, count := sortCount(nums)
	return count
}
// test
assertEq(reversePairs([]int{1, 3, 2, 3, 1}), 2)
assertEq(reversePairs([]int{2, 4, 3, 5, 1}), 3)
assertEq(reversePairs([]int{2147483647, 2147483647, 2147483647}), 0)
```

## 218. The Skyline Problem | Hard
related: 253, 850
tags: sweep-line, max-heap

```problem
`buildings[i] = [left, right, height]` are rectangles standing on flat ground. Return the skyline (the outer contour of the union) as a list of key points `[x, y]` sorted by x: each marks the left end of a horizontal segment, and the last point has y = 0. Consecutive points must not have equal heights.

**Example:** `[[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]` → `[[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]`

**Constraints:** up to 10⁴ buildings.
```

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

```go
type heightHeap [][2]int // (height, right edge), tallest first

func (h heightHeap) Len() int           { return len(h) }
func (h heightHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h heightHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *heightHeap) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *heightHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func getSkyline(buildings [][]int) [][]int {
	xs := []int{} // every edge is a candidate key point
	for _, b := range buildings {
		xs = append(xs, b[0], b[1])
	}
	sort.Ints(xs)
	res := [][]int{}
	h := &heightHeap{{0, math.MaxInt}} // the ground never ends
	i := 0
	for k, x := range xs {
		if k > 0 && x == xs[k-1] {
			continue
		}
		for ; i < len(buildings) && buildings[i][0] <= x; i++ { // buildings started by x
			heap.Push(h, [2]int{buildings[i][2], buildings[i][1]})
		}
		for (*h)[0][1] <= x { // lazily drop the tallest while it has ended
			heap.Pop(h)
		}
		if top := (*h)[0][0]; len(res) == 0 || res[len(res)-1][1] != top {
			res = append(res, []int{x, top})
		}
	}
	return res
}
// test
B := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
assertEq(getSkyline(B), [][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}})
assertEq(getSkyline([][]int{{0, 2, 3}, {2, 5, 3}}), [][]int{{0, 3}, {5, 0}})
```

## 1732. Find the Highest Altitude | Easy
related: 724, 303
tags: prefix-sum

```problem
A biker starts at altitude 0; `gain[i]` is the net altitude change between points i and i + 1. Return the highest altitude reached.

**Example:** `[-5,1,5,0,-7]` → `1` (altitudes 0, −5, −4, 1, 1, −6)
```

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

```go
func largestAltitude(gain []int) int {
	alt, best := 0, 0 // the start (altitude 0) counts too
	for _, g := range gain {
		alt += g
		best = max(best, alt)
	}
	return best
}
// test
assertEq(largestAltitude([]int{-5, 1, 5, 0, -7}), 1)
assertEq(largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}), 0)
assertEq(largestAltitude([]int{3}), 3)
```

## 724. Find Pivot Index | Easy
related: 1732, 238, 560
tags: prefix-sum

```problem
Return the leftmost index where the sum of the elements strictly to its left equals the sum of the elements strictly to its right (an empty side sums to 0), or −1 if there is none.

**Example:** `[1,7,3,6,5,6]` → `3` (1 + 7 + 3 = 5 + 6); `[1,2,3]` → `-1`

**Constraints:** n ≤ 10⁴.
```

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

```go
func pivotIndex(nums []int) int {
	total, left := 0, 0
	for _, x := range nums {
		total += x
	}
	for i, x := range nums {
		if left == total-left-x { // left sum equals right sum
			return i
		}
		left += x
	}
	return -1
}
// test
assertEq(pivotIndex([]int{1, 7, 3, 6, 5, 6}), 3)
assertEq(pivotIndex([]int{1, 2, 3}), -1)
assertEq(pivotIndex([]int{2, 1, -1}), 0)
assertEq(pivotIndex([]int{0}), 0)
assertEq(pivotIndex([]int{-1, -1, 0, 1, 1, 0}), 5)
```
