# Arrays & Hashing

**When to use:** you need O(1) lookups ("have I seen X?", "how many times?", "where was X?"), grouping by a key, or counting.

- **Hash set**: membership / duplicates.
- **Hash map**: value → index / count / group.
- **Prefix sums + map**: "number of subarrays with sum = k" in O(n): `sum(i..j) = pre[j+1] - pre[i]`.
- **In-place index tricks**: when values are in `[1..n]`, use the array itself as the hash table.

```python
# prefix-sum + hashmap template: count subarrays whose sum == k
count, pre, seen = 0, 0, {0: 1}
for x in nums:
    pre += x
    count += seen.get(pre - k, 0)
    seen[pre] = seen.get(pre, 0) + 1
```

## 1. Two Sum | Easy
related: 167, 15, 18, 560
tags: blind75, hash-map

Scan once. For each `x`, the partner we need is `target - x`. If we've already seen it, we're done; otherwise remember `x → index`.

```viz
nums = [2, 7, 11, 15], target = 9
i=0  x=2   need 7   seen {}       -> store {2:0}
i=1  x=7   need 2   seen {2:0}    -> hit!  answer [0, 1]
```

**Complexity:** O(n) time, O(n) space.

```python
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        seen = {}  # value -> index
        for i, x in enumerate(nums):
            if target - x in seen:
                return [seen[target - x], i]
            seen[x] = i
        return []
# test
assert Solution().twoSum([2, 7, 11, 15], 9) == [0, 1]
assert Solution().twoSum([3, 2, 4], 6) == [1, 2]
```

```java
class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> seen = new HashMap<>(); // value -> index
        for (int i = 0; i < nums.length; i++) {
            Integer j = seen.get(target - nums[i]);
            if (j != null) return new int[]{j, i};
            seen.put(nums[i], i);
        }
        return new int[0];
    }
}
// test
check(new Solution().twoSum(new int[]{2, 7, 11, 15}, 9), new int[]{0, 1});
check(new Solution().twoSum(new int[]{3, 2, 4}, 6), new int[]{1, 2});
```

## 217. Contains Duplicate | Easy
related: 242, 128
tags: blind75, hash-set

Insert into a set; if an element is already there, it's a duplicate.

**Complexity:** O(n) time, O(n) space (sorting gives O(n log n) / O(1)).

```python
class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        seen = set()
        for x in nums:
            if x in seen:
                return True
            seen.add(x)
        return False
# test
assert Solution().containsDuplicate([1, 2, 3, 1]) is True
assert Solution().containsDuplicate([1, 2, 3, 4]) is False
```

```java
class Solution {
    public boolean containsDuplicate(int[] nums) {
        Set<Integer> seen = new HashSet<>();
        for (int x : nums) if (!seen.add(x)) return true; // add() returns false if present
        return false;
    }
}
// test
check(new Solution().containsDuplicate(new int[]{1, 2, 3, 1}), true);
check(new Solution().containsDuplicate(new int[]{1, 2, 3, 4}), false);
```

## 242. Valid Anagram | Easy
related: 49, 438, 567
tags: blind75, counting

Anagrams have identical letter counts. Count up for `s`, down for `t`; every bucket must end at 0.

**Complexity:** O(n) time, O(1) space (26 letters).

```python
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        cnt = [0] * 26
        for a, b in zip(s, t):
            cnt[ord(a) - 97] += 1
            cnt[ord(b) - 97] -= 1
        return all(c == 0 for c in cnt)
# test
assert Solution().isAnagram("anagram", "nagaram") is True
assert Solution().isAnagram("rat", "car") is False
```

```java
class Solution {
    public boolean isAnagram(String s, String t) {
        if (s.length() != t.length()) return false;
        int[] cnt = new int[26];
        for (int i = 0; i < s.length(); i++) {
            cnt[s.charAt(i) - 'a']++;
            cnt[t.charAt(i) - 'a']--;
        }
        for (int c : cnt) if (c != 0) return false;
        return true;
    }
}
// test
check(new Solution().isAnagram("anagram", "nagaram"), true);
check(new Solution().isAnagram("rat", "car"), false);
```

## 49. Group Anagrams | Medium
related: 242, 438
tags: blind75, hash-map

Every anagram shares a **canonical key**: the sorted word (O(k log k)) or the 26-letter count tuple (O(k)). Group words by key.

```viz
"eat" -> key "aet" ┐
"tea" -> key "aet" ├─> ["eat","tea","ate"]
"ate" -> key "aet" ┘
"tan" -> key "ant" ──> ["tan","nat"]   ...
```

**Complexity:** O(n·k) with count keys (n words, length k).

```python
class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        groups = defaultdict(list)
        for w in strs:
            key = [0] * 26
            for ch in w:
                key[ord(ch) - 97] += 1
            groups[tuple(key)].append(w)
        return list(groups.values())
# test
res = Solution().groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"])
assert sorted(sorted(g) for g in res) == [["ate", "eat", "tea"], ["bat"], ["nat", "tan"]]
```

```java
class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> groups = new HashMap<>();
        for (String w : strs) {
            char[] c = w.toCharArray();
            Arrays.sort(c);                       // canonical key
            groups.computeIfAbsent(new String(c), k -> new ArrayList<>()).add(w);
        }
        return new ArrayList<>(groups.values());
    }
}
// test
List<List<String>> r = new Solution().groupAnagrams(new String[]{"eat", "tea", "tan", "ate", "nat", "bat"});
r.forEach(Collections::sort);
r.sort((a, b) -> a.get(0).compareTo(b.get(0)));
check(r, "[[ate, eat, tea], [bat], [nat, tan]]");
```

## 347. Top K Frequent Elements | Medium
related: 215, 692, 973
tags: blind75, bucket-sort, heap

Count, then **bucket sort by frequency**: `bucket[f]` holds numbers appearing `f` times (f ≤ n). Walk buckets from high to low until k numbers are collected. (A size-k min-heap gives O(n log k).)

```viz
nums = [1,1,1,2,2,3]   counts {1:3, 2:2, 3:1}
bucket: 0:[] 1:[3] 2:[2] 3:[1] ...   read from the right -> [1, 2]
```

**Complexity:** O(n) time, O(n) space.

```python
class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        cnt = Counter(nums)
        bucket = [[] for _ in range(len(nums) + 1)]
        for x, f in cnt.items():
            bucket[f].append(x)
        res = []
        for f in range(len(bucket) - 1, 0, -1):
            for x in bucket[f]:
                res.append(x)
                if len(res) == k:
                    return res
        return res
# test
assert sorted(Solution().topKFrequent([1, 1, 1, 2, 2, 3], 2)) == [1, 2]
assert Solution().topKFrequent([1], 1) == [1]
```

```java
class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int x : nums) cnt.merge(x, 1, Integer::sum);
        List<List<Integer>> bucket = new ArrayList<>();
        for (int i = 0; i <= nums.length; i++) bucket.add(new ArrayList<>());
        for (Map.Entry<Integer, Integer> e : cnt.entrySet()) bucket.get(e.getValue()).add(e.getKey());
        int[] res = new int[k];
        int i = 0;
        for (int f = nums.length; f > 0 && i < k; f--)
            for (int x : bucket.get(f)) if (i < k) res[i++] = x;
        return res;
    }
}
// test
int[] t = new Solution().topKFrequent(new int[]{1, 1, 1, 2, 2, 3}, 2);
Arrays.sort(t);
check(t, new int[]{1, 2});
```

## 238. Product of Array Except Self | Medium
related: 42, 152
tags: blind75, prefix-product

`ans[i] = (product of everything left of i) × (product of everything right of i)`. Fill left products in one pass, then multiply by a running right product in a reverse pass. No division needed.

```viz
nums     [ 1,  2,  3,  4]
left     [ 1,  1,  2,  6]     (product before i)
right    [24, 12,  4,  1]     (product after i)
answer   [24, 12,  8,  6]
```

**Complexity:** O(n) time, O(1) extra space (output excluded).

```python
class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        n = len(nums)
        ans = [1] * n
        for i in range(1, n):
            ans[i] = ans[i - 1] * nums[i - 1]
        right = 1
        for i in range(n - 1, -1, -1):
            ans[i] *= right
            right *= nums[i]
        return ans
# test
assert Solution().productExceptSelf([1, 2, 3, 4]) == [24, 12, 8, 6]
assert Solution().productExceptSelf([-1, 1, 0, -3, 3]) == [0, 0, 9, 0, 0]
```

```java
class Solution {
    public int[] productExceptSelf(int[] nums) {
        int n = nums.length;
        int[] ans = new int[n];
        ans[0] = 1;
        for (int i = 1; i < n; i++) ans[i] = ans[i - 1] * nums[i - 1];
        int right = 1;
        for (int i = n - 1; i >= 0; i--) {
            ans[i] *= right;
            right *= nums[i];
        }
        return ans;
    }
}
// test
check(new Solution().productExceptSelf(new int[]{1, 2, 3, 4}), new int[]{24, 12, 8, 6});
check(new Solution().productExceptSelf(new int[]{-1, 1, 0, -3, 3}), new int[]{0, 0, 9, 0, 0});
```

## 36. Valid Sudoku | Medium
related: 37
tags: hash-set, matrix

Keep a set for each row, column and 3×3 box. Box index = `(r // 3) * 3 + c // 3`. Any repeat → invalid.

**Complexity:** O(81) = O(1).

```python
class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        rows = [set() for _ in range(9)]
        cols = [set() for _ in range(9)]
        boxes = [set() for _ in range(9)]
        for r in range(9):
            for c in range(9):
                v = board[r][c]
                if v == ".":
                    continue
                b = (r // 3) * 3 + c // 3
                if v in rows[r] or v in cols[c] or v in boxes[b]:
                    return False
                rows[r].add(v); cols[c].add(v); boxes[b].add(v)
        return True
# test
B = ["53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1",
     "7...2...6", ".6....28.", "...419..5", "....8..79"]
board = [list(r) for r in B]
assert Solution().isValidSudoku(board) is True
board[0][0] = "8"  # column 0 already has an 8
assert Solution().isValidSudoku(board) is False
```

```java
class Solution {
    public boolean isValidSudoku(char[][] board) {
        boolean[][] row = new boolean[9][9], col = new boolean[9][9], box = new boolean[9][9];
        for (int r = 0; r < 9; r++)
            for (int c = 0; c < 9; c++) {
                if (board[r][c] == '.') continue;
                int v = board[r][c] - '1', b = (r / 3) * 3 + c / 3;
                if (row[r][v] || col[c][v] || box[b][v]) return false;
                row[r][v] = col[c][v] = box[b][v] = true;
            }
        return true;
    }
}
// test
char[][] b = Stream.of("53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1",
        "7...2...6", ".6....28.", "...419..5", "....8..79").map(String::toCharArray).toArray(char[][]::new);
check(new Solution().isValidSudoku(b), true);
b[0][0] = '8';
check(new Solution().isValidSudoku(b), false);
```

## 271. Encode and Decode Strings | Medium
related: 49
tags: blind75, string, premium

Any delimiter could appear inside a string, so prefix each string with its **length** and a separator: `"4#neet4#code"`. To decode, read digits up to `#`, then take exactly that many characters.

**Complexity:** O(total length).

```python
class Codec:
    def encode(self, strs: List[str]) -> str:
        return "".join(f"{len(s)}#{s}" for s in strs)

    def decode(self, s: str) -> List[str]:
        res, i = [], 0
        while i < len(s):
            j = s.index("#", i)
            n = int(s[i:j])
            res.append(s[j + 1:j + 1 + n])
            i = j + 1 + n
        return res
# test
c = Codec()
for data in (["neet", "code", "love", "you"], ["we", "say", ":", "yes", "#4#", ""], []):
    assert c.decode(c.encode(data)) == data
```

```java
class Codec {
    public String encode(List<String> strs) {
        StringBuilder sb = new StringBuilder();
        for (String s : strs) sb.append(s.length()).append('#').append(s);
        return sb.toString();
    }

    public List<String> decode(String s) {
        List<String> res = new ArrayList<>();
        int i = 0;
        while (i < s.length()) {
            int j = s.indexOf('#', i);
            int n = Integer.parseInt(s.substring(i, j));
            res.add(s.substring(j + 1, j + 1 + n));
            i = j + 1 + n;
        }
        return res;
    }
}
// test
Codec c = new Codec();
List<String> data = Arrays.asList("we", "say", ":", "yes", "#4#", "");
check(c.decode(c.encode(data)), data);
```

## 128. Longest Consecutive Sequence | Medium
related: 217
tags: blind75, hash-set

Put everything in a set. Only start counting from numbers that begin a run (`x - 1` not in set), then walk `x+1, x+2, …`. Each number is visited at most twice → O(n).

```viz
set {100, 4, 200, 1, 3, 2}
1   (0 missing -> start)  1,2,3,4  -> length 4
100 (99 missing -> start) 100      -> length 1
2   (1 present -> skip, not a start)
```

**Complexity:** O(n) time, O(n) space.

```python
class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        s = set(nums)
        best = 0
        for x in s:
            if x - 1 not in s:          # x starts a run
                y = x
                while y + 1 in s:
                    y += 1
                best = max(best, y - x + 1)
        return best
# test
assert Solution().longestConsecutive([100, 4, 200, 1, 3, 2]) == 4
assert Solution().longestConsecutive([0, 3, 7, 2, 5, 8, 4, 6, 0, 1]) == 9
```

```java
class Solution {
    public int longestConsecutive(int[] nums) {
        Set<Integer> s = new HashSet<>();
        for (int x : nums) s.add(x);
        int best = 0;
        for (int x : s) {
            if (s.contains(x - 1)) continue;   // not the start of a run
            int y = x;
            while (s.contains(y + 1)) y++;
            best = Math.max(best, y - x + 1);
        }
        return best;
    }
}
// test
check(new Solution().longestConsecutive(new int[]{100, 4, 200, 1, 3, 2}), 4);
check(new Solution().longestConsecutive(new int[]{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}), 9);
```

## 169. Majority Element | Easy
related: 347
tags: boyer-moore

**Boyer–Moore voting:** keep a candidate and a counter. Same value → +1, different → −1; at 0 adopt the current value. The majority (> n/2) can't be cancelled out.

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        cand, cnt = None, 0
        for x in nums:
            if cnt == 0:
                cand = x
            cnt += 1 if x == cand else -1
        return cand
# test
assert Solution().majorityElement([3, 2, 3]) == 3
assert Solution().majorityElement([2, 2, 1, 1, 1, 2, 2]) == 2
```

```java
class Solution {
    public int majorityElement(int[] nums) {
        int cand = 0, cnt = 0;
        for (int x : nums) {
            if (cnt == 0) cand = x;
            cnt += (x == cand) ? 1 : -1;
        }
        return cand;
    }
}
// test
check(new Solution().majorityElement(new int[]{3, 2, 3}), 3);
check(new Solution().majorityElement(new int[]{2, 2, 1, 1, 1, 2, 2}), 2);
```

## 560. Subarray Sum Equals K | Medium
related: 1, 525, 437, 209
tags: prefix-sum, hash-map

Subarray `(i, j]` sums to `k` iff `pre[j] - pre[i] = k`. While scanning, count how many earlier prefixes equal `pre - k`. Works with negatives (sliding window doesn't).

```viz
nums=[1,2,3] k=3   prefix: 0 -> 1 -> 3 -> 6
pre=3: need 0 (seen once)  -> +1   [1,2]
pre=6: need 3 (seen once)  -> +1   [3]          total 2
```

**Complexity:** O(n) time, O(n) space.

```python
class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        count, pre = 0, 0
        seen = defaultdict(int, {0: 1})
        for x in nums:
            pre += x
            count += seen[pre - k]
            seen[pre] += 1
        return count
# test
assert Solution().subarraySum([1, 1, 1], 2) == 2
assert Solution().subarraySum([1, 2, 3], 3) == 2
assert Solution().subarraySum([1, -1, 0], 0) == 3
```

```java
class Solution {
    public int subarraySum(int[] nums, int k) {
        Map<Integer, Integer> seen = new HashMap<>();
        seen.put(0, 1);
        int count = 0, pre = 0;
        for (int x : nums) {
            pre += x;
            count += seen.getOrDefault(pre - k, 0);
            seen.merge(pre, 1, Integer::sum);
        }
        return count;
    }
}
// test
check(new Solution().subarraySum(new int[]{1, 1, 1}, 2), 2);
check(new Solution().subarraySum(new int[]{1, 2, 3}, 3), 2);
check(new Solution().subarraySum(new int[]{1, -1, 0}, 0), 3);
```

## 41. First Missing Positive | Hard
related: 268, 287
tags: in-place-hashing

The answer is in `[1, n+1]`. Use the array as a hash table: **cyclic-swap** each value `v ∈ [1, n]` to index `v-1`. Then the first index `i` with `nums[i] != i+1` gives `i+1`.

```viz
[3, 4, -1, 1]
swap 3 -> idx2:  [-1, 4, 3, 1]
swap 4 -> idx3:  [-1, 1, 3, 4]
swap 1 -> idx0:  [ 1,-1, 3, 4]
scan: idx1 holds -1  -> answer 2
```

**Complexity:** O(n) time (each swap fixes one slot), O(1) space.

```python
class Solution:
    def firstMissingPositive(self, nums: List[int]) -> int:
        n = len(nums)
        for i in range(n):
            while 1 <= nums[i] <= n and nums[nums[i] - 1] != nums[i]:
                j = nums[i] - 1
                nums[i], nums[j] = nums[j], nums[i]
        for i in range(n):
            if nums[i] != i + 1:
                return i + 1
        return n + 1
# test
assert Solution().firstMissingPositive([1, 2, 0]) == 3
assert Solution().firstMissingPositive([3, 4, -1, 1]) == 2
assert Solution().firstMissingPositive([7, 8, 9, 11, 12]) == 1
```

```java
class Solution {
    public int firstMissingPositive(int[] nums) {
        int n = nums.length;
        for (int i = 0; i < n; i++) {
            while (nums[i] >= 1 && nums[i] <= n && nums[nums[i] - 1] != nums[i]) {
                int j = nums[i] - 1, t = nums[i];
                nums[i] = nums[j];
                nums[j] = t;
            }
        }
        for (int i = 0; i < n; i++) if (nums[i] != i + 1) return i + 1;
        return n + 1;
    }
}
// test
check(new Solution().firstMissingPositive(new int[]{1, 2, 0}), 3);
check(new Solution().firstMissingPositive(new int[]{3, 4, -1, 1}), 2);
check(new Solution().firstMissingPositive(new int[]{7, 8, 9, 11, 12}), 1);
```

## 189. Rotate Array | Medium
related: 48, 61
tags: reverse-trick

Rotating right by k = reverse all, then reverse the first k, then reverse the rest.

```viz
[1 2 3 4 5 6 7], k=3
reverse all     [7 6 5 4 3 2 1]
reverse [0,3)   [5 6 7 | 4 3 2 1]
reverse [3,7)   [5 6 7 | 1 2 3 4]
```

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        def rev(i, j):
            while i < j:
                nums[i], nums[j] = nums[j], nums[i]
                i, j = i + 1, j - 1
        n = len(nums)
        k %= n
        rev(0, n - 1); rev(0, k - 1); rev(k, n - 1)
# test
a = [1, 2, 3, 4, 5, 6, 7]; Solution().rotate(a, 3)
assert a == [5, 6, 7, 1, 2, 3, 4]
b = [-1, -100, 3, 99]; Solution().rotate(b, 2)
assert b == [3, 99, -1, -100]
```

```java
class Solution {
    public void rotate(int[] nums, int k) {
        int n = nums.length;
        k %= n;
        rev(nums, 0, n - 1);
        rev(nums, 0, k - 1);
        rev(nums, k, n - 1);
    }

    private void rev(int[] a, int i, int j) {
        while (i < j) { int t = a[i]; a[i++] = a[j]; a[j--] = t; }
    }
}
// test
int[] a = {1, 2, 3, 4, 5, 6, 7};
new Solution().rotate(a, 3);
check(a, new int[]{5, 6, 7, 1, 2, 3, 4});
```

## 88. Merge Sorted Array | Easy
related: 21, 977
tags: two-pointers

Fill `nums1` **from the back**: the largest remaining element goes to position `k`. Working backwards never overwrites unread data.

**Complexity:** O(m + n) time, O(1) space.

```python
class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        i, j, k = m - 1, n - 1, m + n - 1
        while j >= 0:
            if i >= 0 and nums1[i] > nums2[j]:
                nums1[k] = nums1[i]; i -= 1
            else:
                nums1[k] = nums2[j]; j -= 1
            k -= 1
# test
a = [1, 2, 3, 0, 0, 0]; Solution().merge(a, 3, [2, 5, 6], 3)
assert a == [1, 2, 2, 3, 5, 6]
b = [0]; Solution().merge(b, 0, [1], 1)
assert b == [1]
```

```java
class Solution {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        int i = m - 1, j = n - 1, k = m + n - 1;
        while (j >= 0) {
            if (i >= 0 && nums1[i] > nums2[j]) nums1[k--] = nums1[i--];
            else nums1[k--] = nums2[j--];
        }
    }
}
// test
int[] a = {1, 2, 3, 0, 0, 0};
new Solution().merge(a, 3, new int[]{2, 5, 6}, 3);
check(a, new int[]{1, 2, 2, 3, 5, 6});
```

## 26. Remove Duplicates from Sorted Array | Easy
related: 283, 88
tags: two-pointers

Slow pointer `k` marks the end of the unique prefix; fast pointer scans. Copy whenever a new value appears.

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        k = 0
        for x in nums:
            if k == 0 or x != nums[k - 1]:
                nums[k] = x
                k += 1
        return k
# test
a = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]
k = Solution().removeDuplicates(a)
assert k == 5 and a[:k] == [0, 1, 2, 3, 4]
```

```java
class Solution {
    public int removeDuplicates(int[] nums) {
        int k = 0;
        for (int x : nums) if (k == 0 || x != nums[k - 1]) nums[k++] = x;
        return k;
    }
}
// test
int[] a = {0, 0, 1, 1, 1, 2, 2, 3, 3, 4};
int k = new Solution().removeDuplicates(a);
check(Arrays.copyOf(a, k), new int[]{0, 1, 2, 3, 4});
```

## 14. Longest Common Prefix | Easy
related: 208
tags: string

Vertical scan: compare column `i` of every string with the first string; stop at the first mismatch or end of any string.

**Complexity:** O(total characters).

```python
class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        first = strs[0]
        for i, ch in enumerate(first):
            for s in strs[1:]:
                if i == len(s) or s[i] != ch:
                    return first[:i]
        return first
# test
assert Solution().longestCommonPrefix(["flower", "flow", "flight"]) == "fl"
assert Solution().longestCommonPrefix(["dog", "racecar", "car"]) == ""
```

```java
class Solution {
    public String longestCommonPrefix(String[] strs) {
        String first = strs[0];
        for (int i = 0; i < first.length(); i++)
            for (String s : strs)
                if (i == s.length() || s.charAt(i) != first.charAt(i)) return first.substring(0, i);
        return first;
    }
}
// test
check(new Solution().longestCommonPrefix(new String[]{"flower", "flow", "flight"}), "fl");
check(new Solution().longestCommonPrefix(new String[]{"dog", "racecar", "car"}), "");
```

## 205. Isomorphic Strings | Easy
related: 242
tags: hash-map

Need a **bijection**: map `s→t` and `t→s`. A conflict in either direction → false.

**Complexity:** O(n).

```python
class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        st, ts = {}, {}
        for a, b in zip(s, t):
            if st.setdefault(a, b) != b or ts.setdefault(b, a) != a:
                return False
        return True
# test
assert Solution().isIsomorphic("egg", "add") is True
assert Solution().isIsomorphic("foo", "bar") is False
assert Solution().isIsomorphic("badc", "baba") is False
```

```java
class Solution {
    public boolean isIsomorphic(String s, String t) {
        int[] st = new int[256], ts = new int[256]; // store last index+1 seen
        for (int i = 0; i < s.length(); i++) {
            char a = s.charAt(i), b = t.charAt(i);
            if (st[a] != ts[b]) return false;   // must have been paired together before (or never)
            st[a] = ts[b] = i + 1;
        }
        return true;
    }
}
// test
check(new Solution().isIsomorphic("egg", "add"), true);
check(new Solution().isIsomorphic("foo", "bar"), false);
check(new Solution().isIsomorphic("badc", "baba"), false);
```

## 525. Contiguous Array | Medium
related: 560
tags: prefix-sum, hash-map

Treat 0 as −1. A subarray with equal 0s and 1s has sum 0 ⇔ two equal prefix sums. Store the **first** index of each prefix; the longest span between equal prefixes wins.

```viz
nums   0   1   1   0   1   1
as±1  -1  +1  +1  -1  +1  +1
pre  0 -1   0   1   0   1   2      pre=0 first at -1, again at idx 3 -> len 4
```

**Complexity:** O(n) time, O(n) space.

```python
class Solution:
    def findMaxLength(self, nums: List[int]) -> int:
        first = {0: -1}
        pre = best = 0
        for i, x in enumerate(nums):
            pre += 1 if x else -1
            if pre in first:
                best = max(best, i - first[pre])
            else:
                first[pre] = i
        return best
# test
assert Solution().findMaxLength([0, 1]) == 2
assert Solution().findMaxLength([0, 1, 1, 0, 1, 1]) == 4
```

```java
class Solution {
    public int findMaxLength(int[] nums) {
        Map<Integer, Integer> first = new HashMap<>();
        first.put(0, -1);
        int pre = 0, best = 0;
        for (int i = 0; i < nums.length; i++) {
            pre += nums[i] == 1 ? 1 : -1;
            Integer j = first.putIfAbsent(pre, i);
            if (j != null) best = Math.max(best, i - j);
        }
        return best;
    }
}
// test
check(new Solution().findMaxLength(new int[]{0, 1}), 2);
check(new Solution().findMaxLength(new int[]{0, 1, 1, 0, 1, 1}), 4);
```
