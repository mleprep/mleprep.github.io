# Binary Search

**When to use:** sorted data, or any **monotonic predicate** `ok(x)` (false…false true…true). "Minimize the maximum" / "smallest capacity such that …" → **binary search on the answer**.

One template to rule them all — find the first index where `ok` is true:

```python
lo, hi = 0, n            # answer in [lo, hi]; hi = "none found"
while lo < hi:
    mid = (lo + hi) // 2
    if ok(mid): hi = mid     # mid could be the answer
    else:       lo = mid + 1
return lo
```

```viz
ok:   F  F  F  T  T  T
      lo       ^first T       hi
```

Pitfalls: in Java use `lo + (hi - lo) / 2`; make sure each branch shrinks the range.

## 704. Binary Search | Easy
related: 35, 34
tags: classic

Classic closed interval `[lo, hi]`.

**Complexity:** O(log n).

```python
class Solution:
    def search(self, nums: List[int], target: int) -> int:
        lo, hi = 0, len(nums) - 1
        while lo <= hi:
            mid = (lo + hi) // 2
            if nums[mid] == target:
                return mid
            if nums[mid] < target:
                lo = mid + 1
            else:
                hi = mid - 1
        return -1
# test
assert Solution().search([-1, 0, 3, 5, 9, 12], 9) == 4
assert Solution().search([-1, 0, 3, 5, 9, 12], 2) == -1
```

```java
class Solution {
    public int search(int[] nums, int target) {
        int lo = 0, hi = nums.length - 1;
        while (lo <= hi) {
            int mid = lo + (hi - lo) / 2;
            if (nums[mid] == target) return mid;
            if (nums[mid] < target) lo = mid + 1; else hi = mid - 1;
        }
        return -1;
    }
}
// test
check(new Solution().search(new int[]{-1, 0, 3, 5, 9, 12}, 9), 4);
check(new Solution().search(new int[]{-1, 0, 3, 5, 9, 12}, 2), -1);
```

## 35. Search Insert Position | Easy
related: 704, 34
tags: lower-bound

First index with `nums[i] >= target` (lower bound). Uses the template with `hi = n`.

**Complexity:** O(log n).

```python
class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        lo, hi = 0, len(nums)
        while lo < hi:
            mid = (lo + hi) // 2
            if nums[mid] >= target:
                hi = mid
            else:
                lo = mid + 1
        return lo
# test
assert Solution().searchInsert([1, 3, 5, 6], 5) == 2
assert Solution().searchInsert([1, 3, 5, 6], 2) == 1
assert Solution().searchInsert([1, 3, 5, 6], 7) == 4
```

```java
class Solution {
    public int searchInsert(int[] nums, int target) {
        int lo = 0, hi = nums.length;
        while (lo < hi) {
            int mid = lo + (hi - lo) / 2;
            if (nums[mid] >= target) hi = mid; else lo = mid + 1;
        }
        return lo;
    }
}
// test
check(new Solution().searchInsert(new int[]{1, 3, 5, 6}, 2), 1);
check(new Solution().searchInsert(new int[]{1, 3, 5, 6}, 7), 4);
```

## 34. Find First and Last Position of Element in Sorted Array | Medium
related: 35, 704
tags: lower-bound

`first = lowerBound(target)`, `last = lowerBound(target + 1) − 1`. Check `first` actually holds `target`.

**Complexity:** O(log n).

```python
class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        def lower(t):
            lo, hi = 0, len(nums)
            while lo < hi:
                mid = (lo + hi) // 2
                if nums[mid] >= t:
                    hi = mid
                else:
                    lo = mid + 1
            return lo

        a = lower(target)
        if a == len(nums) or nums[a] != target:
            return [-1, -1]
        return [a, lower(target + 1) - 1]
# test
assert Solution().searchRange([5, 7, 7, 8, 8, 10], 8) == [3, 4]
assert Solution().searchRange([5, 7, 7, 8, 8, 10], 6) == [-1, -1]
assert Solution().searchRange([], 0) == [-1, -1]
```

```java
class Solution {
    public int[] searchRange(int[] nums, int target) {
        int a = lower(nums, target);
        if (a == nums.length || nums[a] != target) return new int[]{-1, -1};
        return new int[]{a, lower(nums, target + 1) - 1};
    }

    private int lower(int[] nums, int t) {
        int lo = 0, hi = nums.length;
        while (lo < hi) {
            int mid = lo + (hi - lo) / 2;
            if (nums[mid] >= t) hi = mid; else lo = mid + 1;
        }
        return lo;
    }
}
// test
check(new Solution().searchRange(new int[]{5, 7, 7, 8, 8, 10}, 8), new int[]{3, 4});
check(new Solution().searchRange(new int[]{5, 7, 7, 8, 8, 10}, 6), new int[]{-1, -1});
```

## 74. Search a 2D Matrix | Medium
related: 704, 240
tags: matrix

Rows are sorted and each row starts after the previous ends → it's one sorted array of length `m·n`. Index `k` ↦ `(k // n, k % n)`.

**Complexity:** O(log(m·n)).

```python
class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        m, n = len(matrix), len(matrix[0])
        lo, hi = 0, m * n - 1
        while lo <= hi:
            mid = (lo + hi) // 2
            v = matrix[mid // n][mid % n]
            if v == target:
                return True
            if v < target:
                lo = mid + 1
            else:
                hi = mid - 1
        return False
# test
M = [[1, 3, 5, 7], [10, 11, 16, 20], [23, 30, 34, 60]]
assert Solution().searchMatrix(M, 3) is True
assert Solution().searchMatrix(M, 13) is False
```

```java
class Solution {
    public boolean searchMatrix(int[][] matrix, int target) {
        int m = matrix.length, n = matrix[0].length, lo = 0, hi = m * n - 1;
        while (lo <= hi) {
            int mid = lo + (hi - lo) / 2, v = matrix[mid / n][mid % n];
            if (v == target) return true;
            if (v < target) lo = mid + 1; else hi = mid - 1;
        }
        return false;
    }
}
// test
int[][] M = {{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}};
check(new Solution().searchMatrix(M, 3), true);
check(new Solution().searchMatrix(M, 13), false);
```

## 875. Koko Eating Bananas | Medium
related: 1011, 410
tags: search-on-answer

`hours(speed)` is decreasing in speed → find the smallest speed in `[1, max(piles)]` with `hours ≤ h`.

```viz
piles [3,6,7,11], h=8
speed: 1  2  3  4  5 ...  11
ok?    F  F  F  T  T ...  T     -> 4
```

**Complexity:** O(n log max).

```python
class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        lo, hi = 1, max(piles)
        while lo < hi:
            k = (lo + hi) // 2
            if sum((p + k - 1) // k for p in piles) <= h:
                hi = k
            else:
                lo = k + 1
        return lo
# test
assert Solution().minEatingSpeed([3, 6, 7, 11], 8) == 4
assert Solution().minEatingSpeed([30, 11, 23, 4, 20], 5) == 30
assert Solution().minEatingSpeed([30, 11, 23, 4, 20], 6) == 23
```

```java
class Solution {
    public int minEatingSpeed(int[] piles, int h) {
        int lo = 1, hi = 0;
        for (int p : piles) hi = Math.max(hi, p);
        while (lo < hi) {
            int k = lo + (hi - lo) / 2;
            long hours = 0;
            for (int p : piles) hours += (p + k - 1) / k;
            if (hours <= h) hi = k; else lo = k + 1;
        }
        return lo;
    }
}
// test
check(new Solution().minEatingSpeed(new int[]{3, 6, 7, 11}, 8), 4);
check(new Solution().minEatingSpeed(new int[]{30, 11, 23, 4, 20}, 6), 23);
```

## 153. Find Minimum in Rotated Sorted Array | Medium
related: 33, 154
tags: blind75, rotated

Compare `nums[mid]` with `nums[hi]`: if `mid` is bigger, the drop (minimum) is to the right; otherwise `mid` could be the min, so `hi = mid`.

```viz
[4 5 6 7 0 1 2]
 lo    mid   hi   7 > 2  -> lo = mid+1
         lo m hi  1 < 2  -> hi = mid ... -> 0
```

**Complexity:** O(log n).

```python
class Solution:
    def findMin(self, nums: List[int]) -> int:
        lo, hi = 0, len(nums) - 1
        while lo < hi:
            mid = (lo + hi) // 2
            if nums[mid] > nums[hi]:
                lo = mid + 1
            else:
                hi = mid
        return nums[lo]
# test
assert Solution().findMin([3, 4, 5, 1, 2]) == 1
assert Solution().findMin([4, 5, 6, 7, 0, 1, 2]) == 0
assert Solution().findMin([11, 13, 15, 17]) == 11
```

```java
class Solution {
    public int findMin(int[] nums) {
        int lo = 0, hi = nums.length - 1;
        while (lo < hi) {
            int mid = lo + (hi - lo) / 2;
            if (nums[mid] > nums[hi]) lo = mid + 1; else hi = mid;
        }
        return nums[lo];
    }
}
// test
check(new Solution().findMin(new int[]{4, 5, 6, 7, 0, 1, 2}), 0);
check(new Solution().findMin(new int[]{11, 13, 15, 17}), 11);
```

## 33. Search in Rotated Sorted Array | Medium
related: 153, 81
tags: blind75, rotated

At least one half around `mid` is sorted. If `nums[lo] ≤ nums[mid]` the left half is sorted: target in `[nums[lo], nums[mid])` → go left, else right. Mirror for the right half.

**Complexity:** O(log n).

```python
class Solution:
    def search(self, nums: List[int], target: int) -> int:
        lo, hi = 0, len(nums) - 1
        while lo <= hi:
            mid = (lo + hi) // 2
            if nums[mid] == target:
                return mid
            if nums[lo] <= nums[mid]:                    # left half sorted
                if nums[lo] <= target < nums[mid]:
                    hi = mid - 1
                else:
                    lo = mid + 1
            else:                                        # right half sorted
                if nums[mid] < target <= nums[hi]:
                    lo = mid + 1
                else:
                    hi = mid - 1
        return -1
# test
assert Solution().search([4, 5, 6, 7, 0, 1, 2], 0) == 4
assert Solution().search([4, 5, 6, 7, 0, 1, 2], 3) == -1
assert Solution().search([1], 0) == -1
```

```java
class Solution {
    public int search(int[] nums, int target) {
        int lo = 0, hi = nums.length - 1;
        while (lo <= hi) {
            int mid = lo + (hi - lo) / 2;
            if (nums[mid] == target) return mid;
            if (nums[lo] <= nums[mid]) {
                if (nums[lo] <= target && target < nums[mid]) hi = mid - 1; else lo = mid + 1;
            } else {
                if (nums[mid] < target && target <= nums[hi]) lo = mid + 1; else hi = mid - 1;
            }
        }
        return -1;
    }
}
// test
check(new Solution().search(new int[]{4, 5, 6, 7, 0, 1, 2}, 0), 4);
check(new Solution().search(new int[]{4, 5, 6, 7, 0, 1, 2}, 3), -1);
```

## 162. Find Peak Element | Medium
related: 852
tags: slope

If `nums[mid] < nums[mid+1]` we're on an uphill slope → a peak exists to the right; else a peak is at `mid` or left. (Edges count as −∞.)

**Complexity:** O(log n).

```python
class Solution:
    def findPeakElement(self, nums: List[int]) -> int:
        lo, hi = 0, len(nums) - 1
        while lo < hi:
            mid = (lo + hi) // 2
            if nums[mid] < nums[mid + 1]:
                lo = mid + 1
            else:
                hi = mid
        return lo
# test
assert Solution().findPeakElement([1, 2, 3, 1]) == 2
assert Solution().findPeakElement([1, 2, 1, 3, 5, 6, 4]) in (1, 5)
```

```java
class Solution {
    public int findPeakElement(int[] nums) {
        int lo = 0, hi = nums.length - 1;
        while (lo < hi) {
            int mid = lo + (hi - lo) / 2;
            if (nums[mid] < nums[mid + 1]) lo = mid + 1; else hi = mid;
        }
        return lo;
    }
}
// test
check(new Solution().findPeakElement(new int[]{1, 2, 3, 1}), 2);
```

## 981. Time Based Key-Value Store | Medium
related: 704
tags: design, upper-bound

Timestamps per key arrive increasing → append to a list. `get` = binary search for the **last** timestamp ≤ t.

**Complexity:** O(1) set, O(log n) get.

```python
class TimeMap:
    def __init__(self):
        self.data = defaultdict(list)  # key -> [(ts, value)]

    def set(self, key: str, value: str, timestamp: int) -> None:
        self.data[key].append((timestamp, value))

    def get(self, key: str, timestamp: int) -> str:
        arr = self.data.get(key, [])
        i = bisect.bisect_right(arr, (timestamp, chr(0x10FFFF)))  # first ts > timestamp
        return arr[i - 1][1] if i else ""
# test
tm = TimeMap()
tm.set("foo", "bar", 1)
assert tm.get("foo", 1) == "bar" and tm.get("foo", 3) == "bar"
tm.set("foo", "bar2", 4)
assert tm.get("foo", 4) == "bar2" and tm.get("foo", 5) == "bar2" and tm.get("foo", 0) == ""
```

```java
class TimeMap {
    private final Map<String, TreeMap<Integer, String>> data = new HashMap<>();

    public void set(String key, String value, int timestamp) {
        data.computeIfAbsent(key, k -> new TreeMap<>()).put(timestamp, value);
    }

    public String get(String key, int timestamp) {
        TreeMap<Integer, String> t = data.get(key);
        if (t == null) return "";
        Map.Entry<Integer, String> e = t.floorEntry(timestamp); // largest ts <= timestamp
        return e == null ? "" : e.getValue();
    }
}
// test
TimeMap tm = new TimeMap();
tm.set("foo", "bar", 1);
check(tm.get("foo", 3), "bar");
tm.set("foo", "bar2", 4);
check(tm.get("foo", 5), "bar2");
check(tm.get("foo", 0), "");
```

## 4. Median of Two Sorted Arrays | Hard
related: 295
tags: partition

Binary search a **cut** in the shorter array `A` at `i`, with cut `j = half − i` in `B`. The cut is correct when `A[i-1] ≤ B[j]` and `B[j-1] ≤ A[i]`; then the median comes from the four border values.

```viz
A:  1  3 | 8  9          left part: {1,3,2,4,5}   right: {8,9,7,10,...}
B:  2  4  5 | 7  10      maxLeft = 5, minRight = 7
```

**Complexity:** O(log min(m, n)).

```python
class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        A, B = (nums1, nums2) if len(nums1) <= len(nums2) else (nums2, nums1)
        m, n = len(A), len(B)
        half = (m + n + 1) // 2
        lo, hi = 0, m
        while True:
            i = (lo + hi) // 2
            j = half - i
            al = A[i - 1] if i else float("-inf")
            ar = A[i] if i < m else float("inf")
            bl = B[j - 1] if j else float("-inf")
            br = B[j] if j < n else float("inf")
            if al <= br and bl <= ar:
                if (m + n) % 2:
                    return float(max(al, bl))
                return (max(al, bl) + min(ar, br)) / 2
            if al > br:
                hi = i - 1
            else:
                lo = i + 1
# test
assert Solution().findMedianSortedArrays([1, 3], [2]) == 2.0
assert Solution().findMedianSortedArrays([1, 2], [3, 4]) == 2.5
assert Solution().findMedianSortedArrays([], [1]) == 1.0
```

```java
class Solution {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int[] A = nums1.length <= nums2.length ? nums1 : nums2, B = A == nums1 ? nums2 : nums1;
        int m = A.length, n = B.length, half = (m + n + 1) / 2, lo = 0, hi = m;
        while (true) {
            int i = (lo + hi) / 2, j = half - i;
            int al = i > 0 ? A[i - 1] : Integer.MIN_VALUE, ar = i < m ? A[i] : Integer.MAX_VALUE;
            int bl = j > 0 ? B[j - 1] : Integer.MIN_VALUE, br = j < n ? B[j] : Integer.MAX_VALUE;
            if (al <= br && bl <= ar) {
                if ((m + n) % 2 == 1) return Math.max(al, bl);
                return (Math.max(al, bl) + (double) Math.min(ar, br)) / 2;
            }
            if (al > br) hi = i - 1; else lo = i + 1;
        }
    }
}
// test
check(new Solution().findMedianSortedArrays(new int[]{1, 3}, new int[]{2}), 2.0);
check(new Solution().findMedianSortedArrays(new int[]{1, 2}, new int[]{3, 4}), 2.5);
```

## 1011. Capacity To Ship Packages Within D Days | Medium
related: 875, 410
tags: search-on-answer

Capacity in `[max(w), sum(w)]`. `days(cap)` (greedy fill) decreases with cap → smallest cap with `days ≤ D`.

**Complexity:** O(n log sum).

```python
class Solution:
    def shipWithinDays(self, weights: List[int], days: int) -> int:
        def need(cap):
            d, load = 1, 0
            for w in weights:
                if load + w > cap:
                    d, load = d + 1, 0
                load += w
            return d

        lo, hi = max(weights), sum(weights)
        while lo < hi:
            mid = (lo + hi) // 2
            if need(mid) <= days:
                hi = mid
            else:
                lo = mid + 1
        return lo
# test
assert Solution().shipWithinDays([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 5) == 15
assert Solution().shipWithinDays([3, 2, 2, 4, 1, 4], 3) == 6
```

```java
class Solution {
    public int shipWithinDays(int[] weights, int days) {
        int lo = 0, hi = 0;
        for (int w : weights) { lo = Math.max(lo, w); hi += w; }
        while (lo < hi) {
            int mid = lo + (hi - lo) / 2;
            if (need(weights, mid) <= days) hi = mid; else lo = mid + 1;
        }
        return lo;
    }

    private int need(int[] weights, int cap) {
        int d = 1, load = 0;
        for (int w : weights) {
            if (load + w > cap) { d++; load = 0; }
            load += w;
        }
        return d;
    }
}
// test
check(new Solution().shipWithinDays(new int[]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5), 15);
```

## 410. Split Array Largest Sum | Hard
related: 1011, 875
tags: search-on-answer

Same as shipping packages: minimize the largest piece sum when splitting into `k` parts. Binary search the answer; greedily count parts needed.

**Complexity:** O(n log sum).

```python
class Solution:
    def splitArray(self, nums: List[int], k: int) -> int:
        def parts(cap):
            cnt, s = 1, 0
            for x in nums:
                if s + x > cap:
                    cnt, s = cnt + 1, 0
                s += x
            return cnt

        lo, hi = max(nums), sum(nums)
        while lo < hi:
            mid = (lo + hi) // 2
            if parts(mid) <= k:
                hi = mid
            else:
                lo = mid + 1
        return lo
# test
assert Solution().splitArray([7, 2, 5, 10, 8], 2) == 18
assert Solution().splitArray([1, 2, 3, 4, 5], 2) == 9
```

```java
class Solution {
    public int splitArray(int[] nums, int k) {
        long lo = 0, hi = 0;
        for (int x : nums) { lo = Math.max(lo, x); hi += x; }
        while (lo < hi) {
            long mid = (lo + hi) / 2;
            int cnt = 1;
            long s = 0;
            for (int x : nums) {
                if (s + x > mid) { cnt++; s = 0; }
                s += x;
            }
            if (cnt <= k) hi = mid; else lo = mid + 1;
        }
        return (int) lo;
    }
}
// test
check(new Solution().splitArray(new int[]{7, 2, 5, 10, 8}, 2), 18);
```

## 69. Sqrt(x) | Easy | sqrtx
related: 367, 50
tags: search-on-answer

Largest `m` with `m*m ≤ x`. Search for the first `m` with `m*m > x`, answer `m − 1`.

**Complexity:** O(log x).

```python
class Solution:
    def mySqrt(self, x: int) -> int:
        lo, hi = 0, x + 1          # first m with m*m > x
        while lo < hi:
            m = (lo + hi) // 2
            if m * m > x:
                hi = m
            else:
                lo = m + 1
        return lo - 1
# test
assert Solution().mySqrt(4) == 2
assert Solution().mySqrt(8) == 2
assert Solution().mySqrt(0) == 0
```

```java
class Solution {
    public int mySqrt(int x) {
        long lo = 0, hi = (long) x + 1;
        while (lo < hi) {
            long m = (lo + hi) / 2;
            if (m * m > x) hi = m; else lo = m + 1;
        }
        return (int) (lo - 1);
    }
}
// test
check(new Solution().mySqrt(8), 2);
check(new Solution().mySqrt(2147395599), 46339);
```
