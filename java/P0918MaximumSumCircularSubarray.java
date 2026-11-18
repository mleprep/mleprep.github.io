// 918. Maximum Sum Circular Subarray (Medium)
// https://leetcode.com/problems/maximum-sum-circular-subarray/
// Run (Java 11+): java P0918MaximumSumCircularSubarray.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0918MaximumSumCircularSubarray {
    public static void main(String[] args) {
        check(new Solution().maxSubarraySumCircular(new int[]{5, -3, 5}), 10);
        check(new Solution().maxSubarraySumCircular(new int[]{-3, -2, -3}), -2);
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
    public int maxSubarraySumCircular(int[] nums) {
        int curMax = 0, curMin = 0, bestMax = nums[0], bestMin = nums[0], total = 0;
        for (int x : nums) {
            curMax = Math.max(x, curMax + x);
            bestMax = Math.max(bestMax, curMax);
            curMin = Math.min(x, curMin + x);
            bestMin = Math.min(bestMin, curMin);
            total += x;
        }
        return bestMax < 0 ? bestMax : Math.max(bestMax, total - bestMin);
    }
}
