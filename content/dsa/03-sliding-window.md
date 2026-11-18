# Sliding Window

**When to use:** contiguous subarray/substring with a constraint ("longest/shortest/count of windows such that …") where expanding/shrinking the window changes the constraint **monotonically**.

- **Variable window**: grow `r` every step; while the window is invalid, shrink `l`. Answer is updated when valid.
- **Fixed window** of size k: add `a[r]`, remove `a[r-k]`.
- Needs non-negative numbers for sum constraints — with negatives use prefix sums (see 560).

```python
l = 0
for r, x in enumerate(a):
    add(x)
    while invalid():
        remove(a[l]); l += 1
    best = max(best, r - l + 1)
```

## 121. Best Time to Buy and Sell Stock | Easy
related: 122, 309, 123, 53
tags: blind75, running-min

Track the cheapest price so far; each day, profit = price − min so far.

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        lo, best = float("inf"), 0
        for p in prices:
            lo = min(lo, p)
            best = max(best, p - lo)
        return best
# test
assert Solution().maxProfit([7, 1, 5, 3, 6, 4]) == 5
assert Solution().maxProfit([7, 6, 4, 3, 1]) == 0
```

```java
class Solution {
    public int maxProfit(int[] prices) {
        int lo = Integer.MAX_VALUE, best = 0;
        for (int p : prices) {
            lo = Math.min(lo, p);
            best = Math.max(best, p - lo);
        }
        return best;
    }
}
// test
check(new Solution().maxProfit(new int[]{7, 1, 5, 3, 6, 4}), 5);
check(new Solution().maxProfit(new int[]{7, 6, 4, 3, 1}), 0);
```

## 3. Longest Substring Without Repeating Characters | Medium
related: 424, 76, 159
tags: blind75, variable-window

Keep the last index of each char. When `s[r]` was seen inside the window, jump `l` past that index.

```viz
a b c a b c b b
l     r          'a' last at 0 -> l = 1   window "bca"
best = 3
```

**Complexity:** O(n) time, O(min(n, alphabet)) space.

```python
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        last = {}
        l = best = 0
        for r, ch in enumerate(s):
            if last.get(ch, -1) >= l:
                l = last[ch] + 1
            last[ch] = r
            best = max(best, r - l + 1)
        return best
# test
assert Solution().lengthOfLongestSubstring("abcabcbb") == 3
assert Solution().lengthOfLongestSubstring("bbbbb") == 1
assert Solution().lengthOfLongestSubstring("pwwkew") == 3
assert Solution().lengthOfLongestSubstring("abba") == 2
```

```java
class Solution {
    public int lengthOfLongestSubstring(String s) {
        int[] last = new int[128];
        Arrays.fill(last, -1);
        int l = 0, best = 0;
        for (int r = 0; r < s.length(); r++) {
            char c = s.charAt(r);
            if (last[c] >= l) l = last[c] + 1;
            last[c] = r;
            best = Math.max(best, r - l + 1);
        }
        return best;
    }
}
// test
check(new Solution().lengthOfLongestSubstring("abcabcbb"), 3);
check(new Solution().lengthOfLongestSubstring("pwwkew"), 3);
check(new Solution().lengthOfLongestSubstring("abba"), 2);
```

## 424. Longest Repeating Character Replacement | Medium
related: 3, 1004
tags: blind75, variable-window

A window is fixable if `len − (count of its most frequent char) ≤ k`. Grow `r`; if invalid, slide `l` by one. `maxf` never needs to decrease: the answer only improves when a larger `maxf` appears.

**Complexity:** O(n) time, O(26) space.

```python
class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        cnt = Counter()
        l = maxf = 0
        for r, ch in enumerate(s):
            cnt[ch] += 1
            maxf = max(maxf, cnt[ch])
            if r - l + 1 - maxf > k:
                cnt[s[l]] -= 1
                l += 1
        return len(s) - l
# test
assert Solution().characterReplacement("ABAB", 2) == 4
assert Solution().characterReplacement("AABABBA", 1) == 4
```

```java
class Solution {
    public int characterReplacement(String s, int k) {
        int[] cnt = new int[26];
        int l = 0, maxf = 0;
        for (int r = 0; r < s.length(); r++) {
            maxf = Math.max(maxf, ++cnt[s.charAt(r) - 'A']);
            if (r - l + 1 - maxf > k) cnt[s.charAt(l++) - 'A']--;
        }
        return s.length() - l; // window never shrinks, so its final size is the best size
    }
}
// test
check(new Solution().characterReplacement("ABAB", 2), 4);
check(new Solution().characterReplacement("AABABBA", 1), 4);
```

## 567. Permutation in String | Medium
related: 438, 76, 242
tags: fixed-window, counting

Slide a window of size `len(s1)` over `s2` and compare letter counts. Track `matches` = number of letters whose counts agree (0..26) to make each slide O(1).

**Complexity:** O(n) time, O(26) space.

```python
class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        k = len(s1)
        if k > len(s2):
            return False
        need, win = Counter(s1), Counter(s2[:k])
        if need == win:
            return True
        for r in range(k, len(s2)):
            win[s2[r]] += 1
            win[s2[r - k]] -= 1
            if win[s2[r - k]] == 0:
                del win[s2[r - k]]
            if win == need:
                return True
        return False
# test
assert Solution().checkInclusion("ab", "eidbaooo") is True
assert Solution().checkInclusion("ab", "eidboaoo") is False
```

```java
class Solution {
    public boolean checkInclusion(String s1, String s2) {
        int k = s1.length(), n = s2.length();
        if (k > n) return false;
        int[] diff = new int[26];              // diff = count(window) - count(s1)
        for (int i = 0; i < k; i++) { diff[s1.charAt(i) - 'a']--; diff[s2.charAt(i) - 'a']++; }
        int zeros = 0;
        for (int d : diff) if (d == 0) zeros++;
        for (int r = k; ; r++) {
            if (zeros == 26) return true;
            if (r == n) return false;
            int in = s2.charAt(r) - 'a', out = s2.charAt(r - k) - 'a';
            if (diff[in] == 0) zeros--; if (++diff[in] == 0) zeros++;
            if (diff[out] == 0) zeros--; if (--diff[out] == 0) zeros++;
        }
    }
}
// test
check(new Solution().checkInclusion("ab", "eidbaooo"), true);
check(new Solution().checkInclusion("ab", "eidboaoo"), false);
```

## 76. Minimum Window Substring | Hard
related: 3, 567, 438
tags: blind75, variable-window

Expand `r` until the window covers all of `t` (`missing == 0`), then shrink `l` as far as possible while still covering, recording the smallest window.

```viz
s = ADOBECODEBANC, t = ABC
[ADOBEC]ODEBANC      covers -> shrink: drop A -> missing A
 ADOBEC[ODEBA]NC     ...
 ADOBECODE[BANC]     covers, len 4  <- answer "BANC"
```

**Complexity:** O(|s| + |t|).

```python
class Solution:
    def minWindow(self, s: str, t: str) -> str:
        need = Counter(t)
        missing = len(t)
        l = start = 0
        best = float("inf")
        for r, ch in enumerate(s):
            if need[ch] > 0:
                missing -= 1
            need[ch] -= 1
            while missing == 0:
                if r - l + 1 < best:
                    start, best = l, r - l + 1
                need[s[l]] += 1
                if need[s[l]] > 0:
                    missing += 1
                l += 1
        return "" if best == float("inf") else s[start:start + best]
# test
assert Solution().minWindow("ADOBECODEBANC", "ABC") == "BANC"
assert Solution().minWindow("a", "aa") == ""
```

```java
class Solution {
    public String minWindow(String s, String t) {
        int[] need = new int[128];
        for (char c : t.toCharArray()) need[c]++;
        int missing = t.length(), l = 0, start = 0, best = Integer.MAX_VALUE;
        for (int r = 0; r < s.length(); r++) {
            if (need[s.charAt(r)]-- > 0) missing--;
            while (missing == 0) {
                if (r - l + 1 < best) { best = r - l + 1; start = l; }
                if (++need[s.charAt(l++)] > 0) missing++;
            }
        }
        return best == Integer.MAX_VALUE ? "" : s.substring(start, start + best);
    }
}
// test
check(new Solution().minWindow("ADOBECODEBANC", "ABC"), "BANC");
check(new Solution().minWindow("a", "aa"), "");
```

## 239. Sliding Window Maximum | Hard
related: 155, 1438, 862
tags: monotonic-deque

Keep a deque of indices with **decreasing** values. New element pops smaller ones from the back (they can never be a max again). Front is the max; drop it when it leaves the window.

```viz
nums 1 3 -1 -3 5 3 6 7, k=3
r=2  deque [3,-1]       max 3
r=3  deque [3,-1,-3]    max 3
r=4  5 pops all -> [5]  max 5
```

**Complexity:** O(n) time (each index pushed/popped once), O(k) space.

```python
class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        dq, res = deque(), []
        for i, x in enumerate(nums):
            while dq and nums[dq[-1]] <= x:
                dq.pop()
            dq.append(i)
            if dq[0] <= i - k:
                dq.popleft()
            if i >= k - 1:
                res.append(nums[dq[0]])
        return res
# test
assert Solution().maxSlidingWindow([1, 3, -1, -3, 5, 3, 6, 7], 3) == [3, 3, 5, 5, 6, 7]
assert Solution().maxSlidingWindow([1], 1) == [1]
```

```java
class Solution {
    public int[] maxSlidingWindow(int[] nums, int k) {
        Deque<Integer> dq = new ArrayDeque<>();
        int[] res = new int[nums.length - k + 1];
        for (int i = 0; i < nums.length; i++) {
            while (!dq.isEmpty() && nums[dq.peekLast()] <= nums[i]) dq.pollLast();
            dq.addLast(i);
            if (dq.peekFirst() <= i - k) dq.pollFirst();
            if (i >= k - 1) res[i - k + 1] = nums[dq.peekFirst()];
        }
        return res;
    }
}
// test
check(new Solution().maxSlidingWindow(new int[]{1, 3, -1, -3, 5, 3, 6, 7}, 3), new int[]{3, 3, 5, 5, 6, 7});
```

## 209. Minimum Size Subarray Sum | Medium
related: 76, 560, 862
tags: variable-window

All positive → expand `r`, and while `sum ≥ target` record the length and shrink from `l`.

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        l = s = 0
        best = float("inf")
        for r, x in enumerate(nums):
            s += x
            while s >= target:
                best = min(best, r - l + 1)
                s -= nums[l]
                l += 1
        return 0 if best == float("inf") else best
# test
assert Solution().minSubArrayLen(7, [2, 3, 1, 2, 4, 3]) == 2
assert Solution().minSubArrayLen(11, [1, 1, 1, 1, 1, 1, 1, 1]) == 0
```

```java
class Solution {
    public int minSubArrayLen(int target, int[] nums) {
        int l = 0, s = 0, best = Integer.MAX_VALUE;
        for (int r = 0; r < nums.length; r++) {
            s += nums[r];
            while (s >= target) {
                best = Math.min(best, r - l + 1);
                s -= nums[l++];
            }
        }
        return best == Integer.MAX_VALUE ? 0 : best;
    }
}
// test
check(new Solution().minSubArrayLen(7, new int[]{2, 3, 1, 2, 4, 3}), 2);
```

## 438. Find All Anagrams in a String | Medium
related: 567, 242, 49
tags: fixed-window, counting

Fixed window of size `len(p)`; compare 26-letter counts at every position (O(26) per step), collecting start indices that match.

**Complexity:** O(26·n).

```python
class Solution:
    def findAnagrams(self, s: str, p: str) -> List[int]:
        k, res = len(p), []
        need, win = [0] * 26, [0] * 26
        for ch in p:
            need[ord(ch) - 97] += 1
        for r, ch in enumerate(s):
            win[ord(ch) - 97] += 1
            if r >= k:
                win[ord(s[r - k]) - 97] -= 1
            if win == need:
                res.append(r - k + 1)
        return res
# test
assert Solution().findAnagrams("cbaebabacd", "abc") == [0, 6]
assert Solution().findAnagrams("abab", "ab") == [0, 1, 2]
```

```java
class Solution {
    public List<Integer> findAnagrams(String s, String p) {
        int k = p.length();
        int[] need = new int[26], win = new int[26];
        for (char c : p.toCharArray()) need[c - 'a']++;
        List<Integer> res = new ArrayList<>();
        for (int r = 0; r < s.length(); r++) {
            win[s.charAt(r) - 'a']++;
            if (r >= k) win[s.charAt(r - k) - 'a']--;
            if (Arrays.equals(win, need)) res.add(r - k + 1);
        }
        return res;
    }
}
// test
check(new Solution().findAnagrams("cbaebabacd", "abc"), "[0, 6]");
check(new Solution().findAnagrams("abab", "ab"), "[0, 1, 2]");
```

## 1004. Max Consecutive Ones III | Medium
related: 424, 487
tags: variable-window

Longest window containing at most `k` zeros. Grow `r`; when zeros exceed `k`, move `l` until one zero leaves.

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def longestOnes(self, nums: List[int], k: int) -> int:
        l = zeros = best = 0
        for r, x in enumerate(nums):
            zeros += x == 0
            while zeros > k:
                zeros -= nums[l] == 0
                l += 1
            best = max(best, r - l + 1)
        return best
# test
assert Solution().longestOnes([1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0], 2) == 6
assert Solution().longestOnes([0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1], 3) == 10
```

```java
class Solution {
    public int longestOnes(int[] nums, int k) {
        int l = 0, zeros = 0, best = 0;
        for (int r = 0; r < nums.length; r++) {
            if (nums[r] == 0) zeros++;
            while (zeros > k) if (nums[l++] == 0) zeros--;
            best = Math.max(best, r - l + 1);
        }
        return best;
    }
}
// test
check(new Solution().longestOnes(new int[]{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2), 6);
```

## 713. Subarray Product Less Than K | Medium
related: 209, 560
tags: variable-window, counting

For each right end `r`, shrink until `product < k`; then every subarray ending at `r` and starting in `[l, r]` qualifies → add `r − l + 1`.

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def numSubarrayProductLessThanK(self, nums: List[int], k: int) -> int:
        if k <= 1:
            return 0
        prod, l, count = 1, 0, 0
        for r, x in enumerate(nums):
            prod *= x
            while prod >= k:
                prod //= nums[l]
                l += 1
            count += r - l + 1
        return count
# test
assert Solution().numSubarrayProductLessThanK([10, 5, 2, 6], 100) == 8
assert Solution().numSubarrayProductLessThanK([1, 2, 3], 0) == 0
```

```java
class Solution {
    public int numSubarrayProductLessThanK(int[] nums, int k) {
        if (k <= 1) return 0;
        int prod = 1, l = 0, count = 0;
        for (int r = 0; r < nums.length; r++) {
            prod *= nums[r];
            while (prod >= k) prod /= nums[l++];
            count += r - l + 1;
        }
        return count;
    }
}
// test
check(new Solution().numSubarrayProductLessThanK(new int[]{10, 5, 2, 6}, 100), 8);
```
