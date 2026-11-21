# Greedy

**When to use:** a locally optimal choice can be proven to never hurt (exchange argument), usually after **sorting** or while tracking a running best/reach.

Typical shapes
- **Kadane**: best subarray ending here = `max(x, best_here + x)`.
- **Reach**: track the farthest index reachable so far (jump games).
- **Sort then sweep**: by end time, by height, by deadline.
- **Two passes** (left→right, right→left) when each element has constraints from both sides (Candy).

If you can't argue why greedy is safe, fall back to DP.

## 53. Maximum Subarray | Medium
related: 152, 918, 121
tags: blind75, kadane

```problem
Return the largest sum of any non-empty contiguous subarray of `nums`.

**Example:** `[-2,1,-3,4,-1,2,1,-5,4]` → `6` (`[4,-1,2,1]`); `[-1]` → `-1`

**Constraints:** n ≤ 10⁵. Follow-up: a divide-and-conquer solution.
```

**Kadane**: the best subarray ending at `i` either extends the previous one or starts fresh at `i`.

```viz
nums  -2  1  -3  4  -1  2  1  -5  4
cur   -2  1  -2  4   3  5  6   1  5
best  -2  1   1  4   4  5  6   6  6
```

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        best = cur = nums[0]
        for x in nums[1:]:
            cur = max(x, cur + x)
            best = max(best, cur)
        return best
# test
assert Solution().maxSubArray([-2, 1, -3, 4, -1, 2, 1, -5, 4]) == 6
assert Solution().maxSubArray([-1]) == -1
```

```java
class Solution {
    public int maxSubArray(int[] nums) {
        int best = nums[0], cur = nums[0];
        for (int i = 1; i < nums.length; i++) {
            cur = Math.max(nums[i], cur + nums[i]);
            best = Math.max(best, cur);
        }
        return best;
    }
}
// test
check(new Solution().maxSubArray(new int[]{-2, 1, -3, 4, -1, 2, 1, -5, 4}), 6);
```

```go
func maxSubArray(nums []int) int {
	best, cur := nums[0], nums[0] // cur = best sum ending here (Kadane)
	for _, x := range nums[1:] {
		cur = max(x, cur+x)
		best = max(best, cur)
	}
	return best
}
// test
assertEq(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), 6)
assertEq(maxSubArray([]int{-1}), -1)
```


## 918. Maximum Sum Circular Subarray | Medium
related: 53
tags: kadane

```problem
`nums` is circular: its end wraps around to its start. Return the largest sum of a non-empty subarray that uses each element at most once.

**Example:** `[1,-2,3,-2]` → `3`; `[5,-3,5]` → `10` (5 + 5 across the wrap)

**Constraints:** n ≤ 3·10⁴.
```

A wrap-around subarray = total − (a minimum subarray in the middle). Answer = `max(maxKadane, total − minKadane)`, except when all numbers are negative (then the "wrap" would be empty) → `maxKadane`.

**Complexity:** O(n).

```python
class Solution:
    def maxSubarraySumCircular(self, nums: List[int]) -> int:
        cur_max = cur_min = 0
        best_max, best_min = nums[0], nums[0]
        for x in nums:
            cur_max = max(x, cur_max + x)
            best_max = max(best_max, cur_max)
            cur_min = min(x, cur_min + x)
            best_min = min(best_min, cur_min)
        if best_max < 0:
            return best_max
        return max(best_max, sum(nums) - best_min)
# test
assert Solution().maxSubarraySumCircular([1, -2, 3, -2]) == 3
assert Solution().maxSubarraySumCircular([5, -3, 5]) == 10
assert Solution().maxSubarraySumCircular([-3, -2, -3]) == -2
```

```java
class Solution {
    public int maxSubarraySumCircular(int[] nums) {
        int curMax = 0, curMin = 0, bestMax = nums[0], bestMin = nums[0], total = 0;
        for (int x : nums) {
            curMax = Math.max(x, curMax + x);
            bestMax = Math.max(bestMax, curMax);
            curMin = Math.min(x, curMin + x);
            bestMin = Math.min(bestMin, curMin);
            total += x;
        }
        return bestMax < 0 ? bestMax : Math.max(bestMax, total - bestMin);
    }
}
// test
check(new Solution().maxSubarraySumCircular(new int[]{5, -3, 5}), 10);
check(new Solution().maxSubarraySumCircular(new int[]{-3, -2, -3}), -2);
```

```go
func maxSubarraySumCircular(nums []int) int {
	curMax, curMin, total := 0, 0, 0
	bestMax, bestMin := nums[0], nums[0]
	for _, x := range nums {
		curMax = max(x, curMax+x)
		bestMax = max(bestMax, curMax)
		curMin = min(x, curMin+x)
		bestMin = min(bestMin, curMin)
		total += x
	}
	if bestMax < 0 {
		return bestMax // all negative: the wrap-around case would be empty
	}
	return max(bestMax, total-bestMin) // a wrapping subarray = total minus a middle one
}
// test
assertEq(maxSubarraySumCircular([]int{1, -2, 3, -2}), 3)
assertEq(maxSubarraySumCircular([]int{5, -3, 5}), 10)
assertEq(maxSubarraySumCircular([]int{-3, -2, -3}), -2)
```


## 55. Jump Game | Medium
related: 45
tags: blind75, reach

```problem
You start at index 0, and `nums[i]` is the maximum jump length from index i. Decide whether you can reach the last index.

**Example:** `[2,3,1,1,4]` → `true`; `[3,2,1,0,4]` → `false`

**Constraints:** n ≤ 10⁴.
```

Track the farthest reachable index; if you ever stand beyond it, you're stuck.

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def canJump(self, nums: List[int]) -> bool:
        reach = 0
        for i, x in enumerate(nums):
            if i > reach:
                return False
            reach = max(reach, i + x)
        return True
# test
assert Solution().canJump([2, 3, 1, 1, 4]) is True
assert Solution().canJump([3, 2, 1, 0, 4]) is False
```

```java
class Solution {
    public boolean canJump(int[] nums) {
        int reach = 0;
        for (int i = 0; i < nums.length; i++) {
            if (i > reach) return false;
            reach = Math.max(reach, i + nums[i]);
        }
        return true;
    }
}
// test
check(new Solution().canJump(new int[]{2, 3, 1, 1, 4}), true);
check(new Solution().canJump(new int[]{3, 2, 1, 0, 4}), false);
```

```go
func canJump(nums []int) bool {
	reach := 0 // farthest index reachable so far
	for i, x := range nums {
		if i > reach {
			return false
		}
		reach = max(reach, i+x)
	}
	return true
}
// test
assertEq(canJump([]int{2, 3, 1, 1, 4}), true)
assertEq(canJump([]int{3, 2, 1, 0, 4}), false)
```


## 45. Jump Game II | Medium
related: 55
tags: reach, implicit-bfs

```problem
You start at index 0, and `nums[i]` is the maximum jump length from index i. Return the minimum number of jumps to reach the last index (always possible).

**Example:** `[2,3,1,1,4]` → `2` (0 → 1 → 4)

**Constraints:** n ≤ 10⁴.
```

BFS by "levels": the current jump covers `[.., end]`; while scanning it, track the farthest next reach. When you hit `end`, you must jump: `jumps += 1`, `end = farthest`.

```viz
nums 2 3 1 1 4
     [0]           jump 1 covers 1..2
       [1  2]      farthest = 4  -> jump 2 reaches the end
```

**Complexity:** O(n).

```python
class Solution:
    def jump(self, nums: List[int]) -> int:
        jumps = end = far = 0
        for i in range(len(nums) - 1):
            far = max(far, i + nums[i])
            if i == end:
                jumps += 1
                end = far
        return jumps
# test
assert Solution().jump([2, 3, 1, 1, 4]) == 2
assert Solution().jump([2, 3, 0, 1, 4]) == 2
assert Solution().jump([0]) == 0
```

```java
class Solution {
    public int jump(int[] nums) {
        int jumps = 0, end = 0, far = 0;
        for (int i = 0; i < nums.length - 1; i++) {
            far = Math.max(far, i + nums[i]);
            if (i == end) { jumps++; end = far; }
        }
        return jumps;
    }
}
// test
check(new Solution().jump(new int[]{2, 3, 1, 1, 4}), 2);
```

```go
func jump(nums []int) int {
	jumps, end, far := 0, 0, 0 // BFS by levels: [.., end] is reachable with `jumps` jumps
	for i := 0; i < len(nums)-1; i++ {
		far = max(far, i+nums[i])
		if i == end {
			jumps++
			end = far
		}
	}
	return jumps
}
// test
assertEq(jump([]int{2, 3, 1, 1, 4}), 2)
assertEq(jump([]int{2, 3, 0, 1, 4}), 2)
assertEq(jump([]int{0}), 0)
```


## 122. Best Time to Buy and Sell Stock II | Medium
related: 121, 309, 714
tags: greedy

```problem
`prices[i]` is a stock's price on day i. Buy and sell as often as you like, holding at most one share at a time (you may sell and buy again on the same day). Return the maximum profit.

**Example:** `[7,1,5,3,6,4]` → `7` (buy 1 sell 5, buy 3 sell 6); `[1,2,3,4,5]` → `4`
```

Unlimited transactions → collect every positive day-to-day increase.

**Complexity:** O(n).

```python
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        return sum(max(0, b - a) for a, b in zip(prices, prices[1:]))
# test
assert Solution().maxProfit([7, 1, 5, 3, 6, 4]) == 7
assert Solution().maxProfit([1, 2, 3, 4, 5]) == 4
```

```java
class Solution {
    public int maxProfit(int[] prices) {
        int profit = 0;
        for (int i = 1; i < prices.length; i++) profit += Math.max(0, prices[i] - prices[i - 1]);
        return profit;
    }
}
// test
check(new Solution().maxProfit(new int[]{7, 1, 5, 3, 6, 4}), 7);
```

```go
func maxProfit(prices []int) int {
	total := 0
	for i := 1; i < len(prices); i++ {
		total += max(0, prices[i]-prices[i-1]) // collect every rise
	}
	return total
}
// test
assertEq(maxProfit([]int{7, 1, 5, 3, 6, 4}), 7)
assertEq(maxProfit([]int{1, 2, 3, 4, 5}), 4)
```


## 134. Gas Station | Medium
related: 55
tags: greedy

```problem
Gas stations stand on a circular route; station i has `gas[i]` fuel, and driving from it to station i + 1 costs `cost[i]`. Starting with an empty tank, return the station from which you can drive around the whole circuit once, or −1 if none works. The answer is unique when it exists.

**Example:** `gas = [1,2,3,4,5], cost = [3,4,5,1,2]` → `3`

**Constraints:** n ≤ 10⁵.
```

If total gas < total cost → impossible. Otherwise, scan with a running tank; whenever it drops below 0, no station in the current segment can be the start → restart from `i + 1`.

**Complexity:** O(n).

```python
class Solution:
    def canCompleteCircuit(self, gas: List[int], cost: List[int]) -> int:
        if sum(gas) < sum(cost):
            return -1
        start = tank = 0
        for i, (g, c) in enumerate(zip(gas, cost)):
            tank += g - c
            if tank < 0:
                start, tank = i + 1, 0
        return start
# test
assert Solution().canCompleteCircuit([1, 2, 3, 4, 5], [3, 4, 5, 1, 2]) == 3
assert Solution().canCompleteCircuit([2, 3, 4], [3, 4, 3]) == -1
```

```java
class Solution {
    public int canCompleteCircuit(int[] gas, int[] cost) {
        int total = 0, tank = 0, start = 0;
        for (int i = 0; i < gas.length; i++) {
            total += gas[i] - cost[i];
            tank += gas[i] - cost[i];
            if (tank < 0) { start = i + 1; tank = 0; }
        }
        return total < 0 ? -1 : start;
    }
}
// test
check(new Solution().canCompleteCircuit(new int[]{1, 2, 3, 4, 5}, new int[]{3, 4, 5, 1, 2}), 3);
```

```go
func canCompleteCircuit(gas []int, cost []int) int {
	total, tank, start := 0, 0, 0
	for i := range gas {
		total += gas[i] - cost[i]
		tank += gas[i] - cost[i]
		if tank < 0 { // cannot reach i+1 from start, nor from any station in between
			start, tank = i+1, 0
		}
	}
	if total < 0 {
		return -1
	}
	return start
}
// test
assertEq(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}), 3)
assertEq(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}), -1)
```


## 135. Candy | Hard
related: 42
tags: two-pass

```problem
Children stand in a line with `ratings`. Give each child at least one candy, and a child with a higher rating than a neighbour must get more candies than that neighbour. Return the minimum total candies.

**Example:** `[1,0,2]` → `5` (2, 1, 2); `[1,2,2]` → `4` (1, 2, 1)

**Constraints:** n ≤ 2·10⁴.
```

Left→right pass satisfies "higher than left neighbour"; right→left pass satisfies "higher than right neighbour" while keeping the max.

```viz
ratings  1 0 2
L->R     1 1 2
R->L     2 1 2     total 5
```

**Complexity:** O(n).

```python
class Solution:
    def candy(self, ratings: List[int]) -> int:
        n = len(ratings)
        c = [1] * n
        for i in range(1, n):
            if ratings[i] > ratings[i - 1]:
                c[i] = c[i - 1] + 1
        for i in range(n - 2, -1, -1):
            if ratings[i] > ratings[i + 1]:
                c[i] = max(c[i], c[i + 1] + 1)
        return sum(c)
# test
assert Solution().candy([1, 0, 2]) == 5
assert Solution().candy([1, 2, 2]) == 4
```

```java
class Solution {
    public int candy(int[] ratings) {
        int n = ratings.length, total = 0;
        int[] c = new int[n];
        Arrays.fill(c, 1);
        for (int i = 1; i < n; i++) if (ratings[i] > ratings[i - 1]) c[i] = c[i - 1] + 1;
        for (int i = n - 2; i >= 0; i--) if (ratings[i] > ratings[i + 1]) c[i] = Math.max(c[i], c[i + 1] + 1);
        for (int x : c) total += x;
        return total;
    }
}
// test
check(new Solution().candy(new int[]{1, 0, 2}), 5);
check(new Solution().candy(new int[]{1, 2, 2}), 4);
```

```go
func candy(ratings []int) int {
	n := len(ratings)
	c := make([]int, n)
	for i := range c {
		c[i] = 1
	}
	for i := 1; i < n; i++ { // left neighbour rule
		if ratings[i] > ratings[i-1] {
			c[i] = c[i-1] + 1
		}
	}
	total := c[n-1]
	for i := n - 2; i >= 0; i-- { // right neighbour rule
		if ratings[i] > ratings[i+1] {
			c[i] = max(c[i], c[i+1]+1)
		}
		total += c[i]
	}
	return total
}
// test
assertEq(candy([]int{1, 0, 2}), 5)
assertEq(candy([]int{1, 2, 2}), 4)
```


## 846. Hand of Straights | Medium
related: 659
tags: sorted-counter

```problem
Decide whether the cards in `hand` can be split into groups of `groupSize` cards whose values are consecutive (like 2, 3, 4).

**Example:** `hand = [1,2,3,6,2,3,4,7,8], groupSize = 3` → `true` (`[1,2,3]`, `[2,3,4]`, `[6,7,8]`)

**Constraints:** up to 10⁴ cards.
```

Always start a group at the **smallest** remaining card (it can't belong anywhere else). Consume `x, x+1, …, x+k−1` from the counter.

**Complexity:** O(n log n).

```python
class Solution:
    def isNStraightHand(self, hand: List[int], groupSize: int) -> bool:
        if len(hand) % groupSize:
            return False
        cnt = Counter(hand)
        for x in sorted(cnt):
            c = cnt[x]
            if c == 0:
                continue
            for y in range(x, x + groupSize):
                if cnt[y] < c:
                    return False
                cnt[y] -= c
        return True
# test
assert Solution().isNStraightHand([1, 2, 3, 6, 2, 3, 4, 7, 8], 3) is True
assert Solution().isNStraightHand([1, 2, 3, 4, 5], 4) is False
```

```java
class Solution {
    public boolean isNStraightHand(int[] hand, int groupSize) {
        if (hand.length % groupSize != 0) return false;
        TreeMap<Integer, Integer> cnt = new TreeMap<>();
        for (int x : hand) cnt.merge(x, 1, Integer::sum);
        for (int x : new ArrayList<>(cnt.keySet())) {
            int c = cnt.get(x);
            if (c == 0) continue;
            for (int y = x; y < x + groupSize; y++) {
                int have = cnt.getOrDefault(y, 0);
                if (have < c) return false;
                cnt.put(y, have - c);
            }
        }
        return true;
    }
}
// test
check(new Solution().isNStraightHand(new int[]{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3), true);
check(new Solution().isNStraightHand(new int[]{1, 2, 3, 4, 5}, 4), false);
```

```go
func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	cnt := map[int]int{}
	for _, x := range hand {
		cnt[x]++
	}
	keys := make([]int, 0, len(cnt))
	for x := range cnt {
		keys = append(keys, x)
	}
	sort.Ints(keys)
	for _, x := range keys { // the smallest card left must start a group
		c := cnt[x]
		if c == 0 {
			continue
		}
		for y := x; y < x+groupSize; y++ {
			if cnt[y] < c {
				return false
			}
			cnt[y] -= c
		}
	}
	return true
}
// test
assertEq(isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3), true)
assertEq(isNStraightHand([]int{1, 2, 3, 4, 5}, 4), false)
```


## 1899. Merge Triplets to Form Target Triplet | Medium
related: 55
tags: greedy

```problem
You may repeatedly pick two triplets and replace the second with their element-wise maximum. Decide whether `target` can end up as one of the `triplets`.

**Example:** `triplets = [[2,5,3],[1,8,4],[1,7,5]], target = [2,7,5]` → `true`

**Constraints:** up to 10⁵ triplets.
```

Ignore any triplet with a value exceeding the target in some coordinate (it would poison the max). Among the rest, check that each target coordinate is matched by at least one.

**Complexity:** O(n).

```python
class Solution:
    def mergeTriplets(self, triplets: List[List[int]], target: List[int]) -> bool:
        good = set()
        for t in triplets:
            if all(t[i] <= target[i] for i in range(3)):
                good |= {i for i in range(3) if t[i] == target[i]}
        return len(good) == 3
# test
assert Solution().mergeTriplets([[2, 5, 3], [1, 8, 4], [1, 7, 5]], [2, 7, 5]) is True
assert Solution().mergeTriplets([[3, 4, 5], [4, 5, 6]], [3, 2, 5]) is False
```

```java
class Solution {
    public boolean mergeTriplets(int[][] triplets, int[] target) {
        boolean[] good = new boolean[3];
        for (int[] t : triplets) {
            if (t[0] > target[0] || t[1] > target[1] || t[2] > target[2]) continue;
            for (int i = 0; i < 3; i++) if (t[i] == target[i]) good[i] = true;
        }
        return good[0] && good[1] && good[2];
    }
}
// test
check(new Solution().mergeTriplets(new int[][]{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}, new int[]{2, 7, 5}), true);
```

```go
func mergeTriplets(triplets [][]int, target []int) bool {
	good := [3]bool{}
	for _, t := range triplets {
		if t[0] > target[0] || t[1] > target[1] || t[2] > target[2] {
			continue // merging it would overshoot
		}
		for i := 0; i < 3; i++ {
			if t[i] == target[i] {
				good[i] = true
			}
		}
	}
	return good[0] && good[1] && good[2]
}
// test
assertEq(mergeTriplets([][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}, []int{2, 7, 5}), true)
assertEq(mergeTriplets([][]int{{3, 4, 5}, {4, 5, 6}}, []int{3, 2, 5}), false)
```


## 763. Partition Labels | Medium
related: 56
tags: last-occurrence

```problem
Split `s` into as many parts as possible so that each letter appears in at most one part. Return the part sizes in order.

**Example:** `"ababcbacadefegdehijhklij"` → `[9,7,8]`
```

Record each letter's last index. Extend the current part's end to the max last-index of letters seen; when `i == end`, cut.

```viz
a b a b c b a c a | d e f e g d e | h i j h k l i j
last(a)=8 -> part ends at 8 once all letters inside end by 8
```

**Complexity:** O(n).

```python
class Solution:
    def partitionLabels(self, s: str) -> List[int]:
        last = {ch: i for i, ch in enumerate(s)}
        res, start, end = [], 0, 0
        for i, ch in enumerate(s):
            end = max(end, last[ch])
            if i == end:
                res.append(end - start + 1)
                start = i + 1
        return res
# test
assert Solution().partitionLabels("ababcbacadefegdehijhklij") == [9, 7, 8]
assert Solution().partitionLabels("eccbbbbdec") == [10]
```

```java
class Solution {
    public List<Integer> partitionLabels(String s) {
        int[] last = new int[26];
        for (int i = 0; i < s.length(); i++) last[s.charAt(i) - 'a'] = i;
        List<Integer> res = new ArrayList<>();
        int start = 0, end = 0;
        for (int i = 0; i < s.length(); i++) {
            end = Math.max(end, last[s.charAt(i) - 'a']);
            if (i == end) { res.add(end - start + 1); start = i + 1; }
        }
        return res;
    }
}
// test
check(new Solution().partitionLabels("ababcbacadefegdehijhklij"), "[9, 7, 8]");
```

```go
func partitionLabels(s string) []int {
	last := [26]int{}
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}
	res := []int{}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		end = max(end, last[s[i]-'a']) // the part must reach every letter's last occurrence
		if i == end {
			res = append(res, end-start+1)
			start = i + 1
		}
	}
	return res
}
// test
assertEq(partitionLabels("ababcbacadefegdehijhklij"), []int{9, 7, 8})
assertEq(partitionLabels("eccbbbbdec"), []int{10})
```


## 678. Valid Parenthesis String | Medium
related: 20, 32
tags: range-tracking

```problem
`s` contains `(`, `)` and `*`, where each `*` can act as `(`, as `)` or as nothing. Decide whether `s` can be a valid parentheses string.

**Example:** `"(*)"` → `true`; `"(*))"` → `true`; `")("` → `false`

**Constraints:** length ≤ 100.
```

Track the **range** `[lo, hi]` of possible open counts. `(` → both +1; `)` → both −1; `*` → `lo−1, hi+1`. If `hi < 0` fail; clamp `lo ≥ 0`. Valid iff `lo == 0` at the end.

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def checkValidString(self, s: str) -> bool:
        lo = hi = 0
        for ch in s:
            if ch == "(":
                lo, hi = lo + 1, hi + 1
            elif ch == ")":
                lo, hi = lo - 1, hi - 1
            else:
                lo, hi = lo - 1, hi + 1
            if hi < 0:
                return False
            lo = max(lo, 0)
        return lo == 0
# test
assert Solution().checkValidString("()") is True
assert Solution().checkValidString("(*)") is True
assert Solution().checkValidString("(*))") is True
assert Solution().checkValidString("(((*)") is False
```

```java
class Solution {
    public boolean checkValidString(String s) {
        int lo = 0, hi = 0;
        for (char c : s.toCharArray()) {
            if (c == '(') { lo++; hi++; }
            else if (c == ')') { lo--; hi--; }
            else { lo--; hi++; }
            if (hi < 0) return false;
            lo = Math.max(lo, 0);
        }
        return lo == 0;
    }
}
// test
check(new Solution().checkValidString("(*))"), true);
check(new Solution().checkValidString("(((*)"), false);
```

```go
func checkValidString(s string) bool {
	lo, hi := 0, 0 // range of possible open-bracket counts
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			lo, hi = lo+1, hi+1
		case ')':
			lo, hi = lo-1, hi-1
		default: // '*' is ')', '' or '('
			lo, hi = lo-1, hi+1
		}
		if hi < 0 {
			return false
		}
		lo = max(lo, 0)
	}
	return lo == 0
}
// test
assertEq(checkValidString("()"), true)
assertEq(checkValidString("(*)"), true)
assertEq(checkValidString("(*))"), true)
assertEq(checkValidString("(((*)"), false)
```


## 406. Queue Reconstruction by Height | Medium
related: 315
tags: sort-then-insert

```problem
People in a queue were shuffled; `people[i] = [h, k]` says person i has height `h` and exactly `k` people in front of them are at least as tall. Reconstruct and return the queue.

**Example:** `[[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]` → `[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]`

**Constraints:** up to 2000 people.
```

Sort by height **descending**, then `k` ascending. Insert each person at index `k`: everyone already placed is taller or equal, so exactly `k` of them are in front, and shorter people inserted later don't affect it.

```viz
sorted: [7,0] [7,1] [6,1] [5,0] [5,2] [4,4]
insert: [7,0]
        [7,0] [7,1]
        [7,0] [6,1] [7,1]
        [5,0] [7,0] [6,1] [7,1]
        [5,0] [7,0] [5,2] [6,1] [7,1]
        [5,0] [7,0] [5,2] [6,1] [4,4] [7,1]
```

**Complexity:** O(n²).

```python
class Solution:
    def reconstructQueue(self, people: List[List[int]]) -> List[List[int]]:
        res = []
        for h, k in sorted(people, key=lambda p: (-p[0], p[1])):
            res.insert(k, [h, k])
        return res
# test
P = [[7, 0], [4, 4], [7, 1], [5, 0], [6, 1], [5, 2]]
assert Solution().reconstructQueue(P) == [[5, 0], [7, 0], [5, 2], [6, 1], [4, 4], [7, 1]]
assert Solution().reconstructQueue([[6, 0], [5, 0], [4, 0], [3, 2], [2, 2], [1, 4]]) == [[4, 0], [5, 0], [2, 2], [3, 2], [1, 4], [6, 0]]
```

```java
class Solution {
    public int[][] reconstructQueue(int[][] people) {
        Arrays.sort(people, (a, b) -> a[0] != b[0] ? b[0] - a[0] : a[1] - b[1]);
        List<int[]> res = new ArrayList<>();
        for (int[] p : people) res.add(p[1], p);
        return res.toArray(new int[0][]);
    }
}
// test
int[][] P = {{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}};
check(new Solution().reconstructQueue(P), new int[][]{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}});
```

```go
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool { // tallest first; fewer people in front first
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	res := [][]int{}
	for _, p := range people { // everyone already placed is at least as tall: insert at index k
		k := p[1]
		res = append(res, p) // grow by one, then shift right from k
		copy(res[k+1:], res[k:])
		res[k] = p
	}
	return res
}
// test
P := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
assertEq(reconstructQueue(P), [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}})
assertEq(reconstructQueue([][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}}), [][]int{{4, 0}, {5, 0}, {2, 2}, {3, 2}, {1, 4}, {6, 0}})
```



## 455. Assign Cookies | Easy
related: 881
tags: sort, two-pointers

```problem
Child i is content with a cookie of size at least `g[i]`; cookie j has size `s[j]`. Each child gets at most one cookie. Return the maximum number of content children.

**Example:** `g = [1,2,3], s = [1,1]` → `1`; `g = [1,2], s = [1,2,3]` → `2`
```

Sort both. Give the smallest cookie that satisfies the least greedy child; a cookie too small for this child is too small for everyone after.

**Complexity:** O(n log n).

```python
class Solution:
    def findContentChildren(self, g: List[int], s: List[int]) -> int:
        g.sort(); s.sort()
        child = 0
        for size in s:
            if child < len(g) and size >= g[child]:
                child += 1
        return child
# test
assert Solution().findContentChildren([1, 2, 3], [1, 1]) == 1
assert Solution().findContentChildren([1, 2], [1, 2, 3]) == 2
```

```java
class Solution {
    public int findContentChildren(int[] g, int[] s) {
        Arrays.sort(g); Arrays.sort(s);
        int child = 0;
        for (int size : s) if (child < g.length && size >= g[child]) child++;
        return child;
    }
}
// test
check(new Solution().findContentChildren(new int[]{1, 2}, new int[]{1, 2, 3}), 2);
```

```go
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	child := 0
	for _, size := range s { // each cookie goes to the least greedy child it satisfies
		if child < len(g) && size >= g[child] {
			child++
		}
	}
	return child
}
// test
assertEq(findContentChildren([]int{1, 2, 3}, []int{1, 1}), 1)
assertEq(findContentChildren([]int{1, 2}, []int{1, 2, 3}), 2)
```


## 860. Lemonade Change | Easy
related: 455
tags: simulation

```problem
Lemonade costs $5. Customers in line each pay with a $5, $10 or $20 bill, and you start with no change. Decide whether you can give every customer correct change, in order.

**Example:** `[5,5,5,10,20]` → `true`; `[5,5,10,10,20]` → `false`
```

Track 5s and 10s. For a 20, prefer giving `10 + 5` (keeps 5s, which are more flexible) over `5 + 5 + 5`.

**Complexity:** O(n).

```python
class Solution:
    def lemonadeChange(self, bills: List[int]) -> bool:
        five = ten = 0
        for b in bills:
            if b == 5:
                five += 1
            elif b == 10:
                five, ten = five - 1, ten + 1
            elif ten:
                five, ten = five - 1, ten - 1
            else:
                five -= 3
            if five < 0:
                return False
        return True
# test
assert Solution().lemonadeChange([5, 5, 5, 10, 20]) is True
assert Solution().lemonadeChange([5, 5, 10, 10, 20]) is False
```

```java
class Solution {
    public boolean lemonadeChange(int[] bills) {
        int five = 0, ten = 0;
        for (int b : bills) {
            if (b == 5) five++;
            else if (b == 10) { five--; ten++; }
            else if (ten > 0) { five--; ten--; }
            else five -= 3;
            if (five < 0) return false;
        }
        return true;
    }
}
// test
check(new Solution().lemonadeChange(new int[]{5, 5, 10, 10, 20}), false);
```

```go
func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, b := range bills {
		switch {
		case b == 5:
			five++
		case b == 10:
			five, ten = five-1, ten+1
		case ten > 0: // give 10 + 5 for a 20: keep fives, they are more useful
			five, ten = five-1, ten-1
		default:
			five -= 3
		}
		if five < 0 {
			return false
		}
	}
	return true
}
// test
assertEq(lemonadeChange([]int{5, 5, 5, 10, 20}), true)
assertEq(lemonadeChange([]int{5, 5, 10, 10, 20}), false)
```


## 1029. Two City Scheduling | Medium
related: 406
tags: sort, exchange-argument

```problem
A company will interview 2n people; `costs[i] = [aCost, bCost]` is the cost of flying person i to city A or city B. Send exactly n people to each city at minimum total cost.

**Example:** `[[10,20],[30,200],[400,50],[30,20]]` → `110`
```

Send everyone to A, then switch the n people with the most negative `costB − costA` to B. Equivalently, sort by `costA − costB`: first half → A, second half → B.

**Complexity:** O(n log n).

```python
class Solution:
    def twoCitySchedCost(self, costs: List[List[int]]) -> int:
        costs.sort(key=lambda c: c[0] - c[1])
        n = len(costs) // 2
        return sum(a for a, _ in costs[:n]) + sum(b for _, b in costs[n:])
# test
assert Solution().twoCitySchedCost([[10, 20], [30, 200], [400, 50], [30, 20]]) == 110
assert Solution().twoCitySchedCost([[259, 770], [448, 54], [926, 667], [184, 139], [840, 118], [577, 469]]) == 1859
```

```java
class Solution {
    public int twoCitySchedCost(int[][] costs) {
        Arrays.sort(costs, (x, y) -> (x[0] - x[1]) - (y[0] - y[1]));
        int n = costs.length / 2, total = 0;
        for (int i = 0; i < costs.length; i++) total += i < n ? costs[i][0] : costs[i][1];
        return total;
    }
}
// test
check(new Solution().twoCitySchedCost(new int[][]{{10, 20}, {30, 200}, {400, 50}, {30, 20}}), 110);
```

```go
func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool { // most to gain from city A first
		return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
	})
	n, total := len(costs)/2, 0
	for i, c := range costs {
		if i < n {
			total += c[0]
		} else {
			total += c[1]
		}
	}
	return total
}
// test
assertEq(twoCitySchedCost([][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}), 110)
assertEq(twoCitySchedCost([][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}), 1859)
```


## 649. Dota2 Senate | Medium
related: 1834
tags: queue, simulation

```problem
Senators from two parties, `'R'` (Radiant) and `'D'` (Dire), act in the order of the string, round after round. On their turn, a senator still in play either removes one opposing senator's rights for good or, if only their own party is left, declares victory. Assuming everyone plays optimally, return `"Radiant"` or `"Dire"`.

**Example:** `"RD"` → `"Radiant"`; `"RDD"` → `"Dire"`

**Constraints:** length ≤ 10⁴.
```

Two queues of senator indices. The earlier senator bans the other; the winner re-enters with index `+ n` (next round).

**Complexity:** O(n).

```python
class Solution:
    def predictPartyVictory(self, senate: str) -> str:
        n = len(senate)
        r = deque(i for i, c in enumerate(senate) if c == "R")
        d = deque(i for i, c in enumerate(senate) if c == "D")
        while r and d:
            a, b = r.popleft(), d.popleft()
            if a < b:
                r.append(a + n)
            else:
                d.append(b + n)
        return "Radiant" if r else "Dire"
# test
assert Solution().predictPartyVictory("RD") == "Radiant"
assert Solution().predictPartyVictory("RDD") == "Dire"
```

```java
class Solution {
    public String predictPartyVictory(String senate) {
        int n = senate.length();
        Deque<Integer> r = new ArrayDeque<>(), d = new ArrayDeque<>();
        for (int i = 0; i < n; i++) (senate.charAt(i) == 'R' ? r : d).add(i);
        while (!r.isEmpty() && !d.isEmpty()) {
            int a = r.poll(), b = d.poll();
            if (a < b) r.add(a + n); else d.add(b + n);
        }
        return r.isEmpty() ? "Dire" : "Radiant";
    }
}
// test
check(new Solution().predictPartyVictory("RDD"), "Dire");
```

```go
func predictPartyVictory(senate string) string {
	n := len(senate)
	r, d := []int{}, []int{} // turn indices of each party's active senators
	for i := 0; i < n; i++ {
		if senate[i] == 'R' {
			r = append(r, i)
		} else {
			d = append(d, i)
		}
	}
	for len(r) > 0 && len(d) > 0 {
		a, b := r[0], d[0]
		r, d = r[1:], d[1:]
		if a < b { // the earlier senator bans the other and votes again next round
			r = append(r, a+n)
		} else {
			d = append(d, b+n)
		}
	}
	if len(r) > 0 {
		return "Radiant"
	}
	return "Dire"
}
// test
assertEq(predictPartyVictory("RD"), "Radiant")
assertEq(predictPartyVictory("RDD"), "Dire")
```


## 659. Split Array into Consecutive Subsequences | Medium
related: 846
tags: greedy, hash-map

```problem
`nums` is sorted. Decide whether it can be split into subsequences that each consist of consecutive integers (each one more than the previous) and have length at least 3.

**Example:** `[1,2,3,3,4,5]` → `true` (1,2,3 and 3,4,5); `[1,2,3,4,4,5]` → `false`

**Constraints:** n ≤ 10⁴.
```

Greedy: each number first tries to **extend** a sequence ending at `x−1` (`ends[x−1] > 0`); otherwise it must **start** a new sequence `x, x+1, x+2` (needs those available). If neither works → false.

**Complexity:** O(n).

```python
class Solution:
    def isPossible(self, nums: List[int]) -> bool:
        left, ends = Counter(nums), Counter()
        for x in nums:
            if not left[x]:
                continue
            left[x] -= 1
            if ends[x - 1]:
                ends[x - 1] -= 1
                ends[x] += 1
            elif left[x + 1] and left[x + 2]:
                left[x + 1] -= 1
                left[x + 2] -= 1
                ends[x + 2] += 1
            else:
                return False
        return True
# test
assert Solution().isPossible([1, 2, 3, 3, 4, 5]) is True
assert Solution().isPossible([1, 2, 3, 3, 4, 4, 5, 5]) is True
assert Solution().isPossible([1, 2, 3, 4, 4, 5]) is False
```

```java
class Solution {
    public boolean isPossible(int[] nums) {
        Map<Integer, Integer> left = new HashMap<>(), ends = new HashMap<>();
        for (int x : nums) left.merge(x, 1, Integer::sum);
        for (int x : nums) {
            if (left.getOrDefault(x, 0) == 0) continue;
            left.merge(x, -1, Integer::sum);
            if (ends.getOrDefault(x - 1, 0) > 0) {
                ends.merge(x - 1, -1, Integer::sum);
                ends.merge(x, 1, Integer::sum);
            } else if (left.getOrDefault(x + 1, 0) > 0 && left.getOrDefault(x + 2, 0) > 0) {
                left.merge(x + 1, -1, Integer::sum);
                left.merge(x + 2, -1, Integer::sum);
                ends.merge(x + 2, 1, Integer::sum);
            } else return false;
        }
        return true;
    }
}
// test
check(new Solution().isPossible(new int[]{1, 2, 3, 3, 4, 4, 5, 5}), true);
check(new Solution().isPossible(new int[]{1, 2, 3, 4, 4, 5}), false);
```

```go
func isPossible(nums []int) bool {
	left, ends := map[int]int{}, map[int]int{} // unused counts; subsequences ending at x
	for _, x := range nums {
		left[x]++
	}
	for _, x := range nums {
		if left[x] == 0 {
			continue
		}
		left[x]--
		if ends[x-1] > 0 { // extend an existing subsequence
			ends[x-1]--
			ends[x]++
		} else if left[x+1] > 0 && left[x+2] > 0 { // or start a new one of length 3
			left[x+1]--
			left[x+2]--
			ends[x+2]++
		} else {
			return false
		}
	}
	return true
}
// test
assertEq(isPossible([]int{1, 2, 3, 3, 4, 5}), true)
assertEq(isPossible([]int{1, 2, 3, 3, 4, 4, 5, 5}), true)
assertEq(isPossible([]int{1, 2, 3, 4, 4, 5}), false)
```


## 1326. Minimum Number of Taps to Open to Water a Garden | Hard
related: 45, 1024
tags: greedy, reach

```problem
A garden spans [0, n] on a line, with taps at points 0…n; tap i waters [i − ranges[i], i + ranges[i]]. Return the fewest taps that must be open to water the entire garden, or −1 if it can't be done.

**Example:** `n = 5, ranges = [3,4,1,1,0,0]` → `1` (tap 1 covers [−3, 5])

**Constraints:** n ≤ 10⁴.
```

Convert each tap to an interval and record, for every left end, the farthest right reach. Then it's **Jump Game II** over `[0, n]`: greedily extend the covered range level by level; a gap → −1.

**Complexity:** O(n).

```python
class Solution:
    def minTaps(self, n: int, ranges: List[int]) -> int:
        reach = [0] * (n + 1)
        for i, r in enumerate(ranges):
            lo, hi = max(0, i - r), min(n, i + r)
            reach[lo] = max(reach[lo], hi)
        taps = end = far = 0
        for i in range(n):
            far = max(far, reach[i])
            if i == end:
                if far <= i:
                    return -1
                taps += 1
                end = far
        return taps
# test
assert Solution().minTaps(5, [3, 4, 1, 1, 0, 0]) == 1
assert Solution().minTaps(3, [0, 0, 0, 0]) == -1
assert Solution().minTaps(7, [1, 2, 1, 0, 2, 1, 0, 1]) == 3
```

```java
class Solution {
    public int minTaps(int n, int[] ranges) {
        int[] reach = new int[n + 1];
        for (int i = 0; i <= n; i++) {
            int lo = Math.max(0, i - ranges[i]), hi = Math.min(n, i + ranges[i]);
            reach[lo] = Math.max(reach[lo], hi);
        }
        int taps = 0, end = 0, far = 0;
        for (int i = 0; i < n; i++) {
            far = Math.max(far, reach[i]);
            if (i == end) {
                if (far <= i) return -1;
                taps++;
                end = far;
            }
        }
        return taps;
    }
}
// test
check(new Solution().minTaps(7, new int[]{1, 2, 1, 0, 2, 1, 0, 1}), 3);
check(new Solution().minTaps(3, new int[]{0, 0, 0, 0}), -1);
```

```go
func minTaps(n int, ranges []int) int {
	reach := make([]int, n+1) // reach[lo] = farthest point a tap starting at lo covers
	for i, r := range ranges {
		lo, hi := max(0, i-r), min(n, i+r)
		reach[lo] = max(reach[lo], hi)
	}
	taps, end, far := 0, 0, 0 // then it is Jump Game II
	for i := 0; i < n; i++ {
		far = max(far, reach[i])
		if i == end {
			if far <= i {
				return -1 // a gap nothing covers
			}
			taps++
			end = far
		}
	}
	return taps
}
// test
assertEq(minTaps(5, []int{3, 4, 1, 1, 0, 0}), 1)
assertEq(minTaps(3, []int{0, 0, 0, 0}), -1)
assertEq(minTaps(7, []int{1, 2, 1, 0, 2, 1, 0, 1}), 3)
```


## 605. Can Place Flowers | Easy
related: 495
tags: greedy

```problem
`flowerbed` holds 0 (empty) and 1 (planted), and no two flowers may be adjacent. Decide whether `n` new flowers can be planted without breaking the rule.

**Example:** `flowerbed = [1,0,0,0,1]`: `n = 1` → `true`; `n = 2` → `false`
```

Planting as **early** as possible never hurts (an earlier flower blocks at most what a later one would). Pad both ends with 0 and plant wherever a spot and both neighbours are empty.

**Complexity:** O(n).

```python
class Solution:
    def canPlaceFlowers(self, flowerbed: List[int], n: int) -> bool:
        bed = [0] + flowerbed + [0]
        for i in range(1, len(bed) - 1):
            if bed[i - 1] == bed[i] == bed[i + 1] == 0:
                bed[i] = 1
                n -= 1
        return n <= 0
# test
assert Solution().canPlaceFlowers([1, 0, 0, 0, 1], 1) is True
assert Solution().canPlaceFlowers([1, 0, 0, 0, 1], 2) is False
assert Solution().canPlaceFlowers([0], 1) is True
assert Solution().canPlaceFlowers([0, 0, 1, 0, 0], 2) is True
assert Solution().canPlaceFlowers([1, 0, 1, 0, 1], 0) is True
```

```java
class Solution {
    public boolean canPlaceFlowers(int[] flowerbed, int n) {
        int len = flowerbed.length;
        for (int i = 0; i < len && n > 0; i++) {
            boolean leftFree = i == 0 || flowerbed[i - 1] == 0, rightFree = i == len - 1 || flowerbed[i + 1] == 0;
            if (flowerbed[i] == 0 && leftFree && rightFree) { flowerbed[i] = 1; n--; }
        }
        return n <= 0;
    }
}
// test
check(new Solution().canPlaceFlowers(new int[]{1, 0, 0, 0, 1}, 1), true);
check(new Solution().canPlaceFlowers(new int[]{1, 0, 0, 0, 1}, 2), false);
```

```go
func canPlaceFlowers(flowerbed []int, n int) bool {
	bed := append(append([]int{0}, flowerbed...), 0) // pad both ends with empty plots
	for i := 1; i < len(bed)-1; i++ {
		if bed[i-1] == 0 && bed[i] == 0 && bed[i+1] == 0 { // planting as early as possible is optimal
			bed[i] = 1
			n--
		}
	}
	return n <= 0
}
// test
assertEq(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1), true)
assertEq(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2), false)
assertEq(canPlaceFlowers([]int{0}, 1), true)
assertEq(canPlaceFlowers([]int{0, 0, 1, 0, 0}, 2), true)
assertEq(canPlaceFlowers([]int{1, 0, 1, 0, 1}, 0), true)
```


## 334. Increasing Triplet Subsequence | Medium
related: 300
tags: greedy, lis

```problem
Decide whether there are indices i < j < k with nums[i] < nums[j] < nums[k].

**Example:** `[2,1,5,0,4,6]` → `true`; `[5,4,3,2,1]` → `false`

**Constraints:** n ≤ 5·10⁵. Follow-up: O(n) time and O(1) space.
```

Keep the smallest possible **first** element and the smallest possible **second** element of an increasing pair seen so far. Any number bigger than both completes a triplet. (Updating `first` after `second` is set is fine: `second` still certifies that some smaller element came before it.)

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def increasingTriplet(self, nums: List[int]) -> bool:
        first = second = float("inf")
        for x in nums:
            if x <= first:
                first = x
            elif x <= second:
                second = x
            else:
                return True
        return False
# test
assert Solution().increasingTriplet([1, 2, 3, 4, 5]) is True
assert Solution().increasingTriplet([5, 4, 3, 2, 1]) is False
assert Solution().increasingTriplet([2, 1, 5, 0, 4, 6]) is True
assert Solution().increasingTriplet([20, 100, 10, 12, 5, 13]) is True
assert Solution().increasingTriplet([1, 1, 1]) is False
assert Solution().increasingTriplet([1, 5, 0, 4, 1, 3]) is True
```

```java
class Solution {
    public boolean increasingTriplet(int[] nums) {
        int first = Integer.MAX_VALUE, second = Integer.MAX_VALUE;
        for (int x : nums) {
            if (x <= first) first = x;
            else if (x <= second) second = x;
            else return true;
        }
        return false;
    }
}
// test
check(new Solution().increasingTriplet(new int[]{2, 1, 5, 0, 4, 6}), true);
check(new Solution().increasingTriplet(new int[]{5, 4, 3, 2, 1}), false);
```

```go
func increasingTriplet(nums []int) bool {
	first, second := math.MaxInt, math.MaxInt // smallest end of increasing runs of length 1 and 2
	for _, x := range nums {
		if x <= first {
			first = x
		} else if x <= second {
			second = x
		} else {
			return true
		}
	}
	return false
}
// test
assertEq(increasingTriplet([]int{1, 2, 3, 4, 5}), true)
assertEq(increasingTriplet([]int{5, 4, 3, 2, 1}), false)
assertEq(increasingTriplet([]int{2, 1, 5, 0, 4, 6}), true)
assertEq(increasingTriplet([]int{20, 100, 10, 12, 5, 13}), true)
assertEq(increasingTriplet([]int{1, 1, 1}), false)
assertEq(increasingTriplet([]int{1, 5, 0, 4, 1, 3}), true)
```


## 68. Text Justification | Hard
related: 6, 1592
tags: simulation, string

```problem
Format `words` into lines of exactly `maxWidth` characters. Pack as many words per line as fit (greedily), then pad with spaces so each line is fully justified, giving extra spaces to the leftmost gaps first. The last line, and any line with a single word, is left-justified and padded with spaces on the right.

**Example:** `["This","is","an","example","of","text","justification."], maxWidth = 16` → `["This    is    an","example  of text","justification.  "]`
```

Greedy line packing: a word fits if `letters + len(word) + (number of words already on the line)` ≤ maxWidth (one space minimum per gap). When the line is full, distribute `maxWidth − letters` spaces over the gaps **round-robin from the left** (left gaps get the extras); a single-word line is left-justified. The last line is left-justified and padded.

```viz
maxWidth 16: "This" "is" "an"  letters 8, gaps 2, spaces 8 -> 4 + 4
"This    is    an"
"example  of text"   letters 13, gaps 2, spaces 3 -> 2 + 1
"justification.  "   last line: single spaces, pad right
```

**Complexity:** O(total characters).

```python
class Solution:
    def fullJustify(self, words: List[str], maxWidth: int) -> List[str]:
        res, line, letters = [], [], 0
        for w in words:
            if letters + len(w) + len(line) > maxWidth:
                gaps = len(line) - 1 or 1
                for i in range(maxWidth - letters):
                    line[i % gaps] += " "
                res.append("".join(line))
                line, letters = [], 0
            line.append(w)
            letters += len(w)
        res.append(" ".join(line).ljust(maxWidth))
        return res
# test
assert Solution().fullJustify(["This", "is", "an", "example", "of", "text", "justification."], 16) == [
    "This    is    an", "example  of text", "justification.  "]
assert Solution().fullJustify(["What", "must", "be", "acknowledgment", "shall", "be"], 16) == [
    "What   must   be", "acknowledgment  ", "shall be        "]
assert Solution().fullJustify(["Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
                               "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"], 20) == [
    "Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is",
    "everything  else  we", "do                  "]
assert Solution().fullJustify(["a"], 1) == ["a"]
```

```java
class Solution {
    public List<String> fullJustify(String[] words, int maxWidth) {
        List<String> res = new ArrayList<>();
        int i = 0;
        while (i < words.length) {
            int j = i, letters = 0;
            while (j < words.length && letters + words[j].length() + (j - i) <= maxWidth) letters += words[j++].length();
            StringBuilder sb = new StringBuilder(words[i]);
            int gaps = j - i - 1;
            if (j == words.length || gaps == 0) {              // last line or single word: left-justify
                for (int k = i + 1; k < j; k++) sb.append(' ').append(words[k]);
                sb.append(" ".repeat(maxWidth - sb.length()));
            } else {
                int spaces = maxWidth - letters, each = spaces / gaps, extra = spaces % gaps;
                for (int k = i + 1; k < j; k++) sb.append(" ".repeat(each + (k - i <= extra ? 1 : 0))).append(words[k]);
            }
            res.add(sb.toString());
            i = j;
        }
        return res;
    }
}
// test
check(new Solution().fullJustify(new String[]{"This", "is", "an", "example", "of", "text", "justification."}, 16),
      List.of("This    is    an", "example  of text", "justification.  "));
check(new Solution().fullJustify(new String[]{"What", "must", "be", "acknowledgment", "shall", "be"}, 16),
      List.of("What   must   be", "acknowledgment  ", "shall be        "));
```

```go
func fullJustify(words []string, maxWidth int) []string {
	res := []string{}
	line, letters := []string{}, 0
	for _, w := range words {
		if letters+len(w)+len(line) > maxWidth { // w does not fit: justify the current line
			gaps := max(len(line)-1, 1)
			for i := 0; i < maxWidth-letters; i++ {
				line[i%gaps] += " " // round-robin puts extra spaces on the left
			}
			res = append(res, strings.Join(line, ""))
			line, letters = []string{}, 0
		}
		line = append(line, w)
		letters += len(w)
	}
	last := strings.Join(line, " ") // the last line is left-justified
	return append(res, last+strings.Repeat(" ", maxWidth-len(last)))
}
// test
assertEq(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16), []string{
	"This    is    an", "example  of text", "justification.  "})
assertEq(fullJustify([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16), []string{
	"What   must   be", "acknowledgment  ", "shall be        "})
assertEq(fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
	"to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20), []string{
	"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is",
	"everything  else  we", "do                  "})
assertEq(fullJustify([]string{"a"}, 1), []string{"a"})
```


## 1871. Jump Game VII | Medium
related: 55, 45, 1696
tags: sliding-window, prefix-sum

```problem
`s` is a binary string starting with `'0'`. From index i you can jump to any j with i + minJump ≤ j ≤ i + maxJump and `s[j] == '0'`. Decide whether you can reach the last index from index 0.

**Example:** `s = "011010", minJump = 2, maxJump = 3` → `true`

**Constraints:** length ≤ 10⁵.
```

`i` is reachable iff `s[i] == '0'` and **some** reachable index lies in `[i − maxJump, i − minJump]`. Keep a running count of reachable indices inside that window (add the index entering at `i − minJump`, drop the one leaving at `i − maxJump − 1`) — O(1) per position instead of scanning the window.

**Complexity:** O(n).

```python
class Solution:
    def canReach(self, s: str, minJump: int, maxJump: int) -> bool:
        n = len(s)
        ok = [False] * n
        ok[0] = True
        window = 0
        for i in range(1, n):
            if i >= minJump and ok[i - minJump]:
                window += 1
            if i > maxJump and ok[i - maxJump - 1]:
                window -= 1
            ok[i] = window > 0 and s[i] == "0"
        return ok[-1]
# test
assert Solution().canReach("011010", 2, 3) is True
assert Solution().canReach("01101110", 2, 3) is False
assert Solution().canReach("00", 1, 1) is True
assert Solution().canReach("01", 1, 1) is False
assert Solution().canReach("0000000000", 8, 8) is False
```

```java
class Solution {
    public boolean canReach(String s, int minJump, int maxJump) {
        int n = s.length(), window = 0;
        boolean[] ok = new boolean[n];
        ok[0] = true;
        for (int i = 1; i < n; i++) {
            if (i >= minJump && ok[i - minJump]) window++;
            if (i > maxJump && ok[i - maxJump - 1]) window--;
            ok[i] = window > 0 && s.charAt(i) == '0';
        }
        return ok[n - 1];
    }
}
// test
check(new Solution().canReach("011010", 2, 3), true);
check(new Solution().canReach("01101110", 2, 3), false);
```

```go
func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	ok := make([]bool, n) // ok[i]: index i is reachable
	ok[0] = true
	window := 0 // reachable indices in [i-maxJump, i-minJump]
	for i := 1; i < n; i++ {
		if i >= minJump && ok[i-minJump] {
			window++
		}
		if i > maxJump && ok[i-maxJump-1] {
			window--
		}
		ok[i] = window > 0 && s[i] == '0'
	}
	return ok[n-1]
}
// test
assertEq(canReach("011010", 2, 3), true)
assertEq(canReach("01101110", 2, 3), false)
assertEq(canReach("00", 1, 1), true)
assertEq(canReach("01", 1, 1), false)
assertEq(canReach("0000000000", 8, 8), false)
```

