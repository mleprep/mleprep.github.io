# Math & Geometry

**Matrix tricks**
- Rotate 90° clockwise = **transpose** then **reverse each row**.
- Spiral = shrink four boundaries `top, bottom, left, right`.
- O(1)-space marking: reuse the first row/column as flags.

**Number tricks**
- Fast power: square-and-multiply on the bits of the exponent → O(log n).
- Cycle detection on sequences (happy number) = Floyd, like linked lists.
- Sieve of Eratosthenes for all primes < n in O(n log log n).
- Geometry: use integer arithmetic — represent slopes as reduced `(dy, dx)` pairs, never floats.

```viz
rotate:  1 2 3     transpose  1 4 7    reverse rows  7 4 1
         4 5 6    ---------> 2 5 8   ------------>  8 5 2
         7 8 9                3 6 9                  9 6 3
```

## 48. Rotate Image | Medium
related: 54, 189
tags: blind75, matrix

Transpose (swap `m[i][j]` with `m[j][i]`), then reverse every row. In place.

**Complexity:** O(n²) time, O(1) space.

```python
class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        n = len(matrix)
        for i in range(n):
            for j in range(i + 1, n):
                matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        for row in matrix:
            row.reverse()
# test
m = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]; Solution().rotate(m)
assert m == [[7, 4, 1], [8, 5, 2], [9, 6, 3]]
m = [[5, 1, 9, 11], [2, 4, 8, 10], [13, 3, 6, 7], [15, 14, 12, 16]]; Solution().rotate(m)
assert m == [[15, 13, 2, 5], [14, 3, 4, 1], [12, 6, 8, 9], [16, 7, 10, 11]]
m = [[1]]; Solution().rotate(m)
assert m == [[1]]
```

```java
class Solution {
    public void rotate(int[][] matrix) {
        int n = matrix.length;
        for (int i = 0; i < n; i++)
            for (int j = i + 1; j < n; j++) { int t = matrix[i][j]; matrix[i][j] = matrix[j][i]; matrix[j][i] = t; }
        for (int[] row : matrix)
            for (int l = 0, r = n - 1; l < r; l++, r--) { int t = row[l]; row[l] = row[r]; row[r] = t; }
    }
}
// test
int[][] m = {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}};
new Solution().rotate(m);
check(m, new int[][]{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}});
```

## 54. Spiral Matrix | Medium
related: 59, 48
tags: blind75, matrix, boundaries

Walk right along `top`, down along `right`, left along `bottom`, up along `left`, shrinking each boundary after use. Guard the last two walks for single rows/columns.

```viz
 1 →  2 →  3 →  4
                 ↓
 5 →  6 →  7     8        order: 1 2 3 4 8 12 11 10 9 5 6 7
 ↑               ↓
 9 ← 10 ← 11 ← 12
```

**Complexity:** O(m·n).

```python
class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        res = []
        top, bottom, left, right = 0, len(matrix) - 1, 0, len(matrix[0]) - 1
        while top <= bottom and left <= right:
            res += [matrix[top][c] for c in range(left, right + 1)]
            res += [matrix[r][right] for r in range(top + 1, bottom + 1)]
            if top < bottom and left < right:
                res += [matrix[bottom][c] for c in range(right - 1, left - 1, -1)]
                res += [matrix[r][left] for r in range(bottom - 1, top, -1)]
            top, bottom, left, right = top + 1, bottom - 1, left + 1, right - 1
        return res
# test
assert Solution().spiralOrder([[1, 2, 3], [4, 5, 6], [7, 8, 9]]) == [1, 2, 3, 6, 9, 8, 7, 4, 5]
assert Solution().spiralOrder([[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12]]) == [1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7]
```

```java
class Solution {
    public List<Integer> spiralOrder(int[][] matrix) {
        List<Integer> res = new ArrayList<>();
        int top = 0, bottom = matrix.length - 1, left = 0, right = matrix[0].length - 1;
        while (top <= bottom && left <= right) {
            for (int c = left; c <= right; c++) res.add(matrix[top][c]);
            for (int r = top + 1; r <= bottom; r++) res.add(matrix[r][right]);
            if (top < bottom && left < right) {
                for (int c = right - 1; c >= left; c--) res.add(matrix[bottom][c]);
                for (int r = bottom - 1; r > top; r--) res.add(matrix[r][left]);
            }
            top++; bottom--; left++; right--;
        }
        return res;
    }
}
// test
check(new Solution().spiralOrder(new int[][]{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}), "[1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7]");
```

## 73. Set Matrix Zeroes | Medium
related: 289
tags: blind75, matrix, in-place

Use row 0 and column 0 as zero-flags for their rows/columns; remember separately whether row 0 / col 0 themselves need zeroing.

**Complexity:** O(m·n) time, O(1) space.

```python
class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        R, C = len(matrix), len(matrix[0])
        row0 = any(matrix[0][c] == 0 for c in range(C))
        col0 = any(matrix[r][0] == 0 for r in range(R))
        for r in range(1, R):
            for c in range(1, C):
                if matrix[r][c] == 0:
                    matrix[r][0] = matrix[0][c] = 0
        for r in range(1, R):
            for c in range(1, C):
                if matrix[r][0] == 0 or matrix[0][c] == 0:
                    matrix[r][c] = 0
        if row0:
            for c in range(C):
                matrix[0][c] = 0
        if col0:
            for r in range(R):
                matrix[r][0] = 0
# test
m = [[0, 1, 2, 0], [3, 4, 5, 2], [1, 3, 1, 5]]; Solution().setZeroes(m)
assert m == [[0, 0, 0, 0], [0, 4, 5, 0], [0, 3, 1, 0]]
m = [[1, 1, 1], [1, 0, 1], [1, 1, 1]]; Solution().setZeroes(m)
assert m == [[1, 0, 1], [0, 0, 0], [1, 0, 1]]
```

```java
class Solution {
    public void setZeroes(int[][] matrix) {
        int R = matrix.length, C = matrix[0].length;
        boolean row0 = false, col0 = false;
        for (int c = 0; c < C; c++) if (matrix[0][c] == 0) row0 = true;
        for (int r = 0; r < R; r++) if (matrix[r][0] == 0) col0 = true;
        for (int r = 1; r < R; r++)
            for (int c = 1; c < C; c++) if (matrix[r][c] == 0) matrix[r][0] = matrix[0][c] = 0;
        for (int r = 1; r < R; r++)
            for (int c = 1; c < C; c++) if (matrix[r][0] == 0 || matrix[0][c] == 0) matrix[r][c] = 0;
        if (row0) Arrays.fill(matrix[0], 0);
        if (col0) for (int r = 0; r < R; r++) matrix[r][0] = 0;
    }
}
// test
int[][] m = {{1, 1, 1}, {1, 0, 1}, {1, 1, 1}};
new Solution().setZeroes(m);
check(m, new int[][]{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}});
```

## 202. Happy Number | Easy
related: 141, 258
tags: cycle-detection

Repeatedly summing squared digits either reaches 1 or loops. Detect the loop with Floyd (or a seen-set).

**Complexity:** O(log n) per step; the sequence quickly drops below 243.

```python
class Solution:
    def isHappy(self, n: int) -> bool:
        def nxt(x):
            return sum(int(d) ** 2 for d in str(x))

        slow, fast = n, nxt(n)
        while fast != 1 and slow != fast:
            slow, fast = nxt(slow), nxt(nxt(fast))
        return fast == 1
# test
assert Solution().isHappy(19) is True
assert Solution().isHappy(2) is False
```

```java
class Solution {
    public boolean isHappy(int n) {
        int slow = n, fast = next(n);
        while (fast != 1 && slow != fast) { slow = next(slow); fast = next(next(fast)); }
        return fast == 1;
    }

    private int next(int x) {
        int s = 0;
        for (; x > 0; x /= 10) s += (x % 10) * (x % 10);
        return s;
    }
}
// test
check(new Solution().isHappy(19), true);
check(new Solution().isHappy(2), false);
```

## 66. Plus One | Easy
related: 2, 67
tags: carry

Walk from the last digit: a 9 becomes 0 and carries; any other digit increments and we're done. All 9s → prepend 1.

**Complexity:** O(n).

```python
class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        for i in range(len(digits) - 1, -1, -1):
            if digits[i] < 9:
                digits[i] += 1
                return digits
            digits[i] = 0
        return [1] + digits
# test
assert Solution().plusOne([1, 2, 3]) == [1, 2, 4]
assert Solution().plusOne([9, 9]) == [1, 0, 0]
```

```java
class Solution {
    public int[] plusOne(int[] digits) {
        for (int i = digits.length - 1; i >= 0; i--) {
            if (digits[i] < 9) { digits[i]++; return digits; }
            digits[i] = 0;
        }
        int[] res = new int[digits.length + 1];
        res[0] = 1;
        return res;
    }
}
// test
check(new Solution().plusOne(new int[]{9, 9}), new int[]{1, 0, 0});
```

## 50. Pow(x, n) | Medium
related: 69, 372
tags: fast-power

Binary exponentiation: `x^n = (x²)^(n/2)` (· x if n is odd). Negative n → `1 / x^(−n)` (use a long for `−Integer.MIN_VALUE`).

```viz
x^13, 13 = 1101b
bit 1: res *= x       x -> x^2
bit 0:                x^2 -> x^4
bit 1: res *= x^4     x^4 -> x^8
bit 1: res *= x^8     => x^(1+4+8) = x^13
```

**Complexity:** O(log n).

```python
class Solution:
    def myPow(self, x: float, n: int) -> float:
        if n < 0:
            x, n = 1 / x, -n
        res = 1.0
        while n:
            if n & 1:
                res *= x
            x *= x
            n >>= 1
        return res
# test
assert abs(Solution().myPow(2.0, 10) - 1024.0) < 1e-9
assert abs(Solution().myPow(2.1, 3) - 9.261) < 1e-9
assert abs(Solution().myPow(2.0, -2) - 0.25) < 1e-9
```

```java
class Solution {
    public double myPow(double x, int n) {
        long e = n;
        if (e < 0) { x = 1 / x; e = -e; }
        double res = 1;
        for (; e > 0; e >>= 1) {
            if ((e & 1) == 1) res *= x;
            x *= x;
        }
        return res;
    }
}
// test
check(new Solution().myPow(2.0, 10), 1024.0);
check(new Solution().myPow(2.0, -2), 0.25);
check(new Solution().myPow(1.0, Integer.MIN_VALUE), 1.0);
```

## 43. Multiply Strings | Medium
related: 2, 415
tags: grade-school

Digit `i` of `a` times digit `j` of `b` lands at positions `i + j` (carry) and `i + j + 1`. Accumulate into an array of length `m + n`, then strip leading zeros.

```viz
      1 2 3
    ×   4 5
   ---------
 pos: 0 1 2 3 4      3*5 -> pos[3,4], 2*5 -> pos[2,3], ...
 result 5 5 3 5      = 5535
```

**Complexity:** O(m·n).

```python
class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        m, n = len(num1), len(num2)
        pos = [0] * (m + n)
        for i in range(m - 1, -1, -1):
            for j in range(n - 1, -1, -1):
                s = int(num1[i]) * int(num2[j]) + pos[i + j + 1]
                pos[i + j + 1] = s % 10
                pos[i + j] += s // 10
        res = "".join(map(str, pos)).lstrip("0")
        return res or "0"
# test
assert Solution().multiply("2", "3") == "6"
assert Solution().multiply("123", "456") == "56088"
assert Solution().multiply("0", "52") == "0"
```

```java
class Solution {
    public String multiply(String num1, String num2) {
        int m = num1.length(), n = num2.length();
        int[] pos = new int[m + n];
        for (int i = m - 1; i >= 0; i--)
            for (int j = n - 1; j >= 0; j--) {
                int s = (num1.charAt(i) - '0') * (num2.charAt(j) - '0') + pos[i + j + 1];
                pos[i + j + 1] = s % 10;
                pos[i + j] += s / 10;
            }
        StringBuilder sb = new StringBuilder();
        for (int d : pos) if (sb.length() > 0 || d != 0) sb.append(d);
        return sb.length() == 0 ? "0" : sb.toString();
    }
}
// test
check(new Solution().multiply("123", "456"), "56088");
check(new Solution().multiply("0", "52"), "0");
```

## 13. Roman to Integer | Easy
related: 12
tags: string

Add each symbol's value, but **subtract** it when a larger symbol follows (IV, IX, XL, …).

**Complexity:** O(n).

```python
class Solution:
    def romanToInt(self, s: str) -> int:
        val = {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
        total = 0
        for i, ch in enumerate(s):
            v = val[ch]
            total += -v if i + 1 < len(s) and val[s[i + 1]] > v else v
        return total
# test
assert Solution().romanToInt("III") == 3
assert Solution().romanToInt("LVIII") == 58
assert Solution().romanToInt("MCMXCIV") == 1994
```

```java
class Solution {
    public int romanToInt(String s) {
        int total = 0;
        for (int i = 0; i < s.length(); i++) {
            int v = val(s.charAt(i));
            total += (i + 1 < s.length() && val(s.charAt(i + 1)) > v) ? -v : v;
        }
        return total;
    }

    private int val(char c) {
        switch (c) {
            case 'I': return 1; case 'V': return 5; case 'X': return 10; case 'L': return 50;
            case 'C': return 100; case 'D': return 500; default: return 1000;
        }
    }
}
// test
check(new Solution().romanToInt("MCMXCIV"), 1994);
```

## 12. Integer to Roman | Medium
related: 13
tags: greedy

Greedy over the 13 symbol values in descending order, **including** the subtractive pairs (900 CM, 400 CD, 90 XC, …).

**Complexity:** O(1) (bounded output).

```python
class Solution:
    def intToRoman(self, num: int) -> str:
        table = [(1000, "M"), (900, "CM"), (500, "D"), (400, "CD"), (100, "C"), (90, "XC"),
                 (50, "L"), (40, "XL"), (10, "X"), (9, "IX"), (5, "V"), (4, "IV"), (1, "I")]
        res = []
        for v, sym in table:
            k, num = divmod(num, v)
            res.append(sym * k)
        return "".join(res)
# test
assert Solution().intToRoman(3749) == "MMMDCCXLIX"
assert Solution().intToRoman(58) == "LVIII"
assert Solution().intToRoman(1994) == "MCMXCIV"
```

```java
class Solution {
    public String intToRoman(int num) {
        int[] vals = {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1};
        String[] syms = {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"};
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < vals.length; i++)
            while (num >= vals[i]) { num -= vals[i]; sb.append(syms[i]); }
        return sb.toString();
    }
}
// test
check(new Solution().intToRoman(1994), "MCMXCIV");
```

## 204. Count Primes | Medium
related: 263
tags: sieve

Sieve of Eratosthenes: for each prime `p` with `p² < n`, cross out `p², p²+p, …`.

**Complexity:** O(n log log n).

```python
class Solution:
    def countPrimes(self, n: int) -> int:
        if n < 3:
            return 0
        is_p = bytearray([1]) * n
        is_p[0] = is_p[1] = 0
        for p in range(2, int(n ** 0.5) + 1):
            if is_p[p]:
                is_p[p * p::p] = bytearray(len(range(p * p, n, p)))
        return sum(is_p)
# test
assert Solution().countPrimes(10) == 4
assert Solution().countPrimes(0) == 0
assert Solution().countPrimes(100) == 25
```

```java
class Solution {
    public int countPrimes(int n) {
        if (n < 3) return 0;
        boolean[] composite = new boolean[n];
        int count = 0;
        for (int p = 2; p < n; p++) {
            if (composite[p]) continue;
            count++;
            for (long q = (long) p * p; q < n; q += p) composite[(int) q] = true;
        }
        return count;
    }
}
// test
check(new Solution().countPrimes(10), 4);
check(new Solution().countPrimes(100), 25);
```

## 149. Max Points on a Line | Hard
related: 2013
tags: geometry, hash-map, gcd

For each anchor point, group all other points by **reduced slope** `(dy/g, dx/g)` with a normalized sign. The largest group + 1 (the anchor) is the best line through that anchor.

**Complexity:** O(n²).

```python
class Solution:
    def maxPoints(self, points: List[List[int]]) -> int:
        best = 1
        for i, (x1, y1) in enumerate(points):
            slopes = defaultdict(int)
            for x2, y2 in points[i + 1:]:
                dx, dy = x2 - x1, y2 - y1
                g = math.gcd(dx, dy)
                dx, dy = dx // g, dy // g
                if dx < 0 or (dx == 0 and dy < 0):
                    dx, dy = -dx, -dy
                slopes[(dx, dy)] += 1
                best = max(best, slopes[(dx, dy)] + 1)
        return best
# test
assert Solution().maxPoints([[1, 1], [2, 2], [3, 3]]) == 3
assert Solution().maxPoints([[1, 1], [3, 2], [5, 3], [4, 1], [2, 3], [1, 4]]) == 4
```

```java
class Solution {
    public int maxPoints(int[][] points) {
        int best = 1;
        for (int i = 0; i < points.length; i++) {
            Map<String, Integer> slopes = new HashMap<>();
            for (int j = i + 1; j < points.length; j++) {
                int dx = points[j][0] - points[i][0], dy = points[j][1] - points[i][1];
                int g = gcd(Math.abs(dx), Math.abs(dy));
                dx /= g; dy /= g;
                if (dx < 0 || (dx == 0 && dy < 0)) { dx = -dx; dy = -dy; }
                best = Math.max(best, slopes.merge(dx + "/" + dy, 1, Integer::sum) + 1);
            }
        }
        return best;
    }

    private int gcd(int a, int b) { return b == 0 ? a : gcd(b, a % b); }
}
// test
check(new Solution().maxPoints(new int[][]{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}), 4);
```

## 2013. Detect Squares | Medium
related: 149
tags: design, hash-map

Count points. For a query `(x, y)`, each stored point `(px, py)` with `|px − x| = |py − y| ≠ 0` is a diagonal corner; the square count is `cnt[diag] · cnt[(x, py)] · cnt[(px, y)]`.

```viz
(x,py) ---- (px,py)      query (x,y), diagonal (px,py)
  |            |         need the other two corners
(x,y)  ---- (px,y)
```

**Complexity:** add O(1), count O(#distinct points).

```python
class DetectSquares:
    def __init__(self):
        self.cnt = Counter()

    def add(self, point: List[int]) -> None:
        self.cnt[tuple(point)] += 1

    def count(self, point: List[int]) -> int:
        x, y = point
        total = 0
        for (px, py), c in list(self.cnt.items()):
            if abs(px - x) != abs(py - y) or px == x:
                continue
            total += c * self.cnt[(x, py)] * self.cnt[(px, y)]
        return total
# test
ds = DetectSquares()
for p in ([3, 10], [11, 2], [3, 2]):
    ds.add(p)
assert ds.count([11, 10]) == 1
assert ds.count([14, 8]) == 0
ds.add([11, 2])
assert ds.count([11, 10]) == 2
```

```java
class DetectSquares {
    private final Map<Long, Integer> cnt = new HashMap<>();
    private final List<int[]> pts = new ArrayList<>();

    private long key(int x, int y) { return (long) x * 2001 + y; }

    public void add(int[] point) {
        if (cnt.merge(key(point[0], point[1]), 1, Integer::sum) == 1) pts.add(point);
    }

    public int count(int[] point) {
        int x = point[0], y = point[1], total = 0;
        for (int[] p : pts) {
            if (Math.abs(p[0] - x) != Math.abs(p[1] - y) || p[0] == x) continue;
            total += cnt.get(key(p[0], p[1])) * cnt.getOrDefault(key(x, p[1]), 0) * cnt.getOrDefault(key(p[0], y), 0);
        }
        return total;
    }
}
// test
DetectSquares ds = new DetectSquares();
ds.add(new int[]{3, 10}); ds.add(new int[]{11, 2}); ds.add(new int[]{3, 2});
check(ds.count(new int[]{11, 10}), 1);
ds.add(new int[]{11, 2});
check(ds.count(new int[]{11, 10}), 2);
```


## 9. Palindrome Number | Easy
related: 125, 7
tags: math

Negative numbers and numbers ending in 0 (except 0) aren't palindromes. Reverse only the **second half** of the digits and compare with the first half — no overflow risk.

**Complexity:** O(log x).

```python
class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0 or (x % 10 == 0 and x != 0):
            return False
        rev = 0
        while x > rev:
            x, d = divmod(x, 10)
            rev = rev * 10 + d
        return x == rev or x == rev // 10
# test
assert Solution().isPalindrome(121) is True
assert Solution().isPalindrome(-121) is False
assert Solution().isPalindrome(10) is False
assert Solution().isPalindrome(1221) is True
```

```java
class Solution {
    public boolean isPalindrome(int x) {
        if (x < 0 || (x % 10 == 0 && x != 0)) return false;
        int rev = 0;
        while (x > rev) { rev = rev * 10 + x % 10; x /= 10; }
        return x == rev || x == rev / 10;
    }
}
// test
check(new Solution().isPalindrome(121), true);
check(new Solution().isPalindrome(10), false);
```

## 67. Add Binary | Easy
related: 2, 415
tags: carry

Add from the right with a carry, like decimal addition in base 2.

**Complexity:** O(max(m, n)).

```python
class Solution:
    def addBinary(self, a: str, b: str) -> str:
        i, j, carry, out = len(a) - 1, len(b) - 1, 0, []
        while i >= 0 or j >= 0 or carry:
            s = carry + (int(a[i]) if i >= 0 else 0) + (int(b[j]) if j >= 0 else 0)
            out.append(str(s % 2))
            carry = s // 2
            i, j = i - 1, j - 1
        return "".join(reversed(out))
# test
assert Solution().addBinary("11", "1") == "100"
assert Solution().addBinary("1010", "1011") == "10101"
```

```java
class Solution {
    public String addBinary(String a, String b) {
        StringBuilder sb = new StringBuilder();
        for (int i = a.length() - 1, j = b.length() - 1, carry = 0; i >= 0 || j >= 0 || carry > 0; i--, j--) {
            int s = carry + (i >= 0 ? a.charAt(i) - '0' : 0) + (j >= 0 ? b.charAt(j) - '0' : 0);
            sb.append(s % 2);
            carry = s / 2;
        }
        return sb.reverse().toString();
    }
}
// test
check(new Solution().addBinary("1010", "1011"), "10101");
```

## 172. Factorial Trailing Zeroes | Medium
related: 204
tags: math

Each trailing zero needs a factor 10 = 2 · 5, and 2s are plentiful, so count factors of 5: `n/5 + n/25 + n/125 + …`.

**Complexity:** O(log n).

```python
class Solution:
    def trailingZeroes(self, n: int) -> int:
        count = 0
        while n:
            n //= 5
            count += n
        return count
# test
assert Solution().trailingZeroes(3) == 0
assert Solution().trailingZeroes(5) == 1
assert Solution().trailingZeroes(100) == 24
```

```java
class Solution {
    public int trailingZeroes(int n) {
        int count = 0;
        while (n > 0) { n /= 5; count += n; }
        return count;
    }
}
// test
check(new Solution().trailingZeroes(100), 24);
```

## 59. Spiral Matrix II | Medium
related: 54
tags: matrix, boundaries

Same four boundaries as Spiral Matrix, but write `1..n²` instead of reading.

**Complexity:** O(n²).

```python
class Solution:
    def generateMatrix(self, n: int) -> List[List[int]]:
        m = [[0] * n for _ in range(n)]
        top, bottom, left, right, k = 0, n - 1, 0, n - 1, 1
        while top <= bottom:
            for c in range(left, right + 1):
                m[top][c] = k; k += 1
            for r in range(top + 1, bottom + 1):
                m[r][right] = k; k += 1
            for c in range(right - 1, left - 1, -1):
                if top < bottom:
                    m[bottom][c] = k; k += 1
            for r in range(bottom - 1, top, -1):
                if left < right:
                    m[r][left] = k; k += 1
            top, bottom, left, right = top + 1, bottom - 1, left + 1, right - 1
        return m
# test
assert Solution().generateMatrix(3) == [[1, 2, 3], [8, 9, 4], [7, 6, 5]]
assert Solution().generateMatrix(1) == [[1]]
```

```java
class Solution {
    public int[][] generateMatrix(int n) {
        int[][] m = new int[n][n];
        int top = 0, bottom = n - 1, left = 0, right = n - 1, k = 1;
        while (top <= bottom) {
            for (int c = left; c <= right; c++) m[top][c] = k++;
            for (int r = top + 1; r <= bottom; r++) m[r][right] = k++;
            if (top < bottom) for (int c = right - 1; c >= left; c--) m[bottom][c] = k++;
            if (left < right) for (int r = bottom - 1; r > top; r--) m[r][left] = k++;
            top++; bottom--; left++; right--;
        }
        return m;
    }
}
// test
check(new Solution().generateMatrix(3), new int[][]{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}});
```

## 8. String to Integer (atoi) | Medium
related: 7, 65
tags: parsing, overflow

Skip spaces, read an optional sign, read digits until a non-digit; clamp to `[−2³¹, 2³¹−1]` (check before multiplying in Java).

**Complexity:** O(n).

```python
class Solution:
    def myAtoi(self, s: str) -> int:
        i, n = 0, len(s)
        while i < n and s[i] == " ":
            i += 1
        sign = 1
        if i < n and s[i] in "+-":
            sign = -1 if s[i] == "-" else 1
            i += 1
        num = 0
        while i < n and s[i].isdigit():
            num = num * 10 + int(s[i])
            i += 1
        return max(-2 ** 31, min(2 ** 31 - 1, sign * num))
# test
assert Solution().myAtoi("42") == 42
assert Solution().myAtoi("   -042") == -42
assert Solution().myAtoi("1337c0d3") == 1337
assert Solution().myAtoi("0-1") == 0
assert Solution().myAtoi("-91283472332") == -2147483648
```

```java
class Solution {
    public int myAtoi(String s) {
        int i = 0, n = s.length(), sign = 1, num = 0;
        while (i < n && s.charAt(i) == ' ') i++;
        if (i < n && (s.charAt(i) == '+' || s.charAt(i) == '-')) sign = s.charAt(i++) == '-' ? -1 : 1;
        while (i < n && Character.isDigit(s.charAt(i))) {
            int d = s.charAt(i++) - '0';
            if (num > (Integer.MAX_VALUE - d) / 10) return sign == 1 ? Integer.MAX_VALUE : Integer.MIN_VALUE;
            num = num * 10 + d;
        }
        return sign * num;
    }
}
// test
check(new Solution().myAtoi("   -042"), -42);
check(new Solution().myAtoi("-91283472332"), -2147483648);
check(new Solution().myAtoi("2147483648"), 2147483647);
```

## 29. Divide Two Integers | Medium
related: 50, 371
tags: bit-shifting

No `*`, `/`, `%`: subtract the **largest doubled divisor** (`divisor << k`) that fits, adding `1 << k` to the quotient. Work with negatives (or longs) to dodge the `−2³¹` overflow; the only overflow case is `−2³¹ / −1`.

**Complexity:** O(log² n) (O(log n) with a precomputed shift list).

```python
class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        if dividend == -2 ** 31 and divisor == -1:
            return 2 ** 31 - 1
        neg = (dividend < 0) != (divisor < 0)
        a, b, q = abs(dividend), abs(divisor), 0
        while a >= b:
            k = 0
            while a >= b << (k + 1):
                k += 1
            a -= b << k
            q += 1 << k
        return -q if neg else q
# test
assert Solution().divide(10, 3) == 3
assert Solution().divide(7, -3) == -2
assert Solution().divide(-2147483648, -1) == 2147483647
assert Solution().divide(-2147483648, 2) == -1073741824
```

```java
class Solution {
    public int divide(int dividend, int divisor) {
        if (dividend == Integer.MIN_VALUE && divisor == -1) return Integer.MAX_VALUE;
        boolean neg = (dividend < 0) != (divisor < 0);
        long a = Math.abs((long) dividend), b = Math.abs((long) divisor), q = 0;
        while (a >= b) {
            int k = 0;
            while (a >= (b << (k + 1))) k++;
            a -= b << k;
            q += 1L << k;
        }
        return (int) (neg ? -q : q);
    }
}
// test
check(new Solution().divide(7, -3), -2);
check(new Solution().divide(-2147483648, -1), 2147483647);
```

## 470. Implement Rand10() Using Rand7() | Medium
related: 528
tags: rejection-sampling

`(rand7() − 1) · 7 + rand7()` is uniform on 1..49. Accept values ≤ 40 (a multiple of 10) and return `(v − 1) % 10 + 1`; otherwise retry. Expected ~2.45 calls to rand7 per rand10.

**Complexity:** O(1) expected.

```python
class Solution:
    def rand10(self) -> int:
        while True:
            v = (rand7() - 1) * 7 + rand7()     # rand7 provided by LeetCode
            if v <= 40:
                return (v - 1) % 10 + 1
# test
random.seed(7)
def rand7(): return random.randint(1, 7)
cnt = Counter(Solution().rand10() for _ in range(20000))
assert set(cnt) == set(range(1, 11)) and all(1800 < c < 2200 for c in cnt.values())
```

```java
class SolBase { // provided by LeetCode
    private final Random rnd = new Random(7);
    int rand7() { return rnd.nextInt(7) + 1; }
}

class Solution extends SolBase {
    public int rand10() {
        while (true) {
            int v = (rand7() - 1) * 7 + rand7();
            if (v <= 40) return (v - 1) % 10 + 1;
        }
    }
}
// test
Solution s = new Solution();
int[] cnt = new int[11];
for (int i = 0; i < 20000; i++) cnt[s.rand10()]++;
boolean ok = true;
for (int d = 1; d <= 10; d++) ok &= cnt[d] > 1800 && cnt[d] < 2200;
check(ok, true);
```
