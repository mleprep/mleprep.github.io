// 974. Subarray Sums Divisible by K (Medium)
// https://leetcode.com/problems/subarray-sums-divisible-by-k/
// Run (Java 11+): java P0974SubarraySumsDivisibleByK.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0974SubarraySumsDivisibleByK {
    public static void main(String[] args) {
        check(new Solution().subarraysDivByK(new int[]{4, 5, 0, -2, -3, 1}, 5), 7);
        if (fails > 0) { System.out.println(fails + " test(s) FAILED"); System.exit(1); }
        System.out.println("OK");
    }

    static int fails = 0;

    static void check(Object got, Object expected) {
        String g = str(got), e = str(expected);
        if (!g.equals(e)) { fails++; System.out.println("FAIL: got " + g + ", expected " + e); }
    }

    static String str(Object o) {
        if (o == null) return "null";
        if (o instanceof int[]) return Arrays.toString((int[]) o);
        if (o instanceof long[]) return Arrays.toString((long[]) o);
        if (o instanceof double[]) return Arrays.toString((double[]) o);
        if (o instanceof char[]) return Arrays.toString((char[]) o);
        if (o instanceof boolean[]) return Arrays.toString((boolean[]) o);
        if (o instanceof Object[]) return Arrays.deepToString((Object[]) o);
        return String.valueOf(o);
    }
}

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
