# Backtracking

**When to use:** "generate all …" (subsets, permutations, combinations, partitions, boards). Explore a decision tree depth-first: **choose → recurse → un-choose**.

```python
def backtrack(start, path):
    if is_solution(path):
        res.append(path[:])           # copy!
        return                         # (subsets: record at every node, no return)
    for i in range(start, len(choices)):
        if i > start and choices[i] == choices[i - 1]:
            continue                   # skip duplicates (sorted input)
        path.append(choices[i])
        backtrack(i + 1, path)         # i  -> reuse allowed;  i+1 -> each used once
        path.pop()
```

```viz
subsets of [1,2,3]               []
                       /          |        \
                    [1]          [2]       [3]
                   /   \          |
               [1,2]  [1,3]     [2,3]
                 |
              [1,2,3]
```

| Problem type | start index | dedupe | used[] |
|---|---|---|---|
| Subsets / combinations | `i + 1` | sort + skip `i > start` dup | – |
| Combination sum (reuse) | `i` | – | – |
| Permutations | `0` | sort + skip if prev dup unused | yes |

## 78. Subsets | Medium
related: 90, 77, 46
tags: subsets

Every node of the decision tree is a subset — record the path on entry.

**Complexity:** O(n · 2ⁿ).

```python
class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        res, path = [], []

        def bt(start):
            res.append(path[:])
            for i in range(start, len(nums)):
                path.append(nums[i])
                bt(i + 1)
                path.pop()

        bt(0)
        return res
# test
assert sorted(map(sorted, Solution().subsets([1, 2, 3]))) == sorted(map(sorted, [[], [1], [1, 2], [1, 2, 3], [1, 3], [2], [2, 3], [3]]))
assert sorted(map(sorted, Solution().subsets([0]))) == [[], [0]]
```

```java
class Solution {
    public List<List<Integer>> subsets(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        bt(nums, 0, new ArrayList<>(), res);
        return res;
    }

    private void bt(int[] nums, int start, List<Integer> path, List<List<Integer>> res) {
        res.add(new ArrayList<>(path));
        for (int i = start; i < nums.length; i++) {
            path.add(nums[i]);
            bt(nums, i + 1, path, res);
            path.remove(path.size() - 1);
        }
    }
}
// test
check(new Solution().subsets(new int[]{1, 2, 3}), "[[], [1], [1, 2], [1, 2, 3], [1, 3], [2], [2, 3], [3]]");
```

```go
func subsets(nums []int) [][]int {
	res, path := [][]int{}, []int{}
	var bt func(start int)
	bt = func(start int) {
		res = append(res, append([]int{}, path...)) // every node of the tree is a subset
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			bt(i + 1)
			path = path[:len(path)-1]
		}
	}
	bt(0)
	return res
}
// test
assertUnorderedDeep(subsets([]int{1, 2, 3}), [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}})
assertUnorderedDeep(subsets([]int{0}), [][]int{{}, {0}})
```

## 90. Subsets II | Medium
related: 78, 40, 47
tags: subsets, dedupe

Sort; at the same depth skip a value equal to its left neighbour (`i > start`), so each distinct choice branches once.

**Complexity:** O(n · 2ⁿ).

```python
class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        res, path = [], []

        def bt(start):
            res.append(path[:])
            for i in range(start, len(nums)):
                if i > start and nums[i] == nums[i - 1]:
                    continue
                path.append(nums[i])
                bt(i + 1)
                path.pop()

        bt(0)
        return res
# test
assert sorted(map(sorted, Solution().subsetsWithDup([1, 2, 2]))) == sorted(map(sorted, [[], [1], [1, 2], [1, 2, 2], [2], [2, 2]]))
assert sorted(map(sorted, Solution().subsetsWithDup([0]))) == [[], [0]]
assert sorted(map(sorted, Solution().subsetsWithDup([4, 4, 1, 4]))) == [[], [1], [1, 4], [1, 4, 4], [1, 4, 4, 4], [4], [4, 4], [4, 4, 4]]
```

```java
class Solution {
    public List<List<Integer>> subsetsWithDup(int[] nums) {
        Arrays.sort(nums);
        List<List<Integer>> res = new ArrayList<>();
        bt(nums, 0, new ArrayList<>(), res);
        return res;
    }

    private void bt(int[] nums, int start, List<Integer> path, List<List<Integer>> res) {
        res.add(new ArrayList<>(path));
        for (int i = start; i < nums.length; i++) {
            if (i > start && nums[i] == nums[i - 1]) continue;
            path.add(nums[i]);
            bt(nums, i + 1, path, res);
            path.remove(path.size() - 1);
        }
    }
}
// test
check(new Solution().subsetsWithDup(new int[]{1, 2, 2}), "[[], [1], [1, 2], [1, 2, 2], [2], [2, 2]]");
```

```go
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res, path := [][]int{}, []int{}
	var bt func(start int)
	bt = func(start int) {
		res = append(res, append([]int{}, path...))
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] { // same value at the same depth: skip
				continue
			}
			path = append(path, nums[i])
			bt(i + 1)
			path = path[:len(path)-1]
		}
	}
	bt(0)
	return res
}
// test
assertUnorderedDeep(subsetsWithDup([]int{1, 2, 2}), [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}})
assertUnorderedDeep(subsetsWithDup([]int{0}), [][]int{{}, {0}})
assertUnorderedDeep(subsetsWithDup([]int{4, 4, 1, 4}), [][]int{{}, {1}, {1, 4}, {1, 4, 4}, {1, 4, 4, 4}, {4}, {4, 4}, {4, 4, 4}})
```

## 39. Combination Sum | Medium
related: 40, 216, 377, 322
tags: blind75, combinations

Unlimited reuse → recurse with the **same** `i`. Sorting lets us break as soon as a candidate exceeds the remainder.

**Complexity:** exponential, bounded by O(n^(target/min)).

```python
class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates.sort()
        res, path = [], []

        def bt(start, rest):
            if rest == 0:
                res.append(path[:])
                return
            for i in range(start, len(candidates)):
                c = candidates[i]
                if c > rest:
                    break
                path.append(c)
                bt(i, rest - c)          # i: same number may be reused
                path.pop()

        bt(0, target)
        return res
# test
assert sorted(map(sorted, Solution().combinationSum([2, 3, 6, 7], 7))) == sorted(map(sorted, [[2, 2, 3], [7]]))
assert sorted(map(sorted, Solution().combinationSum([2, 3, 5], 8))) == sorted(map(sorted, [[2, 2, 2, 2], [2, 3, 3], [3, 5]]))
```

```java
class Solution {
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        Arrays.sort(candidates);
        List<List<Integer>> res = new ArrayList<>();
        bt(candidates, 0, target, new ArrayList<>(), res);
        return res;
    }

    private void bt(int[] c, int start, int rest, List<Integer> path, List<List<Integer>> res) {
        if (rest == 0) { res.add(new ArrayList<>(path)); return; }
        for (int i = start; i < c.length && c[i] <= rest; i++) {
            path.add(c[i]);
            bt(c, i, rest - c[i], path, res);
            path.remove(path.size() - 1);
        }
    }
}
// test
check(new Solution().combinationSum(new int[]{2, 3, 6, 7}, 7), "[[2, 2, 3], [7]]");
```

```go
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res, path := [][]int{}, []int{}
	var bt func(start, rest int)
	bt = func(start, rest int) {
		if rest == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := start; i < len(candidates) && candidates[i] <= rest; i++ {
			path = append(path, candidates[i])
			bt(i, rest-candidates[i]) // i: the same number may be reused
			path = path[:len(path)-1]
		}
	}
	bt(0, target)
	return res
}
// test
assertUnorderedDeep(combinationSum([]int{2, 3, 6, 7}, 7), [][]int{{2, 2, 3}, {7}})
assertUnorderedDeep(combinationSum([]int{2, 3, 5}, 8), [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}})
```

## 40. Combination Sum II | Medium
related: 39, 90
tags: combinations, dedupe

Each number used once (`i + 1`) and the input has duplicates → sort and skip equal neighbours at the same depth.

**Complexity:** O(2ⁿ).

```python
class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates.sort()
        res, path = [], []

        def bt(start, rest):
            if rest == 0:
                res.append(path[:])
                return
            for i in range(start, len(candidates)):
                if i > start and candidates[i] == candidates[i - 1]:
                    continue
                if candidates[i] > rest:
                    break
                path.append(candidates[i])
                bt(i + 1, rest - candidates[i])
                path.pop()

        bt(0, target)
        return res
# test
assert sorted(map(sorted, Solution().combinationSum2([10, 1, 2, 7, 6, 1, 5], 8))) == sorted(map(sorted, [[1, 1, 6], [1, 2, 5], [1, 7], [2, 6]]))
assert sorted(map(sorted, Solution().combinationSum2([2, 5, 2, 1, 2], 5))) == [[1, 2, 2], [5]]
assert Solution().combinationSum2([1], 2) == []
```

```java
class Solution {
    public List<List<Integer>> combinationSum2(int[] candidates, int target) {
        Arrays.sort(candidates);
        List<List<Integer>> res = new ArrayList<>();
        bt(candidates, 0, target, new ArrayList<>(), res);
        return res;
    }

    private void bt(int[] c, int start, int rest, List<Integer> path, List<List<Integer>> res) {
        if (rest == 0) { res.add(new ArrayList<>(path)); return; }
        for (int i = start; i < c.length && c[i] <= rest; i++) {
            if (i > start && c[i] == c[i - 1]) continue;
            path.add(c[i]);
            bt(c, i + 1, rest - c[i], path, res);
            path.remove(path.size() - 1);
        }
    }
}
// test
check(new Solution().combinationSum2(new int[]{10, 1, 2, 7, 6, 1, 5}, 8), "[[1, 1, 6], [1, 2, 5], [1, 7], [2, 6]]");
```

```go
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res, path := [][]int{}, []int{}
	var bt func(start, rest int)
	bt = func(start, rest int) {
		if rest == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := start; i < len(candidates) && candidates[i] <= rest; i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			bt(i+1, rest-candidates[i])
			path = path[:len(path)-1]
		}
	}
	bt(0, target)
	return res
}
// test
assertUnorderedDeep(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8), [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}})
assertUnorderedDeep(combinationSum2([]int{2, 5, 2, 1, 2}, 5), [][]int{{1, 2, 2}, {5}})
assertEq(combinationSum2([]int{1}, 2), [][]int{})
```

## 46. Permutations | Medium
related: 47, 78
tags: permutations

Each level picks any unused element (`used[]` array), depth n = one permutation.

**Complexity:** O(n · n!).

```python
class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        res, path, used = [], [], [False] * len(nums)

        def bt():
            if len(path) == len(nums):
                res.append(path[:])
                return
            for i, x in enumerate(nums):
                if not used[i]:
                    used[i] = True
                    path.append(x)
                    bt()
                    path.pop()
                    used[i] = False

        bt()
        return res
# test
assert sorted(Solution().permute([1, 2, 3])) == sorted([[1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]])
assert sorted(Solution().permute([0, 1])) == [[0, 1], [1, 0]]
assert Solution().permute([1]) == [[1]]
```

```java
class Solution {
    public List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        bt(nums, new boolean[nums.length], new ArrayList<>(), res);
        return res;
    }

    private void bt(int[] nums, boolean[] used, List<Integer> path, List<List<Integer>> res) {
        if (path.size() == nums.length) { res.add(new ArrayList<>(path)); return; }
        for (int i = 0; i < nums.length; i++) {
            if (used[i]) continue;
            used[i] = true; path.add(nums[i]);
            bt(nums, used, path, res);
            used[i] = false; path.remove(path.size() - 1);
        }
    }
}
// test
check(new Solution().permute(new int[]{1, 2, 3}).size(), 6);
check(new Solution().permute(new int[]{0, 1}), "[[0, 1], [1, 0]]");
```

```go
func permute(nums []int) [][]int {
	res, path, used := [][]int{}, []int{}, make([]bool, len(nums))
	var bt func()
	bt = func() {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i, x := range nums {
			if !used[i] {
				used[i] = true
				path = append(path, x)
				bt()
				path = path[:len(path)-1]
				used[i] = false
			}
		}
	}
	bt()
	return res
}
// test
assertUnordered(permute([]int{1, 2, 3}), [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}})
assertUnordered(permute([]int{0, 1}), [][]int{{0, 1}, {1, 0}})
assertEq(permute([]int{1}), [][]int{{1}})
```

## 47. Permutations II | Medium
related: 46, 90
tags: permutations, dedupe

Sort; skip `nums[i]` if it equals `nums[i-1]` **and `nums[i-1]` is not used** — i.e. only the first copy of a value may start a branch at each depth.

**Complexity:** O(n · n!).

```python
class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        res, path, used = [], [], [False] * len(nums)

        def bt():
            if len(path) == len(nums):
                res.append(path[:])
                return
            for i in range(len(nums)):
                if used[i] or (i and nums[i] == nums[i - 1] and not used[i - 1]):
                    continue
                used[i] = True
                path.append(nums[i])
                bt()
                path.pop()
                used[i] = False

        bt()
        return res
# test
assert sorted(Solution().permuteUnique([1, 1, 2])) == sorted([[1, 1, 2], [1, 2, 1], [2, 1, 1]])
assert sorted(Solution().permuteUnique([1, 2, 3])) == [[1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]]
assert Solution().permuteUnique([2, 2]) == [[2, 2]]
```

```java
class Solution {
    public List<List<Integer>> permuteUnique(int[] nums) {
        Arrays.sort(nums);
        List<List<Integer>> res = new ArrayList<>();
        bt(nums, new boolean[nums.length], new ArrayList<>(), res);
        return res;
    }

    private void bt(int[] nums, boolean[] used, List<Integer> path, List<List<Integer>> res) {
        if (path.size() == nums.length) { res.add(new ArrayList<>(path)); return; }
        for (int i = 0; i < nums.length; i++) {
            if (used[i] || (i > 0 && nums[i] == nums[i - 1] && !used[i - 1])) continue;
            used[i] = true; path.add(nums[i]);
            bt(nums, used, path, res);
            used[i] = false; path.remove(path.size() - 1);
        }
    }
}
// test
check(new Solution().permuteUnique(new int[]{1, 1, 2}), "[[1, 1, 2], [1, 2, 1], [2, 1, 1]]");
```

```go
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res, path, used := [][]int{}, []int{}, make([]bool, len(nums))
	var bt func()
	bt = func() {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := range nums {
			// equal values are used left to right only, so each arrangement appears once
			if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			bt()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	bt()
	return res
}
// test
assertUnordered(permuteUnique([]int{1, 1, 2}), [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}})
assertUnordered(permuteUnique([]int{1, 2, 3}), [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}})
assertEq(permuteUnique([]int{2, 2}), [][]int{{2, 2}})
```

## 77. Combinations | Medium
related: 78, 39
tags: combinations, pruning

Choose k of `1..n`. Prune: only start values that leave enough numbers to finish (`i ≤ n − (k − len(path)) + 1`).

**Complexity:** O(k · C(n, k)).

```python
class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        res, path = [], []

        def bt(start):
            if len(path) == k:
                res.append(path[:])
                return
            for i in range(start, n - (k - len(path)) + 2):
                path.append(i)
                bt(i + 1)
                path.pop()

        bt(1)
        return res
# test
assert sorted(map(sorted, Solution().combine(4, 2))) == sorted(map(sorted, [[1, 2], [1, 3], [1, 4], [2, 3], [2, 4], [3, 4]]))
assert sorted(map(sorted, Solution().combine(1, 1))) == sorted(map(sorted, [[1]]))
```

```java
class Solution {
    public List<List<Integer>> combine(int n, int k) {
        List<List<Integer>> res = new ArrayList<>();
        bt(n, k, 1, new ArrayList<>(), res);
        return res;
    }

    private void bt(int n, int k, int start, List<Integer> path, List<List<Integer>> res) {
        if (path.size() == k) { res.add(new ArrayList<>(path)); return; }
        for (int i = start; i <= n - (k - path.size()) + 1; i++) {
            path.add(i);
            bt(n, k, i + 1, path, res);
            path.remove(path.size() - 1);
        }
    }
}
// test
check(new Solution().combine(4, 2), "[[1, 2], [1, 3], [1, 4], [2, 3], [2, 4], [3, 4]]");
```

```go
func combine(n int, k int) [][]int {
	res, path := [][]int{}, []int{}
	var bt func(start int)
	bt = func(start int) {
		if len(path) == k {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := start; i <= n-(k-len(path))+1; i++ { // leave room for the remaining picks
			path = append(path, i)
			bt(i + 1)
			path = path[:len(path)-1]
		}
	}
	bt(1)
	return res
}
// test
assertUnorderedDeep(combine(4, 2), [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}})
assertUnorderedDeep(combine(1, 1), [][]int{{1}})
```

## 17. Letter Combinations of a Phone Number | Medium
related: 22, 78
tags: cartesian-product

Depth = digit index; branch over that digit's letters.

**Complexity:** O(4ⁿ · n).

```python
class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if not digits:
            return []
        pad = {"2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz"}
        res = []

        def bt(i, cur):
            if i == len(digits):
                res.append(cur)
                return
            for ch in pad[digits[i]]:
                bt(i + 1, cur + ch)

        bt(0, "")
        return res
# test
assert sorted(Solution().letterCombinations("23")) == sorted(["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"])
assert sorted(Solution().letterCombinations("")) == sorted([])
```

```java
class Solution {
    private static final String[] PAD = {"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"};

    public List<String> letterCombinations(String digits) {
        List<String> res = new ArrayList<>();
        if (!digits.isEmpty()) bt(digits, 0, new StringBuilder(), res);
        return res;
    }

    private void bt(String d, int i, StringBuilder cur, List<String> res) {
        if (i == d.length()) { res.add(cur.toString()); return; }
        for (char c : PAD[d.charAt(i) - '0'].toCharArray()) {
            cur.append(c);
            bt(d, i + 1, cur, res);
            cur.deleteCharAt(cur.length() - 1);
        }
    }
}
// test
check(new Solution().letterCombinations("23"), "[ad, ae, af, bd, be, bf, cd, ce, cf]");
```

```go
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	pad := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res := []string{}
	var bt func(i int, cur string)
	bt = func(i int, cur string) {
		if i == len(digits) {
			res = append(res, cur)
			return
		}
		for _, ch := range pad[digits[i]-'0'] {
			bt(i+1, cur+string(ch))
		}
	}
	bt(0, "")
	return res
}
// test
assertUnordered(letterCombinations("23"), []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"})
assertEq(letterCombinations(""), []string{})
```

## 22. Generate Parentheses | Medium
related: 20, 17
tags: constrained-generation

Add `(` while `open < n`; add `)` while `close < open`. These two rules generate exactly the valid strings.

```viz
              ""
          "("
       "(("      "()"
   "(()"  ...    "()("
  "(())"         "()()"
```

**Complexity:** O(Catalan(n) · n) ≈ O(4ⁿ/√n).

```python
class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        res = []

        def bt(cur, op, cl):
            if len(cur) == 2 * n:
                res.append(cur)
                return
            if op < n:
                bt(cur + "(", op + 1, cl)
            if cl < op:
                bt(cur + ")", op, cl + 1)

        bt("", 0, 0)
        return res
# test
assert sorted(Solution().generateParenthesis(3)) == sorted(["((()))", "(()())", "(())()", "()(())", "()()()"])
assert Solution().generateParenthesis(1) == ["()"]
assert sorted(Solution().generateParenthesis(4)) == ['(((())))', '((()()))', '((())())', '((()))()', '(()(()))', '(()()())', '(()())()', '(())(())', '(())()()', '()((()))', '()(()())', '()(())()', '()()(())', '()()()()']
```

```java
class Solution {
    public List<String> generateParenthesis(int n) {
        List<String> res = new ArrayList<>();
        bt(n, new StringBuilder(), 0, 0, res);
        return res;
    }

    private void bt(int n, StringBuilder cur, int op, int cl, List<String> res) {
        if (cur.length() == 2 * n) { res.add(cur.toString()); return; }
        if (op < n) { cur.append('('); bt(n, cur, op + 1, cl, res); cur.deleteCharAt(cur.length() - 1); }
        if (cl < op) { cur.append(')'); bt(n, cur, op, cl + 1, res); cur.deleteCharAt(cur.length() - 1); }
    }
}
// test
check(new Solution().generateParenthesis(3), "[((())), (()()), (())(), ()(()), ()()()]");
```

```go
func generateParenthesis(n int) []string {
	res := []string{}
	var bt func(cur string, open, closed int)
	bt = func(cur string, open, closed int) {
		if len(cur) == 2*n {
			res = append(res, cur)
			return
		}
		if open < n {
			bt(cur+"(", open+1, closed)
		}
		if closed < open { // only close what is open
			bt(cur+")", open, closed+1)
		}
	}
	bt("", 0, 0)
	return res
}
// test
assertUnordered(generateParenthesis(3), []string{"((()))", "(()())", "(())()", "()(())", "()()()"})
assertEq(generateParenthesis(1), []string{"()"})
assertUnordered(generateParenthesis(4), []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"})
```

## 79. Word Search | Medium
related: 212, 200
tags: blind75, grid-dfs

DFS from every cell matching `word[0]`; mark the cell visited (overwrite with `#`) while it's on the path, restore on return.

**Complexity:** O(R·C·3^L).

```python
class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        R, C = len(board), len(board[0])

        def dfs(r, c, i):
            if i == len(word):
                return True
            if not (0 <= r < R and 0 <= c < C) or board[r][c] != word[i]:
                return False
            board[r][c] = "#"
            found = dfs(r + 1, c, i + 1) or dfs(r - 1, c, i + 1) or dfs(r, c + 1, i + 1) or dfs(r, c - 1, i + 1)
            board[r][c] = word[i]
            return found

        return any(dfs(r, c, 0) for r in range(R) for c in range(C))
# test
B = [list("ABCE"), list("SFCS"), list("ADEE")]
assert Solution().exist(B, "ABCCED") is True
assert Solution().exist(B, "SEE") is True
assert Solution().exist(B, "ABCB") is False
```

```java
class Solution {
    public boolean exist(char[][] board, String word) {
        for (int r = 0; r < board.length; r++)
            for (int c = 0; c < board[0].length; c++)
                if (dfs(board, word, r, c, 0)) return true;
        return false;
    }

    private boolean dfs(char[][] b, String w, int r, int c, int i) {
        if (i == w.length()) return true;
        if (r < 0 || c < 0 || r >= b.length || c >= b[0].length || b[r][c] != w.charAt(i)) return false;
        b[r][c] = '#';
        boolean found = dfs(b, w, r + 1, c, i + 1) || dfs(b, w, r - 1, c, i + 1)
                || dfs(b, w, r, c + 1, i + 1) || dfs(b, w, r, c - 1, i + 1);
        b[r][c] = w.charAt(i);
        return found;
    }
}
// test
char[][] B = {"ABCE".toCharArray(), "SFCS".toCharArray(), "ADEE".toCharArray()};
check(new Solution().exist(B, "ABCCED"), true);
check(new Solution().exist(B, "ABCB"), false);
```

```go
func exist(board [][]byte, word string) bool {
	R, C := len(board), len(board[0])
	var dfs func(r, c, i int) bool
	dfs = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}
		if r < 0 || r >= R || c < 0 || c >= C || board[r][c] != word[i] {
			return false
		}
		board[r][c] = '#' // mark as used on this path
		found := dfs(r+1, c, i+1) || dfs(r-1, c, i+1) || dfs(r, c+1, i+1) || dfs(r, c-1, i+1)
		board[r][c] = word[i]
		return found
	}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}
// test
B := byteGrid("ABCE", "SFCS", "ADEE")
assertEq(exist(B, "ABCCED"), true)
assertEq(exist(B, "SEE"), true)
assertEq(exist(B, "ABCB"), false)
```

## 131. Palindrome Partitioning | Medium
related: 132, 5, 93
tags: partition

At index `i`, try every end `j` where `s[i..j]` is a palindrome, recurse from `j+1`. (Precompute a palindrome table to make checks O(1).)

**Complexity:** O(n · 2ⁿ).

```python
class Solution:
    def partition(self, s: str) -> List[List[str]]:
        res, path = [], []

        def bt(i):
            if i == len(s):
                res.append(path[:])
                return
            for j in range(i, len(s)):
                piece = s[i:j + 1]
                if piece == piece[::-1]:
                    path.append(piece)
                    bt(j + 1)
                    path.pop()

        bt(0)
        return res
# test
assert sorted(Solution().partition("aab")) == sorted([["a", "a", "b"], ["aa", "b"]])
assert Solution().partition("a") == [["a"]]
assert sorted(Solution().partition("aba")) == [["a", "b", "a"], ["aba"]]
```

```java
class Solution {
    public List<List<String>> partition(String s) {
        int n = s.length();
        boolean[][] pal = new boolean[n][n];
        for (int i = n - 1; i >= 0; i--)
            for (int j = i; j < n; j++)
                pal[i][j] = s.charAt(i) == s.charAt(j) && (j - i < 2 || pal[i + 1][j - 1]);
        List<List<String>> res = new ArrayList<>();
        bt(s, 0, pal, new ArrayList<>(), res);
        return res;
    }

    private void bt(String s, int i, boolean[][] pal, List<String> path, List<List<String>> res) {
        if (i == s.length()) { res.add(new ArrayList<>(path)); return; }
        for (int j = i; j < s.length(); j++) {
            if (!pal[i][j]) continue;
            path.add(s.substring(i, j + 1));
            bt(s, j + 1, pal, path, res);
            path.remove(path.size() - 1);
        }
    }
}
// test
check(new Solution().partition("aab"), "[[a, a, b], [aa, b]]");
```

```go
func partition(s string) [][]string {
	isPal := func(i, j int) bool {
		for ; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}
	res, path := [][]string{}, []string{}
	var bt func(i int)
	bt = func(i int) {
		if i == len(s) {
			res = append(res, append([]string{}, path...))
			return
		}
		for j := i; j < len(s); j++ {
			if isPal(i, j) {
				path = append(path, s[i:j+1])
				bt(j + 1)
				path = path[:len(path)-1]
			}
		}
	}
	bt(0)
	return res
}
// test
assertUnordered(partition("aab"), [][]string{{"a", "a", "b"}, {"aa", "b"}})
assertEq(partition("a"), [][]string{{"a"}})
assertUnordered(partition("aba"), [][]string{{"a", "b", "a"}, {"aba"}})
```

## 93. Restore IP Addresses | Medium
related: 131
tags: partition

Split into exactly 4 parts, each 1–3 digits, value ≤ 255, no leading zeros. Prune when the remaining length can't fit the remaining parts.

**Complexity:** O(3⁴) — constant.

```python
class Solution:
    def restoreIpAddresses(self, s: str) -> List[str]:
        res = []

        def bt(i, parts):
            if len(parts) == 4:
                if i == len(s):
                    res.append(".".join(parts))
                return
            rem = 4 - len(parts)
            if not rem <= len(s) - i <= 3 * rem:
                return
            for L in range(1, 4):
                p = s[i:i + L]
                if len(p) < L or (p[0] == "0" and L > 1) or int(p) > 255:
                    break
                bt(i + L, parts + [p])

        bt(0, [])
        return res
# test
assert sorted(Solution().restoreIpAddresses("25525511135")) == sorted(["255.255.11.135", "255.255.111.35"])
assert sorted(Solution().restoreIpAddresses("0000")) == sorted(["0.0.0.0"])
assert sorted(Solution().restoreIpAddresses("101023")) == sorted(["1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"])
```

```java
class Solution {
    public List<String> restoreIpAddresses(String s) {
        List<String> res = new ArrayList<>();
        bt(s, 0, new ArrayList<>(), res);
        return res;
    }

    private void bt(String s, int i, List<String> parts, List<String> res) {
        int rem = 4 - parts.size(), left = s.length() - i;
        if (rem == 0) { if (left == 0) res.add(String.join(".", parts)); return; }
        if (left < rem || left > 3 * rem) return;
        for (int L = 1; L <= 3 && i + L <= s.length(); L++) {
            String p = s.substring(i, i + L);
            if ((p.charAt(0) == '0' && L > 1) || Integer.parseInt(p) > 255) break;
            parts.add(p);
            bt(s, i + L, parts, res);
            parts.remove(parts.size() - 1);
        }
    }
}
// test
check(new Solution().restoreIpAddresses("25525511135"), "[255.255.11.135, 255.255.111.35]");
check(new Solution().restoreIpAddresses("0000"), "[0.0.0.0]");
```

```go
func restoreIpAddresses(s string) []string {
	res := []string{}
	var bt func(i int, parts []string)
	bt = func(i int, parts []string) {
		if len(parts) == 4 {
			if i == len(s) {
				res = append(res, strings.Join(parts, "."))
			}
			return
		}
		rem := 4 - len(parts)
		if left := len(s) - i; left < rem || left > 3*rem { // too few or too many digits left
			return
		}
		for L := 1; L <= 3 && i+L <= len(s); L++ {
			p := s[i : i+L]
			if v, _ := strconv.Atoi(p); (p[0] == '0' && L > 1) || v > 255 {
				break
			}
			bt(i+L, append(parts, p))
		}
	}
	bt(0, []string{})
	return res
}
// test
assertUnordered(restoreIpAddresses("25525511135"), []string{"255.255.11.135", "255.255.111.35"})
assertUnordered(restoreIpAddresses("0000"), []string{"0.0.0.0"})
assertUnordered(restoreIpAddresses("101023"), []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"})
```

## 51. N-Queens | Hard
related: 52, 37
tags: constraint-search

Place one queen per row. Track used columns and both diagonals (`r − c` and `r + c`) in sets for O(1) conflict checks.

```viz
n=4   . Q . .        cols {1,3,0,2}
      . . . Q        diag r-c, anti r+c all distinct
      Q . . .
      . . Q .
```

**Complexity:** O(n!).

```python
class Solution:
    def solveNQueens(self, n: int) -> List[List[str]]:
        res, cols, d1, d2 = [], set(), set(), set()
        board = [["."] * n for _ in range(n)]

        def bt(r):
            if r == n:
                res.append(["".join(row) for row in board])
                return
            for c in range(n):
                if c in cols or r - c in d1 or r + c in d2:
                    continue
                cols.add(c); d1.add(r - c); d2.add(r + c)
                board[r][c] = "Q"
                bt(r + 1)
                board[r][c] = "."
                cols.remove(c); d1.remove(r - c); d2.remove(r + c)

        bt(0)
        return res
# test
assert sorted(Solution().solveNQueens(4)) == sorted([[".Q..", "...Q", "Q...", "..Q."], ["..Q.", "Q...", "...Q", ".Q.."]])
assert len(Solution().solveNQueens(8)) == 92
```

```java
class Solution {
    public List<List<String>> solveNQueens(int n) {
        List<List<String>> res = new ArrayList<>();
        char[][] board = new char[n][n];
        for (char[] row : board) Arrays.fill(row, '.');
        bt(0, n, board, new boolean[n], new boolean[2 * n], new boolean[2 * n], res);
        return res;
    }

    private void bt(int r, int n, char[][] b, boolean[] cols, boolean[] d1, boolean[] d2, List<List<String>> res) {
        if (r == n) {
            List<String> sol = new ArrayList<>();
            for (char[] row : b) sol.add(new String(row));
            res.add(sol);
            return;
        }
        for (int c = 0; c < n; c++) {
            if (cols[c] || d1[r - c + n] || d2[r + c]) continue;
            cols[c] = d1[r - c + n] = d2[r + c] = true;
            b[r][c] = 'Q';
            bt(r + 1, n, b, cols, d1, d2, res);
            b[r][c] = '.';
            cols[c] = d1[r - c + n] = d2[r + c] = false;
        }
    }
}
// test
check(new Solution().solveNQueens(4), "[[.Q.., ...Q, Q..., ..Q.], [..Q., Q..., ...Q, .Q..]]");
check(new Solution().solveNQueens(8).size(), 92);
```

```go
func solveNQueens(n int) [][]string {
	res := [][]string{}
	cols, d1, d2 := map[int]bool{}, map[int]bool{}, map[int]bool{} // columns and both diagonals in use
	board := make([][]byte, n)
	for i := range board {
		board[i] = []byte(strings.Repeat(".", n))
	}
	var bt func(r int)
	bt = func(r int) {
		if r == n {
			rows := make([]string, n)
			for i, row := range board {
				rows[i] = string(row)
			}
			res = append(res, rows)
			return
		}
		for c := 0; c < n; c++ {
			if cols[c] || d1[r-c] || d2[r+c] {
				continue
			}
			cols[c], d1[r-c], d2[r+c] = true, true, true
			board[r][c] = 'Q'
			bt(r + 1)
			board[r][c] = '.'
			cols[c], d1[r-c], d2[r+c] = false, false, false
		}
	}
	bt(0)
	return res
}
// test
assertUnordered(solveNQueens(4), [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}})
assertEq(len(solveNQueens(8)), 92)
```

## 37. Sudoku Solver | Hard
related: 36, 51
tags: constraint-search

Track used digits per row/col/box (bitmasks or sets). Find an empty cell, try every allowed digit, recurse; undo on failure. (Heuristic: fill the cell with the fewest candidates first.)

**Complexity:** exponential worst case; fast in practice.

```python
class Solution:
    def solveSudoku(self, board: List[List[str]]) -> None:
        rows = [set() for _ in range(9)]
        cols = [set() for _ in range(9)]
        boxes = [set() for _ in range(9)]
        empty = []
        for r in range(9):
            for c in range(9):
                v = board[r][c]
                if v == ".":
                    empty.append((r, c))
                else:
                    rows[r].add(v); cols[c].add(v); boxes[r // 3 * 3 + c // 3].add(v)

        def bt(k):
            if k == len(empty):
                return True
            r, c = empty[k]
            b = r // 3 * 3 + c // 3
            for v in "123456789":
                if v in rows[r] or v in cols[c] or v in boxes[b]:
                    continue
                board[r][c] = v
                rows[r].add(v); cols[c].add(v); boxes[b].add(v)
                if bt(k + 1):
                    return True
                rows[r].remove(v); cols[c].remove(v); boxes[b].remove(v)
            board[r][c] = "."
            return False

        bt(0)
# test
B = ["53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1", "7...2...6", ".6....28.", "...419..5", "....8..79"]
board = [list(r) for r in B]
Solution().solveSudoku(board)
assert ["".join(r) for r in board][0] == "534678912" and ["".join(r) for r in board][8] == "345286179"
assert ["".join(r) for r in board] == ['534678912', '672195348', '198342567', '859761423', '426853791', '713924856', '961537284', '287419635', '345286179']
```

```java
class Solution {
    private final boolean[][] row = new boolean[9][10], col = new boolean[9][10], box = new boolean[9][10];

    public void solveSudoku(char[][] board) {
        for (int r = 0; r < 9; r++)
            for (int c = 0; c < 9; c++)
                if (board[r][c] != '.') mark(r, c, board[r][c] - '0', true);
        bt(board, 0);
    }

    private void mark(int r, int c, int d, boolean v) { row[r][d] = col[c][d] = box[r / 3 * 3 + c / 3][d] = v; }

    private boolean bt(char[][] b, int k) {
        if (k == 81) return true;
        int r = k / 9, c = k % 9;
        if (b[r][c] != '.') return bt(b, k + 1);
        for (int d = 1; d <= 9; d++) {
            if (row[r][d] || col[c][d] || box[r / 3 * 3 + c / 3][d]) continue;
            b[r][c] = (char) ('0' + d);
            mark(r, c, d, true);
            if (bt(b, k + 1)) return true;
            mark(r, c, d, false);
        }
        b[r][c] = '.';
        return false;
    }
}
// test
char[][] b = Stream.of("53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1",
        "7...2...6", ".6....28.", "...419..5", "....8..79").map(String::toCharArray).toArray(char[][]::new);
new Solution().solveSudoku(b);
check(new String(b[0]) + new String(b[8]), "534678912345286179");
```

```go
func solveSudoku(board [][]byte) {
	var rows, cols, boxes [9][10]bool
	empty := [][2]int{}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				empty = append(empty, [2]int{r, c})
			} else {
				v := board[r][c] - '0'
				rows[r][v], cols[c][v], boxes[r/3*3+c/3][v] = true, true, true
			}
		}
	}
	var bt func(k int) bool
	bt = func(k int) bool {
		if k == len(empty) {
			return true
		}
		r, c := empty[k][0], empty[k][1]
		b := r/3*3 + c/3
		for v := byte(1); v <= 9; v++ {
			if rows[r][v] || cols[c][v] || boxes[b][v] {
				continue
			}
			board[r][c] = '0' + v
			rows[r][v], cols[c][v], boxes[b][v] = true, true, true
			if bt(k + 1) {
				return true
			}
			rows[r][v], cols[c][v], boxes[b][v] = false, false, false
		}
		board[r][c] = '.'
		return false
	}
	bt(0)
}
// test
board := byteGrid("53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1", "7...2...6", ".6....28.", "...419..5", "....8..79")
solveSudoku(board)
assertEq(board, byteGrid("534678912", "672195348", "198342567", "859761423", "426853791", "713924856", "961537284", "287419635", "345286179"))
```


## 216. Combination Sum III | Medium
related: 39, 40, 77
tags: combinations

Choose k distinct digits from 1–9 summing to n: combinations with `i + 1` and pruning when the remainder goes negative.

**Complexity:** O(C(9, k)).

```python
class Solution:
    def combinationSum3(self, k: int, n: int) -> List[List[int]]:
        res, path = [], []

        def bt(start, rest):
            if len(path) == k:
                if rest == 0:
                    res.append(path[:])
                return
            for d in range(start, 10):
                if d > rest:
                    break
                path.append(d)
                bt(d + 1, rest - d)
                path.pop()

        bt(1, n)
        return res
# test
assert sorted(map(sorted, Solution().combinationSum3(3, 7))) == sorted(map(sorted, [[1, 2, 4]]))
assert sorted(map(sorted, Solution().combinationSum3(3, 9))) == sorted(map(sorted, [[1, 2, 6], [1, 3, 5], [2, 3, 4]]))
assert sorted(map(sorted, Solution().combinationSum3(4, 1))) == sorted(map(sorted, []))
```

```java
class Solution {
    public List<List<Integer>> combinationSum3(int k, int n) {
        List<List<Integer>> res = new ArrayList<>();
        bt(k, 1, n, new ArrayList<>(), res);
        return res;
    }

    private void bt(int k, int start, int rest, List<Integer> path, List<List<Integer>> res) {
        if (path.size() == k) { if (rest == 0) res.add(new ArrayList<>(path)); return; }
        for (int d = start; d <= 9 && d <= rest; d++) {
            path.add(d);
            bt(k, d + 1, rest - d, path, res);
            path.remove(path.size() - 1);
        }
    }
}
// test
check(new Solution().combinationSum3(3, 9), "[[1, 2, 6], [1, 3, 5], [2, 3, 4]]");
```

```go
func combinationSum3(k int, n int) [][]int {
	res, path := [][]int{}, []int{}
	var bt func(start, rest int)
	bt = func(start, rest int) {
		if len(path) == k {
			if rest == 0 {
				res = append(res, append([]int{}, path...))
			}
			return
		}
		for d := start; d <= 9 && d <= rest; d++ {
			path = append(path, d)
			bt(d+1, rest-d)
			path = path[:len(path)-1]
		}
	}
	bt(1, n)
	return res
}
// test
assertUnorderedDeep(combinationSum3(3, 7), [][]int{{1, 2, 4}})
assertUnorderedDeep(combinationSum3(3, 9), [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}})
assertEq(combinationSum3(4, 1), [][]int{})
```

## 784. Letter Case Permutation | Medium
related: 78, 17
tags: branching

Each letter branches into lower/upper; digits don't branch. Building iteratively: for each letter, double the result list.

**Complexity:** O(2^L · n).

```python
class Solution:
    def letterCasePermutation(self, s: str) -> List[str]:
        res = [""]
        for ch in s:
            if ch.isalpha():
                res = [p + c for p in res for c in (ch.lower(), ch.upper())]
            else:
                res = [p + ch for p in res]
        return res
# test
assert sorted(Solution().letterCasePermutation("a1b2")) == ["A1B2", "A1b2", "a1B2", "a1b2"]
assert Solution().letterCasePermutation("3z4") == ["3z4", "3Z4"]
```

```java
class Solution {
    public List<String> letterCasePermutation(String s) {
        List<String> res = new ArrayList<>();
        bt(s.toCharArray(), 0, res);
        return res;
    }

    private void bt(char[] c, int i, List<String> res) {
        if (i == c.length) { res.add(new String(c)); return; }
        if (Character.isLetter(c[i])) {
            c[i] = Character.toLowerCase(c[i]); bt(c, i + 1, res);
            c[i] = Character.toUpperCase(c[i]); bt(c, i + 1, res);
        } else bt(c, i + 1, res);
    }
}
// test
check(new Solution().letterCasePermutation("3z4"), "[3z4, 3Z4]");
```

```go
func letterCasePermutation(s string) []string {
	res := []string{""}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		next := []string{}
		for _, p := range res {
			if unicode.IsLetter(rune(ch)) { // branch: lower then upper
				next = append(next, p+strings.ToLower(string(ch)), p+strings.ToUpper(string(ch)))
			} else {
				next = append(next, p+string(ch))
			}
		}
		res = next
	}
	return res
}
// test
assertUnordered(letterCasePermutation("a1b2"), []string{"A1B2", "A1b2", "a1B2", "a1b2"})
assertEq(letterCasePermutation("3z4"), []string{"3z4", "3Z4"})
```

## 52. N-Queens II | Hard
related: 51
tags: bitmask, constraint-search

Count placements only. Bitmasks for columns and both diagonals: `free = ~(cols | d1 | d2) & full`; take the lowest bit each time; shift the diagonals by one when moving to the next row.

**Complexity:** O(n!), very fast in practice.

```python
class Solution:
    def totalNQueens(self, n: int) -> int:
        full = (1 << n) - 1

        def bt(cols, d1, d2):
            if cols == full:
                return 1
            count, free = 0, ~(cols | d1 | d2) & full
            while free:
                bit = free & -free
                free -= bit
                count += bt(cols | bit, ((d1 | bit) << 1) & full, (d2 | bit) >> 1)
            return count

        return bt(0, 0, 0)
# test
assert Solution().totalNQueens(4) == 2
assert Solution().totalNQueens(1) == 1
assert Solution().totalNQueens(8) == 92
```

```java
class Solution {
    private int full;

    public int totalNQueens(int n) {
        full = (1 << n) - 1;
        return bt(0, 0, 0);
    }

    private int bt(int cols, int d1, int d2) {
        if (cols == full) return 1;
        int count = 0;
        for (int free = ~(cols | d1 | d2) & full; free != 0; ) {
            int bit = free & -free;
            free -= bit;
            count += bt(cols | bit, ((d1 | bit) << 1) & full, (d2 | bit) >> 1);
        }
        return count;
    }
}
// test
check(new Solution().totalNQueens(8), 92);
```

```go
func totalNQueens(n int) int {
	full := 1<<n - 1
	var bt func(cols, d1, d2 int) int // bitmasks of attacked columns in the current row
	bt = func(cols, d1, d2 int) int {
		if cols == full {
			return 1
		}
		count := 0
		for free := ^(cols | d1 | d2) & full; free != 0; {
			bit := free & -free // lowest free column
			free -= bit
			count += bt(cols|bit, (d1|bit)<<1&full, (d2|bit)>>1)
		}
		return count
	}
	return bt(0, 0, 0)
}
// test
assertEq(totalNQueens(4), 2)
assertEq(totalNQueens(1), 1)
assertEq(totalNQueens(8), 92)
```

## 140. Word Break II | Hard
related: 139, 472
tags: memoization, backtracking

`sentences(i)` = all segmentations of `s[i:]`: for every dictionary word starting at `i`, prepend it to each sentence of the rest. Memoize by `i`.

**Complexity:** exponential output in the worst case; memoization avoids recomputation.

```python
class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> List[str]:
        words = set(wordDict)

        @cache
        def go(i):
            if i == len(s):
                return [""]
            out = []
            for j in range(i + 1, len(s) + 1):
                if s[i:j] in words:
                    out += [s[i:j] + (" " + rest if rest else "") for rest in go(j)]
            return out

        return go(0)
# test
assert sorted(Solution().wordBreak("catsanddog", ["cat", "cats", "and", "sand", "dog"])) == ["cat sand dog", "cats and dog"]
assert sorted(Solution().wordBreak("pineapplepenapple", ["apple", "pen", "applepen", "pine", "pineapple"])) == [
    "pine apple pen apple", "pine applepen apple", "pineapple pen apple"]
assert Solution().wordBreak("catsandog", ["cats", "dog", "sand", "and", "cat"]) == []
```

```java
class Solution {
    private final Map<Integer, List<String>> memo = new HashMap<>();

    public List<String> wordBreak(String s, List<String> wordDict) {
        return go(s, 0, new HashSet<>(wordDict));
    }

    private List<String> go(String s, int i, Set<String> words) {
        if (memo.containsKey(i)) return memo.get(i);
        List<String> out = new ArrayList<>();
        if (i == s.length()) out.add("");
        for (int j = i + 1; j <= s.length(); j++) {
            String w = s.substring(i, j);
            if (!words.contains(w)) continue;
            for (String rest : go(s, j, words)) out.add(rest.isEmpty() ? w : w + " " + rest);
        }
        memo.put(i, out);
        return out;
    }
}
// test
List<String> r = new Solution().wordBreak("catsanddog", List.of("cat", "cats", "and", "sand", "dog"));
Collections.sort(r);
check(r, "[cat sand dog, cats and dog]");
```

```go
func wordBreak(s string, wordDict []string) []string {
	words := map[string]bool{}
	for _, w := range wordDict {
		words[w] = true
	}
	memo := map[int][]string{}
	var gen func(i int) []string // all sentences for s[i:]
	gen = func(i int) []string {
		if i == len(s) {
			return []string{""}
		}
		if out, ok := memo[i]; ok {
			return out
		}
		out := []string{}
		for j := i + 1; j <= len(s); j++ {
			if w := s[i:j]; words[w] {
				for _, rest := range gen(j) {
					if rest == "" {
						out = append(out, w)
					} else {
						out = append(out, w+" "+rest)
					}
				}
			}
		}
		memo[i] = out
		return out
	}
	return gen(0)
}
// test
assertUnordered(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}), []string{"cat sand dog", "cats and dog"})
assertUnordered(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}), []string{
	"pine apple pen apple", "pine applepen apple", "pineapple pen apple"})
assertEq(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}), []string{})
```

## 282. Expression Add Operators | Hard
related: 224, 227
tags: backtracking, expression

Try every split of the digits and every operator. Track the running `value` and the last **operand** `prev`, so `*` can undo it: `value − prev + prev·cur`. Disallow numbers with leading zeros.

```viz
"123", target 6
1+2+3 = 6 ✓     1*2*3 = 6 ✓
1+2*3: value 3, prev 2 -> 3 - 2 + 2*3 = 7
```

**Complexity:** O(4ⁿ · n).

```python
class Solution:
    def addOperators(self, num: str, target: int) -> List[str]:
        res = []

        def bt(i, expr, value, prev):
            if i == len(num):
                if value == target:
                    res.append(expr)
                return
            for j in range(i + 1, len(num) + 1):
                s = num[i:j]
                if len(s) > 1 and s[0] == "0":
                    break
                cur = int(s)
                if i == 0:
                    bt(j, s, cur, cur)
                else:
                    bt(j, expr + "+" + s, value + cur, cur)
                    bt(j, expr + "-" + s, value - cur, -cur)
                    bt(j, expr + "*" + s, value - prev + prev * cur, prev * cur)

        bt(0, "", 0, 0)
        return res
# test
assert sorted(Solution().addOperators("123", 6)) == ["1*2*3", "1+2+3"]
assert sorted(Solution().addOperators("232", 8)) == ["2*3+2", "2+3*2"]
assert Solution().addOperators("3456237490", 9191) == []
assert sorted(Solution().addOperators("105", 5)) == ["1*0+5", "10-5"]
```

```java
class Solution {
    public List<String> addOperators(String num, int target) {
        List<String> res = new ArrayList<>();
        bt(num, target, 0, "", 0, 0, res);
        return res;
    }

    private void bt(String num, int target, int i, String expr, long value, long prev, List<String> res) {
        if (i == num.length()) { if (value == target) res.add(expr); return; }
        for (int j = i + 1; j <= num.length(); j++) {
            String s = num.substring(i, j);
            if (s.length() > 1 && s.charAt(0) == '0') break;
            long cur = Long.parseLong(s);
            if (i == 0) bt(num, target, j, s, cur, cur, res);
            else {
                bt(num, target, j, expr + "+" + s, value + cur, cur, res);
                bt(num, target, j, expr + "-" + s, value - cur, -cur, res);
                bt(num, target, j, expr + "*" + s, value - prev + prev * cur, prev * cur, res);
            }
        }
    }
}
// test
List<String> r = new Solution().addOperators("232", 8);
Collections.sort(r);
check(r, "[2*3+2, 2+3*2]");
```

```go
func addOperators(num string, target int) []string {
	res := []string{}
	// prev is the last multiplicative term, so "*" can undo it: value - prev + prev*cur
	var bt func(i int, expr string, value, prev int)
	bt = func(i int, expr string, value, prev int) {
		if i == len(num) {
			if value == target {
				res = append(res, expr)
			}
			return
		}
		for j := i + 1; j <= len(num); j++ {
			s := num[i:j]
			if len(s) > 1 && s[0] == '0' { // no leading zeros
				break
			}
			cur, _ := strconv.Atoi(s)
			if i == 0 {
				bt(j, s, cur, cur)
			} else {
				bt(j, expr+"+"+s, value+cur, cur)
				bt(j, expr+"-"+s, value-cur, -cur)
				bt(j, expr+"*"+s, value-prev+prev*cur, prev*cur)
			}
		}
	}
	bt(0, "", 0, 0)
	return res
}
// test
assertUnordered(addOperators("123", 6), []string{"1*2*3", "1+2+3"})
assertUnordered(addOperators("232", 8), []string{"2*3+2", "2+3*2"})
assertEq(addOperators("3456237490", 9191), []string{})
assertUnordered(addOperators("105", 5), []string{"1*0+5", "10-5"})
```

## 301. Remove Invalid Parentheses | Hard
related: 1249, 20
tags: bfs, backtracking

**BFS by number of removals**: level 0 = the string, level d = all strings with d characters removed. The first level that contains valid strings is the answer (minimum removals); collect all of them, deduped with a set.

**Complexity:** exponential worst case; fine for n ≤ 25.

```python
class Solution:
    def removeInvalidParentheses(self, s: str) -> List[str]:
        def valid(t):
            bal = 0
            for ch in t:
                if ch == "(":
                    bal += 1
                elif ch == ")":
                    bal -= 1
                    if bal < 0:
                        return False
            return bal == 0

        level = {s}
        while True:
            good = [t for t in level if valid(t)]
            if good:
                return good
            level = {t[:i] + t[i + 1:] for t in level for i in range(len(t)) if t[i] in "()"}
# test
assert sorted(Solution().removeInvalidParentheses("()())()")) == ["(())()", "()()()"]
assert sorted(Solution().removeInvalidParentheses("(a)())()")) == ["(a())()", "(a)()()"]
assert Solution().removeInvalidParentheses(")(") == [""]
```

```java
class Solution {
    public List<String> removeInvalidParentheses(String s) {
        Set<String> level = new HashSet<>(List.of(s));
        while (true) {
            List<String> good = new ArrayList<>();
            for (String t : level) if (valid(t)) good.add(t);
            if (!good.isEmpty()) return good;
            Set<String> next = new HashSet<>();
            for (String t : level)
                for (int i = 0; i < t.length(); i++)
                    if (t.charAt(i) == '(' || t.charAt(i) == ')') next.add(t.substring(0, i) + t.substring(i + 1));
            level = next;
        }
    }

    private boolean valid(String t) {
        int bal = 0;
        for (char c : t.toCharArray()) {
            if (c == '(') bal++;
            else if (c == ')' && --bal < 0) return false;
        }
        return bal == 0;
    }
}
// test
List<String> r = new Solution().removeInvalidParentheses("()())()");
Collections.sort(r);
check(r, "[(())(), ()()()]");
```

```go
func removeInvalidParentheses(s string) []string {
	valid := func(t string) bool {
		bal := 0
		for i := 0; i < len(t); i++ {
			if t[i] == '(' {
				bal++
			} else if t[i] == ')' {
				if bal--; bal < 0 {
					return false
				}
			}
		}
		return bal == 0
	}
	level := map[string]bool{s: true}
	for { // BFS by number of removals: the first level with a valid string wins
		good := []string{}
		for t := range level {
			if valid(t) {
				good = append(good, t)
			}
		}
		if len(good) > 0 {
			return good
		}
		next := map[string]bool{}
		for t := range level {
			for i := 0; i < len(t); i++ {
				if t[i] == '(' || t[i] == ')' {
					next[t[:i]+t[i+1:]] = true
				}
			}
		}
		level = next
	}
}
// test
assertUnordered(removeInvalidParentheses("()())()"), []string{"(())()", "()()()"})
assertUnordered(removeInvalidParentheses("(a)())()"), []string{"(a())()", "(a)()()"})
assertEq(removeInvalidParentheses(")("), []string{""})
```

## 698. Partition to K Equal Sum Subsets | Medium
related: 416, 473
tags: backtracking, pruning

Target = total / k. Place numbers (sorted **descending**, so failures happen early) into k buckets. Prune: never put a number into a bucket that would overflow, and if placing into an **empty** bucket failed, other empty buckets will fail too.

**Complexity:** O(kⁿ) worst case, heavily pruned.

```python
class Solution:
    def canPartitionKSubsets(self, nums: List[int], k: int) -> bool:
        total = sum(nums)
        if total % k:
            return False
        target = total // k
        nums.sort(reverse=True)
        if nums[0] > target:
            return False
        buckets = [0] * k

        def bt(i):
            if i == len(nums):
                return True
            for b in range(k):
                if buckets[b] + nums[i] <= target:
                    buckets[b] += nums[i]
                    if bt(i + 1):
                        return True
                    buckets[b] -= nums[i]
                if buckets[b] == 0:
                    break                     # symmetric empty buckets
            return False

        return bt(0)
# test
assert Solution().canPartitionKSubsets([4, 3, 2, 3, 5, 2, 1], 4) is True
assert Solution().canPartitionKSubsets([1, 2, 3, 4], 3) is False
```

```java
class Solution {
    public boolean canPartitionKSubsets(int[] nums, int k) {
        int total = 0;
        for (int x : nums) total += x;
        if (total % k != 0) return false;
        Arrays.sort(nums);
        int target = total / k;
        if (nums[nums.length - 1] > target) return false;
        return bt(nums, nums.length - 1, new int[k], target);   // largest first
    }

    private boolean bt(int[] nums, int i, int[] buckets, int target) {
        if (i < 0) return true;
        for (int b = 0; b < buckets.length; b++) {
            if (buckets[b] + nums[i] <= target) {
                buckets[b] += nums[i];
                if (bt(nums, i - 1, buckets, target)) return true;
                buckets[b] -= nums[i];
            }
            if (buckets[b] == 0) break;
        }
        return false;
    }
}
// test
check(new Solution().canPartitionKSubsets(new int[]{4, 3, 2, 3, 5, 2, 1}, 4), true);
check(new Solution().canPartitionKSubsets(new int[]{1, 2, 3, 4}, 3), false);
```

```go
func canPartitionKSubsets(nums []int, k int) bool {
	total := 0
	for _, x := range nums {
		total += x
	}
	if total%k != 0 {
		return false
	}
	target := total / k
	sort.Sort(sort.Reverse(sort.IntSlice(nums))) // big numbers first fail fast
	if nums[0] > target {
		return false
	}
	buckets := make([]int, k)
	var bt func(i int) bool
	bt = func(i int) bool {
		if i == len(nums) {
			return true
		}
		for b := 0; b < k; b++ {
			if buckets[b]+nums[i] <= target {
				buckets[b] += nums[i]
				if bt(i + 1) {
					return true
				}
				buckets[b] -= nums[i]
			}
			if buckets[b] == 0 { // empty buckets are interchangeable
				break
			}
		}
		return false
	}
	return bt(0)
}
// test
assertEq(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4), true)
assertEq(canPartitionKSubsets([]int{1, 2, 3, 4}, 3), false)
```

## 1863. Sum of All Subset XOR Totals | Easy
related: 78, 136
tags: subsets, bit-trick

Backtracking over all 2ⁿ subsets works (n ≤ 12). The elegant O(n) answer: a bit that is set in **any** element is set in exactly half of all subset XORs, so the sum is `(OR of all nums) · 2^(n−1)`.

**Complexity:** O(n).

```python
class Solution:
    def subsetXORSum(self, nums: List[int]) -> int:
        return functools.reduce(lambda x, y: x | y, nums) << (len(nums) - 1)
# test
assert Solution().subsetXORSum([1, 3]) == 6
assert Solution().subsetXORSum([5, 1, 6]) == 28
assert Solution().subsetXORSum([3, 4, 5, 6, 7, 8]) == 480
nums = [9, 2, 7, 12, 1]
brute = sum(functools.reduce(lambda x, y: x ^ y, c, 0) for r in range(len(nums) + 1) for c in itertools.combinations(nums, r))
assert Solution().subsetXORSum(nums) == brute
```

```java
class Solution {
    public int subsetXORSum(int[] nums) {
        int or = 0;
        for (int x : nums) or |= x;
        return or << (nums.length - 1);
    }
}
// test
check(new Solution().subsetXORSum(new int[]{5, 1, 6}), 28);
check(new Solution().subsetXORSum(new int[]{3, 4, 5, 6, 7, 8}), 480);
```

```go
func subsetXORSum(nums []int) int {
	or := 0
	for _, x := range nums {
		or |= x
	}
	// each bit set in some number is set in exactly half of the 2^n subset XORs
	return or << (len(nums) - 1)
}
// test
assertEq(subsetXORSum([]int{1, 3}), 6)
assertEq(subsetXORSum([]int{5, 1, 6}), 28)
assertEq(subsetXORSum([]int{3, 4, 5, 6, 7, 8}), 480)
nums := []int{9, 2, 7, 12, 1}
brute := 0
for mask := 0; mask < 1<<len(nums); mask++ {
	x := 0
	for i, v := range nums {
		if mask>>i&1 == 1 {
			x ^= v
		}
	}
	brute += x
}
assertEq(subsetXORSum(nums), brute)
```

## 473. Matchsticks to Square | Medium
related: 698, 416
tags: backtracking, pruning

Each side must equal `total / 4`. Place sticks **longest first** into one of four sides, backtracking on failure. Two prunings make it fast: skip a side if the stick overflows it, and if a stick doesn't work on an **empty** side, don't try the other empty sides (they're interchangeable).

**Complexity:** O(4ⁿ) worst case, tiny in practice with pruning (n ≤ 15).

```python
class Solution:
    def makesquare(self, matchsticks: List[int]) -> bool:
        total = sum(matchsticks)
        if total % 4:
            return False
        side = total // 4
        sticks = sorted(matchsticks, reverse=True)
        if sticks[0] > side:
            return False
        sides = [0] * 4

        def place(i):
            if i == len(sticks):
                return True                    # every side <= side and the total is 4·side -> all equal
            for k in range(4):
                if sides[k] + sticks[i] <= side:
                    sides[k] += sticks[i]
                    if place(i + 1):
                        return True
                    sides[k] -= sticks[i]
                if sides[k] == 0:
                    break
            return False

        return place(0)
# test
assert Solution().makesquare([1, 1, 2, 2, 2]) is True
assert Solution().makesquare([3, 3, 3, 3, 4]) is False
assert Solution().makesquare([5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3]) is True
assert Solution().makesquare([10, 6, 5, 5, 5, 3, 3, 3, 2, 2, 2, 2]) is True
assert Solution().makesquare([1, 1, 1]) is False
```

```java
class Solution {
    public boolean makesquare(int[] matchsticks) {
        int total = 0;
        for (int m : matchsticks) total += m;
        if (total % 4 != 0) return false;
        Integer[] sticks = Arrays.stream(matchsticks).boxed().sorted(Collections.reverseOrder()).toArray(Integer[]::new);
        if (sticks[0] > total / 4) return false;
        return place(sticks, 0, new int[4], total / 4);
    }

    private boolean place(Integer[] sticks, int i, int[] sides, int side) {
        if (i == sticks.length) return true;
        for (int k = 0; k < 4; k++) {
            if (sides[k] + sticks[i] <= side) {
                sides[k] += sticks[i];
                if (place(sticks, i + 1, sides, side)) return true;
                sides[k] -= sticks[i];
            }
            if (sides[k] == 0) break;
        }
        return false;
    }
}
// test
check(new Solution().makesquare(new int[]{1, 1, 2, 2, 2}), true);
check(new Solution().makesquare(new int[]{3, 3, 3, 3, 4}), false);
```

```go
func makesquare(matchsticks []int) bool {
	total := 0
	for _, x := range matchsticks {
		total += x
	}
	if total%4 != 0 {
		return false
	}
	side := total / 4
	sticks := append([]int{}, matchsticks...)
	sort.Sort(sort.Reverse(sort.IntSlice(sticks)))
	if sticks[0] > side {
		return false
	}
	var sides [4]int
	var place func(i int) bool
	place = func(i int) bool {
		if i == len(sticks) { // every side <= side and the total is 4*side, so all are equal
			return true
		}
		for k := 0; k < 4; k++ {
			if sides[k]+sticks[i] <= side {
				sides[k] += sticks[i]
				if place(i + 1) {
					return true
				}
				sides[k] -= sticks[i]
			}
			if sides[k] == 0 {
				break
			}
		}
		return false
	}
	return place(0)
}
// test
assertEq(makesquare([]int{1, 1, 2, 2, 2}), true)
assertEq(makesquare([]int{3, 3, 3, 3, 4}), false)
assertEq(makesquare([]int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}), true)
assertEq(makesquare([]int{10, 6, 5, 5, 5, 3, 3, 3, 2, 2, 2, 2}), true)
assertEq(makesquare([]int{1, 1, 1}), false)
```
