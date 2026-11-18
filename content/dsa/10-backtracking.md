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
assert Solution().subsets([1, 2, 3]) == [[], [1], [1, 2], [1, 2, 3], [1, 3], [2], [2, 3], [3]]
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
assert Solution().subsetsWithDup([1, 2, 2]) == [[], [1], [1, 2], [1, 2, 2], [2], [2, 2]]
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
assert Solution().combinationSum([2, 3, 6, 7], 7) == [[2, 2, 3], [7]]
assert Solution().combinationSum([2, 3, 5], 8) == [[2, 2, 2, 2], [2, 3, 3], [3, 5]]
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
assert Solution().combinationSum2([10, 1, 2, 7, 6, 1, 5], 8) == [[1, 1, 6], [1, 2, 5], [1, 7], [2, 6]]
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
assert Solution().permute([1, 2, 3]) == [[1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]]
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
assert Solution().permuteUnique([1, 1, 2]) == [[1, 1, 2], [1, 2, 1], [2, 1, 1]]
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
assert Solution().combine(4, 2) == [[1, 2], [1, 3], [1, 4], [2, 3], [2, 4], [3, 4]]
assert Solution().combine(1, 1) == [[1]]
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
assert Solution().letterCombinations("23") == ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]
assert Solution().letterCombinations("") == []
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
assert Solution().generateParenthesis(3) == ["((()))", "(()())", "(())()", "()(())", "()()()"]
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
assert Solution().partition("aab") == [["a", "a", "b"], ["aa", "b"]]
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
assert Solution().restoreIpAddresses("25525511135") == ["255.255.11.135", "255.255.111.35"]
assert Solution().restoreIpAddresses("0000") == ["0.0.0.0"]
assert Solution().restoreIpAddresses("101023") == ["1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"]
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
assert Solution().solveNQueens(4) == [[".Q..", "...Q", "Q...", "..Q."], ["..Q.", "Q...", "...Q", ".Q.."]]
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
