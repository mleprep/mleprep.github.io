// 983. Minimum Cost For Tickets (Medium)
// https://leetcode.com/problems/minimum-cost-for-tickets/
// Run (Java 11+): java P0983MinimumCostForTickets.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0983MinimumCostForTickets {
    public static void main(String[] args) {
        check(new Solution().mincostTickets(new int[]{1, 4, 6, 7, 8, 20}, new int[]{2, 7, 15}), 11);
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
    public int mincostTickets(int[] days, int[] costs) {
        int last = days[days.length - 1];
        boolean[] travel = new boolean[last + 1];
        for (int d : days) travel[d] = true;
        int[] dp = new int[last + 1];
        for (int d = 1; d <= last; d++) {
            if (!travel[d]) { dp[d] = dp[d - 1]; continue; }
            dp[d] = Math.min(dp[d - 1] + costs[0],
                    Math.min(dp[Math.max(0, d - 7)] + costs[1], dp[Math.max(0, d - 30)] + costs[2]));
        }
        return dp[last];
    }
}
