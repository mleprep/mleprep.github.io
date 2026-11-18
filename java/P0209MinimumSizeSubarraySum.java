// 209. Minimum Size Subarray Sum (Medium)
// https://leetcode.com/problems/minimum-size-subarray-sum/
// Run (Java 11+): java P0209MinimumSizeSubarraySum.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0209MinimumSizeSubarraySum {
    public static void main(String[] args) {
        check(new Solution().minSubArrayLen(7, new int[]{2, 3, 1, 2, 4, 3}), 2);
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
    public int minSubArrayLen(int target, int[] nums) {
        int l = 0, s = 0, best = Integer.MAX_VALUE;
        for (int r = 0; r < nums.length; r++) {
            s += nums[r];
            while (s >= target) {
                best = Math.min(best, r - l + 1);
                s -= nums[l++];
            }
        }
        return best == Integer.MAX_VALUE ? 0 : best;
    }
}
