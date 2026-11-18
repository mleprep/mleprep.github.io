// 309. Best Time to Buy and Sell Stock with Cooldown (Medium)
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
// Run (Java 11+): java P0309BestTimeToBuyAndSellStockWithCooldown.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0309BestTimeToBuyAndSellStockWithCooldown {
    public static void main(String[] args) {
        check(new Solution().maxProfit(new int[]{1, 2, 3, 0, 2}), 3);
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
        int hold = Integer.MIN_VALUE / 2, sold = 0, rest = 0;
        for (int p : prices) {
            int h = Math.max(hold, rest - p), s = hold + p, r = Math.max(rest, sold);
            hold = h; sold = s; rest = r;
        }
        return Math.max(sold, rest);
    }
}
