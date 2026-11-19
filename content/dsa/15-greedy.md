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

## 918. Maximum Sum Circular Subarray | Medium
related: 53
tags: kadane

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

## 55. Jump Game | Medium
related: 45
tags: blind75, reach

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

## 45. Jump Game II | Medium
related: 55
tags: reach, implicit-bfs

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

## 122. Best Time to Buy and Sell Stock II | Medium
related: 121, 309, 714
tags: greedy

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

## 134. Gas Station | Medium
related: 55
tags: greedy

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

## 135. Candy | Hard
related: 42
tags: two-pass

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

## 846. Hand of Straights | Medium
related: 659
tags: sorted-counter

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

## 1899. Merge Triplets to Form Target Triplet | Medium
related: 55
tags: greedy

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

## 763. Partition Labels | Medium
related: 56
tags: last-occurrence

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

## 678. Valid Parenthesis String | Medium
related: 20, 32
tags: range-tracking

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

## 406. Queue Reconstruction by Height | Medium
related: 315
tags: sort-then-insert

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


## 455. Assign Cookies | Easy
related: 881
tags: sort, two-pointers

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

## 860. Lemonade Change | Easy
related: 455
tags: simulation

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

## 1029. Two City Scheduling | Medium
related: 406
tags: sort, exchange-argument

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

## 649. Dota2 Senate | Medium
related: 1834
tags: queue, simulation

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

## 659. Split Array into Consecutive Subsequences | Medium
related: 846
tags: greedy, hash-map

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

## 1326. Minimum Number of Taps to Open to Water a Garden | Hard
related: 45, 1024
tags: greedy, reach

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
