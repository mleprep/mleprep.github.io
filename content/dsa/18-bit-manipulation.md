# Bit Manipulation

| Trick | Meaning |
|---|---|
| `x & (x - 1)` | clear the lowest set bit (count bits, power-of-two test) |
| `x & -x` | isolate the lowest set bit |
| `a ^ a = 0`, `a ^ 0 = a` | XOR cancels pairs → find the unique element |
| `(x >> i) & 1` | read bit i |
| `x | (1 << i)`, `x & ~(1 << i)` | set / clear bit i |

**Python caveat:** ints are unbounded — emulate 32-bit with `& 0xFFFFFFFF` and convert back with `x if x < 2**31 else x - 2**32`.

```viz
x       = 1 0 1 1 0 0
x - 1   = 1 0 1 0 1 1
x&(x-1) = 1 0 1 0 0 0      lowest 1 removed
```

## 136. Single Number | Easy
related: 137, 268, 260
tags: xor

XOR everything: pairs cancel, the single number remains.

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        return reduce(lambda a, b: a ^ b, nums)
# test
assert Solution().singleNumber([4, 1, 2, 1, 2]) == 4
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

## 137. Single Number II | Medium
related: 136
tags: bit-counting

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

## 191. Number of 1 Bits | Easy
related: 338, 190
tags: blind75, brian-kernighan

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

## 338. Counting Bits | Easy
related: 191
tags: blind75, dp

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

## 190. Reverse Bits | Easy
related: 191, 7
tags: blind75

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

## 268. Missing Number | Easy
related: 41, 136, 287
tags: blind75, xor

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

## 371. Sum of Two Integers | Medium
related: 67, 2
tags: blind75, adder

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

## 7. Reverse Integer | Medium
related: 190, 8
tags: blind75, overflow

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


## 231. Power of Two | Easy
related: 191, 338
tags: bit-trick

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

## 461. Hamming Distance | Easy
related: 191
tags: xor, popcount

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

## 260. Single Number III | Medium
related: 136, 137
tags: xor, partition

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
assert Solution().singleNumber([1, 2, 1, 3, 2, 5]) == [3, 5]
assert Solution().singleNumber([-1, 0]) == [-1, 0]
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

## 201. Bitwise AND of Numbers Range | Medium
related: 191
tags: common-prefix

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

## 89. Gray Code | Medium
related: 784
tags: bit-trick

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
g = Solution().grayCode(3)
assert g == [0, 1, 3, 2, 6, 7, 5, 4]
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
