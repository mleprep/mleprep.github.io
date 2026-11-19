// 188. Best Time to Buy and Sell Stock IV (Hard)
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
// Run (Java 11+): java P0188BestTimeToBuyAndSellStockIv.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0188BestTimeToBuyAndSellStockIv {
    public static void main(String[] args) {
        check(new Solution().maxProfit(2, new int[]{3, 2, 6, 5, 0, 3}), 7);
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
    public int maxProfit(int k, int[] prices) {
        int n = prices.length;
        if (k >= n / 2) {
            int s = 0;
            for (int i = 1; i < n; i++) s += Math.max(0, prices[i] - prices[i - 1]);
            return s;
        }
        int[] buy = new int[k + 1], sell = new int[k + 1];
        Arrays.fill(buy, Integer.MIN_VALUE / 2);
        for (int p : prices)
            for (int j = 1; j <= k; j++) {
                buy[j] = Math.max(buy[j], sell[j - 1] - p);
                sell[j] = Math.max(sell[j], buy[j] + p);
            }
        return sell[k];
    }
}
