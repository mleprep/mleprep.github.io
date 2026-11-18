// 494. Target Sum (Medium)
// https://leetcode.com/problems/target-sum/
// Run (Java 11+): java P0494TargetSum.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0494TargetSum {
    public static void main(String[] args) {
        check(new Solution().findTargetSumWays(new int[]{1, 1, 1, 1, 1}, 3), 5);
        check(new Solution().findTargetSumWays(new int[]{0, 0, 1}, 1), 4);
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
    public int findTargetSumWays(int[] nums, int target) {
        int total = 0;
        for (int x : nums) total += x;
        if (Math.abs(target) > total || (total + target) % 2 != 0) return 0;
        int P = (total + target) / 2;
        int[] dp = new int[P + 1];
        dp[0] = 1;
        for (int x : nums)
            for (int s = P; s >= x; s--) dp[s] += dp[s - x];
        return dp[P];
    }
}
