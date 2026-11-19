// 1235. Maximum Profit in Job Scheduling (Hard)
// https://leetcode.com/problems/maximum-profit-in-job-scheduling/
// Run (Java 11+): java P1235MaximumProfitInJobScheduling.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1235MaximumProfitInJobScheduling {
    public static void main(String[] args) {
        check(new Solution().jobScheduling(new int[]{1, 2, 3, 3}, new int[]{3, 4, 5, 6}, new int[]{50, 10, 40, 70}), 120);
        check(new Solution().jobScheduling(new int[]{1, 2, 3, 4, 6}, new int[]{3, 5, 10, 6, 9}, new int[]{20, 20, 100, 70, 60}), 150);
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
    public int jobScheduling(int[] startTime, int[] endTime, int[] profit) {
        int n = startTime.length;
        Integer[] idx = new Integer[n];
        for (int i = 0; i < n; i++) idx[i] = i;
        Arrays.sort(idx, (a, b) -> endTime[a] - endTime[b]);
        int[] dp = new int[n + 1];
        for (int i = 1; i <= n; i++) {
            int job = idx[i - 1], s = startTime[job];
            int lo = 0, hi = i - 1;                     // count of jobs among first i-1 ending <= s
            while (lo < hi) {
                int mid = (lo + hi) / 2;
                if (endTime[idx[mid]] <= s) lo = mid + 1; else hi = mid;
            }
            dp[i] = Math.max(dp[i - 1], dp[lo] + profit[job]);
        }
        return dp[n];
    }
}
