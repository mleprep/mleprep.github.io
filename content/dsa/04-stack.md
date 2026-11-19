# Stack & Monotonic Stack

**When to use:** matching/nesting (parentheses, decode), expression evaluation, "undo" behaviour, and **next greater / smaller element** questions.

**Monotonic stack**: keep indices whose values are increasing (or decreasing). When a new value breaks the order, pop — the popped index has just found its *next greater/smaller* element. Each index is pushed and popped once → O(n).

```python
# next greater element to the right, for every index
res, st = [-1] * n, []          # st holds indices with decreasing values
for i, x in enumerate(a):
    while st and a[st[-1]] < x:
        res[st.pop()] = x
    st.append(i)
```

## 20. Valid Parentheses | Easy
related: 22, 32, 394
tags: blind75

Push expected closers; on a closer, it must match the top.

**Complexity:** O(n).

```python
class Solution:
    def isValid(self, s: str) -> bool:
        pair = {")": "(", "]": "[", "}": "{"}
        st = []
        for ch in s:
            if ch in pair:
                if not st or st.pop() != pair[ch]:
                    return False
            else:
                st.append(ch)
        return not st
# test
assert Solution().isValid("()[]{}") is True
assert Solution().isValid("(]") is False
assert Solution().isValid("([)]") is False
assert Solution().isValid("{[]}") is True
```

```java
class Solution {
    public boolean isValid(String s) {
        Deque<Character> st = new ArrayDeque<>();
        for (char c : s.toCharArray()) {
            if (c == '(') st.push(')');
            else if (c == '[') st.push(']');
            else if (c == '{') st.push('}');
            else if (st.isEmpty() || st.pop() != c) return false;
        }
        return st.isEmpty();
    }
}
// test
check(new Solution().isValid("()[]{}"), true);
check(new Solution().isValid("([)]"), false);
```

## 155. Min Stack | Medium
related: 239, 716
tags: design

Store `(value, min so far)` pairs; the minimum is always at the top in O(1).

**Complexity:** O(1) per operation.

```python
class MinStack:
    def __init__(self):
        self.st = []  # (val, current min)

    def push(self, val: int) -> None:
        self.st.append((val, min(val, self.st[-1][1]) if self.st else val))

    def pop(self) -> None:
        self.st.pop()

    def top(self) -> int:
        return self.st[-1][0]

    def getMin(self) -> int:
        return self.st[-1][1]
# test
m = MinStack()
m.push(-2); m.push(0); m.push(-3)
assert m.getMin() == -3
m.pop()
assert m.top() == 0 and m.getMin() == -2
```

```java
class MinStack {
    private final Deque<int[]> st = new ArrayDeque<>(); // {val, min}

    public void push(int val) {
        st.push(new int[]{val, st.isEmpty() ? val : Math.min(val, st.peek()[1])});
    }
    public void pop() { st.pop(); }
    public int top() { return st.peek()[0]; }
    public int getMin() { return st.peek()[1]; }
}
// test
MinStack m = new MinStack();
m.push(-2); m.push(0); m.push(-3);
check(m.getMin(), -3);
m.pop();
check(m.top(), 0);
check(m.getMin(), -2);
```

## 150. Evaluate Reverse Polish Notation | Medium
related: 224, 227
tags: expression

Push numbers; an operator pops `b` then `a` and pushes `a op b`. Division truncates toward zero.

**Complexity:** O(n).

```python
class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        st = []
        for t in tokens:
            if t in "+-*/":
                b, a = st.pop(), st.pop()
                if t == "+": st.append(a + b)
                elif t == "-": st.append(a - b)
                elif t == "*": st.append(a * b)
                else: st.append(int(a / b))  # truncate toward 0
            else:
                st.append(int(t))
        return st[0]
# test
assert Solution().evalRPN(["2", "1", "+", "3", "*"]) == 9
assert Solution().evalRPN(["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]) == 22
```

```java
class Solution {
    public int evalRPN(String[] tokens) {
        Deque<Integer> st = new ArrayDeque<>();
        for (String t : tokens) {
            switch (t) {
                case "+": st.push(st.pop() + st.pop()); break;
                case "-": { int b = st.pop(), a = st.pop(); st.push(a - b); break; }
                case "*": st.push(st.pop() * st.pop()); break;
                case "/": { int b = st.pop(), a = st.pop(); st.push(a / b); break; }
                default: st.push(Integer.parseInt(t));
            }
        }
        return st.pop();
    }
}
// test
check(new Solution().evalRPN(new String[]{"2", "1", "+", "3", "*"}), 9);
check(new Solution().evalRPN(new String[]{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}), 22);
```

## 739. Daily Temperatures | Medium
related: 496, 503, 901
tags: monotonic-stack

Decreasing stack of indices. A warmer day pops every colder index and fills its wait = `i − j`.

```viz
temps 73 74 75 71 69 72 76 73
i=5 (72): pops 69(idx4)->1, 71(idx3)->2     stack [75, 72]
i=6 (76): pops 72->1, 75->4                 stack [76]
```

**Complexity:** O(n).

```python
class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        res, st = [0] * len(temperatures), []
        for i, t in enumerate(temperatures):
            while st and temperatures[st[-1]] < t:
                j = st.pop()
                res[j] = i - j
            st.append(i)
        return res
# test
assert Solution().dailyTemperatures([73, 74, 75, 71, 69, 72, 76, 73]) == [1, 1, 4, 2, 1, 1, 0, 0]
assert Solution().dailyTemperatures([30, 40, 50, 60]) == [1, 1, 1, 0]
assert Solution().dailyTemperatures([60, 50, 40]) == [0, 0, 0]
```

```java
class Solution {
    public int[] dailyTemperatures(int[] t) {
        int[] res = new int[t.length];
        Deque<Integer> st = new ArrayDeque<>();
        for (int i = 0; i < t.length; i++) {
            while (!st.isEmpty() && t[st.peek()] < t[i]) {
                int j = st.pop();
                res[j] = i - j;
            }
            st.push(i);
        }
        return res;
    }
}
// test
check(new Solution().dailyTemperatures(new int[]{73, 74, 75, 71, 69, 72, 76, 73}), new int[]{1, 1, 4, 2, 1, 1, 0, 0});
```

## 853. Car Fleet | Medium
related: 739
tags: monotonic-stack, sort

Sort cars by position, closest to target first. Compute arrival time. A car behind that would arrive **earlier or equal** catches up and merges into the fleet ahead; otherwise it starts a new fleet.

**Complexity:** O(n log n).

```python
class Solution:
    def carFleet(self, target: int, position: List[int], speed: List[int]) -> int:
        fleets, slowest = 0, 0.0
        for p, s in sorted(zip(position, speed), reverse=True):
            t = (target - p) / s
            if t > slowest:          # can't catch the fleet ahead
                fleets += 1
                slowest = t
        return fleets
# test
assert Solution().carFleet(12, [10, 8, 0, 5, 3], [2, 4, 1, 1, 3]) == 3
assert Solution().carFleet(100, [0, 2, 4], [4, 2, 1]) == 1
```

```java
class Solution {
    public int carFleet(int target, int[] position, int[] speed) {
        int n = position.length;
        Integer[] idx = new Integer[n];
        for (int i = 0; i < n; i++) idx[i] = i;
        Arrays.sort(idx, (a, b) -> position[b] - position[a]);
        int fleets = 0;
        double slowest = 0;
        for (int i : idx) {
            double t = (double) (target - position[i]) / speed[i];
            if (t > slowest) { fleets++; slowest = t; }
        }
        return fleets;
    }
}
// test
check(new Solution().carFleet(12, new int[]{10, 8, 0, 5, 3}, new int[]{2, 4, 1, 1, 3}), 3);
```

## 84. Largest Rectangle in Histogram | Hard
related: 85, 42, 739
tags: monotonic-stack

Increasing stack of indices. When bar `i` is lower than the top, the top bar's rectangle is bounded: right edge `i`, left edge = new top after popping. Append a sentinel 0 to flush the stack.

```viz
heights 2 1 5 6 2 3
at i=4 (h=2): pop 6 -> width 1 -> 6
              pop 5 -> width 2 (idx 2..3) -> 10   <- max
```

**Complexity:** O(n).

```python
class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        st, best = [], 0              # increasing stack of indices
        for i, h in enumerate(heights + [0]):
            while st and heights[st[-1]] >= h:
                top = heights[st.pop()]
                left = st[-1] if st else -1
                best = max(best, top * (i - left - 1))
            st.append(i)
        return best
# test
assert Solution().largestRectangleArea([2, 1, 5, 6, 2, 3]) == 10
assert Solution().largestRectangleArea([2, 4]) == 4
```

```java
class Solution {
    public int largestRectangleArea(int[] heights) {
        Deque<Integer> st = new ArrayDeque<>();
        int best = 0, n = heights.length;
        for (int i = 0; i <= n; i++) {
            int h = i == n ? 0 : heights[i];
            while (!st.isEmpty() && heights[st.peek()] >= h) {
                int top = heights[st.pop()];
                int left = st.isEmpty() ? -1 : st.peek();
                best = Math.max(best, top * (i - left - 1));
            }
            st.push(i);
        }
        return best;
    }
}
// test
check(new Solution().largestRectangleArea(new int[]{2, 1, 5, 6, 2, 3}), 10);
check(new Solution().largestRectangleArea(new int[]{2, 4}), 4);
```

## 85. Maximal Rectangle | Hard
related: 84, 221
tags: monotonic-stack, matrix

Treat each row as the floor of a histogram: `h[c]` = consecutive 1s ending at this row. Run **Largest Rectangle in Histogram** on every row.

```viz
row 0  1 0 1 0 0      h = 1 0 1 0 0
row 1  1 0 1 1 1      h = 2 0 2 1 1
row 2  1 1 1 1 1      h = 3 1 3 2 2   -> histogram max 6
```

**Complexity:** O(rows × cols).

```python
class Solution:
    def maximalRectangle(self, matrix: List[List[str]]) -> int:
        if not matrix:
            return 0
        h = [0] * len(matrix[0])
        best = 0
        for row in matrix:
            for c, v in enumerate(row):
                h[c] = h[c] + 1 if v == "1" else 0
            st = []
            for i, x in enumerate(h + [0]):
                while st and h[st[-1]] >= x:
                    top = h[st.pop()]
                    best = max(best, top * (i - (st[-1] if st else -1) - 1))
                st.append(i)
        return best
# test
M = ["10100", "10111", "11111", "10010"]
assert Solution().maximalRectangle([list(r) for r in M]) == 6
assert Solution().maximalRectangle([["0"]]) == 0
assert Solution().maximalRectangle([["1"]]) == 1
assert Solution().maximalRectangle([list(r) for r in ["0110", "1111", "1110"]]) == 6
```

```java
class Solution {
    public int maximalRectangle(char[][] matrix) {
        int cols = matrix[0].length, best = 0;
        int[] h = new int[cols];
        for (char[] row : matrix) {
            for (int c = 0; c < cols; c++) h[c] = row[c] == '1' ? h[c] + 1 : 0;
            Deque<Integer> st = new ArrayDeque<>();
            for (int i = 0; i <= cols; i++) {
                int x = i == cols ? 0 : h[i];
                while (!st.isEmpty() && h[st.peek()] >= x) {
                    int top = h[st.pop()];
                    best = Math.max(best, top * (i - (st.isEmpty() ? -1 : st.peek()) - 1));
                }
                st.push(i);
            }
        }
        return best;
    }
}
// test
char[][] m = Stream.of("10100", "10111", "11111", "10010").map(String::toCharArray).toArray(char[][]::new);
check(new Solution().maximalRectangle(m), 6);
```

## 496. Next Greater Element I | Easy
related: 503, 739
tags: monotonic-stack

Compute next-greater for every value of `nums2` with a decreasing stack and store it in a map; answer queries from `nums1`.

**Complexity:** O(n + m).

```python
class Solution:
    def nextGreaterElement(self, nums1: List[int], nums2: List[int]) -> List[int]:
        nxt, st = {}, []
        for x in nums2:
            while st and st[-1] < x:
                nxt[st.pop()] = x
            st.append(x)
        return [nxt.get(x, -1) for x in nums1]
# test
assert Solution().nextGreaterElement([4, 1, 2], [1, 3, 4, 2]) == [-1, 3, -1]
assert Solution().nextGreaterElement([2, 4], [1, 2, 3, 4]) == [3, -1]
```

```java
class Solution {
    public int[] nextGreaterElement(int[] nums1, int[] nums2) {
        Map<Integer, Integer> nxt = new HashMap<>();
        Deque<Integer> st = new ArrayDeque<>();
        for (int x : nums2) {
            while (!st.isEmpty() && st.peek() < x) nxt.put(st.pop(), x);
            st.push(x);
        }
        int[] res = new int[nums1.length];
        for (int i = 0; i < nums1.length; i++) res[i] = nxt.getOrDefault(nums1[i], -1);
        return res;
    }
}
// test
check(new Solution().nextGreaterElement(new int[]{4, 1, 2}, new int[]{1, 3, 4, 2}), new int[]{-1, 3, -1});
```

## 503. Next Greater Element II | Medium
related: 496, 739
tags: monotonic-stack, circular

Circular array → iterate `2n` indices using `i % n`; only push during the first pass.

**Complexity:** O(n).

```python
class Solution:
    def nextGreaterElements(self, nums: List[int]) -> List[int]:
        n = len(nums)
        res, st = [-1] * n, []
        for i in range(2 * n):
            x = nums[i % n]
            while st and nums[st[-1]] < x:
                res[st.pop()] = x
            if i < n:
                st.append(i)
        return res
# test
assert Solution().nextGreaterElements([1, 2, 1]) == [2, -1, 2]
assert Solution().nextGreaterElements([1, 2, 3, 4, 3]) == [2, 3, 4, -1, 4]
```

```java
class Solution {
    public int[] nextGreaterElements(int[] nums) {
        int n = nums.length;
        int[] res = new int[n];
        Arrays.fill(res, -1);
        Deque<Integer> st = new ArrayDeque<>();
        for (int i = 0; i < 2 * n; i++) {
            int x = nums[i % n];
            while (!st.isEmpty() && nums[st.peek()] < x) res[st.pop()] = x;
            if (i < n) st.push(i);
        }
        return res;
    }
}
// test
check(new Solution().nextGreaterElements(new int[]{1, 2, 3, 4, 3}), new int[]{2, 3, 4, -1, 4});
```

## 394. Decode String | Medium
related: 20, 224
tags: nesting

On `[` push `(current string, repeat count)` and reset; on `]` pop and set `cur = prev + cur * k`.

```viz
3[a2[c]]
'3' k=3   '[' push ("",3)   'a' cur="a"
'2' k=2   '[' push ("a",2)  'c' cur="c"
']' cur = "a" + "c"*2 = "acc"
']' cur = ""  + "acc"*3 = "accaccacc"
```

**Complexity:** O(output length).

```python
class Solution:
    def decodeString(self, s: str) -> str:
        st, cur, k = [], "", 0
        for ch in s:
            if ch.isdigit():
                k = k * 10 + int(ch)
            elif ch == "[":
                st.append((cur, k))
                cur, k = "", 0
            elif ch == "]":
                prev, times = st.pop()
                cur = prev + cur * times
            else:
                cur += ch
        return cur
# test
assert Solution().decodeString("3[a]2[bc]") == "aaabcbc"
assert Solution().decodeString("3[a2[c]]") == "accaccacc"
assert Solution().decodeString("2[abc]3[cd]ef") == "abcabccdcdcdef"
```

```java
class Solution {
    public String decodeString(String s) {
        Deque<StringBuilder> strs = new ArrayDeque<>();
        Deque<Integer> counts = new ArrayDeque<>();
        StringBuilder cur = new StringBuilder();
        int k = 0;
        for (char c : s.toCharArray()) {
            if (Character.isDigit(c)) k = k * 10 + (c - '0');
            else if (c == '[') { strs.push(cur); counts.push(k); cur = new StringBuilder(); k = 0; }
            else if (c == ']') {
                StringBuilder prev = strs.pop();
                prev.append(cur.toString().repeat(counts.pop()));
                cur = prev;
            } else cur.append(c);
        }
        return cur.toString();
    }
}
// test
check(new Solution().decodeString("3[a2[c]]"), "accaccacc");
check(new Solution().decodeString("2[abc]3[cd]ef"), "abcabccdcdcdef");
```

## 71. Simplify Path | Medium
related: 20
tags: string

Split on `/`. Skip `""` and `"."`; `".."` pops; anything else pushes. Join with `/`.

**Complexity:** O(n).

```python
class Solution:
    def simplifyPath(self, path: str) -> str:
        st = []
        for part in path.split("/"):
            if part == "..":
                if st:
                    st.pop()
            elif part and part != ".":
                st.append(part)
        return "/" + "/".join(st)
# test
assert Solution().simplifyPath("/home/") == "/home"
assert Solution().simplifyPath("/../") == "/"
assert Solution().simplifyPath("/home//foo/") == "/home/foo"
assert Solution().simplifyPath("/a/./b/../../c/") == "/c"
```

```java
class Solution {
    public String simplifyPath(String path) {
        Deque<String> st = new ArrayDeque<>();
        for (String p : path.split("/")) {
            if (p.equals("..")) { if (!st.isEmpty()) st.pollLast(); }
            else if (!p.isEmpty() && !p.equals(".")) st.addLast(p);
        }
        return "/" + String.join("/", st);
    }
}
// test
check(new Solution().simplifyPath("/a/./b/../../c/"), "/c");
check(new Solution().simplifyPath("/../"), "/");
```

## 224. Basic Calculator | Hard
related: 227, 150, 394
tags: expression

Only `+ - ( )`. Keep running `res` and `sign`. On `(`, push `(res, sign)` and restart; on `)`, `res = saved_res + saved_sign * res`.

**Complexity:** O(n).

```python
class Solution:
    def calculate(self, s: str) -> int:
        res, num, sign, st = 0, 0, 1, []
        for ch in s:
            if ch.isdigit():
                num = num * 10 + int(ch)
            elif ch in "+-":
                res += sign * num
                num, sign = 0, (1 if ch == "+" else -1)
            elif ch == "(":
                st.append((res, sign))
                res, sign = 0, 1
            elif ch == ")":
                res += sign * num
                num = 0
                prev, psign = st.pop()
                res = prev + psign * res
        return res + sign * num
# test
assert Solution().calculate("1 + 1") == 2
assert Solution().calculate(" 2-1 + 2 ") == 3
assert Solution().calculate("(1+(4+5+2)-3)+(6+8)") == 23
assert Solution().calculate("-(2+3)") == -5
```

```java
class Solution {
    public int calculate(String s) {
        int res = 0, num = 0, sign = 1;
        Deque<Integer> st = new ArrayDeque<>();
        for (char c : s.toCharArray()) {
            if (Character.isDigit(c)) num = num * 10 + (c - '0');
            else if (c == '+' || c == '-') { res += sign * num; num = 0; sign = c == '+' ? 1 : -1; }
            else if (c == '(') { st.push(res); st.push(sign); res = 0; sign = 1; }
            else if (c == ')') {
                res += sign * num; num = 0;
                res = st.pop() * res + st.pop();   // pop sign, then previous result
            }
        }
        return res + sign * num;
    }
}
// test
check(new Solution().calculate("(1+(4+5+2)-3)+(6+8)"), 23);
check(new Solution().calculate("-(2+3)"), -5);
```

## 227. Basic Calculator II | Medium
related: 224, 150
tags: expression

No parentheses, but `* /` bind tighter. When an operator (or the end) is reached, apply the **previous** operator: `+x` / `-x` push, `*` / `/` combine with the top. Sum the stack.

**Complexity:** O(n).

```python
class Solution:
    def calculate(self, s: str) -> int:
        st, num, op = [], 0, "+"
        for i, ch in enumerate(s):
            if ch.isdigit():
                num = num * 10 + int(ch)
            if ch in "+-*/" or i == len(s) - 1:
                if op == "+": st.append(num)
                elif op == "-": st.append(-num)
                elif op == "*": st.append(st.pop() * num)
                else: st.append(int(st.pop() / num))
                op, num = ch, 0
        return sum(st)
# test
assert Solution().calculate("3+2*2") == 7
assert Solution().calculate(" 3/2 ") == 1
assert Solution().calculate(" 3+5 / 2 ") == 5
assert Solution().calculate("14-3/2") == 13
```

```java
class Solution {
    public int calculate(String s) {
        Deque<Integer> st = new ArrayDeque<>();
        int num = 0;
        char op = '+';
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (Character.isDigit(c)) num = num * 10 + (c - '0');
            if ("+-*/".indexOf(c) >= 0 || i == s.length() - 1) {
                if (op == '+') st.push(num);
                else if (op == '-') st.push(-num);
                else if (op == '*') st.push(st.pop() * num);
                else st.push(st.pop() / num);
                op = c; num = 0;
            }
        }
        int res = 0;
        for (int x : st) res += x;
        return res;
    }
}
// test
check(new Solution().calculate("3+2*2"), 7);
check(new Solution().calculate("14-3/2"), 13);
```

## 901. Online Stock Span | Medium
related: 739, 496
tags: monotonic-stack, design

Stack of `(price, span)` with decreasing prices. A new price absorbs (pops) all smaller-or-equal prices and adds their spans.

**Complexity:** amortized O(1) per call.

```python
class StockSpanner:
    def __init__(self):
        self.st = []  # (price, span)

    def next(self, price: int) -> int:
        span = 1
        while self.st and self.st[-1][0] <= price:
            span += self.st.pop()[1]
        self.st.append((price, span))
        return span
# test
s = StockSpanner()
assert [s.next(p) for p in [100, 80, 60, 70, 60, 75, 85]] == [1, 1, 1, 2, 1, 4, 6]
s2 = StockSpanner()
assert [s2.next(p) for p in [31, 41, 48, 59, 79]] == [1, 2, 3, 4, 5]
```

```java
class StockSpanner {
    private final Deque<int[]> st = new ArrayDeque<>(); // {price, span}

    public int next(int price) {
        int span = 1;
        while (!st.isEmpty() && st.peek()[0] <= price) span += st.pop()[1];
        st.push(new int[]{price, span});
        return span;
    }
}
// test
StockSpanner sp = new StockSpanner();
int[] out = new int[7], prices = {100, 80, 60, 70, 60, 75, 85};
for (int i = 0; i < 7; i++) out[i] = sp.next(prices[i]);
check(out, new int[]{1, 1, 1, 2, 1, 4, 6});
```


## 1249. Minimum Remove to Make Valid Parentheses | Medium
related: 20, 301, 921
tags: stack

Stack the indices of unmatched `(`; a `)` with an empty stack is invalid. Afterwards, every index still on the stack is an invalid `(`. Remove all invalid indices.

**Complexity:** O(n).

```python
class Solution:
    def minRemoveToMakeValid(self, s: str) -> str:
        st, bad = [], set()
        for i, ch in enumerate(s):
            if ch == "(":
                st.append(i)
            elif ch == ")":
                if st:
                    st.pop()
                else:
                    bad.add(i)
        bad.update(st)
        return "".join(ch for i, ch in enumerate(s) if i not in bad)
# test
assert Solution().minRemoveToMakeValid("lee(t(c)o)de)") == "lee(t(c)o)de"
assert Solution().minRemoveToMakeValid("a)b(c)d") == "ab(c)d"
assert Solution().minRemoveToMakeValid("))((") == ""
```

```java
class Solution {
    public String minRemoveToMakeValid(String s) {
        Deque<Integer> st = new ArrayDeque<>();
        boolean[] bad = new boolean[s.length()];
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (c == '(') st.push(i);
            else if (c == ')') { if (st.isEmpty()) bad[i] = true; else st.pop(); }
        }
        for (int i : st) bad[i] = true;
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < s.length(); i++) if (!bad[i]) sb.append(s.charAt(i));
        return sb.toString();
    }
}
// test
check(new Solution().minRemoveToMakeValid("lee(t(c)o)de)"), "lee(t(c)o)de");
check(new Solution().minRemoveToMakeValid("))(("), "");
```

## 735. Asteroid Collision | Medium
related: 20
tags: stack, simulation

Only a right-mover (stack top > 0) followed by a left-mover (x < 0) collide. While that happens, the smaller explodes; equal sizes both explode.

**Complexity:** O(n).

```python
class Solution:
    def asteroidCollision(self, asteroids: List[int]) -> List[int]:
        st = []
        for x in asteroids:
            alive = True
            while alive and x < 0 and st and st[-1] > 0:
                if st[-1] < -x:
                    st.pop()
                    continue
                if st[-1] == -x:
                    st.pop()
                alive = False
            if alive:
                st.append(x)
        return st
# test
assert Solution().asteroidCollision([5, 10, -5]) == [5, 10]
assert Solution().asteroidCollision([8, -8]) == []
assert Solution().asteroidCollision([10, 2, -5]) == [10]
assert Solution().asteroidCollision([-2, -1, 1, 2]) == [-2, -1, 1, 2]
```

```java
class Solution {
    public int[] asteroidCollision(int[] asteroids) {
        Deque<Integer> st = new ArrayDeque<>();
        for (int x : asteroids) {
            boolean alive = true;
            while (alive && x < 0 && !st.isEmpty() && st.peekLast() > 0) {
                int top = st.peekLast();
                if (top < -x) { st.pollLast(); continue; }
                if (top == -x) st.pollLast();
                alive = false;
            }
            if (alive) st.addLast(x);
        }
        return st.stream().mapToInt(Integer::intValue).toArray();
    }
}
// test
check(new Solution().asteroidCollision(new int[]{10, 2, -5}), new int[]{10});
check(new Solution().asteroidCollision(new int[]{-2, -1, 1, 2}), new int[]{-2, -1, 1, 2});
```

## 402. Remove K Digits | Medium
related: 316, 321
tags: monotonic-stack, greedy

To minimize the number, a larger digit followed by a smaller one should go. Keep an **increasing** stack; pop while the top is bigger than the incoming digit and removals remain. Remove leftovers from the end; strip leading zeros.

```viz
num 1432219, k=3
1 -> 14 -> 13 (pop 4) -> 12 (pop 3) -> 122 -> 121 (pop 2) -> 1219
```

**Complexity:** O(n).

```python
class Solution:
    def removeKdigits(self, num: str, k: int) -> str:
        st = []
        for d in num:
            while k and st and st[-1] > d:
                st.pop()
                k -= 1
            st.append(d)
        if k:
            st = st[:-k]
        return "".join(st).lstrip("0") or "0"
# test
assert Solution().removeKdigits("1432219", 3) == "1219"
assert Solution().removeKdigits("10200", 1) == "200"
assert Solution().removeKdigits("10", 2) == "0"
```

```java
class Solution {
    public String removeKdigits(String num, int k) {
        StringBuilder st = new StringBuilder();
        for (char d : num.toCharArray()) {
            while (k > 0 && st.length() > 0 && st.charAt(st.length() - 1) > d) { st.deleteCharAt(st.length() - 1); k--; }
            st.append(d);
        }
        st.setLength(st.length() - k);
        int i = 0;
        while (i < st.length() && st.charAt(i) == '0') i++;
        return i == st.length() ? "0" : st.substring(i);
    }
}
// test
check(new Solution().removeKdigits("1432219", 3), "1219");
check(new Solution().removeKdigits("10200", 1), "200");
check(new Solution().removeKdigits("10", 2), "0");
```

## 316. Remove Duplicate Letters | Medium
related: 402, 1081
tags: monotonic-stack, greedy

Smallest lexicographic result with each letter once. Keep an increasing stack; pop the top if it's larger than the current letter **and appears again later** (we can re-add it). Skip letters already in the stack.

**Complexity:** O(n).

```python
class Solution:
    def removeDuplicateLetters(self, s: str) -> str:
        last = {ch: i for i, ch in enumerate(s)}
        st, used = [], set()
        for i, ch in enumerate(s):
            if ch in used:
                continue
            while st and st[-1] > ch and last[st[-1]] > i:
                used.remove(st.pop())
            st.append(ch)
            used.add(ch)
        return "".join(st)
# test
assert Solution().removeDuplicateLetters("bcabc") == "abc"
assert Solution().removeDuplicateLetters("cbacdcbc") == "acdb"
```

```java
class Solution {
    public String removeDuplicateLetters(String s) {
        int[] last = new int[26];
        for (int i = 0; i < s.length(); i++) last[s.charAt(i) - 'a'] = i;
        boolean[] used = new boolean[26];
        StringBuilder st = new StringBuilder();
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (used[c - 'a']) continue;
            while (st.length() > 0 && st.charAt(st.length() - 1) > c && last[st.charAt(st.length() - 1) - 'a'] > i) {
                used[st.charAt(st.length() - 1) - 'a'] = false;
                st.deleteCharAt(st.length() - 1);
            }
            st.append(c);
            used[c - 'a'] = true;
        }
        return st.toString();
    }
}
// test
check(new Solution().removeDuplicateLetters("cbacdcbc"), "acdb");
```

## 456. 132 Pattern | Medium
related: 496
tags: monotonic-stack

Scan from the right keeping a decreasing stack of candidates for "3" and `two` = the best "2" (largest value popped, i.e. smaller than some later-left "3"). If the current number is less than `two`, it's a valid "1".

**Complexity:** O(n).

```python
class Solution:
    def find132pattern(self, nums: List[int]) -> bool:
        st, two = [], float("-inf")
        for x in reversed(nums):
            if x < two:
                return True
            while st and st[-1] < x:
                two = st.pop()          # x acts as "3", popped value as "2"
            st.append(x)
        return False
# test
assert Solution().find132pattern([1, 2, 3, 4]) is False
assert Solution().find132pattern([3, 1, 4, 2]) is True
assert Solution().find132pattern([-1, 3, 2, 0]) is True
```

```java
class Solution {
    public boolean find132pattern(int[] nums) {
        Deque<Integer> st = new ArrayDeque<>();
        long two = Long.MIN_VALUE;
        for (int i = nums.length - 1; i >= 0; i--) {
            if (nums[i] < two) return true;
            while (!st.isEmpty() && st.peek() < nums[i]) two = st.pop();
            st.push(nums[i]);
        }
        return false;
    }
}
// test
check(new Solution().find132pattern(new int[]{3, 1, 4, 2}), true);
check(new Solution().find132pattern(new int[]{1, 2, 3, 4}), false);
```

## 907. Sum of Subarray Minimums | Medium
related: 84, 2104
tags: monotonic-stack, contribution

Count, for each element, how many subarrays it is the minimum of: `left[i]` = distance to the previous **smaller** element, `right[i]` = distance to the next **smaller-or-equal** (asymmetric to avoid double-counting ties). Contribution = `arr[i] · left · right`.

**Complexity:** O(n).

```python
class Solution:
    def sumSubarrayMins(self, arr: List[int]) -> int:
        MOD, n = 10 ** 9 + 7, len(arr)
        left, right, st = [0] * n, [0] * n, []
        for i in range(n):
            while st and arr[st[-1]] > arr[i]:
                st.pop()
            left[i] = i - (st[-1] if st else -1)
            st.append(i)
        st = []
        for i in range(n - 1, -1, -1):
            while st and arr[st[-1]] >= arr[i]:
                st.pop()
            right[i] = (st[-1] if st else n) - i
            st.append(i)
        return sum(a * l * r for a, l, r in zip(arr, left, right)) % MOD
# test
assert Solution().sumSubarrayMins([3, 1, 2, 4]) == 17
assert Solution().sumSubarrayMins([11, 81, 94, 43, 3]) == 444
```

```java
class Solution {
    public int sumSubarrayMins(int[] arr) {
        int n = arr.length;
        long MOD = 1_000_000_007L, res = 0;
        int[] left = new int[n], right = new int[n];
        Deque<Integer> st = new ArrayDeque<>();
        for (int i = 0; i < n; i++) {
            while (!st.isEmpty() && arr[st.peek()] > arr[i]) st.pop();
            left[i] = i - (st.isEmpty() ? -1 : st.peek());
            st.push(i);
        }
        st.clear();
        for (int i = n - 1; i >= 0; i--) {
            while (!st.isEmpty() && arr[st.peek()] >= arr[i]) st.pop();
            right[i] = (st.isEmpty() ? n : st.peek()) - i;
            st.push(i);
        }
        for (int i = 0; i < n; i++) res = (res + (long) arr[i] * left[i] % MOD * right[i]) % MOD;
        return (int) res;
    }
}
// test
check(new Solution().sumSubarrayMins(new int[]{3, 1, 2, 4}), 17);
check(new Solution().sumSubarrayMins(new int[]{11, 81, 94, 43, 3}), 444);
```

## 1047. Remove All Adjacent Duplicates In String | Easy
related: 1209, 20
tags: stack

Use the output as a stack: if the incoming char equals the top, pop; else push.

**Complexity:** O(n).

```python
class Solution:
    def removeDuplicates(self, s: str) -> str:
        st = []
        for ch in s:
            if st and st[-1] == ch:
                st.pop()
            else:
                st.append(ch)
        return "".join(st)
# test
assert Solution().removeDuplicates("abbaca") == "ca"
assert Solution().removeDuplicates("azxxzy") == "ay"
```

```java
class Solution {
    public String removeDuplicates(String s) {
        StringBuilder st = new StringBuilder();
        for (char c : s.toCharArray()) {
            if (st.length() > 0 && st.charAt(st.length() - 1) == c) st.deleteCharAt(st.length() - 1);
            else st.append(c);
        }
        return st.toString();
    }
}
// test
check(new Solution().removeDuplicates("azxxzy"), "ay");
```

## 225. Implement Stack using Queues | Easy
related: 232
tags: design, queue

One queue: after pushing `x`, rotate the previous `n−1` elements behind it so `x` is at the front. Push O(n), pop/top O(1).

**Complexity:** push O(n), others O(1).

```python
class MyStack:
    def __init__(self):
        self.q = deque()

    def push(self, x: int) -> None:
        self.q.append(x)
        for _ in range(len(self.q) - 1):
            self.q.append(self.q.popleft())

    def pop(self) -> int:
        return self.q.popleft()

    def top(self) -> int:
        return self.q[0]

    def empty(self) -> bool:
        return not self.q
# test
s = MyStack()
s.push(1); s.push(2)
assert s.top() == 2 and s.pop() == 2 and not s.empty() and s.pop() == 1 and s.empty()
s2 = MyStack(); s2.push(5); s2.push(6); s2.push(7)
assert s2.pop() == 7 and s2.top() == 6 and s2.pop() == 6 and s2.top() == 5 and not s2.empty()
```

```java
class MyStack {
    private final Deque<Integer> q = new ArrayDeque<>();

    public void push(int x) {
        q.addLast(x);
        for (int i = 1; i < q.size(); i++) q.addLast(q.pollFirst());
    }

    public int pop() { return q.pollFirst(); }
    public int top() { return q.peekFirst(); }
    public boolean empty() { return q.isEmpty(); }
}
// test
MyStack s = new MyStack();
s.push(1); s.push(2);
check(s.top(), 2);
check(s.pop(), 2);
check(s.empty(), false);
```
