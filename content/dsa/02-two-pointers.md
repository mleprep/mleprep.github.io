# Two Pointers

**When to use:** sorted arrays, pairs/triplets with a target sum, palindromes, in-place partitioning.

- **Opposite ends** (`l=0, r=n-1`): move the side that can only improve the answer (sorted sum search, container area, palindrome check).
- **Same direction** (slow/fast): slow marks the "kept" prefix, fast scans (remove/move elements in place).
- Sorting first often turns an O(n²) search into O(n log n).

```python
l, r = 0, len(a) - 1
while l < r:
    s = a[l] + a[r]
    if s == target: return [l, r]
    if s < target: l += 1     # need bigger
    else: r -= 1              # need smaller
```

## 125. Valid Palindrome | Easy
related: 680, 5
tags: blind75

Pointers from both ends, skip non-alphanumerics, compare lowercase.

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def isPalindrome(self, s: str) -> bool:
        l, r = 0, len(s) - 1
        while l < r:
            if not s[l].isalnum():
                l += 1
            elif not s[r].isalnum():
                r -= 1
            elif s[l].lower() != s[r].lower():
                return False
            else:
                l, r = l + 1, r - 1
        return True
# test
assert Solution().isPalindrome("A man, a plan, a canal: Panama") is True
assert Solution().isPalindrome("race a car") is False
```

```java
class Solution {
    public boolean isPalindrome(String s) {
        int l = 0, r = s.length() - 1;
        while (l < r) {
            char a = s.charAt(l), b = s.charAt(r);
            if (!Character.isLetterOrDigit(a)) l++;
            else if (!Character.isLetterOrDigit(b)) r--;
            else if (Character.toLowerCase(a) != Character.toLowerCase(b)) return false;
            else { l++; r--; }
        }
        return true;
    }
}
// test
check(new Solution().isPalindrome("A man, a plan, a canal: Panama"), true);
check(new Solution().isPalindrome("race a car"), false);
```

## 167. Two Sum II - Input Array Is Sorted | Medium
related: 1, 15
tags: opposite-ends

Sorted → if the sum is too small move `l` right, too big move `r` left. Each step discards one element that can't be in any answer.

```viz
[2, 7, 11, 15]  target 9
 l          r   2+15=17 > 9  -> r--
 l      r       2+11=13 > 9  -> r--
 l  r           2+7 = 9      -> [1, 2] (1-indexed)
```

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        l, r = 0, len(numbers) - 1
        while l < r:
            s = numbers[l] + numbers[r]
            if s == target:
                return [l + 1, r + 1]
            if s < target:
                l += 1
            else:
                r -= 1
        return []
# test
assert Solution().twoSum([2, 7, 11, 15], 9) == [1, 2]
assert Solution().twoSum([-1, 0], -1) == [1, 2]
```

```java
class Solution {
    public int[] twoSum(int[] numbers, int target) {
        int l = 0, r = numbers.length - 1;
        while (l < r) {
            int s = numbers[l] + numbers[r];
            if (s == target) return new int[]{l + 1, r + 1};
            if (s < target) l++; else r--;
        }
        return new int[0];
    }
}
// test
check(new Solution().twoSum(new int[]{2, 7, 11, 15}, 9), new int[]{1, 2});
```

## 15. 3Sum | Medium
related: 1, 167, 16, 18
tags: blind75, sort

Sort. Fix `nums[i]`, then run Two Sum II on the suffix for `-nums[i]`. Skip equal neighbours (for `i` and after a hit) to avoid duplicate triplets. Stop early when `nums[i] > 0`.

```viz
sorted: [-4, -1, -1, 0, 1, 2]
i=-1:    l=-1 ............ r=2   -1+(-1)+2 = 0  ✓ -> [-1,-1,2]
         l=0 ......r=1           -1+0+1   = 0  ✓ -> [-1,0,1]
i=-1 (dup) skip
```

**Complexity:** O(n²) time, O(1) extra (besides sort).

```python
class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        res = []
        for i in range(len(nums) - 2):
            if nums[i] > 0:
                break
            if i and nums[i] == nums[i - 1]:
                continue
            l, r = i + 1, len(nums) - 1
            while l < r:
                s = nums[i] + nums[l] + nums[r]
                if s < 0:
                    l += 1
                elif s > 0:
                    r -= 1
                else:
                    res.append([nums[i], nums[l], nums[r]])
                    l += 1
                    while l < r and nums[l] == nums[l - 1]:
                        l += 1
                    r -= 1
        return res
# test
assert sorted(map(sorted, Solution().threeSum([-1, 0, 1, 2, -1, -4]))) == sorted(map(sorted, [[-1, -1, 2], [-1, 0, 1]]))
assert sorted(map(sorted, Solution().threeSum([0, 0, 0]))) == sorted(map(sorted, [[0, 0, 0]]))
```

```java
class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);
        List<List<Integer>> res = new ArrayList<>();
        for (int i = 0; i < nums.length - 2 && nums[i] <= 0; i++) {
            if (i > 0 && nums[i] == nums[i - 1]) continue;
            int l = i + 1, r = nums.length - 1;
            while (l < r) {
                int s = nums[i] + nums[l] + nums[r];
                if (s < 0) l++;
                else if (s > 0) r--;
                else {
                    res.add(Arrays.asList(nums[i], nums[l], nums[r]));
                    l++;
                    while (l < r && nums[l] == nums[l - 1]) l++;
                    r--;
                }
            }
        }
        return res;
    }
}
// test
check(new Solution().threeSum(new int[]{-1, 0, 1, 2, -1, -4}), "[[-1, -1, 2], [-1, 0, 1]]");
check(new Solution().threeSum(new int[]{0, 0, 0}), "[[0, 0, 0]]");
```

## 11. Container With Most Water | Medium
related: 42
tags: blind75, opposite-ends, greedy

Area = `min(h[l], h[r]) * (r - l)`. Moving the **taller** side can never help (width shrinks, height capped by the shorter one), so always move the shorter side.

```viz
h: 1 8 6 2 5 4 8 3 7
     l             r     min(8,7)*7 = 49  (best)
```

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def maxArea(self, height: List[int]) -> int:
        l, r, best = 0, len(height) - 1, 0
        while l < r:
            best = max(best, min(height[l], height[r]) * (r - l))
            if height[l] < height[r]:
                l += 1
            else:
                r -= 1
        return best
# test
assert Solution().maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7]) == 49
assert Solution().maxArea([1, 1]) == 1
```

```java
class Solution {
    public int maxArea(int[] height) {
        int l = 0, r = height.length - 1, best = 0;
        while (l < r) {
            best = Math.max(best, Math.min(height[l], height[r]) * (r - l));
            if (height[l] < height[r]) l++; else r--;
        }
        return best;
    }
}
// test
check(new Solution().maxArea(new int[]{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49);
```

## 42. Trapping Rain Water | Hard
related: 11, 84, 238
tags: opposite-ends, prefix-max

Water at `i` = `min(maxLeft, maxRight) - h[i]`. With two pointers, the side with the smaller running max is the bottleneck, so its water is already determined: process that side and move inward.

```viz
                        █
                █ ~ ~ ~ █ █ ~ █
            █ ~ █ █ ~ █ █ █ █ █ █    (~ = trapped water)
height    0 1 0 2 1 0 1 3 2 1 2 1
water     . . 1 . 1 2 1 . . 1 . .    total = 6
```

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def trap(self, height: List[int]) -> int:
        l, r = 0, len(height) - 1
        lmax = rmax = water = 0
        while l < r:
            if height[l] < height[r]:
                lmax = max(lmax, height[l])
                water += lmax - height[l]
                l += 1
            else:
                rmax = max(rmax, height[r])
                water += rmax - height[r]
                r -= 1
        return water
# test
assert Solution().trap([0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]) == 6
assert Solution().trap([4, 2, 0, 3, 2, 5]) == 9
```

```java
class Solution {
    public int trap(int[] height) {
        int l = 0, r = height.length - 1, lmax = 0, rmax = 0, water = 0;
        while (l < r) {
            if (height[l] < height[r]) {
                lmax = Math.max(lmax, height[l]);
                water += lmax - height[l++];
            } else {
                rmax = Math.max(rmax, height[r]);
                water += rmax - height[r--];
            }
        }
        return water;
    }
}
// test
check(new Solution().trap(new int[]{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), 6);
check(new Solution().trap(new int[]{4, 2, 0, 3, 2, 5}), 9);
```

## 283. Move Zeroes | Easy
related: 26, 75
tags: fast-slow

Slow pointer `k` = next slot for a non-zero. Swap each non-zero into `k`; zeros drift to the end while relative order is kept.

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        k = 0
        for i in range(len(nums)):
            if nums[i]:
                nums[k], nums[i] = nums[i], nums[k]
                k += 1
# test
a = [0, 1, 0, 3, 12]; Solution().moveZeroes(a)
assert a == [1, 3, 12, 0, 0]
a = [0]; Solution().moveZeroes(a)
assert a == [0]
a = [1, 0, 0, 2]; Solution().moveZeroes(a)
assert a == [1, 2, 0, 0]
```

```java
class Solution {
    public void moveZeroes(int[] nums) {
        int k = 0;
        for (int i = 0; i < nums.length; i++)
            if (nums[i] != 0) { int t = nums[k]; nums[k++] = nums[i]; nums[i] = t; }
    }
}
// test
int[] a = {0, 1, 0, 3, 12};
new Solution().moveZeroes(a);
check(a, new int[]{1, 3, 12, 0, 0});
```

## 16. 3Sum Closest | Medium
related: 15
tags: sort

Same skeleton as 3Sum: sort, fix `i`, two pointers on the rest; track the sum with the smallest `|sum - target|`.

**Complexity:** O(n²).

```python
class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        nums.sort()
        best = nums[0] + nums[1] + nums[2]
        for i in range(len(nums) - 2):
            l, r = i + 1, len(nums) - 1
            while l < r:
                s = nums[i] + nums[l] + nums[r]
                if abs(s - target) < abs(best - target):
                    best = s
                if s < target:
                    l += 1
                elif s > target:
                    r -= 1
                else:
                    return s
        return best
# test
assert Solution().threeSumClosest([-1, 2, 1, -4], 1) == 2
assert Solution().threeSumClosest([0, 0, 0], 1) == 0
```

```java
class Solution {
    public int threeSumClosest(int[] nums, int target) {
        Arrays.sort(nums);
        int best = nums[0] + nums[1] + nums[2];
        for (int i = 0; i < nums.length - 2; i++) {
            int l = i + 1, r = nums.length - 1;
            while (l < r) {
                int s = nums[i] + nums[l] + nums[r];
                if (Math.abs(s - target) < Math.abs(best - target)) best = s;
                if (s < target) l++;
                else if (s > target) r--;
                else return s;
            }
        }
        return best;
    }
}
// test
check(new Solution().threeSumClosest(new int[]{-1, 2, 1, -4}, 1), 2);
```

## 18. 4Sum | Medium
related: 15, 1
tags: sort, k-sum

Generalize: **kSum** recursion reduces to 2Sum with two pointers. Skip duplicates at every level. Use `long` sums in Java to avoid overflow.

**Complexity:** O(n³) for 4Sum (O(n^(k-1)) for kSum).

```python
class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        nums.sort()

        def ksum(start, k, target):
            res = []
            if k == 2:
                l, r = start, len(nums) - 1
                while l < r:
                    s = nums[l] + nums[r]
                    if s < target or (l > start and nums[l] == nums[l - 1]):
                        l += 1
                    elif s > target:
                        r -= 1
                    else:
                        res.append([nums[l], nums[r]])
                        l += 1; r -= 1
                return res
            for i in range(start, len(nums) - k + 1):
                if i > start and nums[i] == nums[i - 1]:
                    continue
                for sub in ksum(i + 1, k - 1, target - nums[i]):
                    res.append([nums[i]] + sub)
            return res

        return ksum(0, 4, target)
# test
assert sorted(map(sorted, Solution().fourSum([1, 0, -1, 0, -2, 2], 0))) == sorted(map(sorted, [[-2, -1, 1, 2], [-2, 0, 0, 2], [-1, 0, 0, 1]]))
assert sorted(map(sorted, Solution().fourSum([2, 2, 2, 2, 2], 8))) == sorted(map(sorted, [[2, 2, 2, 2]]))
```

```java
class Solution {
    public List<List<Integer>> fourSum(int[] nums, int target) {
        Arrays.sort(nums);
        List<List<Integer>> res = new ArrayList<>();
        int n = nums.length;
        for (int i = 0; i < n - 3; i++) {
            if (i > 0 && nums[i] == nums[i - 1]) continue;
            for (int j = i + 1; j < n - 2; j++) {
                if (j > i + 1 && nums[j] == nums[j - 1]) continue;
                int l = j + 1, r = n - 1;
                while (l < r) {
                    long s = (long) nums[i] + nums[j] + nums[l] + nums[r];
                    if (s < target) l++;
                    else if (s > target) r--;
                    else {
                        res.add(Arrays.asList(nums[i], nums[j], nums[l], nums[r]));
                        l++;
                        while (l < r && nums[l] == nums[l - 1]) l++;
                        r--;
                    }
                }
            }
        }
        return res;
    }
}
// test
check(new Solution().fourSum(new int[]{1, 0, -1, 0, -2, 2}, 0), "[[-2, -1, 1, 2], [-2, 0, 0, 2], [-1, 0, 0, 1]]");
```

## 75. Sort Colors | Medium
related: 283
tags: dutch-flag, three-pointers

**Dutch national flag**: `lo` = next slot for 0, `hi` = next slot for 2, `i` scans. Swap 0s to the front, 2s to the back (don't advance `i` after swapping with `hi`, since the swapped-in value is unexamined).

```viz
[2 0 2 1 1 0]
 lo/i       hi   2 -> swap with hi, hi--
 ...
result [0 0 1 1 2 2]   regions: [0..lo) = 0, [lo..i) = 1, (hi..] = 2
```

**Complexity:** O(n) one pass, O(1) space.

```python
class Solution:
    def sortColors(self, nums: List[int]) -> None:
        lo, i, hi = 0, 0, len(nums) - 1
        while i <= hi:
            if nums[i] == 0:
                nums[lo], nums[i] = nums[i], nums[lo]
                lo += 1; i += 1
            elif nums[i] == 2:
                nums[hi], nums[i] = nums[i], nums[hi]
                hi -= 1
            else:
                i += 1
# test
a = [2, 0, 2, 1, 1, 0]; Solution().sortColors(a)
assert a == [0, 0, 1, 1, 2, 2]
a = [2, 0, 1]; Solution().sortColors(a)
assert a == [0, 1, 2]
a = [1, 1, 0, 0, 2, 2, 1]; Solution().sortColors(a)
assert a == [0, 0, 1, 1, 1, 2, 2]
```

```java
class Solution {
    public void sortColors(int[] nums) {
        int lo = 0, i = 0, hi = nums.length - 1;
        while (i <= hi) {
            if (nums[i] == 0) swap(nums, lo++, i++);
            else if (nums[i] == 2) swap(nums, i, hi--);
            else i++;
        }
    }

    private void swap(int[] a, int i, int j) { int t = a[i]; a[i] = a[j]; a[j] = t; }
}
// test
int[] a = {2, 0, 2, 1, 1, 0};
new Solution().sortColors(a);
check(a, new int[]{0, 0, 1, 1, 2, 2});
```

## 680. Valid Palindrome II | Easy
related: 125
tags: opposite-ends

Walk inward; at the first mismatch you may delete one char: check whether `s[l+1..r]` **or** `s[l..r-1]` is a palindrome.

**Complexity:** O(n).

```python
class Solution:
    def validPalindrome(self, s: str) -> bool:
        def pal(l, r):
            while l < r:
                if s[l] != s[r]:
                    return False
                l, r = l + 1, r - 1
            return True

        l, r = 0, len(s) - 1
        while l < r:
            if s[l] != s[r]:
                return pal(l + 1, r) or pal(l, r - 1)
            l, r = l + 1, r - 1
        return True
# test
assert Solution().validPalindrome("aba") is True
assert Solution().validPalindrome("abca") is True
assert Solution().validPalindrome("abc") is False
```

```java
class Solution {
    public boolean validPalindrome(String s) {
        for (int l = 0, r = s.length() - 1; l < r; l++, r--)
            if (s.charAt(l) != s.charAt(r)) return pal(s, l + 1, r) || pal(s, l, r - 1);
        return true;
    }

    private boolean pal(String s, int l, int r) {
        while (l < r) if (s.charAt(l++) != s.charAt(r--)) return false;
        return true;
    }
}
// test
check(new Solution().validPalindrome("abca"), true);
check(new Solution().validPalindrome("abc"), false);
```

## 392. Is Subsequence | Easy
related: 1143, 115
tags: greedy

Greedy match: advance in `s` whenever `t[j]` equals `s[i]`. Success if all of `s` is consumed.
*Follow-up (many `s` queries):* index positions of each char in `t`, then binary-search the next position.

**Complexity:** O(|t|).

```python
class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        i = 0
        for ch in t:
            if i < len(s) and s[i] == ch:
                i += 1
        return i == len(s)
# test
assert Solution().isSubsequence("abc", "ahbgdc") is True
assert Solution().isSubsequence("axc", "ahbgdc") is False
```

```java
class Solution {
    public boolean isSubsequence(String s, String t) {
        int i = 0;
        for (int j = 0; j < t.length() && i < s.length(); j++)
            if (s.charAt(i) == t.charAt(j)) i++;
        return i == s.length();
    }
}
// test
check(new Solution().isSubsequence("abc", "ahbgdc"), true);
check(new Solution().isSubsequence("axc", "ahbgdc"), false);
```


## 977. Squares of a Sorted Array | Easy
related: 88
tags: opposite-ends

The largest square is at one of the ends. Fill the result from the back, taking the larger `|value|` each time.

**Complexity:** O(n).

```python
class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        n = len(nums)
        res = [0] * n
        l, r = 0, n - 1
        for k in range(n - 1, -1, -1):
            if abs(nums[l]) > abs(nums[r]):
                res[k] = nums[l] ** 2; l += 1
            else:
                res[k] = nums[r] ** 2; r -= 1
        return res
# test
assert Solution().sortedSquares([-4, -1, 0, 3, 10]) == [0, 1, 9, 16, 100]
assert Solution().sortedSquares([-7, -3, 2, 3, 11]) == [4, 9, 9, 49, 121]
assert Solution().sortedSquares([-5, -3, -2, -1]) == [1, 4, 9, 25]
```

```java
class Solution {
    public int[] sortedSquares(int[] nums) {
        int n = nums.length, l = 0, r = n - 1;
        int[] res = new int[n];
        for (int k = n - 1; k >= 0; k--) {
            if (Math.abs(nums[l]) > Math.abs(nums[r])) { res[k] = nums[l] * nums[l]; l++; }
            else { res[k] = nums[r] * nums[r]; r--; }
        }
        return res;
    }
}
// test
check(new Solution().sortedSquares(new int[]{-7, -3, 2, 3, 11}), new int[]{4, 9, 9, 49, 121});
```

## 27. Remove Element | Easy
related: 26, 283
tags: fast-slow

Copy every element not equal to `val` to the write pointer `k`.

**Complexity:** O(n).

```python
class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        k = 0
        for x in nums:
            if x != val:
                nums[k] = x
                k += 1
        return k
# test
a = [0, 1, 2, 2, 3, 0, 4, 2]
k = Solution().removeElement(a, 2)
assert k == 5 and sorted(a[:k]) == [0, 0, 1, 3, 4]
a = [3, 2, 2, 3]
k = Solution().removeElement(a, 3)
assert k == 2 and sorted(a[:k]) == [2, 2]
a = [1]
k = Solution().removeElement(a, 1)
assert k == 0
```

```java
class Solution {
    public int removeElement(int[] nums, int val) {
        int k = 0;
        for (int x : nums) if (x != val) nums[k++] = x;
        return k;
    }
}
// test
int[] a = {3, 2, 2, 3};
check(new Solution().removeElement(a, 3), 2);
```

## 80. Remove Duplicates from Sorted Array II | Medium
related: 26
tags: fast-slow

Allow at most two copies: keep `x` if `k < 2` or `x != nums[k-2]` (compare with the element two slots back in the *kept* prefix).

**Complexity:** O(n).

```python
class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        k = 0
        for x in nums:
            if k < 2 or x != nums[k - 2]:
                nums[k] = x
                k += 1
        return k
# test
a = [0, 0, 1, 1, 1, 1, 2, 3, 3]
k = Solution().removeDuplicates(a)
assert a[:k] == [0, 0, 1, 1, 2, 3, 3]
a = [1, 1, 1, 2, 2, 3]
k = Solution().removeDuplicates(a)
assert k == 5 and a[:k] == [1, 1, 2, 2, 3]
```

```java
class Solution {
    public int removeDuplicates(int[] nums) {
        int k = 0;
        for (int x : nums) if (k < 2 || x != nums[k - 2]) nums[k++] = x;
        return k;
    }
}
// test
int[] a = {1, 1, 1, 2, 2, 3};
int k = new Solution().removeDuplicates(a);
check(Arrays.copyOf(a, k), new int[]{1, 1, 2, 2, 3});
```

## 151. Reverse Words in a String | Medium
related: 344
tags: string

Split on whitespace (drops extra spaces), reverse, join. In-place follow-up (char array): reverse the whole string, then reverse each word, then compact spaces.

**Complexity:** O(n).

```python
class Solution:
    def reverseWords(self, s: str) -> str:
        return " ".join(reversed(s.split()))
# test
assert Solution().reverseWords("the sky is blue") == "blue is sky the"
assert Solution().reverseWords("  hello world  ") == "world hello"
assert Solution().reverseWords("a good   example") == "example good a"
```

```java
class Solution {
    public String reverseWords(String s) {
        List<String> words = Arrays.asList(s.trim().split("\\s+"));
        Collections.reverse(words);
        return String.join(" ", words);
    }
}
// test
check(new Solution().reverseWords("  hello world  "), "world hello");
```

## 344. Reverse String | Easy
related: 151, 125
tags: opposite-ends

Swap from both ends towards the middle.

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def reverseString(self, s: List[str]) -> None:
        l, r = 0, len(s) - 1
        while l < r:
            s[l], s[r] = s[r], s[l]
            l, r = l + 1, r - 1
# test
a = list("hello"); Solution().reverseString(a)
assert a == list("olleh")
a = list("Hannah"); Solution().reverseString(a)
assert a == list("hannaH")
a = ["x"]; Solution().reverseString(a)
assert a == ["x"]
```

```java
class Solution {
    public void reverseString(char[] s) {
        for (int l = 0, r = s.length - 1; l < r; l++, r--) { char t = s[l]; s[l] = s[r]; s[r] = t; }
    }
}
// test
char[] a = "hello".toCharArray();
new Solution().reverseString(a);
check(new String(a), "olleh");
```

## 881. Boats to Save People | Medium
related: 167
tags: greedy, opposite-ends

Sort. The heaviest person always takes a boat; pair them with the lightest if they fit together.

**Complexity:** O(n log n).

```python
class Solution:
    def numRescueBoats(self, people: List[int], limit: int) -> int:
        people.sort()
        l, r, boats = 0, len(people) - 1, 0
        while l <= r:
            if people[l] + people[r] <= limit:
                l += 1
            r -= 1
            boats += 1
        return boats
# test
assert Solution().numRescueBoats([3, 2, 2, 1], 3) == 3
assert Solution().numRescueBoats([3, 5, 3, 4], 5) == 4
```

```java
class Solution {
    public int numRescueBoats(int[] people, int limit) {
        Arrays.sort(people);
        int l = 0, r = people.length - 1, boats = 0;
        while (l <= r) {
            if (people[l] + people[r] <= limit) l++;
            r--;
            boats++;
        }
        return boats;
    }
}
// test
check(new Solution().numRescueBoats(new int[]{3, 2, 2, 1}, 3), 3);
```

## 844. Backspace String Compare | Easy
related: 2390, 1047
tags: two-pointers, reverse-scan

A stack works (O(n) space). For O(1) space walk **both strings from the end**: count pending `#`s and skip that many characters; compare the surviving characters one by one.

**Complexity:** O(n + m) time, O(1) space.

```python
class Solution:
    def backspaceCompare(self, s: str, t: str) -> bool:
        def survivors(x):                  # surviving characters, right to left
            skip = 0
            for ch in reversed(x):
                if ch == "#":
                    skip += 1
                elif skip:
                    skip -= 1
                else:
                    yield ch

        return all(a == b for a, b in itertools.zip_longest(survivors(s), survivors(t)))
# test
assert Solution().backspaceCompare("ab#c", "ad#c") is True
assert Solution().backspaceCompare("ab##", "c#d#") is True
assert Solution().backspaceCompare("a#c", "b") is False
assert Solution().backspaceCompare("a##c", "#a#c") is True
assert Solution().backspaceCompare("bxj##tw", "bxo#j##tw") is True
assert Solution().backspaceCompare("ab", "b") is False
```

```java
class Solution {
    public boolean backspaceCompare(String s, String t) {
        int i = s.length() - 1, j = t.length() - 1;
        while (true) {
            i = next(s, i);
            j = next(t, j);
            if (i < 0 || j < 0) return i < 0 && j < 0;
            if (s.charAt(i--) != t.charAt(j--)) return false;
        }
    }

    private int next(String x, int i) {   // index of the next surviving char at or before i, or -1
        for (int skip = 0; i >= 0; i--) {
            if (x.charAt(i) == '#') skip++;
            else if (skip > 0) skip--;
            else break;
        }
        return i;
    }
}
// test
check(new Solution().backspaceCompare("ab##", "c#d#"), true);
check(new Solution().backspaceCompare("a#c", "b"), false);
check(new Solution().backspaceCompare("bxj##tw", "bxo#j##tw"), true);
```

## 1768. Merge Strings Alternately | Easy
related: 88
tags: two-pointers

Take one character from each word in turn; when one runs out, append the rest of the other.

**Complexity:** O(n + m).

```python
class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        return "".join(a + b for a, b in itertools.zip_longest(word1, word2, fillvalue=""))
# test
assert Solution().mergeAlternately("abc", "pqr") == "apbqcr"
assert Solution().mergeAlternately("ab", "pqrs") == "apbqrs"
assert Solution().mergeAlternately("abcd", "pq") == "apbqcd"
```

```java
class Solution {
    public String mergeAlternately(String word1, String word2) {
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < Math.max(word1.length(), word2.length()); i++) {
            if (i < word1.length()) sb.append(word1.charAt(i));
            if (i < word2.length()) sb.append(word2.charAt(i));
        }
        return sb.toString();
    }
}
// test
check(new Solution().mergeAlternately("ab", "pqrs"), "apbqrs");
check(new Solution().mergeAlternately("abcd", "pq"), "apbqcd");
```

## 345. Reverse Vowels of a String | Easy
related: 344, 125
tags: opposite-ends

Opposite-end pointers: advance each pointer past consonants, swap when both sit on vowels (either case).

**Complexity:** O(n).

```python
class Solution:
    def reverseVowels(self, s: str) -> str:
        vowels = set("aeiouAEIOU")
        a, i, j = list(s), 0, len(s) - 1
        while i < j:
            if a[i] not in vowels:
                i += 1
            elif a[j] not in vowels:
                j -= 1
            else:
                a[i], a[j] = a[j], a[i]
                i, j = i + 1, j - 1
        return "".join(a)
# test
assert Solution().reverseVowels("IceCreAm") == "AceCreIm"
assert Solution().reverseVowels("leetcode") == "leotcede"
assert Solution().reverseVowels("bcd") == "bcd"
assert Solution().reverseVowels("aA") == "Aa"
```

```java
class Solution {
    public String reverseVowels(String s) {
        char[] a = s.toCharArray();
        for (int i = 0, j = a.length - 1; i < j; ) {
            if ("aeiouAEIOU".indexOf(a[i]) < 0) i++;
            else if ("aeiouAEIOU".indexOf(a[j]) < 0) j--;
            else { char t = a[i]; a[i++] = a[j]; a[j--] = t; }
        }
        return new String(a);
    }
}
// test
check(new Solution().reverseVowels("IceCreAm"), "AceCreIm");
check(new Solution().reverseVowels("leetcode"), "leotcede");
```

## 443. String Compression | Medium
related: 38, 271
tags: read-write-pointers, in-place

Read pointer finds each run `[i, j)`; write pointer writes the character and, if the run is longer than 1, the digits of its length. The write pointer never overtakes the read pointer, so it works in place.

**Complexity:** O(n) time, O(1) extra space.

```python
class Solution:
    def compress(self, chars: List[str]) -> int:
        w = i = 0
        while i < len(chars):
            j = i
            while j < len(chars) and chars[j] == chars[i]:
                j += 1
            chars[w] = chars[i]
            w += 1
            if j - i > 1:
                for d in str(j - i):
                    chars[w] = d
                    w += 1
            i = j
        return w
# test
a = ["a", "a", "b", "b", "c", "c", "c"]
k = Solution().compress(a)
assert k == 6 and a[:k] == ["a", "2", "b", "2", "c", "3"]
a = ["a"]
k = Solution().compress(a)
assert k == 1 and a[:k] == ["a"]
a = ["a"] + ["b"] * 12
k = Solution().compress(a)
assert k == 4 and a[:k] == ["a", "b", "1", "2"]
```

```java
class Solution {
    public int compress(char[] chars) {
        int w = 0;
        for (int i = 0, j; i < chars.length; i = j) {
            for (j = i; j < chars.length && chars[j] == chars[i]; j++) ;
            chars[w++] = chars[i];
            if (j - i > 1) for (char d : String.valueOf(j - i).toCharArray()) chars[w++] = d;
        }
        return w;
    }
}
// test
char[] a = {'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'};
int k = new Solution().compress(a);
check(new String(a, 0, k), "ab12");
```

## 1679. Max Number of K-Sum Pairs | Medium
related: 1, 167
tags: hash-map, two-sum

Two Sum, counted: scan once; if a waiting partner `k − x` exists, pair them (use it up); otherwise leave `x` waiting. (Sorting + opposite-end pointers also works in O(n log n).)

**Complexity:** O(n).

```python
class Solution:
    def maxOperations(self, nums: List[int], k: int) -> int:
        waiting, ops = Counter(), 0
        for x in nums:
            if waiting[k - x]:
                waiting[k - x] -= 1
                ops += 1
            else:
                waiting[x] += 1
        return ops
# test
assert Solution().maxOperations([1, 2, 3, 4], 5) == 2
assert Solution().maxOperations([3, 1, 3, 4, 3], 6) == 1
assert Solution().maxOperations([2, 2, 2, 2], 4) == 2
assert Solution().maxOperations([1], 2) == 0
```

```java
class Solution {
    public int maxOperations(int[] nums, int k) {
        Map<Integer, Integer> waiting = new HashMap<>();
        int ops = 0;
        for (int x : nums) {
            int c = waiting.getOrDefault(k - x, 0);
            if (c > 0) { waiting.put(k - x, c - 1); ops++; }
            else waiting.merge(x, 1, Integer::sum);
        }
        return ops;
    }
}
// test
check(new Solution().maxOperations(new int[]{1, 2, 3, 4}, 5), 2);
check(new Solution().maxOperations(new int[]{3, 1, 3, 4, 3}, 6), 1);
```
