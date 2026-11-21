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

```problem
`prices[i]` is a stock's price on day i. Buy on one day and sell on a later day; return the maximum profit, or 0 if no profit is possible.

**Example:** `[7,1,5,3,6,4]` → `5` (buy at 1, sell at 6); `[7,6,4,3,1]` → `0`

**Constraints:** n ≤ 10⁵.
```

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

```go
func maxProfit(prices []int) int {
	lo, best := math.MaxInt, 0
	for _, p := range prices {
		lo = min(lo, p) // cheapest buy so far
		best = max(best, p-lo)
	}
	return best
}
// test
assertEq(maxProfit([]int{7, 1, 5, 3, 6, 4}), 5)
assertEq(maxProfit([]int{7, 6, 4, 3, 1}), 0)
```

## 3. Longest Substring Without Repeating Characters | Medium
related: 424, 76, 159
tags: blind75, variable-window

```problem
Return the length of the longest substring of `s` that contains no repeated character.

**Example:** `"abcabcbb"` → `3` (`"abc"`); `"bbbbb"` → `1`

**Constraints:** length ≤ 5·10⁴; letters, digits, symbols and spaces.
```

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

```go
func lengthOfLongestSubstring(s string) int {
	last := map[byte]int{} // char -> last index
	l, best := 0, 0
	for r := 0; r < len(s); r++ {
		if i, ok := last[s[r]]; ok && i >= l {
			l = i + 1 // jump past the previous copy
		}
		last[s[r]] = r
		best = max(best, r-l+1)
	}
	return best
}
// test
assertEq(lengthOfLongestSubstring("abcabcbb"), 3)
assertEq(lengthOfLongestSubstring("bbbbb"), 1)
assertEq(lengthOfLongestSubstring("pwwkew"), 3)
assertEq(lengthOfLongestSubstring("abba"), 2)
```

## 424. Longest Repeating Character Replacement | Medium
related: 3, 1004
tags: blind75, variable-window

```problem
You may change at most `k` characters of `s` (uppercase letters) to any other uppercase letter. Return the length of the longest substring that can then consist of a single repeated letter.

**Example:** `s = "AABABBA", k = 1` → `4`

**Constraints:** length ≤ 10⁵.
```

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

```go
func characterReplacement(s string, k int) int {
	var cnt [26]int
	l, maxf := 0, 0
	for r := 0; r < len(s); r++ {
		cnt[s[r]-'A']++
		maxf = max(maxf, cnt[s[r]-'A'])
		if r-l+1-maxf > k { // the window never shrinks, it only slides
			cnt[s[l]-'A']--
			l++
		}
	}
	return len(s) - l
}
// test
assertEq(characterReplacement("ABAB", 2), 4)
assertEq(characterReplacement("AABABBA", 1), 4)
```

## 567. Permutation in String | Medium
related: 438, 76, 242
tags: fixed-window, counting

```problem
Decide whether some permutation of `s1` appears as a contiguous substring of `s2`.

**Example:** `s1 = "ab", s2 = "eidbaooo"` → `true` (`"ba"`); `s2 = "eidboaoo"` → `false`

**Constraints:** lowercase letters; lengths ≤ 10⁴.
```

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

```go
func checkInclusion(s1 string, s2 string) bool {
	k := len(s1)
	if k > len(s2) {
		return false
	}
	var need, win [26]int
	for i := 0; i < k; i++ {
		need[s1[i]-'a']++
		win[s2[i]-'a']++
	}
	if need == win {
		return true
	}
	for r := k; r < len(s2); r++ {
		win[s2[r]-'a']++
		win[s2[r-k]-'a']--
		if win == need {
			return true
		}
	}
	return false
}
// test
assertEq(checkInclusion("ab", "eidbaooo"), true)
assertEq(checkInclusion("ab", "eidboaoo"), false)
```

## 76. Minimum Window Substring | Hard
related: 3, 567, 438
tags: blind75, variable-window

```problem
Return the shortest substring of `s` that contains every character of `t`, duplicates included (two `a`s in `t` need two `a`s in the window). Return `""` if there is none; the answer is unique when it exists.

**Example:** `s = "ADOBECODEBANC", t = "ABC"` → `"BANC"`

**Constraints:** lengths ≤ 10⁵. Follow-up: O(m + n).
```

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

```go
func minWindow(s string, t string) string {
	var need [128]int
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	missing, l, start, best := len(t), 0, 0, math.MaxInt
	for r := 0; r < len(s); r++ {
		if need[s[r]] > 0 {
			missing--
		}
		need[s[r]]--
		for missing == 0 { // window covers t: shrink it
			if r-l+1 < best {
				start, best = l, r-l+1
			}
			need[s[l]]++
			if need[s[l]] > 0 {
				missing++
			}
			l++
		}
	}
	if best == math.MaxInt {
		return ""
	}
	return s[start : start+best]
}
// test
assertEq(minWindow("ADOBECODEBANC", "ABC"), "BANC")
assertEq(minWindow("a", "aa"), "")
```

## 239. Sliding Window Maximum | Hard
related: 155, 1438, 862
tags: monotonic-deque

```problem
A window of size `k` slides over `nums` one step at a time from the left. Return the maximum of each window position.

**Example:** `nums = [1,3,-1,-3,5,3,6,7], k = 3` → `[3,3,5,5,6,7]`

**Constraints:** n ≤ 10⁵.
```

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

```go
func maxSlidingWindow(nums []int, k int) []int {
	dq, res := []int{}, []int{} // indices, values decreasing
	for i, x := range nums {
		for len(dq) > 0 && nums[dq[len(dq)-1]] <= x {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
		if dq[0] <= i-k {
			dq = dq[1:]
		}
		if i >= k-1 {
			res = append(res, nums[dq[0]])
		}
	}
	return res
}
// test
assertEq(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3), []int{3, 3, 5, 5, 6, 7})
assertEq(maxSlidingWindow([]int{1}, 1), []int{1})
```

## 209. Minimum Size Subarray Sum | Medium
related: 76, 560, 862
tags: variable-window

```problem
Given positive integers `nums`, return the minimal length of a contiguous subarray whose sum is at least `target`, or 0 if none exists.

**Example:** `target = 7, nums = [2,3,1,2,4,3]` → `2` (`[4,3]`)

**Constraints:** n ≤ 10⁵. Follow-up: also an O(n log n) solution.
```

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

```go
func minSubArrayLen(target int, nums []int) int {
	l, s, best := 0, 0, math.MaxInt
	for r, x := range nums {
		s += x
		for s >= target {
			best = min(best, r-l+1)
			s -= nums[l]
			l++
		}
	}
	if best == math.MaxInt {
		return 0
	}
	return best
}
// test
assertEq(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}), 2)
assertEq(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}), 0)
```

## 438. Find All Anagrams in a String | Medium
related: 567, 242, 49
tags: fixed-window, counting

```problem
Return the start indices of every substring of `s` that is an anagram of `p`, in any order.

**Example:** `s = "cbaebabacd", p = "abc"` → `[0,6]`

**Constraints:** lowercase letters; lengths ≤ 3·10⁴.
```

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
assert sorted(Solution().findAnagrams("cbaebabacd", "abc")) == sorted([0, 6])
assert sorted(Solution().findAnagrams("abab", "ab")) == sorted([0, 1, 2])
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

```go
func findAnagrams(s string, p string) []int {
	k, res := len(p), []int{}
	var need, win [26]int
	for i := 0; i < k; i++ {
		need[p[i]-'a']++
	}
	for r := 0; r < len(s); r++ {
		win[s[r]-'a']++
		if r >= k {
			win[s[r-k]-'a']--
		}
		if win == need {
			res = append(res, r-k+1)
		}
	}
	return res
}
// test
assertUnordered(findAnagrams("cbaebabacd", "abc"), []int{0, 6})
assertUnordered(findAnagrams("abab", "ab"), []int{0, 1, 2})
```

## 1004. Max Consecutive Ones III | Medium
related: 424, 487
tags: variable-window

```problem
`nums` is a binary array. If you may flip at most `k` zeros to ones, what is the longest run of consecutive 1s you can get?

**Example:** `nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2` → `6`

**Constraints:** n ≤ 10⁵.
```

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

```go
func longestOnes(nums []int, k int) int {
	l, zeros, best := 0, 0, 0
	for r, x := range nums {
		if x == 0 {
			zeros++
		}
		for zeros > k {
			if nums[l] == 0 {
				zeros--
			}
			l++
		}
		best = max(best, r-l+1)
	}
	return best
}
// test
assertEq(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2), 6)
assertEq(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3), 10)
```

## 713. Subarray Product Less Than K | Medium
related: 209, 560
tags: variable-window, counting

```problem
Given positive integers `nums`, count the contiguous subarrays whose product is strictly less than `k`.

**Example:** `nums = [10,5,2,6], k = 100` → `8`

**Constraints:** n ≤ 3·10⁴; 1 ≤ nums[i] ≤ 1000; 0 ≤ k ≤ 10⁶.
```

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

```go
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	prod, l, count := 1, 0, 0
	for r, x := range nums {
		prod *= x
		for prod >= k {
			prod /= nums[l]
			l++
		}
		count += r - l + 1 // subarrays ending at r
	}
	return count
}
// test
assertEq(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100), 8)
assertEq(numSubarrayProductLessThanK([]int{1, 2, 3}, 0), 0)
```


## 643. Maximum Average Subarray I | Easy
related: 209
tags: fixed-window

```problem
Return the largest average of any contiguous subarray of exactly length `k`.

**Example:** `nums = [1,12,-5,-6,50,3], k = 4` → `12.75`
```

Fixed window of size k: add the new element, subtract the one leaving, track the max sum.

**Complexity:** O(n).

```python
class Solution:
    def findMaxAverage(self, nums: List[int], k: int) -> float:
        s = best = sum(nums[:k])
        for i in range(k, len(nums)):
            s += nums[i] - nums[i - k]
            best = max(best, s)
        return best / k
# test
assert Solution().findMaxAverage([1, 12, -5, -6, 50, 3], 4) == 12.75
assert Solution().findMaxAverage([5], 1) == 5.0
assert Solution().findMaxAverage([0, 4, 0, 3, 2], 1) == 4.0
assert Solution().findMaxAverage([-1, -2, -3], 3) == -2.0
```

```java
class Solution {
    public double findMaxAverage(int[] nums, int k) {
        int s = 0;
        for (int i = 0; i < k; i++) s += nums[i];
        int best = s;
        for (int i = k; i < nums.length; i++) { s += nums[i] - nums[i - k]; best = Math.max(best, s); }
        return (double) best / k;
    }
}
// test
check(new Solution().findMaxAverage(new int[]{1, 12, -5, -6, 50, 3}, 4), 12.75);
```

```go
func findMaxAverage(nums []int, k int) float64 {
	s := 0
	for _, x := range nums[:k] {
		s += x
	}
	best := s
	for i := k; i < len(nums); i++ {
		s += nums[i] - nums[i-k]
		best = max(best, s)
	}
	return float64(best) / float64(k)
}
// test
assertEq(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4), 12.75)
assertEq(findMaxAverage([]int{5}, 1), 5.0)
assertEq(findMaxAverage([]int{0, 4, 0, 3, 2}, 1), 4.0)
assertEq(findMaxAverage([]int{-1, -2, -3}, 3), -2.0)
```

## 904. Fruit Into Baskets | Medium
related: 3, 992
tags: variable-window

```problem
Trees stand in a row; `fruits[i]` is the type of fruit tree i grows. With two baskets that each hold one fruit type (any amount), start at any tree and pick one fruit from every tree moving right, stopping when a fruit fits no basket. Return the most fruit you can pick, i.e. the longest subarray with at most two distinct values.

**Example:** `[1,2,3,2,2]` → `4` (`[2,3,2,2]`)

**Constraints:** n ≤ 10⁵.
```

= longest window with **at most 2 distinct** values. Count fruits in the window; shrink while more than 2 types.

**Complexity:** O(n).

```python
class Solution:
    def totalFruit(self, fruits: List[int]) -> int:
        cnt = defaultdict(int)
        l = best = 0
        for r, f in enumerate(fruits):
            cnt[f] += 1
            while len(cnt) > 2:
                cnt[fruits[l]] -= 1
                if cnt[fruits[l]] == 0:
                    del cnt[fruits[l]]
                l += 1
            best = max(best, r - l + 1)
        return best
# test
assert Solution().totalFruit([1, 2, 1]) == 3
assert Solution().totalFruit([0, 1, 2, 2]) == 3
assert Solution().totalFruit([1, 2, 3, 2, 2]) == 4
```

```java
class Solution {
    public int totalFruit(int[] fruits) {
        Map<Integer, Integer> cnt = new HashMap<>();
        int l = 0, best = 0;
        for (int r = 0; r < fruits.length; r++) {
            cnt.merge(fruits[r], 1, Integer::sum);
            while (cnt.size() > 2) {
                if (cnt.merge(fruits[l], -1, Integer::sum) == 0) cnt.remove(fruits[l]);
                l++;
            }
            best = Math.max(best, r - l + 1);
        }
        return best;
    }
}
// test
check(new Solution().totalFruit(new int[]{1, 2, 3, 2, 2}), 4);
```

```go
func totalFruit(fruits []int) int {
	cnt := map[int]int{}
	l, best := 0, 0
	for r, f := range fruits {
		cnt[f]++
		for len(cnt) > 2 {
			cnt[fruits[l]]--
			if cnt[fruits[l]] == 0 {
				delete(cnt, fruits[l])
			}
			l++
		}
		best = max(best, r-l+1)
	}
	return best
}
// test
assertEq(totalFruit([]int{1, 2, 1}), 3)
assertEq(totalFruit([]int{0, 1, 2, 2}), 3)
assertEq(totalFruit([]int{1, 2, 3, 2, 2}), 4)
```

## 992. Subarrays with K Different Integers | Hard
related: 904, 713
tags: variable-window, at-most-trick

```problem
Count the contiguous subarrays of `nums` that contain exactly `k` distinct values.

**Example:** `nums = [1,2,1,2,3], k = 2` → `7`

**Constraints:** n ≤ 2·10⁴.
```

"Exactly k" is not monotonic, but "**at most** k" is: `exactly(k) = atMost(k) − atMost(k−1)`. `atMost` counts, for each right end, the `r − l + 1` valid windows.

**Complexity:** O(n).

```python
class Solution:
    def subarraysWithKDistinct(self, nums: List[int], k: int) -> int:
        def at_most(k):
            cnt = defaultdict(int)
            l = res = 0
            for r, x in enumerate(nums):
                cnt[x] += 1
                while len(cnt) > k:
                    cnt[nums[l]] -= 1
                    if cnt[nums[l]] == 0:
                        del cnt[nums[l]]
                    l += 1
                res += r - l + 1
            return res

        return at_most(k) - at_most(k - 1)
# test
assert Solution().subarraysWithKDistinct([1, 2, 1, 2, 3], 2) == 7
assert Solution().subarraysWithKDistinct([1, 2, 1, 3, 4], 3) == 3
```

```java
class Solution {
    public int subarraysWithKDistinct(int[] nums, int k) {
        return atMost(nums, k) - atMost(nums, k - 1);
    }

    private int atMost(int[] nums, int k) {
        int[] cnt = new int[nums.length + 1];
        int l = 0, distinct = 0, res = 0;
        for (int r = 0; r < nums.length; r++) {
            if (cnt[nums[r]]++ == 0) distinct++;
            while (distinct > k) if (--cnt[nums[l++]] == 0) distinct--;
            res += r - l + 1;
        }
        return res;
    }
}
// test
check(new Solution().subarraysWithKDistinct(new int[]{1, 2, 1, 2, 3}, 2), 7);
```

```go
func subarraysWithKDistinct(nums []int, k int) int {
	atMost := func(k int) int { // subarrays with at most k distinct values
		cnt := map[int]int{}
		l, res := 0, 0
		for r, x := range nums {
			cnt[x]++
			for len(cnt) > k {
				cnt[nums[l]]--
				if cnt[nums[l]] == 0 {
					delete(cnt, nums[l])
				}
				l++
			}
			res += r - l + 1
		}
		return res
	}
	return atMost(k) - atMost(k-1)
}
// test
assertEq(subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2), 7)
assertEq(subarraysWithKDistinct([]int{1, 2, 1, 3, 4}, 3), 3)
```

## 30. Substring with Concatenation of All Words | Hard
related: 76, 438
tags: fixed-window, hash-map

```problem
`words` are strings of equal length. Return every start index in `s` of a substring that is some concatenation of all the words, each used exactly once, in any order.

**Example:** `s = "barfoothefoobarman", words = ["foo","bar"]` → `[0,9]`

**Constraints:** length of `s` ≤ 10⁴; up to 5000 words of length ≤ 30.
```

All words have length `L`. For each of the `L` offsets, slide a window **word by word**, counting words; when a word is over-used, drop words from the left. A window with exactly `len(words)` words is a match.

**Complexity:** O(L · n).

```python
class Solution:
    def findSubstring(self, s: str, words: List[str]) -> List[int]:
        L, k = len(words[0]), len(words)
        need = Counter(words)
        res = []
        for off in range(L):
            l, cnt, used = off, defaultdict(int), 0
            for r in range(off, len(s) - L + 1, L):
                w = s[r:r + L]
                if w not in need:
                    l, cnt, used = r + L, defaultdict(int), 0
                    continue
                cnt[w] += 1
                used += 1
                while cnt[w] > need[w]:
                    cnt[s[l:l + L]] -= 1
                    used -= 1
                    l += L
                if used == k:
                    res.append(l)
        return sorted(res)
# test
assert sorted(Solution().findSubstring("barfoothefoobarman", ["foo", "bar"])) == sorted([0, 9])
assert sorted(Solution().findSubstring("wordgoodgoodgoodbestword", ["word", "good", "best", "word"])) == sorted([])
assert sorted(Solution().findSubstring("barfoofoobarthefoobarman", ["bar", "foo", "the"])) == sorted([6, 9, 12])
```

```java
class Solution {
    public List<Integer> findSubstring(String s, String[] words) {
        int L = words[0].length(), k = words.length;
        Map<String, Integer> need = new HashMap<>();
        for (String w : words) need.merge(w, 1, Integer::sum);
        List<Integer> res = new ArrayList<>();
        for (int off = 0; off < L; off++) {
            Map<String, Integer> cnt = new HashMap<>();
            int l = off, used = 0;
            for (int r = off; r + L <= s.length(); r += L) {
                String w = s.substring(r, r + L);
                if (!need.containsKey(w)) { cnt.clear(); used = 0; l = r + L; continue; }
                cnt.merge(w, 1, Integer::sum);
                used++;
                while (cnt.get(w) > need.get(w)) {
                    cnt.merge(s.substring(l, l + L), -1, Integer::sum);
                    used--;
                    l += L;
                }
                if (used == k) res.add(l);
            }
        }
        Collections.sort(res);
        return res;
    }
}
// test
check(new Solution().findSubstring("barfoofoobarthefoobarman", new String[]{"bar", "foo", "the"}), "[6, 9, 12]");
```

```go
func findSubstring(s string, words []string) []int {
	L, k := len(words[0]), len(words)
	need := map[string]int{}
	for _, w := range words {
		need[w]++
	}
	res := []int{}
	for off := 0; off < L; off++ { // one sliding window per alignment
		l, used := off, 0
		cnt := map[string]int{}
		for r := off; r+L <= len(s); r += L {
			w := s[r : r+L]
			if need[w] == 0 {
				l, used, cnt = r+L, 0, map[string]int{}
				continue
			}
			cnt[w]++
			used++
			for cnt[w] > need[w] {
				cnt[s[l:l+L]]--
				used--
				l += L
			}
			if used == k {
				res = append(res, l)
			}
		}
	}
	sort.Ints(res)
	return res
}
// test
assertUnordered(findSubstring("barfoothefoobarman", []string{"foo", "bar"}), []int{0, 9})
assertEq(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}), []int{})
assertUnordered(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}), []int{6, 9, 12})
```

## 1423. Maximum Points You Can Obtain from Cards | Medium
related: 209
tags: fixed-window, complement

```problem
Cards lie in a row with `cardPoints[i]` points. Take exactly `k` cards, each from either the left or the right end of the row. Return the maximum total points.

**Example:** `cardPoints = [1,2,3,4,5,6,1], k = 3` → `12` (take 6, 5, 1 from the right)

**Constraints:** n ≤ 10⁵.
```

Taking k cards from the ends leaves a **contiguous** window of `n − k` cards in the middle. Minimize that window's sum; answer = total − min window.

**Complexity:** O(n).

```python
class Solution:
    def maxScore(self, cardPoints: List[int], k: int) -> int:
        n, w = len(cardPoints), len(cardPoints) - k
        s = best = sum(cardPoints[:w])
        for i in range(w, n):
            s += cardPoints[i] - cardPoints[i - w]
            best = min(best, s)
        return sum(cardPoints) - best
# test
assert Solution().maxScore([1, 2, 3, 4, 5, 6, 1], 3) == 12
assert Solution().maxScore([9, 7, 7, 9, 7, 7, 9], 7) == 55
```

```java
class Solution {
    public int maxScore(int[] cardPoints, int k) {
        int n = cardPoints.length, w = n - k, s = 0, total = 0;
        for (int x : cardPoints) total += x;
        for (int i = 0; i < w; i++) s += cardPoints[i];
        int best = s;
        for (int i = w; i < n; i++) { s += cardPoints[i] - cardPoints[i - w]; best = Math.min(best, s); }
        return total - best;
    }
}
// test
check(new Solution().maxScore(new int[]{1, 2, 3, 4, 5, 6, 1}, 3), 12);
```

```go
func maxScore(cardPoints []int, k int) int {
	n, w := len(cardPoints), len(cardPoints)-k // minimise the middle window of size n-k
	total, s := 0, 0
	for i, x := range cardPoints {
		total += x
		if i < w {
			s += x
		}
	}
	best := s
	for i := w; i < n; i++ {
		s += cardPoints[i] - cardPoints[i-w]
		best = min(best, s)
	}
	return total - best
}
// test
assertEq(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3), 12)
assertEq(maxScore([]int{9, 7, 7, 9, 7, 7, 9}, 7), 55)
```

## 1838. Frequency of the Most Frequent Element | Medium
related: 424
tags: sort, variable-window

```problem
One operation adds 1 to any element of `nums`. With at most `k` operations, return the highest possible frequency of a single value.

**Example:** `nums = [1,2,4], k = 5` → `3` (raise 1 and 2 to 4); `nums = [1,4,8,13], k = 5` → `2`

**Constraints:** n ≤ 10⁵.
```

Sort. A window `[l, r]` can all be raised to `nums[r]` iff `nums[r]·len − sum ≤ k`. Shrink when it isn't.

**Complexity:** O(n log n).

```python
class Solution:
    def maxFrequency(self, nums: List[int], k: int) -> int:
        nums.sort()
        l = s = best = 0
        for r, x in enumerate(nums):
            s += x
            while x * (r - l + 1) - s > k:
                s -= nums[l]
                l += 1
            best = max(best, r - l + 1)
        return best
# test
assert Solution().maxFrequency([1, 2, 4], 5) == 3
assert Solution().maxFrequency([1, 4, 8, 13], 5) == 2
```

```java
class Solution {
    public int maxFrequency(int[] nums, int k) {
        Arrays.sort(nums);
        long s = 0;
        int l = 0, best = 0;
        for (int r = 0; r < nums.length; r++) {
            s += nums[r];
            while ((long) nums[r] * (r - l + 1) - s > k) s -= nums[l++];
            best = Math.max(best, r - l + 1);
        }
        return best;
    }
}
// test
check(new Solution().maxFrequency(new int[]{1, 4, 8, 13}, 5), 2);
```

```go
func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	l, s, best := 0, 0, 0
	for r, x := range nums {
		s += x
		for x*(r-l+1)-s > k { // cost to raise the window to x
			s -= nums[l]
			l++
		}
		best = max(best, r-l+1)
	}
	return best
}
// test
assertEq(maxFrequency([]int{1, 2, 4}, 5), 3)
assertEq(maxFrequency([]int{1, 4, 8, 13}, 5), 2)
```

## 219. Contains Duplicate II | Easy
related: 217, 220
tags: fixed-window, hash-set

```problem
Decide whether there are two indices `i ≠ j` with `nums[i] == nums[j]` and |i − j| ≤ k.

**Example:** `nums = [1,2,3,1], k = 3` → `true`; `nums = [1,2,3,1,2,3], k = 2` → `false`
```

Keep a set of the **last k values**. A value already in the set is a duplicate within distance k. Evict `nums[i − k]` once the window holds more than k values.

**Complexity:** O(n) time, O(k) space.

```python
class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        window = set()
        for i, x in enumerate(nums):
            if x in window:
                return True
            window.add(x)
            if len(window) > k:
                window.remove(nums[i - k])
        return False
# test
assert Solution().containsNearbyDuplicate([1, 2, 3, 1], 3) is True
assert Solution().containsNearbyDuplicate([1, 0, 1, 1], 1) is True
assert Solution().containsNearbyDuplicate([1, 2, 3, 1, 2, 3], 2) is False
assert Solution().containsNearbyDuplicate([1, 1], 0) is False
```

```java
class Solution {
    public boolean containsNearbyDuplicate(int[] nums, int k) {
        Set<Integer> window = new HashSet<>();
        for (int i = 0; i < nums.length; i++) {
            if (!window.add(nums[i])) return true;
            if (window.size() > k) window.remove(nums[i - k]);
        }
        return false;
    }
}
// test
check(new Solution().containsNearbyDuplicate(new int[]{1, 0, 1, 1}, 1), true);
check(new Solution().containsNearbyDuplicate(new int[]{1, 2, 3, 1, 2, 3}, 2), false);
```

```go
func containsNearbyDuplicate(nums []int, k int) bool {
	window := map[int]bool{} // the last k values
	for i, x := range nums {
		if window[x] {
			return true
		}
		window[x] = true
		if len(window) > k {
			delete(window, nums[i-k])
		}
	}
	return false
}
// test
assertEq(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3), true)
assertEq(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1), true)
assertEq(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2), false)
assertEq(containsNearbyDuplicate([]int{1, 1}, 0), false)
```

## 1456. Maximum Number of Vowels in a Substring of Given Length | Medium
related: 643, 1004
tags: fixed-window

```problem
Return the maximum number of vowels in any substring of `s` of length exactly `k`.

**Example:** `s = "abciiidef", k = 3` → `3` (`"iii"`)
```

Fixed window of size k: add the entering character, subtract the leaving one, track the maximum.

**Complexity:** O(n).

```python
class Solution:
    def maxVowels(self, s: str, k: int) -> int:
        vowels = set("aeiou")
        cur = best = sum(c in vowels for c in s[:k])
        for i in range(k, len(s)):
            cur += (s[i] in vowels) - (s[i - k] in vowels)
            best = max(best, cur)
        return best
# test
assert Solution().maxVowels("abciiidef", 3) == 3
assert Solution().maxVowels("aeiou", 2) == 2
assert Solution().maxVowels("leetcode", 3) == 2
assert Solution().maxVowels("tryhard", 4) == 1
```

```java
class Solution {
    public int maxVowels(String s, int k) {
        int cur = 0, best = 0;
        for (int i = 0; i < s.length(); i++) {
            if ("aeiou".indexOf(s.charAt(i)) >= 0) cur++;
            if (i >= k && "aeiou".indexOf(s.charAt(i - k)) >= 0) cur--;
            best = Math.max(best, cur);
        }
        return best;
    }
}
// test
check(new Solution().maxVowels("abciiidef", 3), 3);
check(new Solution().maxVowels("leetcode", 3), 2);
```

```go
func maxVowels(s string, k int) int {
	vowel := func(c byte) int {
		if strings.IndexByte("aeiou", c) >= 0 {
			return 1
		}
		return 0
	}
	cur := 0
	for i := 0; i < k; i++ {
		cur += vowel(s[i])
	}
	best := cur
	for i := k; i < len(s); i++ {
		cur += vowel(s[i]) - vowel(s[i-k])
		best = max(best, cur)
	}
	return best
}
// test
assertEq(maxVowels("abciiidef", 3), 3)
assertEq(maxVowels("aeiou", 2), 2)
assertEq(maxVowels("leetcode", 3), 2)
assertEq(maxVowels("tryhard", 4), 1)
```

## 1493. Longest Subarray of 1's After Deleting One Element | Medium
related: 1004, 487
tags: variable-window

```problem
Delete exactly one element from the binary array `nums`, then return the length of the longest non-empty subarray of only 1s (0 if there is none).

**Example:** `[0,1,1,1,0,1,1,0,1]` → `5`; `[1,1,1]` → `2` (you must delete one)
```

Longest window containing **at most one zero**; the deleted element is inside it, so the answer is `window length − 1` (which also handles the all-ones case, where a 1 must be deleted).

**Complexity:** O(n).

```python
class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        lo = zeros = best = 0
        for hi, x in enumerate(nums):
            zeros += x == 0
            while zeros > 1:
                zeros -= nums[lo] == 0
                lo += 1
            best = max(best, hi - lo)          # (hi - lo + 1) - 1 deleted element
        return best
# test
assert Solution().longestSubarray([1, 1, 0, 1]) == 3
assert Solution().longestSubarray([0, 1, 1, 1, 0, 1, 1, 0, 1]) == 5
assert Solution().longestSubarray([1, 1, 1]) == 2
assert Solution().longestSubarray([0, 0, 0]) == 0
```

```java
class Solution {
    public int longestSubarray(int[] nums) {
        int lo = 0, zeros = 0, best = 0;
        for (int hi = 0; hi < nums.length; hi++) {
            if (nums[hi] == 0) zeros++;
            while (zeros > 1) if (nums[lo++] == 0) zeros--;
            best = Math.max(best, hi - lo);
        }
        return best;
    }
}
// test
check(new Solution().longestSubarray(new int[]{0, 1, 1, 1, 0, 1, 1, 0, 1}), 5);
check(new Solution().longestSubarray(new int[]{1, 1, 1}), 2);
```

```go
func longestSubarray(nums []int) int {
	lo, zeros, best := 0, 0, 0
	for hi, x := range nums {
		if x == 0 {
			zeros++
		}
		for zeros > 1 {
			if nums[lo] == 0 {
				zeros--
			}
			lo++
		}
		best = max(best, hi-lo) // (hi - lo + 1) - 1 deleted element
	}
	return best
}
// test
assertEq(longestSubarray([]int{1, 1, 0, 1}), 3)
assertEq(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}), 5)
assertEq(longestSubarray([]int{1, 1, 1}), 2)
assertEq(longestSubarray([]int{0, 0, 0}), 0)
```

## 978. Longest Turbulent Subarray | Medium
related: 376, 53
tags: kadane, state-machine

```problem
A subarray is turbulent if the comparison sign flips between every pair of adjacent elements (bigger, smaller, bigger… or the reverse). Return the length of the longest turbulent subarray of `arr`.

**Example:** `[9,4,2,10,7,8,8,1,9]` → `5` (`[4,2,10,7,8]`)

**Constraints:** n ≤ 4·10⁴.
```

Track the length of the turbulent run ending here whose **last step goes up** and whose **last step goes down**. An up step extends a run that last went down (and vice versa); an equal pair resets both to 1.

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def maxTurbulenceSize(self, arr: List[int]) -> int:
        best = up = down = 1
        for i in range(1, len(arr)):
            if arr[i] > arr[i - 1]:
                up, down = down + 1, 1
            elif arr[i] < arr[i - 1]:
                up, down = 1, up + 1
            else:
                up = down = 1
            best = max(best, up, down)
        return best
# test
assert Solution().maxTurbulenceSize([9, 4, 2, 10, 7, 8, 8, 1, 9]) == 5
assert Solution().maxTurbulenceSize([4, 8, 12, 16]) == 2
assert Solution().maxTurbulenceSize([100]) == 1
assert Solution().maxTurbulenceSize([9, 9]) == 1
```

```java
class Solution {
    public int maxTurbulenceSize(int[] arr) {
        int best = 1, up = 1, down = 1;
        for (int i = 1; i < arr.length; i++) {
            if (arr[i] > arr[i - 1]) { up = down + 1; down = 1; }
            else if (arr[i] < arr[i - 1]) { down = up + 1; up = 1; }
            else up = down = 1;
            best = Math.max(best, Math.max(up, down));
        }
        return best;
    }
}
// test
check(new Solution().maxTurbulenceSize(new int[]{9, 4, 2, 10, 7, 8, 8, 1, 9}), 5);
```

```go
func maxTurbulenceSize(arr []int) int {
	best, up, down := 1, 1, 1 // longest run ending here going up / down
	for i := 1; i < len(arr); i++ {
		switch {
		case arr[i] > arr[i-1]:
			up, down = down+1, 1
		case arr[i] < arr[i-1]:
			up, down = 1, up+1
		default:
			up, down = 1, 1
		}
		best = max(best, up, down)
	}
	return best
}
// test
assertEq(maxTurbulenceSize([]int{9, 4, 2, 10, 7, 8, 8, 1, 9}), 5)
assertEq(maxTurbulenceSize([]int{4, 8, 12, 16}), 2)
assertEq(maxTurbulenceSize([]int{100}), 1)
assertEq(maxTurbulenceSize([]int{9, 9}), 1)
```
