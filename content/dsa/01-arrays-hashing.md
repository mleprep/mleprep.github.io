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
assert sorted(Solution().twoSum([2, 7, 11, 15], 9)) == sorted([0, 1])
assert sorted(Solution().twoSum([3, 2, 4], 6)) == sorted([1, 2])
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

```go
func twoSum(nums []int, target int) []int {
	seen := map[int]int{} // value -> index
	for i, x := range nums {
		if j, ok := seen[target-x]; ok {
			return []int{j, i}
		}
		seen[x] = i
	}
	return nil
}
// test
assertUnordered(twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1})
assertUnordered(twoSum([]int{3, 2, 4}, 6), []int{1, 2})
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

```go
func containsDuplicate(nums []int) bool {
	seen := map[int]bool{}
	for _, x := range nums {
		if seen[x] {
			return true
		}
		seen[x] = true
	}
	return false
}
// test
assertEq(containsDuplicate([]int{1, 2, 3, 1}), true)
assertEq(containsDuplicate([]int{1, 2, 3, 4}), false)
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

```go
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var cnt [26]int
	for i := range s {
		cnt[s[i]-'a']++
		cnt[t[i]-'a']--
	}
	return cnt == [26]int{} // arrays compare by value
}
// test
assertEq(isAnagram("anagram", "nagaram"), true)
assertEq(isAnagram("rat", "car"), false)
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
assert Solution().groupAnagrams([""]) == [[""]]
assert Solution().groupAnagrams(["a"]) == [["a"]]
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

```go
func groupAnagrams(strs []string) [][]string {
	groups := map[[26]int][]string{} // letter counts are a valid map key
	for _, w := range strs {
		var key [26]int
		for _, ch := range w {
			key[ch-'a']++
		}
		groups[key] = append(groups[key], w)
	}
	res := make([][]string, 0, len(groups))
	for _, g := range groups {
		res = append(res, g)
	}
	return res
}
// test
assertUnorderedDeep(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}), [][]string{{"ate", "eat", "tea"}, {"bat"}, {"nat", "tan"}})
assertEq(groupAnagrams([]string{""}), [][]string{{""}})
assertEq(groupAnagrams([]string{"a"}), [][]string{{"a"}})
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

```go
func topKFrequent(nums []int, k int) []int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	bucket := make([][]int, len(nums)+1) // bucket[f] = values seen f times
	for x, f := range cnt {
		bucket[f] = append(bucket[f], x)
	}
	res := []int{}
	for f := len(bucket) - 1; f > 0; f-- {
		for _, x := range bucket[f] {
			res = append(res, x)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
// test
assertUnordered(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2), []int{1, 2})
assertEq(topKFrequent([]int{1}, 1), []int{1})
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

```go
func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	ans[0] = 1
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] * nums[i-1] // product of everything to the left
	}
	right := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] *= right
		right *= nums[i]
	}
	return ans
}
// test
assertEq(productExceptSelf([]int{1, 2, 3, 4}), []int{24, 12, 8, 6})
assertEq(productExceptSelf([]int{-1, 1, 0, -3, 3}), []int{0, 0, 9, 0, 0})
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

```go
func isValidSudoku(board [][]byte) bool {
	var rows, cols, boxes [9][9]bool
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			v, b := board[r][c]-'1', r/3*3+c/3
			if rows[r][v] || cols[c][v] || boxes[b][v] {
				return false
			}
			rows[r][v], cols[c][v], boxes[b][v] = true, true, true
		}
	}
	return true
}
// test
board := byteGrid("53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1",
	"7...2...6", ".6....28.", "...419..5", "....8..79")
assertEq(isValidSudoku(board), true)
board[0][0] = '8' // column 0 already has an 8
assertEq(isValidSudoku(board), false)
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

```go
type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Encode writes each string as "<length>#<string>".
func (codec *Codec) Encode(strs []string) string {
	var sb strings.Builder
	for _, s := range strs {
		sb.WriteString(strconv.Itoa(len(s)) + "#" + s)
	}
	return sb.String()
}

func (codec *Codec) Decode(s string) []string {
	res := []string{}
	for i := 0; i < len(s); {
		j := i + strings.IndexByte(s[i:], '#')
		n, _ := strconv.Atoi(s[i:j])
		res = append(res, s[j+1:j+1+n])
		i = j + 1 + n
	}
	return res
}
// test
c := Constructor()
for _, data := range [][]string{{"neet", "code", "love", "you"}, {"we", "say", ":", "yes", "#4#", ""}, {}} {
	assertEq(c.Decode(c.Encode(data)), data)
}
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

```go
func longestConsecutive(nums []int) int {
	set := map[int]bool{}
	for _, x := range nums {
		set[x] = true
	}
	best := 0
	for x := range set {
		if !set[x-1] { // x starts a run
			y := x
			for set[y+1] {
				y++
			}
			best = max(best, y-x+1)
		}
	}
	return best
}
// test
assertEq(longestConsecutive([]int{100, 4, 200, 1, 3, 2}), 4)
assertEq(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}), 9)
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

```go
func majorityElement(nums []int) int {
	cand, cnt := 0, 0 // Boyer-Moore voting
	for _, x := range nums {
		if cnt == 0 {
			cand = x
		}
		if x == cand {
			cnt++
		} else {
			cnt--
		}
	}
	return cand
}
// test
assertEq(majorityElement([]int{3, 2, 3}), 3)
assertEq(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}), 2)
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

```go
func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	seen := map[int]int{0: 1} // prefix sum -> how many times seen
	for _, x := range nums {
		pre += x
		count += seen[pre-k]
		seen[pre]++
	}
	return count
}
// test
assertEq(subarraySum([]int{1, 1, 1}, 2), 2)
assertEq(subarraySum([]int{1, 2, 3}, 3), 2)
assertEq(subarraySum([]int{1, -1, 0}, 0), 3)
```

## 41. First Missing Positive | Hard
related: 268, 287
tags: in-place

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

```go
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		for nums[i] >= 1 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			j := nums[i] - 1 // put value v at index v-1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	for i, v := range nums {
		if v != i+1 {
			return i + 1
		}
	}
	return n + 1
}
// test
assertEq(firstMissingPositive([]int{1, 2, 0}), 3)
assertEq(firstMissingPositive([]int{3, 4, -1, 1}), 2)
assertEq(firstMissingPositive([]int{7, 8, 9, 11, 12}), 1)
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

```go
func rotate(nums []int, k int) {
	rev := func(i, j int) {
		for ; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	n := len(nums)
	k %= n
	rev(0, n-1)
	rev(0, k-1)
	rev(k, n-1)
}
// test
a := []int{1, 2, 3, 4, 5, 6, 7}
rotate(a, 3)
assertEq(a, []int{5, 6, 7, 1, 2, 3, 4})
b := []int{-1, -100, 3, 99}
rotate(b, 2)
assertEq(b, []int{3, 99, -1, -100})
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

```go
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1 // fill from the back
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}
// test
a := []int{1, 2, 3, 0, 0, 0}
merge(a, 3, []int{2, 5, 6}, 3)
assertEq(a, []int{1, 2, 2, 3, 5, 6})
b := []int{0}
merge(b, 0, []int{1}, 1)
assertEq(b, []int{1})
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
a = [1, 1, 2]
k = Solution().removeDuplicates(a)
assert k == 2 and a[:k] == [1, 2]
a = [7]
k = Solution().removeDuplicates(a)
assert k == 1 and a[:k] == [7]
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

```go
func removeDuplicates(nums []int) int {
	k := 0
	for _, x := range nums {
		if k == 0 || x != nums[k-1] {
			nums[k] = x
			k++
		}
	}
	return k
}
// test
a := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
k := removeDuplicates(a)
assertEq(a[:k], []int{0, 1, 2, 3, 4})
b := []int{1, 1, 2}
k = removeDuplicates(b)
assertEq(b[:k], []int{1, 2})
c := []int{7}
k = removeDuplicates(c)
assertEq(c[:k], []int{7})
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

```go
func longestCommonPrefix(strs []string) string {
	first := strs[0]
	for i := 0; i < len(first); i++ {
		for _, s := range strs[1:] {
			if i == len(s) || s[i] != first[i] {
				return first[:i]
			}
		}
	}
	return first
}
// test
assertEq(longestCommonPrefix([]string{"flower", "flow", "flight"}), "fl")
assertEq(longestCommonPrefix([]string{"dog", "racecar", "car"}), "")
assertEq(longestCommonPrefix([]string{"ab", "abc"}), "ab")
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

```go
func isIsomorphic(s string, t string) bool {
	var st, ts [256]byte // 0 = not mapped yet
	for i := 0; i < len(s); i++ {
		a, b := s[i], t[i]
		if st[a] == 0 && ts[b] == 0 {
			st[a], ts[b] = b, a
		} else if st[a] != b || ts[b] != a {
			return false
		}
	}
	return true
}
// test
assertEq(isIsomorphic("egg", "add"), true)
assertEq(isIsomorphic("foo", "bar"), false)
assertEq(isIsomorphic("badc", "baba"), false)
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

```go
func findMaxLength(nums []int) int {
	first := map[int]int{0: -1} // balance -> first index it was seen
	pre, best := 0, 0
	for i, x := range nums {
		if x == 1 {
			pre++
		} else {
			pre--
		}
		if j, ok := first[pre]; ok {
			best = max(best, i-j)
		} else {
			first[pre] = i
		}
	}
	return best
}
// test
assertEq(findMaxLength([]int{0, 1}), 2)
assertEq(findMaxLength([]int{0, 1, 1, 0, 1, 1}), 4)
```


## 31. Next Permutation | Medium
related: 46, 47
tags: array

From the right, find the first index `i` with `a[i] < a[i+1]` (the "pivot"). Swap it with the rightmost element larger than it, then reverse the suffix (it was descending → make it ascending). No pivot → reverse the whole array.

```viz
1 3 5 4 2      pivot 3 (3 < 5), rightmost > 3 is 4
1 4 5 3 2      swap
1 4 2 3 5      reverse suffix after pivot
```

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        i = len(nums) - 2
        while i >= 0 and nums[i] >= nums[i + 1]:
            i -= 1
        if i >= 0:
            j = len(nums) - 1
            while nums[j] <= nums[i]:
                j -= 1
            nums[i], nums[j] = nums[j], nums[i]
        nums[i + 1:] = reversed(nums[i + 1:])
# test
for a, b in (([1, 2, 3], [1, 3, 2]), ([3, 2, 1], [1, 2, 3]), ([1, 1, 5], [1, 5, 1]), ([1, 3, 5, 4, 2], [1, 4, 2, 3, 5])):
    Solution().nextPermutation(a)
    assert a == b
```

```java
class Solution {
    public void nextPermutation(int[] nums) {
        int i = nums.length - 2;
        while (i >= 0 && nums[i] >= nums[i + 1]) i--;
        if (i >= 0) {
            int j = nums.length - 1;
            while (nums[j] <= nums[i]) j--;
            swap(nums, i, j);
        }
        for (int l = i + 1, r = nums.length - 1; l < r; l++, r--) swap(nums, l, r);
    }

    private void swap(int[] a, int i, int j) { int t = a[i]; a[i] = a[j]; a[j] = t; }
}
// test
int[] a = {1, 3, 5, 4, 2};
new Solution().nextPermutation(a);
check(a, new int[]{1, 4, 2, 3, 5});
int[] b = {3, 2, 1};
new Solution().nextPermutation(b);
check(b, new int[]{1, 2, 3});
```

```go
func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	for l, r := i+1, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
// test
for _, tc := range [][2][]int{{{1, 2, 3}, {1, 3, 2}}, {{3, 2, 1}, {1, 2, 3}}, {{1, 1, 5}, {1, 5, 1}}, {{1, 3, 5, 4, 2}, {1, 4, 2, 3, 5}}} {
	nextPermutation(tc[0])
	assertEq(tc[0], tc[1])
}
```

## 118. Pascal's Triangle | Easy
related: 62
tags: dp

Each row starts and ends with 1; inner values are the sum of the two values above.

**Complexity:** O(n²).

```python
class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        rows = [[1]]
        for _ in range(numRows - 1):
            prev = rows[-1]
            rows.append([1] + [prev[i] + prev[i + 1] for i in range(len(prev) - 1)] + [1])
        return rows
# test
assert Solution().generate(5) == [[1], [1, 1], [1, 2, 1], [1, 3, 3, 1], [1, 4, 6, 4, 1]]
assert Solution().generate(1) == [[1]]
```

```java
class Solution {
    public List<List<Integer>> generate(int numRows) {
        List<List<Integer>> rows = new ArrayList<>();
        for (int r = 0; r < numRows; r++) {
            List<Integer> row = new ArrayList<>();
            for (int c = 0; c <= r; c++)
                row.add(c == 0 || c == r ? 1 : rows.get(r - 1).get(c - 1) + rows.get(r - 1).get(c));
            rows.add(row);
        }
        return rows;
    }
}
// test
check(new Solution().generate(5), "[[1], [1, 1], [1, 2, 1], [1, 3, 3, 1], [1, 4, 6, 4, 1]]");
```

```go
func generate(numRows int) [][]int {
	rows := [][]int{{1}}
	for len(rows) < numRows {
		prev := rows[len(rows)-1]
		row := []int{1}
		for i := 0; i+1 < len(prev); i++ {
			row = append(row, prev[i]+prev[i+1])
		}
		rows = append(rows, append(row, 1))
	}
	return rows
}
// test
assertEq(generate(5), [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}})
assertEq(generate(1), [][]int{{1}})
```

## 179. Largest Number | Medium
related: 1
tags: custom-sort

Sort strings with the comparator "`a` before `b` if `a + b > b + a`". This ordering is transitive, so the concatenation is maximal. Edge case: all zeros → `"0"`.

**Complexity:** O(n log n · L).

```python
class Solution:
    def largestNumber(self, nums: List[int]) -> str:
        s = sorted(map(str, nums), key=functools.cmp_to_key(lambda a, b: -1 if a + b > b + a else 1))
        res = "".join(s)
        return "0" if res[0] == "0" else res
# test
assert Solution().largestNumber([10, 2]) == "210"
assert Solution().largestNumber([3, 30, 34, 5, 9]) == "9534330"
assert Solution().largestNumber([0, 0]) == "0"
```

```java
class Solution {
    public String largestNumber(int[] nums) {
        String[] s = Arrays.stream(nums).mapToObj(String::valueOf).toArray(String[]::new);
        Arrays.sort(s, (a, b) -> (b + a).compareTo(a + b));
        if (s[0].equals("0")) return "0";
        return String.join("", s);
    }
}
// test
check(new Solution().largestNumber(new int[]{3, 30, 34, 5, 9}), "9534330");
check(new Solution().largestNumber(new int[]{0, 0}), "0");
```

```go
func largestNumber(nums []int) string {
	s := make([]string, len(nums))
	for i, x := range nums {
		s[i] = strconv.Itoa(x)
	}
	sort.Slice(s, func(i, j int) bool { return s[i]+s[j] > s[j]+s[i] })
	if s[0] == "0" {
		return "0"
	}
	return strings.Join(s, "")
}
// test
assertEq(largestNumber([]int{10, 2}), "210")
assertEq(largestNumber([]int{3, 30, 34, 5, 9}), "9534330")
assertEq(largestNumber([]int{0, 0}), "0")
```

## 228. Summary Ranges | Easy
related: 128
tags: array

Extend a run while the next number is consecutive; emit `"a->b"` or `"a"`.

**Complexity:** O(n).

```python
class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:
        res, i = [], 0
        while i < len(nums):
            j = i
            while j + 1 < len(nums) and nums[j + 1] == nums[j] + 1:
                j += 1
            res.append(str(nums[i]) if i == j else f"{nums[i]}->{nums[j]}")
            i = j + 1
        return res
# test
assert Solution().summaryRanges([0, 1, 2, 4, 5, 7]) == ["0->2", "4->5", "7"]
assert Solution().summaryRanges([]) == []
```

```java
class Solution {
    public List<String> summaryRanges(int[] nums) {
        List<String> res = new ArrayList<>();
        for (int i = 0, j; i < nums.length; i = j + 1) {
            j = i;
            while (j + 1 < nums.length && nums[j + 1] == nums[j] + 1) j++;
            res.add(i == j ? String.valueOf(nums[i]) : nums[i] + "->" + nums[j]);
        }
        return res;
    }
}
// test
check(new Solution().summaryRanges(new int[]{0, 1, 2, 4, 5, 7}), "[0->2, 4->5, 7]");
```

```go
func summaryRanges(nums []int) []string {
	res := []string{}
	for i := 0; i < len(nums); {
		j := i
		for j+1 < len(nums) && nums[j+1] == nums[j]+1 {
			j++
		}
		if i == j {
			res = append(res, strconv.Itoa(nums[i]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[i], nums[j]))
		}
		i = j + 1
	}
	return res
}
// test
assertEq(summaryRanges([]int{0, 1, 2, 4, 5, 7}), []string{"0->2", "4->5", "7"})
assertEq(summaryRanges([]int{}), []string{})
```

## 274. H-Index | Medium
related: 347
tags: counting-sort

h ≤ n, so bucket citations into `cnt[min(c, n)]`. Walk h from n down, accumulating papers with ≥ h citations; the first h where that count ≥ h is the answer.

**Complexity:** O(n).

```python
class Solution:
    def hIndex(self, citations: List[int]) -> int:
        n = len(citations)
        cnt = [0] * (n + 1)
        for c in citations:
            cnt[min(c, n)] += 1
        papers = 0
        for h in range(n, -1, -1):
            papers += cnt[h]
            if papers >= h:
                return h
        return 0
# test
assert Solution().hIndex([3, 0, 6, 1, 5]) == 3
assert Solution().hIndex([1, 3, 1]) == 1
```

```java
class Solution {
    public int hIndex(int[] citations) {
        int n = citations.length;
        int[] cnt = new int[n + 1];
        for (int c : citations) cnt[Math.min(c, n)]++;
        for (int h = n, papers = 0; h >= 0; h--) {
            papers += cnt[h];
            if (papers >= h) return h;
        }
        return 0;
    }
}
// test
check(new Solution().hIndex(new int[]{3, 0, 6, 1, 5}), 3);
```

```go
func hIndex(citations []int) int {
	n := len(citations)
	cnt := make([]int, n+1) // cnt[c] = papers with c citations (capped at n)
	for _, c := range citations {
		cnt[min(c, n)]++
	}
	papers := 0
	for h := n; h >= 0; h-- {
		papers += cnt[h]
		if papers >= h {
			return h
		}
	}
	return 0
}
// test
assertEq(hIndex([]int{3, 0, 6, 1, 5}), 3)
assertEq(hIndex([]int{1, 3, 1}), 1)
```

## 289. Game of Life | Medium
related: 73
tags: matrix, in-place

Update in place by storing the **next** state in bit 1 while bit 0 keeps the current state. Count live neighbours with `cell & 1`, then shift every cell right by one.

**Complexity:** O(m·n) time, O(1) space.

```python
class Solution:
    def gameOfLife(self, board: List[List[int]]) -> None:
        R, C = len(board), len(board[0])
        for r in range(R):
            for c in range(C):
                live = sum(board[i][j] & 1 for i in range(max(0, r - 1), min(R, r + 2))
                           for j in range(max(0, c - 1), min(C, c + 2))) - (board[r][c] & 1)
                if live == 3 or (live == 2 and board[r][c] & 1):
                    board[r][c] |= 2
        for r in range(R):
            for c in range(C):
                board[r][c] >>= 1
# test
b = [[0, 1, 0], [0, 0, 1], [1, 1, 1], [0, 0, 0]]
Solution().gameOfLife(b)
assert b == [[0, 0, 0], [1, 0, 1], [0, 1, 1], [0, 1, 0]]
b = [[1, 1], [1, 0]]; Solution().gameOfLife(b)
assert b == [[1, 1], [1, 1]]
```

```java
class Solution {
    public void gameOfLife(int[][] board) {
        int R = board.length, C = board[0].length;
        for (int r = 0; r < R; r++)
            for (int c = 0; c < C; c++) {
                int live = -(board[r][c] & 1);
                for (int i = Math.max(0, r - 1); i <= Math.min(R - 1, r + 1); i++)
                    for (int j = Math.max(0, c - 1); j <= Math.min(C - 1, c + 1); j++) live += board[i][j] & 1;
                if (live == 3 || (live == 2 && (board[r][c] & 1) == 1)) board[r][c] |= 2;
            }
        for (int[] row : board)
            for (int c = 0; c < C; c++) row[c] >>= 1;
    }
}
// test
int[][] b = {{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}};
new Solution().gameOfLife(b);
check(b, new int[][]{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}});
```

```go
func gameOfLife(board [][]int) {
	R, C := len(board), len(board[0])
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			live := 0
			for i := max(0, r-1); i < min(R, r+2); i++ {
				for j := max(0, c-1); j < min(C, c+2); j++ {
					live += board[i][j] & 1 // bit 0 = current state
				}
			}
			live -= board[r][c] & 1
			if live == 3 || (live == 2 && board[r][c]&1 == 1) {
				board[r][c] |= 2 // bit 1 = next state
			}
		}
	}
	for r := range board {
		for c := range board[r] {
			board[r][c] >>= 1
		}
	}
}
// test
b := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
gameOfLife(b)
assertEq(b, [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}})
b2 := [][]int{{1, 1}, {1, 0}}
gameOfLife(b2)
assertEq(b2, [][]int{{1, 1}, {1, 1}})
```

## 448. Find All Numbers Disappeared in an Array | Easy
related: 41, 268
tags: in-place

Values are in `[1, n]`: for each value `v`, mark index `v−1` by making it negative. Indices that stay positive are the missing numbers.

**Complexity:** O(n) time, O(1) extra space.

```python
class Solution:
    def findDisappearedNumbers(self, nums: List[int]) -> List[int]:
        for v in nums:
            i = abs(v) - 1
            nums[i] = -abs(nums[i])
        return [i + 1 for i, v in enumerate(nums) if v > 0]
# test
assert Solution().findDisappearedNumbers([4, 3, 2, 7, 8, 2, 3, 1]) == [5, 6]
assert Solution().findDisappearedNumbers([1, 1]) == [2]
```

```java
class Solution {
    public List<Integer> findDisappearedNumbers(int[] nums) {
        for (int v : nums) {
            int i = Math.abs(v) - 1;
            nums[i] = -Math.abs(nums[i]);
        }
        List<Integer> res = new ArrayList<>();
        for (int i = 0; i < nums.length; i++) if (nums[i] > 0) res.add(i + 1);
        return res;
    }
}
// test
check(new Solution().findDisappearedNumbers(new int[]{4, 3, 2, 7, 8, 2, 3, 1}), "[5, 6]");
```

```go
func findDisappearedNumbers(nums []int) []int {
	for _, v := range nums {
		i := abs(v) - 1
		nums[i] = -abs(nums[i]) // mark value i+1 as seen
	}
	res := []int{}
	for i, v := range nums {
		if v > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
// test
assertEq(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}), []int{5, 6})
assertEq(findDisappearedNumbers([]int{1, 1}), []int{2})
```

## 523. Continuous Subarray Sum | Medium
related: 560, 974
tags: prefix-sum, hash-map

A subarray sum is a multiple of k ⇔ two prefix sums have the same remainder mod k. Store the **first** index of each remainder (seed `0 → −1`); a repeat at least 2 positions later wins.

**Complexity:** O(n).

```python
class Solution:
    def checkSubarraySum(self, nums: List[int], k: int) -> bool:
        first = {0: -1}
        pre = 0
        for i, x in enumerate(nums):
            pre = (pre + x) % k
            if pre in first:
                if i - first[pre] >= 2:
                    return True
            else:
                first[pre] = i
        return False
# test
assert Solution().checkSubarraySum([23, 2, 4, 6, 7], 6) is True
assert Solution().checkSubarraySum([23, 2, 6, 4, 7], 13) is False
assert Solution().checkSubarraySum([5, 0, 0, 0], 3) is True
```

```java
class Solution {
    public boolean checkSubarraySum(int[] nums, int k) {
        Map<Integer, Integer> first = new HashMap<>();
        first.put(0, -1);
        int pre = 0;
        for (int i = 0; i < nums.length; i++) {
            pre = (pre + nums[i]) % k;
            Integer j = first.putIfAbsent(pre, i);
            if (j != null && i - j >= 2) return true;
        }
        return false;
    }
}
// test
check(new Solution().checkSubarraySum(new int[]{23, 2, 4, 6, 7}, 6), true);
check(new Solution().checkSubarraySum(new int[]{23, 2, 6, 4, 7}, 13), false);
```

```go
func checkSubarraySum(nums []int, k int) bool {
	first := map[int]int{0: -1} // prefix sum mod k -> first index
	pre := 0
	for i, x := range nums {
		pre = (pre + x) % k
		if j, ok := first[pre]; ok {
			if i-j >= 2 {
				return true
			}
		} else {
			first[pre] = i
		}
	}
	return false
}
// test
assertEq(checkSubarraySum([]int{23, 2, 4, 6, 7}, 6), true)
assertEq(checkSubarraySum([]int{23, 2, 6, 4, 7}, 13), false)
assertEq(checkSubarraySum([]int{5, 0, 0, 0}, 3), true)
```

## 953. Verifying an Alien Dictionary | Easy
related: 269
tags: hash-map

Map each letter to its rank; compare adjacent words at their first differing character. If none differ, the shorter word must come first.

**Complexity:** O(total characters).

```python
class Solution:
    def isAlienSorted(self, words: List[str], order: str) -> bool:
        rank = {ch: i for i, ch in enumerate(order)}
        keys = [[rank[ch] for ch in w] for w in words]
        return all(a <= b for a, b in zip(keys, keys[1:]))   # list comparison = lexicographic
# test
assert Solution().isAlienSorted(["hello", "leetcode"], "hlabcdefgijkmnopqrstuvwxyz") is True
assert Solution().isAlienSorted(["word", "world", "row"], "worldabcefghijkmnpqstuvxyz") is False
assert Solution().isAlienSorted(["apple", "app"], "abcdefghijklmnopqrstuvwxyz") is False
```

```java
class Solution {
    public boolean isAlienSorted(String[] words, String order) {
        int[] rank = new int[26];
        for (int i = 0; i < 26; i++) rank[order.charAt(i) - 'a'] = i;
        for (int w = 0; w + 1 < words.length; w++) {
            String a = words[w], b = words[w + 1];
            int i = 0;
            while (i < a.length() && i < b.length() && a.charAt(i) == b.charAt(i)) i++;
            if (i == a.length() || i == b.length()) { if (a.length() > b.length()) return false; }
            else if (rank[a.charAt(i) - 'a'] > rank[b.charAt(i) - 'a']) return false;
        }
        return true;
    }
}
// test
check(new Solution().isAlienSorted(new String[]{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"), true);
check(new Solution().isAlienSorted(new String[]{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"), false);
```

```go
func isAlienSorted(words []string, order string) bool {
	var rank [26]int
	for i := 0; i < len(order); i++ {
		rank[order[i]-'a'] = i
	}
	less := func(a, b string) bool { // a <= b in the alien order
		for i := 0; i < len(a) && i < len(b); i++ {
			if a[i] != b[i] {
				return rank[a[i]-'a'] < rank[b[i]-'a']
			}
		}
		return len(a) <= len(b)
	}
	for i := 1; i < len(words); i++ {
		if !less(words[i-1], words[i]) {
			return false
		}
	}
	return true
}
// test
assertEq(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"), true)
assertEq(isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz"), false)
assertEq(isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"), false)
```

## 974. Subarray Sums Divisible by K | Medium
related: 560, 523
tags: prefix-sum, counting

Count pairs of prefix sums with equal remainder mod k. Normalize negative remainders (`((x % k) + k) % k` in Java).

**Complexity:** O(n) time, O(k) space.

```python
class Solution:
    def subarraysDivByK(self, nums: List[int], k: int) -> int:
        cnt = [0] * k
        cnt[0] = 1
        pre = res = 0
        for x in nums:
            pre = (pre + x) % k           # Python % is already non-negative
            res += cnt[pre]
            cnt[pre] += 1
        return res
# test
assert Solution().subarraysDivByK([4, 5, 0, -2, -3, 1], 5) == 7
assert Solution().subarraysDivByK([5], 9) == 0
```

```java
class Solution {
    public int subarraysDivByK(int[] nums, int k) {
        int[] cnt = new int[k];
        cnt[0] = 1;
        int pre = 0, res = 0;
        for (int x : nums) {
            pre = ((pre + x) % k + k) % k;
            res += cnt[pre]++;
        }
        return res;
    }
}
// test
check(new Solution().subarraysDivByK(new int[]{4, 5, 0, -2, -3, 1}, 5), 7);
```

```go
func subarraysDivByK(nums []int, k int) int {
	cnt := make([]int, k)
	cnt[0] = 1
	pre, res := 0, 0
	for _, x := range nums {
		pre = ((pre+x)%k + k) % k // Go's % can be negative
		res += cnt[pre]
		cnt[pre]++
	}
	return res
}
// test
assertEq(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5), 7)
assertEq(subarraysDivByK([]int{5}, 9), 0)
```

## 383. Ransom Note | Easy
related: 242, 409
tags: counting

Count the magazine's letters once; every letter of the note must be available that many times.

**Complexity:** O(m + n) time, O(1) space (26 letters).

```python
class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        have = Counter(magazine)
        return all(have[c] >= k for c, k in Counter(ransomNote).items())
# test
assert Solution().canConstruct("a", "b") is False
assert Solution().canConstruct("aa", "ab") is False
assert Solution().canConstruct("aa", "aab") is True
assert Solution().canConstruct("", "x") is True
```

```java
class Solution {
    public boolean canConstruct(String ransomNote, String magazine) {
        int[] cnt = new int[26];
        for (char c : magazine.toCharArray()) cnt[c - 'a']++;
        for (char c : ransomNote.toCharArray()) if (--cnt[c - 'a'] < 0) return false;
        return true;
    }
}
// test
check(new Solution().canConstruct("aa", "ab"), false);
check(new Solution().canConstruct("aa", "aab"), true);
```

```go
func canConstruct(ransomNote string, magazine string) bool {
	var have [26]int
	for _, c := range magazine {
		have[c-'a']++
	}
	for _, c := range ransomNote {
		if have[c-'a']--; have[c-'a'] < 0 {
			return false
		}
	}
	return true
}
// test
assertEq(canConstruct("a", "b"), false)
assertEq(canConstruct("aa", "ab"), false)
assertEq(canConstruct("aa", "aab"), true)
assertEq(canConstruct("", "x"), true)
```

## 409. Longest Palindrome | Easy
related: 383, 5, 266
tags: counting, greedy

A palindrome uses characters in **pairs** (one on each side) plus at most **one** odd character in the middle. Sum the pairs; add 1 if any character is left over.

**Complexity:** O(n).

```python
class Solution:
    def longestPalindrome(self, s: str) -> int:
        pairs = sum(v // 2 * 2 for v in Counter(s).values())
        return pairs + 1 if pairs < len(s) else pairs
# test
assert Solution().longestPalindrome("abccccdd") == 7
assert Solution().longestPalindrome("a") == 1
assert Solution().longestPalindrome("Aa") == 1
assert Solution().longestPalindrome("bb") == 2
```

```java
class Solution {
    public int longestPalindrome(String s) {
        int[] cnt = new int[128];
        for (char c : s.toCharArray()) cnt[c]++;
        int pairs = 0;
        for (int v : cnt) pairs += v / 2 * 2;
        return pairs < s.length() ? pairs + 1 : pairs;
    }
}
// test
check(new Solution().longestPalindrome("abccccdd"), 7);
check(new Solution().longestPalindrome("Aa"), 1);
```

```go
func longestPalindrome(s string) int {
	var cnt [128]int
	for i := 0; i < len(s); i++ {
		cnt[s[i]]++
	}
	pairs := 0
	for _, v := range cnt {
		pairs += v / 2 * 2
	}
	if pairs < len(s) {
		return pairs + 1 // one odd letter can sit in the middle
	}
	return pairs
}
// test
assertEq(longestPalindrome("abccccdd"), 7)
assertEq(longestPalindrome("a"), 1)
assertEq(longestPalindrome("Aa"), 1)
assertEq(longestPalindrome("bb"), 2)
```

## 58. Length of Last Word | Easy
related: 151, 14
tags: string

Scan from the end: skip trailing spaces, then count characters until the next space. No split, O(1) extra space.

**Complexity:** O(n).

```python
class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        i = len(s) - 1
        while i >= 0 and s[i] == " ":
            i -= 1
        j = i
        while j >= 0 and s[j] != " ":
            j -= 1
        return i - j
# test
assert Solution().lengthOfLastWord("Hello World") == 5
assert Solution().lengthOfLastWord("   fly me   to   the moon  ") == 4
assert Solution().lengthOfLastWord("luffy is still joyboy") == 6
assert Solution().lengthOfLastWord("a") == 1
```

```java
class Solution {
    public int lengthOfLastWord(String s) {
        int i = s.length() - 1;
        while (i >= 0 && s.charAt(i) == ' ') i--;
        int j = i;
        while (j >= 0 && s.charAt(j) != ' ') j--;
        return i - j;
    }
}
// test
check(new Solution().lengthOfLastWord("   fly me   to   the moon  "), 4);
check(new Solution().lengthOfLastWord("a"), 1);
```

```go
func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	j := i
	for j >= 0 && s[j] != ' ' {
		j--
	}
	return i - j
}
// test
assertEq(lengthOfLastWord("Hello World"), 5)
assertEq(lengthOfLastWord("   fly me   to   the moon  "), 4)
assertEq(lengthOfLastWord("luffy is still joyboy"), 6)
assertEq(lengthOfLastWord("a"), 1)
```

## 6. Zigzag Conversion | Medium
related: 54, 498
tags: simulation, string

Simulate writing the zigzag: keep one buffer per row and a cursor that moves down, bounces at the last row, moves up, bounces at row 0. Reading the rows top to bottom gives the answer.

```viz
"PAYPALISHIRING", numRows = 4        cursor: 0 1 2 3 2 1 0 1 2 3 ...
row 0:  P     I     N
row 1:  A   L S   I G
row 2:  Y A   H R
row 3:  P     I
-> "PIN" + "ALSIG" + "YAHR" + "PI"
```

**Complexity:** O(n).

```python
class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1 or numRows >= len(s):
            return s
        rows = [[] for _ in range(numRows)]
        r, step = 0, 1
        for ch in s:
            rows[r].append(ch)
            if r == 0:
                step = 1
            elif r == numRows - 1:
                step = -1
            r += step
        return "".join("".join(row) for row in rows)
# test
assert Solution().convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR"
assert Solution().convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI"
assert Solution().convert("A", 1) == "A"
assert Solution().convert("AB", 1) == "AB"
assert Solution().convert("ABC", 2) == "ACB"
```

```java
class Solution {
    public String convert(String s, int numRows) {
        if (numRows == 1 || numRows >= s.length()) return s;
        StringBuilder[] rows = new StringBuilder[numRows];
        for (int i = 0; i < numRows; i++) rows[i] = new StringBuilder();
        int r = 0, step = 1;
        for (char c : s.toCharArray()) {
            rows[r].append(c);
            if (r == 0) step = 1;
            else if (r == numRows - 1) step = -1;
            r += step;
        }
        StringBuilder out = new StringBuilder();
        for (StringBuilder row : rows) out.append(row);
        return out.toString();
    }
}
// test
check(new Solution().convert("PAYPALISHIRING", 3), "PAHNAPLSIIGYIR");
check(new Solution().convert("PAYPALISHIRING", 4), "PINALSIGYAHRPI");
check(new Solution().convert("AB", 1), "AB");
```

```go
func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	rows := make([][]byte, numRows)
	r, step := 0, 1
	for i := 0; i < len(s); i++ {
		rows[r] = append(rows[r], s[i])
		if r == 0 {
			step = 1
		} else if r == numRows-1 {
			step = -1
		}
		r += step
	}
	var sb strings.Builder
	for _, row := range rows {
		sb.Write(row)
	}
	return sb.String()
}
// test
assertEq(convert("PAYPALISHIRING", 3), "PAHNAPLSIIGYIR")
assertEq(convert("PAYPALISHIRING", 4), "PINALSIGYAHRPI")
assertEq(convert("A", 1), "A")
assertEq(convert("AB", 1), "AB")
assertEq(convert("ABC", 2), "ACB")
```

## 28. Find the Index of the First Occurrence in a String | Easy
related: 459, 214, 1392
tags: string, kmp

The brute force (try every start, compare up to m chars) is O(n·m) and fine for the given limits — say so, then offer **KMP** for O(n + m): precompute `lps[i]` = length of the longest proper prefix of `needle[:i+1]` that is also its suffix. On a mismatch, fall back to `lps[k-1]` instead of restarting, so the haystack pointer never moves backwards.

```viz
needle = "aab"  lps = [0, 1, 0]
haystack a a a a a b
           ^ mismatch at needle[2] after "aa" -> k = lps[1] = 1, keep i
```

**Complexity:** O(n + m) time, O(m) space.

```python
class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        lps, k = [0] * len(needle), 0
        for i in range(1, len(needle)):
            while k and needle[i] != needle[k]:
                k = lps[k - 1]
            if needle[i] == needle[k]:
                k += 1
            lps[i] = k
        k = 0
        for i, ch in enumerate(haystack):
            while k and ch != needle[k]:
                k = lps[k - 1]
            if ch == needle[k]:
                k += 1
                if k == len(needle):
                    return i - k + 1
        return -1
# test
assert Solution().strStr("sadbutsad", "sad") == 0
assert Solution().strStr("leetcode", "leeto") == -1
assert Solution().strStr("aaaaab", "aab") == 3
assert Solution().strStr("mississippi", "issip") == 4
assert Solution().strStr("a", "a") == 0
assert Solution().strStr("abc", "c") == 2
```

```java
class Solution {
    public int strStr(String haystack, String needle) {
        int m = needle.length();
        int[] lps = new int[m];
        for (int i = 1, k = 0; i < m; i++) {
            while (k > 0 && needle.charAt(i) != needle.charAt(k)) k = lps[k - 1];
            if (needle.charAt(i) == needle.charAt(k)) k++;
            lps[i] = k;
        }
        for (int i = 0, k = 0; i < haystack.length(); i++) {
            while (k > 0 && haystack.charAt(i) != needle.charAt(k)) k = lps[k - 1];
            if (haystack.charAt(i) == needle.charAt(k) && ++k == m) return i - m + 1;
        }
        return -1;
    }
}
// test
check(new Solution().strStr("sadbutsad", "sad"), 0);
check(new Solution().strStr("aaaaab", "aab"), 3);
check(new Solution().strStr("leetcode", "leeto"), -1);
```

```go
func strStr(haystack string, needle string) int {
	lps, k := make([]int, len(needle)), 0 // KMP failure function
	for i := 1; i < len(needle); i++ {
		for k > 0 && needle[i] != needle[k] {
			k = lps[k-1]
		}
		if needle[i] == needle[k] {
			k++
		}
		lps[i] = k
	}
	k = 0
	for i := 0; i < len(haystack); i++ {
		for k > 0 && haystack[i] != needle[k] {
			k = lps[k-1]
		}
		if haystack[i] == needle[k] {
			k++
			if k == len(needle) {
				return i - k + 1
			}
		}
	}
	return -1
}
// test
assertEq(strStr("sadbutsad", "sad"), 0)
assertEq(strStr("leetcode", "leeto"), -1)
assertEq(strStr("aaaaab", "aab"), 3)
assertEq(strStr("mississippi", "issip"), 4)
assertEq(strStr("a", "a"), 0)
assertEq(strStr("abc", "c"), 2)
```

## 290. Word Pattern | Easy
related: 205, 242
tags: hash-map, bijection

Same idea as Isomorphic Strings: the mapping letter ↔ word must be a **bijection**, so keep both directions and reject any conflict. Also reject different lengths.

**Complexity:** O(n).

```python
class Solution:
    def wordPattern(self, pattern: str, s: str) -> bool:
        words = s.split()
        if len(words) != len(pattern):
            return False
        p2w, w2p = {}, {}
        for p, w in zip(pattern, words):
            if p2w.setdefault(p, w) != w or w2p.setdefault(w, p) != p:
                return False
        return True
# test
assert Solution().wordPattern("abba", "dog cat cat dog") is True
assert Solution().wordPattern("abba", "dog cat cat fish") is False
assert Solution().wordPattern("aaaa", "dog cat cat dog") is False
assert Solution().wordPattern("abba", "dog dog dog dog") is False
assert Solution().wordPattern("aaa", "aa aa aa aa") is False
```

```java
class Solution {
    public boolean wordPattern(String pattern, String s) {
        String[] words = s.split(" ");
        if (words.length != pattern.length()) return false;
        Map<Character, String> p2w = new HashMap<>();
        Map<String, Character> w2p = new HashMap<>();
        for (int i = 0; i < words.length; i++) {
            char p = pattern.charAt(i);
            String w = words[i];
            if (!p2w.computeIfAbsent(p, x -> w).equals(w)) return false;
            if (w2p.computeIfAbsent(w, x -> p).charValue() != p) return false;
        }
        return true;
    }
}
// test
check(new Solution().wordPattern("abba", "dog cat cat dog"), true);
check(new Solution().wordPattern("abba", "dog dog dog dog"), false);
```

```go
func wordPattern(pattern string, s string) bool {
	words := strings.Fields(s)
	if len(words) != len(pattern) {
		return false
	}
	p2w, w2p := map[byte]string{}, map[string]byte{}
	for i, w := range words {
		p := pattern[i]
		if x, ok := p2w[p]; ok && x != w {
			return false
		}
		if x, ok := w2p[w]; ok && x != p {
			return false
		}
		p2w[p], w2p[w] = w, p
	}
	return true
}
// test
assertEq(wordPattern("abba", "dog cat cat dog"), true)
assertEq(wordPattern("abba", "dog cat cat fish"), false)
assertEq(wordPattern("aaaa", "dog cat cat dog"), false)
assertEq(wordPattern("abba", "dog dog dog dog"), false)
assertEq(wordPattern("aaa", "aa aa aa aa"), false)
```

## 1431. Kids With the Greatest Number of Candies | Easy
related: 2215
tags: array

A kid can become the (tied) maximum iff `candies[i] + extra >= max(candies)`.

**Complexity:** O(n).

```python
class Solution:
    def kidsWithCandies(self, candies: List[int], extraCandies: int) -> List[bool]:
        best = max(candies)
        return [c + extraCandies >= best for c in candies]
# test
assert Solution().kidsWithCandies([2, 3, 5, 1, 3], 3) == [True, True, True, False, True]
assert Solution().kidsWithCandies([4, 2, 1, 1, 2], 1) == [True, False, False, False, False]
assert Solution().kidsWithCandies([12, 1, 12], 10) == [True, False, True]
```

```java
class Solution {
    public List<Boolean> kidsWithCandies(int[] candies, int extraCandies) {
        int best = Arrays.stream(candies).max().getAsInt();
        List<Boolean> res = new ArrayList<>();
        for (int c : candies) res.add(c + extraCandies >= best);
        return res;
    }
}
// test
check(new Solution().kidsWithCandies(new int[]{2, 3, 5, 1, 3}, 3), List.of(true, true, true, false, true));
```

```go
func kidsWithCandies(candies []int, extraCandies int) []bool {
	best := 0
	for _, c := range candies {
		best = max(best, c)
	}
	res := make([]bool, len(candies))
	for i, c := range candies {
		res[i] = c+extraCandies >= best
	}
	return res
}
// test
assertEq(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3), []bool{true, true, true, false, true})
assertEq(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1), []bool{true, false, false, false, false})
assertEq(kidsWithCandies([]int{12, 1, 12}, 10), []bool{true, false, true})
```

## 2215. Find the Difference of Two Arrays | Easy
related: 349, 1431
tags: hash-set

Set difference both ways: `set1 − set2` and `set2 − set1` (duplicates collapse automatically). Either list may be returned in any order.

**Complexity:** O(n + m).

```python
class Solution:
    def findDifference(self, nums1: List[int], nums2: List[int]) -> List[List[int]]:
        a, b = set(nums1), set(nums2)
        return [list(a - b), list(b - a)]
# test
assert [sorted(x) for x in Solution().findDifference([1, 2, 3], [2, 4, 6])] == [[1, 3], [4, 6]]
assert [sorted(x) for x in Solution().findDifference([1, 2, 3, 3], [1, 1, 2, 2])] == [[3], []]
assert [sorted(x) for x in Solution().findDifference([-1], [-1])] == [[], []]
```

```java
class Solution {
    public List<List<Integer>> findDifference(int[] nums1, int[] nums2) {
        Set<Integer> a = new HashSet<>(), b = new HashSet<>();
        for (int x : nums1) a.add(x);
        for (int x : nums2) b.add(x);
        List<Integer> onlyA = new ArrayList<>(), onlyB = new ArrayList<>();
        for (int x : a) if (!b.contains(x)) onlyA.add(x);
        for (int x : b) if (!a.contains(x)) onlyB.add(x);
        return List.of(onlyA, onlyB);
    }
}
// test
List<List<Integer>> d = new Solution().findDifference(new int[]{1, 2, 3}, new int[]{2, 4, 6});
for (List<Integer> l : d) Collections.sort(l);
check(d, List.of(List.of(1, 3), List.of(4, 6)));
```

```go
func findDifference(nums1 []int, nums2 []int) [][]int {
	a, b := map[int]bool{}, map[int]bool{}
	for _, x := range nums1 {
		a[x] = true
	}
	for _, x := range nums2 {
		b[x] = true
	}
	minus := func(x, y map[int]bool) []int {
		out := []int{}
		for v := range x {
			if !y[v] {
				out = append(out, v)
			}
		}
		return out
	}
	return [][]int{minus(a, b), minus(b, a)}
}
// test
r := findDifference([]int{1, 2, 3}, []int{2, 4, 6})
assertUnordered(r[0], []int{1, 3})
assertUnordered(r[1], []int{4, 6})
r = findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2})
assertEq(r, [][]int{{3}, {}})
assertEq(findDifference([]int{-1}, []int{-1}), [][]int{{}, {}})
```

## 1207. Unique Number of Occurrences | Easy
related: 1657, 451
tags: counting

Count each value, then check that no two values share a count: the number of distinct counts must equal the number of distinct values.

**Complexity:** O(n).

```python
class Solution:
    def uniqueOccurrences(self, arr: List[int]) -> bool:
        counts = list(Counter(arr).values())
        return len(counts) == len(set(counts))
# test
assert Solution().uniqueOccurrences([1, 2, 2, 1, 1, 3]) is True
assert Solution().uniqueOccurrences([1, 2]) is False
assert Solution().uniqueOccurrences([-3, 0, 1, -3, 1, 1, 1, -3, 10, 0]) is True
```

```java
class Solution {
    public boolean uniqueOccurrences(int[] arr) {
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int x : arr) cnt.merge(x, 1, Integer::sum);
        return new HashSet<>(cnt.values()).size() == cnt.size();
    }
}
// test
check(new Solution().uniqueOccurrences(new int[]{1, 2, 2, 1, 1, 3}), true);
check(new Solution().uniqueOccurrences(new int[]{1, 2}), false);
```

```go
func uniqueOccurrences(arr []int) bool {
	cnt := map[int]int{}
	for _, x := range arr {
		cnt[x]++
	}
	seen := map[int]bool{}
	for _, c := range cnt {
		if seen[c] {
			return false
		}
		seen[c] = true
	}
	return true
}
// test
assertEq(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}), true)
assertEq(uniqueOccurrences([]int{1, 2}), false)
assertEq(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}), true)
```

## 1657. Determine if Two Strings Are Close | Medium
related: 242, 1207
tags: counting

Operation 1 (swap) means order doesn't matter; operation 2 (swap two letters' roles) means the **multiset of counts** can be reassigned among letters already present. So: same **set of letters** and same **sorted list of counts**.

**Complexity:** O(n) (+ sorting 26 counts).

```python
class Solution:
    def closeStrings(self, word1: str, word2: str) -> bool:
        c1, c2 = Counter(word1), Counter(word2)
        return c1.keys() == c2.keys() and sorted(c1.values()) == sorted(c2.values())
# test
assert Solution().closeStrings("abc", "bca") is True
assert Solution().closeStrings("a", "aa") is False
assert Solution().closeStrings("cabbba", "abbccc") is True
assert Solution().closeStrings("uau", "ssx") is False
```

```java
class Solution {
    public boolean closeStrings(String word1, String word2) {
        if (word1.length() != word2.length()) return false;
        int[] a = new int[26], b = new int[26];
        for (char c : word1.toCharArray()) a[c - 'a']++;
        for (char c : word2.toCharArray()) b[c - 'a']++;
        for (int i = 0; i < 26; i++) if ((a[i] == 0) != (b[i] == 0)) return false;
        Arrays.sort(a);
        Arrays.sort(b);
        return Arrays.equals(a, b);
    }
}
// test
check(new Solution().closeStrings("cabbba", "abbccc"), true);
check(new Solution().closeStrings("uau", "ssx"), false);
```

```go
func closeStrings(word1 string, word2 string) bool {
	var c1, c2 [26]int
	for _, c := range word1 {
		c1[c-'a']++
	}
	for _, c := range word2 {
		c2[c-'a']++
	}
	for i := range c1 {
		if (c1[i] == 0) != (c2[i] == 0) { // same set of letters
			return false
		}
	}
	a, b := c1[:], c2[:]
	sort.Ints(a)
	sort.Ints(b)
	return c1 == c2 // same multiset of counts (arrays sorted in place)
}
// test
assertEq(closeStrings("abc", "bca"), true)
assertEq(closeStrings("a", "aa"), false)
assertEq(closeStrings("cabbba", "abbccc"), true)
assertEq(closeStrings("uau", "ssx"), false)
```

## 2352. Equal Row and Column Pairs | Medium
related: 36
tags: hash-map

Hash every row as a tuple and count; then for each column (as a tuple) add how many rows equal it.

**Complexity:** O(n²).

```python
class Solution:
    def equalPairs(self, grid: List[List[int]]) -> int:
        rows = Counter(tuple(r) for r in grid)
        return sum(rows[col] for col in zip(*grid))
# test
assert Solution().equalPairs([[3, 2, 1], [1, 7, 6], [2, 7, 7]]) == 1
assert Solution().equalPairs([[3, 1, 2, 2], [1, 4, 4, 5], [2, 4, 2, 2], [2, 4, 2, 2]]) == 3
assert Solution().equalPairs([[1]]) == 1
```

```java
class Solution {
    public int equalPairs(int[][] grid) {
        int n = grid.length, res = 0;
        Map<String, Integer> rows = new HashMap<>();
        for (int[] r : grid) rows.merge(Arrays.toString(r), 1, Integer::sum);
        for (int c = 0; c < n; c++) {
            int[] col = new int[n];
            for (int r = 0; r < n; r++) col[r] = grid[r][c];
            res += rows.getOrDefault(Arrays.toString(col), 0);
        }
        return res;
    }
}
// test
check(new Solution().equalPairs(new int[][]{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}), 3);
```

```go
func equalPairs(grid [][]int) int {
	n := len(grid)
	rows := map[string]int{}
	for _, r := range grid {
		rows[fmt.Sprint(r)]++
	}
	res := 0
	for c := 0; c < n; c++ {
		col := make([]int, n)
		for r := 0; r < n; r++ {
			col[r] = grid[r][c]
		}
		res += rows[fmt.Sprint(col)]
	}
	return res
}
// test
assertEq(equalPairs([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}), 1)
assertEq(equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}), 3)
assertEq(equalPairs([][]int{{1}}), 1)
```

## 1929. Concatenation of Array | Easy
related: 217, 1
tags: array

Build `nums + nums`: allocate `2n` slots and write each element twice, at `i` and `i + n`.

**Complexity:** O(n).

```python
class Solution:
    def getConcatenation(self, nums: List[int]) -> List[int]:
        return nums + nums
# test
assert Solution().getConcatenation([1, 2, 1]) == [1, 2, 1, 1, 2, 1]
assert Solution().getConcatenation([1, 3, 2, 1]) == [1, 3, 2, 1, 1, 3, 2, 1]
assert Solution().getConcatenation([5]) == [5, 5]
```

```java
class Solution {
    public int[] getConcatenation(int[] nums) {
        int n = nums.length;
        int[] ans = new int[2 * n];
        for (int i = 0; i < n; i++) ans[i] = ans[i + n] = nums[i];
        return ans;
    }
}
// test
check(new Solution().getConcatenation(new int[]{1, 2, 1}), new int[]{1, 2, 1, 1, 2, 1});
```

```go
func getConcatenation(nums []int) []int {
	return append(append([]int{}, nums...), nums...)
}
// test
assertEq(getConcatenation([]int{1, 2, 1}), []int{1, 2, 1, 1, 2, 1})
assertEq(getConcatenation([]int{1, 3, 2, 1}), []int{1, 3, 2, 1, 1, 3, 2, 1})
assertEq(getConcatenation([]int{5}), []int{5, 5})
```

## 912. Sort an Array | Medium
related: 75, 148, 215
tags: merge-sort, divide-and-conquer

Built-in sort is off limits, so implement an O(n log n) sort. **Merge sort**: sort each half recursively, then merge the two sorted halves with two pointers (stable, guaranteed O(n log n), O(n) extra space). Quicksort is in-place but needs a random pivot to avoid O(n²) on sorted/duplicate-heavy input; heap sort is in-place and O(n log n) but not stable.

**Complexity:** O(n log n) time, O(n) space.

```python
class Solution:
    def sortArray(self, nums: List[int]) -> List[int]:
        if len(nums) <= 1:
            return nums
        mid = len(nums) // 2
        left, right = self.sortArray(nums[:mid]), self.sortArray(nums[mid:])
        out, i, j = [], 0, 0
        while i < len(left) and j < len(right):
            if left[i] <= right[j]:
                out.append(left[i])
                i += 1
            else:
                out.append(right[j])
                j += 1
        return out + left[i:] + right[j:]
# test
assert Solution().sortArray([5, 2, 3, 1]) == [1, 2, 3, 5]
assert Solution().sortArray([5, 1, 1, 2, 0, 0]) == [0, 0, 1, 1, 2, 5]
assert Solution().sortArray([1]) == [1]
rnd = random.Random(7)
data = [rnd.randint(-50, 50) for _ in range(300)]
assert Solution().sortArray(data[:]) == sorted(data)
```

```java
class Solution {
    public int[] sortArray(int[] nums) {
        mergeSort(nums, new int[nums.length], 0, nums.length - 1);
        return nums;
    }

    private void mergeSort(int[] a, int[] tmp, int lo, int hi) {
        if (lo >= hi) return;
        int mid = (lo + hi) / 2;
        mergeSort(a, tmp, lo, mid);
        mergeSort(a, tmp, mid + 1, hi);
        int i = lo, j = mid + 1, k = lo;
        while (i <= mid && j <= hi) tmp[k++] = a[i] <= a[j] ? a[i++] : a[j++];
        while (i <= mid) tmp[k++] = a[i++];
        while (j <= hi) tmp[k++] = a[j++];
        System.arraycopy(tmp, lo, a, lo, hi - lo + 1);
    }
}
// test
check(new Solution().sortArray(new int[]{5, 1, 1, 2, 0, 0}), new int[]{0, 0, 1, 1, 2, 5});
```

```go
func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left, right := sortArray(nums[:mid:mid]), sortArray(nums[mid:])
	out := make([]int, 0, len(nums))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			out = append(out, left[i])
			i++
		} else {
			out = append(out, right[j])
			j++
		}
	}
	out = append(out, left[i:]...)
	return append(out, right[j:]...)
}
// test
assertEq(sortArray([]int{5, 2, 3, 1}), []int{1, 2, 3, 5})
assertEq(sortArray([]int{5, 1, 1, 2, 0, 0}), []int{0, 0, 1, 1, 2, 5})
assertEq(sortArray([]int{1}), []int{1})
rnd := rand.New(rand.NewSource(7))
data := make([]int, 300)
for i := range data {
	data[i] = rnd.Intn(101) - 50
}
want := append([]int{}, data...)
sort.Ints(want)
assertEq(sortArray(data), want)
```

## 229. Majority Element II | Medium
related: 169
tags: boyer-moore

At most **two** values can appear more than n/3 times. Extended Boyer–Moore voting keeps two candidates with counters: a match increments its counter, an empty slot adopts the value, otherwise both counters drop (cancelling a triple of distinct values). A second pass verifies the candidates' real counts.

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def majorityElement(self, nums: List[int]) -> List[int]:
        c1, c2, n1, n2 = None, None, 0, 0
        for x in nums:
            if x == c1:
                n1 += 1
            elif x == c2:
                n2 += 1
            elif n1 == 0:
                c1, n1 = x, 1
            elif n2 == 0:
                c2, n2 = x, 1
            else:
                n1, n2 = n1 - 1, n2 - 1
        return [c for c in (c1, c2) if c is not None and nums.count(c) > len(nums) // 3]
# test
assert Solution().majorityElement([3, 2, 3]) == [3]
assert Solution().majorityElement([1]) == [1]
assert sorted(Solution().majorityElement([1, 2])) == [1, 2]
assert sorted(Solution().majorityElement([2, 2, 1, 1, 1, 2, 2])) == [1, 2]
assert Solution().majorityElement([1, 2, 3, 4]) == []
```

```java
class Solution {
    public List<Integer> majorityElement(int[] nums) {
        int c1 = 0, c2 = 1, n1 = 0, n2 = 0;           // distinct placeholders; verified below
        for (int x : nums) {
            if (x == c1) n1++;
            else if (x == c2) n2++;
            else if (n1 == 0) { c1 = x; n1 = 1; }
            else if (n2 == 0) { c2 = x; n2 = 1; }
            else { n1--; n2--; }
        }
        List<Integer> res = new ArrayList<>();
        for (int c : new int[]{c1, c2}) {
            int cnt = 0;
            for (int x : nums) if (x == c) cnt++;
            if (cnt > nums.length / 3) res.add(c);
        }
        return res;
    }
}
// test
List<Integer> me = new Solution().majorityElement(new int[]{2, 2, 1, 1, 1, 2, 2});
Collections.sort(me);
check(me, List.of(1, 2));
check(new Solution().majorityElement(new int[]{3, 2, 3}), List.of(3));
```

```go
func majorityElement(nums []int) []int {
	c1, c2, n1, n2 := 0, 1, 0, 0 // two candidates (distinct initial values)
	for _, x := range nums {
		switch {
		case x == c1:
			n1++
		case x == c2:
			n2++
		case n1 == 0:
			c1, n1 = x, 1
		case n2 == 0:
			c2, n2 = x, 1
		default:
			n1, n2 = n1-1, n2-1
		}
	}
	res := []int{}
	for _, c := range []int{c1, c2} {
		cnt := 0
		for _, x := range nums {
			if x == c {
				cnt++
			}
		}
		if cnt > len(nums)/3 {
			res = append(res, c)
		}
	}
	return res
}
// test
assertEq(majorityElement([]int{3, 2, 3}), []int{3})
assertEq(majorityElement([]int{1}), []int{1})
assertUnordered(majorityElement([]int{1, 2}), []int{1, 2})
assertUnordered(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}), []int{1, 2})
assertEq(majorityElement([]int{1, 2, 3, 4}), []int{})
```
