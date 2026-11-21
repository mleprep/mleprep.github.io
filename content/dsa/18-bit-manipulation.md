# Bit Manipulation

| Trick | Meaning |
|---|---|
| `x & (x - 1)` | clear the lowest set bit (count bits, power-of-two test) |
| `x & -x` | isolate the lowest set bit |
| `a ^ a = 0`, `a ^ 0 = a` | XOR cancels pairs → find the unique element |
| `(x >> i) & 1` | read bit i |
| `x \| (1 << i)`, `x & ~(1 << i)` | set / clear bit i |

**Python caveat:** ints are unbounded — emulate 32-bit with `& 0xFFFFFFFF` and convert back with `x if x < 2**31 else x - 2**32`.

```viz
x       = 1 0 1 1 0 0
x - 1   = 1 0 1 0 1 1
x&(x-1) = 1 0 1 0 0 0      lowest 1 removed
```

## 136. Single Number | Easy
related: 137, 268, 260
tags: xor

```problem
Every value in `nums` appears twice except one, which appears once. Return that one, in O(n) time and O(1) extra space.

**Example:** `[4,1,2,1,2]` → `4`
```

XOR everything: pairs cancel, the single number remains.

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        return reduce(lambda a, b: a ^ b, nums)
# test
assert Solution().singleNumber([4, 1, 2, 1, 2]) == 4
assert Solution().singleNumber([2, 2, 1]) == 1
assert Solution().singleNumber([1]) == 1
```

```java
class Solution {
    public int singleNumber(int[] nums) {
        int x = 0;
        for (int v : nums) x ^= v;
        return x;
    }
}
// test
check(new Solution().singleNumber(new int[]{4, 1, 2, 1, 2}), 4);
```

```go
func singleNumber(nums []int) int {
	x := 0
	for _, v := range nums { // pairs cancel: a ^ a = 0
		x ^= v
	}
	return x
}
// test
assertEq(singleNumber([]int{4, 1, 2, 1, 2}), 4)
assertEq(singleNumber([]int{2, 2, 1}), 1)
assertEq(singleNumber([]int{1}), 1)
```

## 137. Single Number II | Medium
related: 136
tags: bit-counting

```problem
Every value in `nums` appears three times except one, which appears once. Return that one, in O(n) time and O(1) extra space.

**Example:** `[2,2,3,2]` → `3`; `[0,1,0,1,0,1,99]` → `99`
```

Every other number appears 3 times → for each bit position, the count of set bits mod 3 is the single number's bit. (Handle the sign bit for negatives in Python.)

**Complexity:** O(32·n).

```python
class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        res = 0
        for b in range(32):
            if sum((x >> b) & 1 for x in nums) % 3:
                res |= 1 << b
        return res - (1 << 32) if res >= 1 << 31 else res
# test
assert Solution().singleNumber([2, 2, 3, 2]) == 3
assert Solution().singleNumber([0, 1, 0, 1, 0, 1, 99]) == 99
assert Solution().singleNumber([-2, -2, 1, 1, 4, 1, 4, 4, -4, -2]) == -4
```

```java
class Solution {
    public int singleNumber(int[] nums) {
        int res = 0;
        for (int b = 0; b < 32; b++) {
            int cnt = 0;
            for (int x : nums) cnt += (x >> b) & 1;
            if (cnt % 3 != 0) res |= 1 << b;
        }
        return res;
    }
}
// test
check(new Solution().singleNumber(new int[]{2, 2, 3, 2}), 3);
check(new Solution().singleNumber(new int[]{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}), -4);
```

```go
func singleNumber(nums []int) int {
	res := 0
	for b := 0; b < 32; b++ { // each bit's count mod 3 is the single number's bit
		cnt := 0
		for _, x := range nums {
			cnt += (x >> b) & 1
		}
		if cnt%3 != 0 {
			res |= 1 << b
		}
	}
	return int(int32(res)) // bit 31 is the sign of a 32-bit int
}
// test
assertEq(singleNumber([]int{2, 2, 3, 2}), 3)
assertEq(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}), 99)
assertEq(singleNumber([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}), -4)
```

## 191. Number of 1 Bits | Easy
related: 338, 190
tags: blind75, brian-kernighan

```problem
Return the number of 1 bits in the binary form of the positive integer `n` (its Hamming weight).

**Example:** `11` (`1011`) → `3`; `2147483645` → `30`

**Constraints:** follow-up: the function is called many times.
```

Repeatedly clear the lowest set bit with `n &= n − 1`; count iterations.

**Complexity:** O(number of set bits).

```python
class Solution:
    def hammingWeight(self, n: int) -> int:
        count = 0
        while n:
            n &= n - 1
            count += 1
        return count
# test
assert Solution().hammingWeight(11) == 3
assert Solution().hammingWeight(2147483645) == 30
```

```java
class Solution {
    public int hammingWeight(int n) {
        int count = 0;
        for (; n != 0; n &= n - 1) count++;   // works for negative (unsigned) inputs too
        return count;
    }
}
// test
check(new Solution().hammingWeight(11), 3);
check(new Solution().hammingWeight(-3), 31);
```

```go
func hammingWeight(n int) int {
	count := 0
	for ; n != 0; n &= n - 1 { // n & (n-1) clears the lowest set bit
		count++
	}
	return count
}
// test
assertEq(hammingWeight(11), 3)
assertEq(hammingWeight(2147483645), 30)
```

## 338. Counting Bits | Easy
related: 191
tags: blind75, dp

```problem
For every i from 0 to `n`, count the 1 bits in i's binary form, and return the counts as a list.

**Example:** `n = 5` → `[0,1,1,2,1,2]`

**Constraints:** follow-up: O(n) without a built-in popcount.
```

`bits[i] = bits[i >> 1] + (i & 1)` — drop the last bit, reuse the answer.

**Complexity:** O(n).

```python
class Solution:
    def countBits(self, n: int) -> List[int]:
        bits = [0] * (n + 1)
        for i in range(1, n + 1):
            bits[i] = bits[i >> 1] + (i & 1)
        return bits
# test
assert Solution().countBits(5) == [0, 1, 1, 2, 1, 2]
assert Solution().countBits(0) == [0]
assert Solution().countBits(2) == [0, 1, 1]
```

```java
class Solution {
    public int[] countBits(int n) {
        int[] bits = new int[n + 1];
        for (int i = 1; i <= n; i++) bits[i] = bits[i >> 1] + (i & 1);
        return bits;
    }
}
// test
check(new Solution().countBits(5), new int[]{0, 1, 1, 2, 1, 2});
```

```go
func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i>>1] + i&1 // i's bits = (i/2)'s bits + the last bit
	}
	return bits
}
// test
assertEq(countBits(5), []int{0, 1, 1, 2, 1, 2})
assertEq(countBits(0), []int{0})
assertEq(countBits(2), []int{0, 1, 1})
```

## 190. Reverse Bits | Easy
related: 191, 7
tags: blind75

```problem
Reverse the bits of a 32-bit unsigned integer.

**Example:** `43261596` (`00000010100101000001111010011100`) → `964176192` (`00111001011110000010100101000000`)

**Constraints:** follow-up: the function is called many times.
```

Shift the result left and append the lowest bit of `n`, 32 times. Java: use `>>>` (unsigned shift).

**Complexity:** O(32).

```python
class Solution:
    def reverseBits(self, n: int) -> int:
        res = 0
        for _ in range(32):
            res = (res << 1) | (n & 1)
            n >>= 1
        return res
# test
assert Solution().reverseBits(43261596) == 964176192
assert Solution().reverseBits(4294967293) == 3221225471
```

```java
class Solution {
    public int reverseBits(int n) {
        int res = 0;
        for (int i = 0; i < 32; i++) {
            res = (res << 1) | (n & 1);
            n >>>= 1;
        }
        return res;
    }
}
// test
check(new Solution().reverseBits(43261596), 964176192);
check(new Solution().reverseBits(-3), -1073741825);
```

```go
func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = res<<1 | num&1 // shift the lowest bit of num into res
		num >>= 1
	}
	return res
}
// test
assertEq(reverseBits(43261596), 964176192)
assertEq(reverseBits(4294967293), 3221225471)
```

## 268. Missing Number | Easy
related: 41, 136, 287
tags: blind75, xor

```problem
`nums` contains n distinct numbers from the range [0, n], so exactly one is missing. Return it.

**Example:** `[3,0,1]` → `2`; `[9,6,4,2,3,5,7,0,1]` → `8`

**Constraints:** follow-up: O(1) extra space and O(n) time.
```

XOR all indices `0..n` with all values; everything pairs up except the missing number. (Or `n(n+1)/2 − sum`.)

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        x = len(nums)
        for i, v in enumerate(nums):
            x ^= i ^ v
        return x
# test
assert Solution().missingNumber([3, 0, 1]) == 2
assert Solution().missingNumber([9, 6, 4, 2, 3, 5, 7, 0, 1]) == 8
```

```java
class Solution {
    public int missingNumber(int[] nums) {
        int x = nums.length;
        for (int i = 0; i < nums.length; i++) x ^= i ^ nums[i];
        return x;
    }
}
// test
check(new Solution().missingNumber(new int[]{9, 6, 4, 2, 3, 5, 7, 0, 1}), 8);
```

```go
func missingNumber(nums []int) int {
	x := len(nums)
	for i, v := range nums { // xor all indices 0..n and all values: only the missing one survives
		x ^= i ^ v
	}
	return x
}
// test
assertEq(missingNumber([]int{3, 0, 1}), 2)
assertEq(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}), 8)
```

## 371. Sum of Two Integers | Medium
related: 67, 2
tags: blind75, adder

```problem
Return the sum of integers `a` and `b` without using the `+` or `-` operators.

**Example:** `a = 2, b = 3` → `5`

**Constraints:** −1000 ≤ a, b ≤ 1000.
```

Without `+`: `a ^ b` is the sum without carries, `(a & b) << 1` the carries. Repeat until no carry. Python needs a 32-bit mask.

```viz
a = 0101 (5), b = 0011 (3)
xor 0110  carry 0010   -> a=0110 b=0010
xor 0100  carry 0100   -> a=0100 b=0100
xor 0000  carry 1000   -> a=0000 b=1000
xor 1000  carry 0000   -> 8
```

**Complexity:** O(32).

```python
class Solution:
    def getSum(self, a: int, b: int) -> int:
        mask = 0xFFFFFFFF
        while b & mask:
            a, b = (a ^ b) & mask, ((a & b) << 1) & mask
        a &= mask
        return a if a < 1 << 31 else a - (1 << 32)
# test
assert Solution().getSum(1, 2) == 3
assert Solution().getSum(2, 3) == 5
assert Solution().getSum(-1, 1) == 0
assert Solution().getSum(-12, -8) == -20
```

```java
class Solution {
    public int getSum(int a, int b) {
        while (b != 0) {
            int carry = (a & b) << 1;
            a ^= b;
            b = carry;
        }
        return a;
    }
}
// test
check(new Solution().getSum(-12, -8), -20);
check(new Solution().getSum(-1, 1), 0);
```

```go
func getSum(a int, b int) int {
	for b != 0 { // sum without carry, then the carry shifted left; ints wrap, so negatives work
		a, b = a^b, (a&b)<<1
	}
	return a
}
// test
assertEq(getSum(1, 2), 3)
assertEq(getSum(2, 3), 5)
assertEq(getSum(-1, 1), 0)
assertEq(getSum(-12, -8), -20)
```

## 7. Reverse Integer | Medium
related: 190, 8
tags: overflow

```problem
Reverse the digits of a 32-bit signed integer `x`, keeping its sign. Return 0 if the result falls outside [−2³¹, 2³¹ − 1], and assume 64-bit integers aren't available.

**Example:** `123` → `321`; `-123` → `-321`; `120` → `21`
```

Pop digits with `% 10` and push with `* 10`. Before pushing, check that the result stays within `[−2³¹, 2³¹−1]`.

**Complexity:** O(log x).

```python
class Solution:
    def reverse(self, x: int) -> int:
        sign = -1 if x < 0 else 1
        x, res = abs(x), 0
        while x:
            x, d = divmod(x, 10)
            res = res * 10 + d
        res *= sign
        return res if -2 ** 31 <= res <= 2 ** 31 - 1 else 0
# test
assert Solution().reverse(123) == 321
assert Solution().reverse(-123) == -321
assert Solution().reverse(120) == 21
assert Solution().reverse(1534236469) == 0
```

```java
class Solution {
    public int reverse(int x) {
        int res = 0;
        while (x != 0) {
            int d = x % 10;                      // keeps sign in Java
            if (res > Integer.MAX_VALUE / 10 || res < Integer.MIN_VALUE / 10) return 0;
            res = res * 10 + d;
            x /= 10;
        }
        return res;
    }
}
// test
check(new Solution().reverse(-123), -321);
check(new Solution().reverse(1534236469), 0);
```

```go
func reverse(x int) int {
	res := 0
	for ; x != 0; x /= 10 { // Go's % keeps the sign, so negatives just work
		res = res*10 + x%10
	}
	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}
	return res
}
// test
assertEq(reverse(123), 321)
assertEq(reverse(-123), -321)
assertEq(reverse(120), 21)
assertEq(reverse(1534236469), 0)
```


## 231. Power of Two | Easy
related: 191, 338
tags: bit-trick

```problem
Decide whether the integer `n` is a power of two (2ˣ for some integer x ≥ 0).

**Example:** `16` → `true`; `3` → `false`; `1` → `true`

**Constraints:** follow-up: without loops or recursion.
```

A power of two has exactly one set bit: `n > 0 and n & (n − 1) == 0`.

**Complexity:** O(1).

```python
class Solution:
    def isPowerOfTwo(self, n: int) -> bool:
        return n > 0 and n & (n - 1) == 0
# test
assert Solution().isPowerOfTwo(16) is True
assert Solution().isPowerOfTwo(3) is False
assert Solution().isPowerOfTwo(-16) is False
```

```java
class Solution {
    public boolean isPowerOfTwo(int n) { return n > 0 && (n & (n - 1)) == 0; }
}
// test
check(new Solution().isPowerOfTwo(16), true);
check(new Solution().isPowerOfTwo(Integer.MIN_VALUE), false);
```

```go
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0 // exactly one bit set
}
// test
assertEq(isPowerOfTwo(16), true)
assertEq(isPowerOfTwo(3), false)
assertEq(isPowerOfTwo(-16), false)
```

## 461. Hamming Distance | Easy
related: 191
tags: xor, popcount

```problem
Return the Hamming distance between `x` and `y`: the number of bit positions in which they differ.

**Example:** `x = 1, y = 4` → `2` (`001` vs `100`)
```

Differing bits = set bits of `x ^ y`.

**Complexity:** O(1).

```python
class Solution:
    def hammingDistance(self, x: int, y: int) -> int:
        return bin(x ^ y).count("1")
# test
assert Solution().hammingDistance(1, 4) == 2
assert Solution().hammingDistance(3, 1) == 1
```

```java
class Solution {
    public int hammingDistance(int x, int y) { return Integer.bitCount(x ^ y); }
}
// test
check(new Solution().hammingDistance(1, 4), 2);
```

```go
func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y)) // the differing bits
}
// test
assertEq(hammingDistance(1, 4), 2)
assertEq(hammingDistance(3, 1), 1)
```

## 260. Single Number III | Medium
related: 136, 137
tags: xor, partition

```problem
In `nums`, exactly two values appear once and every other value appears exactly twice. Return the two single values, in any order, using O(n) time and O(1) extra space.

**Example:** `[1,2,1,3,2,5]` → `[3,5]`
```

XOR all → `a ^ b` (non-zero). Its lowest set bit `diff = x & −x` is a bit where a and b differ: split numbers by that bit and XOR each group separately.

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def singleNumber(self, nums: List[int]) -> List[int]:
        x = reduce(lambda a, b: a ^ b, nums)
        diff = x & -x
        a = 0
        for v in nums:
            if v & diff:
                a ^= v
        return sorted([a, x ^ a])
# test
assert sorted(Solution().singleNumber([1, 2, 1, 3, 2, 5])) == sorted([3, 5])
assert sorted(Solution().singleNumber([-1, 0])) == sorted([-1, 0])
```

```java
class Solution {
    public int[] singleNumber(int[] nums) {
        int x = 0;
        for (int v : nums) x ^= v;
        int diff = x & -x, a = 0;
        for (int v : nums) if ((v & diff) != 0) a ^= v;
        return new int[]{Math.min(a, x ^ a), Math.max(a, x ^ a)};
    }
}
// test
check(new Solution().singleNumber(new int[]{1, 2, 1, 3, 2, 5}), new int[]{3, 5});
```

```go
func singleNumber(nums []int) []int {
	x := 0
	for _, v := range nums {
		x ^= v // = a ^ b
	}
	diff := x & -x // a bit where a and b differ splits the numbers into two groups
	a := 0
	for _, v := range nums {
		if v&diff != 0 {
			a ^= v
		}
	}
	return []int{a, x ^ a}
}
// test
assertUnordered(singleNumber([]int{1, 2, 1, 3, 2, 5}), []int{3, 5})
assertUnordered(singleNumber([]int{-1, 0}), []int{-1, 0})
```

## 201. Bitwise AND of Numbers Range | Medium
related: 191
tags: common-prefix

```problem
Return the bitwise AND of every integer in the range [left, right], inclusive.

**Example:** `left = 5, right = 7` → `4`; `left = 1, right = 2147483647` → `0`
```

Any bit that flips somewhere in `[left, right]` becomes 0. The result is the **common binary prefix** of left and right: shift both right until equal, then shift back.

**Complexity:** O(32).

```python
class Solution:
    def rangeBitwiseAnd(self, left: int, right: int) -> int:
        shift = 0
        while left != right:
            left, right, shift = left >> 1, right >> 1, shift + 1
        return left << shift
# test
assert Solution().rangeBitwiseAnd(5, 7) == 4
assert Solution().rangeBitwiseAnd(0, 0) == 0
assert Solution().rangeBitwiseAnd(1, 2147483647) == 0
```

```java
class Solution {
    public int rangeBitwiseAnd(int left, int right) {
        while (right > left) right &= right - 1;   // drop lowest set bits until within the prefix
        return right;
    }
}
// test
check(new Solution().rangeBitwiseAnd(5, 7), 4);
check(new Solution().rangeBitwiseAnd(1, 2147483647), 0);
```

```go
func rangeBitwiseAnd(left int, right int) int {
	shift := 0
	for left != right { // the answer is the common binary prefix
		left, right = left>>1, right>>1
		shift++
	}
	return left << shift
}
// test
assertEq(rangeBitwiseAnd(5, 7), 4)
assertEq(rangeBitwiseAnd(0, 0), 0)
assertEq(rangeBitwiseAnd(1, 2147483647), 0)
```

## 89. Gray Code | Medium
related: 784
tags: bit-trick

```problem
Return an n-bit Gray code: a sequence of all 2ⁿ integers in [0, 2ⁿ − 1], starting at 0, in which consecutive values (including the last and the first) differ in exactly one bit. Any valid sequence is accepted.

**Example:** `n = 2` → `[0,1,3,2]`

**Constraints:** n ≤ 16.
```

The i-th Gray code is `i ^ (i >> 1)`: consecutive values differ in exactly one bit, and the sequence wraps around.

```viz
i     000 001 010 011 100 101 110 111
gray  000 001 011 010 110 111 101 100
```

**Complexity:** O(2ⁿ).

```python
class Solution:
    def grayCode(self, n: int) -> List[int]:
        return [i ^ (i >> 1) for i in range(1 << n)]
# test
# any valid Gray code: starts at 0, visits all 2^n values once, neighbours (incl. wrap-around) differ by one bit
for n in (1, 2, 3, 4):
    g = Solution().grayCode(n)
    assert g[0] == 0 and sorted(g) == list(range(1 << n))
    assert all(bin(g[i] ^ g[(i + 1) % len(g)]).count("1") == 1 for i in range(len(g)))
```

```java
class Solution {
    public List<Integer> grayCode(int n) {
        List<Integer> res = new ArrayList<>();
        for (int i = 0; i < 1 << n; i++) res.add(i ^ (i >> 1));
        return res;
    }
}
// test
check(new Solution().grayCode(2), "[0, 1, 3, 2]");
```

```go
func grayCode(n int) []int {
	res := make([]int, 1<<n)
	for i := range res {
		res[i] = i ^ i>>1 // reflected binary code
	}
	return res
}
// test
// any valid Gray code: starts at 0, visits all 2^n values once, neighbours (incl. wrap-around) differ by one bit
for n := 1; n <= 4; n++ {
	g := grayCode(n)
	ok := len(g) == 1<<n && g[0] == 0
	seen := map[int]bool{}
	for i := 0; ok && i < len(g); i++ {
		ok = g[i] >= 0 && g[i] < 1<<n && !seen[g[i]] && bits.OnesCount(uint(g[i]^g[(i+1)%len(g)])) == 1
		seen[g[i]] = true
	}
	assertEq(ok, true)
}
```

## 1318. Minimum Flips to Make a OR b Equal to c | Medium
related: 191, 461
tags: bit-trick

```problem
Return the minimum number of single-bit flips in `a` and `b` that make `a OR b == c`.

**Example:** `a = 2, b = 6, c = 5` → `3`; `a = 4, b = 2, c = 7` → `1`
```

Bits are independent. For each position: if `c`'s bit is 1 we need at least one of `a`, `b` set (1 flip if both are 0); if it's 0, every set bit among `a`, `b` must be flipped. Equivalently `popcount((a | b) ^ c) + popcount(a & b & ~c)`.

**Complexity:** O(log max(a, b, c)).

```python
class Solution:
    def minFlips(self, a: int, b: int, c: int) -> int:
        flips = 0
        while a or b or c:
            x, y, z = a & 1, b & 1, c & 1
            flips += (x | y) ^ 1 if z else x + y
            a, b, c = a >> 1, b >> 1, c >> 1
        return flips
# test
assert Solution().minFlips(2, 6, 5) == 3
assert Solution().minFlips(4, 2, 7) == 1
assert Solution().minFlips(1, 2, 3) == 0
assert Solution().minFlips(8, 3, 5) == 3
assert all(Solution().minFlips(a, b, c) == bin((a | b) ^ c).count("1") + bin(a & b & ~c).count("1")
           for a in range(16) for b in range(16) for c in range(16))
```

```java
class Solution {
    public int minFlips(int a, int b, int c) {
        return Integer.bitCount((a | b) ^ c) + Integer.bitCount(a & b & ~c);
    }
}
// test
check(new Solution().minFlips(2, 6, 5), 3);
check(new Solution().minFlips(4, 2, 7), 1);
```

```go
func minFlips(a int, b int, c int) int {
	flips := 0
	for a > 0 || b > 0 || c > 0 {
		x, y, z := a&1, b&1, c&1
		if z == 1 {
			flips += (x | y) ^ 1 // need at least one 1
		} else {
			flips += x + y // both must be 0
		}
		a, b, c = a>>1, b>>1, c>>1
	}
	return flips
}
// test
assertEq(minFlips(2, 6, 5), 3)
assertEq(minFlips(4, 2, 7), 1)
assertEq(minFlips(1, 2, 3), 0)
assertEq(minFlips(8, 3, 5), 3)
ok := true
for a := 0; a < 16; a++ {
	for b := 0; b < 16; b++ {
		for c := 0; c < 16; c++ {
			ok = ok && minFlips(a, b, c) == bits.OnesCount(uint((a|b)^c))+bits.OnesCount(uint(a&b&^c))
		}
	}
}
assertEq(ok, true)
```

## 3133. Minimum Array End | Medium
related: 1318, 201
tags: bit-trick

```problem
Build a strictly increasing array of `n` positive integers whose bitwise AND is exactly `x`. Return the smallest possible value of its last element.

**Example:** `n = 3, x = 4` → `6` (`[4,5,6]`); `n = 2, x = 7` → `15`

**Constraints:** n, x ≤ 10⁸.
```

Every element must contain all of x's set bits, so elements are x with some of its **zero bits** filled in. The n-th smallest such number puts the binary digits of `n − 1` into x's zero-bit positions (low to high) — "deposit" `n − 1` into the free bits.

```viz
x = 4      = 0b100   free bits: positions 0, 1, 3, 4, ...
n = 3 -> n-1 = 0b10  -> bit 0 <- 0, bit 1 <- 1   -> 0b110 = 6     (array 4, 5, 6)
```

**Complexity:** O(log n + log x).

```python
class Solution:
    def minEnd(self, n: int, x: int) -> int:
        res, k, bit = x, n - 1, 1
        while k:
            if not x & bit:              # free position: place the next bit of n - 1
                if k & 1:
                    res |= bit
                k >>= 1
            bit <<= 1
        return res
# test
assert Solution().minEnd(3, 4) == 6
assert Solution().minEnd(2, 7) == 15
assert Solution().minEnd(1, 5) == 5
assert Solution().minEnd(5, 1) == 9
def brute(n, x):                     # n-th smallest number containing all bits of x
    v, seen = x, 0
    while True:
        if v & x == x:
            seen += 1
            if seen == n:
                return v
        v += 1
assert all(Solution().minEnd(n, x) == brute(n, x) for n in range(1, 25) for x in range(1, 33))
```

```java
class Solution {
    public long minEnd(int n, int x) {
        long res = x, k = n - 1;
        for (long bit = 1; k > 0; bit <<= 1) {
            if ((x & bit) == 0) {
                if ((k & 1) == 1) res |= bit;
                k >>= 1;
            }
        }
        return res;
    }
}
// test
check(new Solution().minEnd(3, 4), 6L);
check(new Solution().minEnd(2, 7), 15L);
```

```go
func minEnd(n int, x int) int64 {
	res, k := int64(x), int64(n-1)
	for bit := int64(1); k > 0; bit <<= 1 {
		if int64(x)&bit == 0 { // free position: place the next bit of n - 1
			res |= (k & 1) * bit
			k >>= 1
		}
	}
	return res
}
// test
assertEq(minEnd(3, 4), 6)
assertEq(minEnd(2, 7), 15)
assertEq(minEnd(1, 5), 5)
assertEq(minEnd(5, 1), 9)
brute := func(n, x int) int { // n-th smallest number containing all bits of x
	for v, seen := x, 0; ; v++ {
		if v&x == x {
			seen++
			if seen == n {
				return v
			}
		}
	}
}
ok := true
for n := 1; n < 25; n++ {
	for x := 1; x < 33; x++ {
		ok = ok && minEnd(n, x) == int64(brute(n, x))
	}
}
assertEq(ok, true)
```
