// 643. Maximum Average Subarray I (Easy)
// https://leetcode.com/problems/maximum-average-subarray-i/
// Run (Java 11+): java P0643MaximumAverageSubarrayI.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0643MaximumAverageSubarrayI {
    public static void main(String[] args) {
        check(new Solution().findMaxAverage(new int[]{1, 12, -5, -6, 50, 3}, 4), 12.75);
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
    public double findMaxAverage(int[] nums, int k) {
        int s = 0;
        for (int i = 0; i < k; i++) s += nums[i];
        int best = s;
        for (int i = k; i < nums.length; i++) { s += nums[i] - nums[i - k]; best = Math.max(best, s); }
        return (double) best / k;
    }
}
