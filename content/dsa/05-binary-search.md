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
tags: 

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


## 278. First Bad Version | Easy
related: 35, 704
tags: lower-bound

`isBadVersion` is monotonic (F…F T…T) → find the first `true` with the lower-bound template, using as few API calls as possible.

**Complexity:** O(log n) API calls.

```python
class Solution:
    def firstBadVersion(self, n: int) -> int:
        lo, hi = 1, n
        while lo < hi:
            mid = (lo + hi) // 2
            if isBadVersion(mid):      # provided by LeetCode
                hi = mid
            else:
                lo = mid + 1
        return lo
# test
BAD = 4
def isBadVersion(v): return v >= BAD
assert Solution().firstBadVersion(5) == 4
BAD = 1
assert Solution().firstBadVersion(1) == 1
```

```java
class VersionControl { // provided by LeetCode
    static int bad;
    boolean isBadVersion(int v) { return v >= bad; }
}

class Solution extends VersionControl {
    public int firstBadVersion(int n) {
        int lo = 1, hi = n;
        while (lo < hi) {
            int mid = lo + (hi - lo) / 2;      // avoids int overflow for n near 2^31
            if (isBadVersion(mid)) hi = mid; else lo = mid + 1;
        }
        return lo;
    }
}
// test
VersionControl.bad = 1702766719;
check(new Solution().firstBadVersion(2126753390), 1702766719);
```

## 852. Peak Index in a Mountain Array | Medium
related: 162, 1095
tags: slope

Find the first index where the array starts descending: `arr[mid] > arr[mid + 1]`.

**Complexity:** O(log n).

```python
class Solution:
    def peakIndexInMountainArray(self, arr: List[int]) -> int:
        lo, hi = 0, len(arr) - 1
        while lo < hi:
            mid = (lo + hi) // 2
            if arr[mid] > arr[mid + 1]:
                hi = mid
            else:
                lo = mid + 1
        return lo
# test
assert Solution().peakIndexInMountainArray([0, 1, 0]) == 1
assert Solution().peakIndexInMountainArray([0, 10, 5, 2]) == 1
assert Solution().peakIndexInMountainArray([3, 5, 3, 2, 0]) == 1
```

```java
class Solution {
    public int peakIndexInMountainArray(int[] arr) {
        int lo = 0, hi = arr.length - 1;
        while (lo < hi) {
            int mid = (lo + hi) / 2;
            if (arr[mid] > arr[mid + 1]) hi = mid; else lo = mid + 1;
        }
        return lo;
    }
}
// test
check(new Solution().peakIndexInMountainArray(new int[]{0, 10, 5, 2}), 1);
```

## 81. Search in Rotated Sorted Array II | Medium
related: 33, 154
tags: rotated, duplicates

Like 33, but duplicates can hide which half is sorted: when `nums[lo] == nums[mid] == nums[hi]`, shrink both ends by one. Worst case degrades to O(n).

**Complexity:** O(log n) average, O(n) worst.

```python
class Solution:
    def search(self, nums: List[int], target: int) -> bool:
        lo, hi = 0, len(nums) - 1
        while lo <= hi:
            mid = (lo + hi) // 2
            if nums[mid] == target:
                return True
            if nums[lo] == nums[mid] == nums[hi]:
                lo, hi = lo + 1, hi - 1
            elif nums[lo] <= nums[mid]:
                if nums[lo] <= target < nums[mid]:
                    hi = mid - 1
                else:
                    lo = mid + 1
            else:
                if nums[mid] < target <= nums[hi]:
                    lo = mid + 1
                else:
                    hi = mid - 1
        return False
# test
assert Solution().search([2, 5, 6, 0, 0, 1, 2], 0) is True
assert Solution().search([2, 5, 6, 0, 0, 1, 2], 3) is False
assert Solution().search([1, 0, 1, 1, 1], 0) is True
```

```java
class Solution {
    public boolean search(int[] nums, int target) {
        int lo = 0, hi = nums.length - 1;
        while (lo <= hi) {
            int mid = (lo + hi) / 2;
            if (nums[mid] == target) return true;
            if (nums[lo] == nums[mid] && nums[mid] == nums[hi]) { lo++; hi--; }
            else if (nums[lo] <= nums[mid]) {
                if (nums[lo] <= target && target < nums[mid]) hi = mid - 1; else lo = mid + 1;
            } else {
                if (nums[mid] < target && target <= nums[hi]) lo = mid + 1; else hi = mid - 1;
            }
        }
        return false;
    }
}
// test
check(new Solution().search(new int[]{1, 0, 1, 1, 1}, 0), true);
check(new Solution().search(new int[]{2, 5, 6, 0, 0, 1, 2}, 3), false);
```

## 240. Search a 2D Matrix II | Medium
related: 74, 378
tags: staircase

Rows and columns sorted (but rows don't continue each other). Start at the **top-right**: if bigger than target move left (the whole column below is bigger too), if smaller move down.

```viz
 1  4  7 [11] 15      target 5: 11 > 5 <-, 7 > 5 <-, 4 < 5 v, 5 found
 2  5  8  12  19
 3  6  9  16  22
```

**Complexity:** O(m + n).

```python
class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        r, c = 0, len(matrix[0]) - 1
        while r < len(matrix) and c >= 0:
            v = matrix[r][c]
            if v == target:
                return True
            if v > target:
                c -= 1
            else:
                r += 1
        return False
# test
M = [[1, 4, 7, 11, 15], [2, 5, 8, 12, 19], [3, 6, 9, 16, 22], [10, 13, 14, 17, 24], [18, 21, 23, 26, 30]]
assert Solution().searchMatrix(M, 5) is True
assert Solution().searchMatrix(M, 20) is False
```

```java
class Solution {
    public boolean searchMatrix(int[][] matrix, int target) {
        int r = 0, c = matrix[0].length - 1;
        while (r < matrix.length && c >= 0) {
            if (matrix[r][c] == target) return true;
            if (matrix[r][c] > target) c--; else r++;
        }
        return false;
    }
}
// test
int[][] M = {{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}};
check(new Solution().searchMatrix(M, 5), true);
check(new Solution().searchMatrix(M, 20), false);
```

## 528. Random Pick with Weight | Medium
related: 384, 380
tags: prefix-sum, upper-bound

Prefix sums turn weights into intervals on `[0, total)`. Draw `r` uniformly, binary-search the first prefix sum **greater than** `r`.

```viz
w = [1, 3]   prefix = [1, 4]
r in [0,1) -> index 0 (25%)     r in [1,4) -> index 1 (75%)
```

**Complexity:** O(n) init, O(log n) per pick.

```python
class Solution:
    def __init__(self, w: List[int]):
        self.pre = list(itertools.accumulate(w))

    def pickIndex(self) -> int:
        r = random.random() * self.pre[-1]
        return bisect.bisect_right(self.pre, r)
# test
random.seed(1)
s = Solution([1, 3])
picks = [s.pickIndex() for _ in range(4000)]
assert set(picks) == {0, 1} and 0.70 < picks.count(1) / 4000 < 0.80
assert Solution([5]).pickIndex() == 0
```

```java
class Solution {
    private final int[] pre;
    private final Random rnd = new Random(1);

    public Solution(int[] w) {
        pre = new int[w.length];
        for (int i = 0, s = 0; i < w.length; i++) pre[i] = s += w[i];
    }

    public int pickIndex() {
        int r = rnd.nextInt(pre[pre.length - 1]);          // r in [0, total)
        int lo = 0, hi = pre.length - 1;
        while (lo < hi) {                                   // first prefix > r
            int mid = (lo + hi) / 2;
            if (pre[mid] > r) hi = mid; else lo = mid + 1;
        }
        return lo;
    }
}
// test
Solution s = new Solution(new int[]{1, 3});
int ones = 0;
for (int i = 0; i < 4000; i++) ones += s.pickIndex();
check(ones > 2800 && ones < 3200, true);
```

## 1539. Kth Missing Positive Number | Easy
related: 41, 268
tags: binary-search

Missing count before index `i` is `arr[i] − (i + 1)` (monotonic). Find the first index where it is ≥ k; the answer is `lo + k`.

**Complexity:** O(log n).

```python
class Solution:
    def findKthPositive(self, arr: List[int], k: int) -> int:
        lo, hi = 0, len(arr)
        while lo < hi:
            mid = (lo + hi) // 2
            if arr[mid] - (mid + 1) >= k:
                hi = mid
            else:
                lo = mid + 1
        return lo + k
# test
assert Solution().findKthPositive([2, 3, 4, 7, 11], 5) == 9
assert Solution().findKthPositive([1, 2, 3, 4], 2) == 6
```

```java
class Solution {
    public int findKthPositive(int[] arr, int k) {
        int lo = 0, hi = arr.length;
        while (lo < hi) {
            int mid = (lo + hi) / 2;
            if (arr[mid] - (mid + 1) >= k) hi = mid; else lo = mid + 1;
        }
        return lo + k;
    }
}
// test
check(new Solution().findKthPositive(new int[]{2, 3, 4, 7, 11}, 5), 9);
```

## 1482. Minimum Number of Days to Make m Bouquets | Medium
related: 875, 1011
tags: search-on-answer

`canMake(day)` (count runs of `k` adjacent bloomed flowers ≥ m) is monotonic in `day` → binary search the day between `min` and `max` bloom day.

**Complexity:** O(n log max).

```python
class Solution:
    def minDays(self, bloomDay: List[int], m: int, k: int) -> int:
        if m * k > len(bloomDay):
            return -1

        def ok(day):
            made = run = 0
            for b in bloomDay:
                run = run + 1 if b <= day else 0
                if run == k:
                    made, run = made + 1, 0
            return made >= m

        lo, hi = min(bloomDay), max(bloomDay)
        while lo < hi:
            mid = (lo + hi) // 2
            if ok(mid):
                hi = mid
            else:
                lo = mid + 1
        return lo
# test
assert Solution().minDays([1, 10, 3, 10, 2], 3, 1) == 3
assert Solution().minDays([1, 10, 3, 10, 2], 3, 2) == -1
assert Solution().minDays([7, 7, 7, 7, 12, 7, 7], 2, 3) == 12
```

```java
class Solution {
    public int minDays(int[] bloomDay, int m, int k) {
        if ((long) m * k > bloomDay.length) return -1;
        int lo = Integer.MAX_VALUE, hi = 0;
        for (int b : bloomDay) { lo = Math.min(lo, b); hi = Math.max(hi, b); }
        while (lo < hi) {
            int mid = lo + (hi - lo) / 2;
            int made = 0, run = 0;
            for (int b : bloomDay) {
                run = b <= mid ? run + 1 : 0;
                if (run == k) { made++; run = 0; }
            }
            if (made >= m) hi = mid; else lo = mid + 1;
        }
        return lo;
    }
}
// test
check(new Solution().minDays(new int[]{7, 7, 7, 7, 12, 7, 7}, 2, 3), 12);
check(new Solution().minDays(new int[]{1, 10, 3, 10, 2}, 3, 2), -1);
```

## 374. Guess Number Higher or Lower | Easy
related: 278, 704
tags: binary-search

Classic binary search where the comparison is the provided `guess(num)` API: `-1` means the pick is lower than `num`, `1` higher, `0` found.

**Complexity:** O(log n) calls.

```python
class Solution:
    def guessNumber(self, n: int) -> int:
        lo, hi = 1, n
        while True:
            mid = (lo + hi) // 2
            g = guess(mid)          # provided by LeetCode
            if g == 0:
                return mid
            if g < 0:
                hi = mid - 1
            else:
                lo = mid + 1
# test
for n, PICK in ((10, 6), (1, 1), (2, 1), (2, 2), (2126753390, 1702766719)):
    def guess(num, pick=PICK):
        return 0 if num == pick else (-1 if pick < num else 1)
    assert Solution().guessNumber(n) == PICK
```

```java
class GuessGame { // provided by LeetCode
    static int pick;
    int guess(int num) { return Integer.compare(pick, num); }
}

class Solution extends GuessGame {
    public int guessNumber(int n) {
        int lo = 1, hi = n;
        while (true) {
            int mid = lo + (hi - lo) / 2;          // avoids int overflow
            int g = guess(mid);
            if (g == 0) return mid;
            if (g < 0) hi = mid - 1; else lo = mid + 1;
        }
    }
}
// test
GuessGame.pick = 1702766719;
check(new Solution().guessNumber(2126753390), 1702766719);
GuessGame.pick = 6;
check(new Solution().guessNumber(10), 6);
```

## 2300. Successful Pairs of Spells and Potions | Medium
related: 35, 1855
tags: lower-bound, sort

Sort the potions once. A pair succeeds when `potion ≥ ⌈success / spell⌉`, so the count for each spell is `m − lower_bound(⌈success / spell⌉)`. Integer ceiling avoids floating-point trouble.

**Complexity:** O((n + m) log m).

```python
class Solution:
    def successfulPairs(self, spells: List[int], potions: List[int], success: int) -> List[int]:
        potions = sorted(potions)
        m = len(potions)
        return [m - bisect.bisect_left(potions, (success + s - 1) // s) for s in spells]
# test
assert Solution().successfulPairs([5, 1, 3], [1, 2, 3, 4, 5], 7) == [4, 0, 3]
assert Solution().successfulPairs([3, 1, 2], [8, 5, 8], 16) == [2, 0, 2]
assert Solution().successfulPairs([100000], [100000], 10000000000) == [1]
```

```java
class Solution {
    public int[] successfulPairs(int[] spells, int[] potions, long success) {
        Arrays.sort(potions);
        int m = potions.length;
        int[] res = new int[spells.length];
        for (int i = 0; i < spells.length; i++) {
            long need = (success + spells[i] - 1) / spells[i];
            int lo = 0, hi = m;                       // first potion >= need
            while (lo < hi) {
                int mid = (lo + hi) / 2;
                if (potions[mid] >= need) hi = mid; else lo = mid + 1;
            }
            res[i] = m - lo;
        }
        return res;
    }
}
// test
check(new Solution().successfulPairs(new int[]{5, 1, 3}, new int[]{1, 2, 3, 4, 5}, 7), new int[]{4, 0, 3});
check(new Solution().successfulPairs(new int[]{100000}, new int[]{100000}, 10000000000L), new int[]{1});
```

## 1095. Find in Mountain Array | Hard
related: 852, 33, 162
tags: binary-search

Only ~100 `get` calls are allowed, so three binary searches: (1) find the **peak** (first `i` with `a[i] > a[i+1]`), (2) search the **ascending** left part, (3) if not found, search the **descending** right part (comparison flipped). Checking the left side first returns the minimum index.

**Complexity:** O(log n) calls.

```python
class Solution:
    def findInMountainArray(self, target: int, mountainArr: 'MountainArray') -> int:
        n = mountainArr.length()
        lo, hi = 0, n - 1
        while lo < hi:
            mid = (lo + hi) // 2
            if mountainArr.get(mid) < mountainArr.get(mid + 1):
                lo = mid + 1
            else:
                hi = mid
        peak = lo

        def search(lo, hi, ascending):
            while lo <= hi:
                mid = (lo + hi) // 2
                v = mountainArr.get(mid)
                if v == target:
                    return mid
                if (v < target) == ascending:
                    lo = mid + 1
                else:
                    hi = mid - 1
            return -1

        i = search(0, peak, True)
        return i if i != -1 else search(peak + 1, n - 1, False)
# test
class MountainArray:  # LeetCode's interface (counts calls; the limit is 100)
    def __init__(self, a): self.a, self.calls = a, 0
    def get(self, i): self.calls += 1; return self.a[i]
    def length(self): return len(self.a)
m = MountainArray([1, 2, 3, 4, 5, 3, 1])
assert Solution().findInMountainArray(3, m) == 2 and m.calls <= 100
assert Solution().findInMountainArray(3, MountainArray([0, 1, 2, 4, 2, 1])) == -1
assert Solution().findInMountainArray(1, MountainArray([0, 1, 2, 4, 2, 1])) == 1
big = MountainArray(list(range(5000)) + list(range(4998, -1, -1)))
assert Solution().findInMountainArray(4000, big) == 4000 and big.calls <= 100
big2 = MountainArray(list(range(5000)) + list(range(4998, 2000, -1)))
assert Solution().findInMountainArray(2001, big2) == 2001
assert Solution().findInMountainArray(2000, MountainArray(list(range(10, 3000)) + [5])) == 1990
```

```java
interface MountainArray { // provided by LeetCode
    int get(int index);
    int length();
}

class Solution {
    public int findInMountainArray(int target, MountainArray m) {
        int n = m.length(), lo = 0, hi = n - 1;
        while (lo < hi) {
            int mid = (lo + hi) / 2;
            if (m.get(mid) < m.get(mid + 1)) lo = mid + 1; else hi = mid;
        }
        int peak = lo, i = search(m, target, 0, peak, true);
        return i != -1 ? i : search(m, target, peak + 1, n - 1, false);
    }

    private int search(MountainArray m, int target, int lo, int hi, boolean asc) {
        while (lo <= hi) {
            int mid = (lo + hi) / 2, v = m.get(mid);
            if (v == target) return mid;
            if ((v < target) == asc) lo = mid + 1; else hi = mid - 1;
        }
        return -1;
    }
}
// test
int[] arr = {1, 2, 3, 4, 5, 3, 1};
MountainArray ma = new MountainArray() {
    public int get(int i) { return arr[i]; }
    public int length() { return arr.length; }
};
check(new Solution().findInMountainArray(3, ma), 2);
check(new Solution().findInMountainArray(6, ma), -1);
```
