// 416. Partition Equal Subset Sum (Medium)
// https://leetcode.com/problems/partition-equal-subset-sum/
// Run (Java 11+): java P0416PartitionEqualSubsetSum.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0416PartitionEqualSubsetSum {
    public static void main(String[] args) {
        check(new Solution().canPartition(new int[]{1, 5, 11, 5}), true);
        check(new Solution().canPartition(new int[]{1, 2, 3, 5}), false);
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
    public boolean canPartition(int[] nums) {
        int total = 0;
        for (int x : nums) total += x;
        if (total % 2 == 1) return false;
        int target = total / 2;
        boolean[] dp = new boolean[target + 1];
        dp[0] = true;
        for (int x : nums)
            for (int s = target; s >= x; s--) dp[s] |= dp[s - x];
        return dp[target];
    }
}
