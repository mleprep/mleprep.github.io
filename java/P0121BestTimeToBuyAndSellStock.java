// 121. Best Time to Buy and Sell Stock (Easy)
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
// Run (Java 11+): java P0121BestTimeToBuyAndSellStock.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0121BestTimeToBuyAndSellStock {
    public static void main(String[] args) {
        check(new Solution().maxProfit(new int[]{7, 1, 5, 3, 6, 4}), 5);
        check(new Solution().maxProfit(new int[]{7, 6, 4, 3, 1}), 0);
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
    public int maxProfit(int[] prices) {
        int lo = Integer.MAX_VALUE, best = 0;
        for (int p : prices) {
            lo = Math.min(lo, p);
            best = Math.max(best, p - lo);
        }
        return best;
    }
}
