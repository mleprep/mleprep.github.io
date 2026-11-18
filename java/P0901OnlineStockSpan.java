// 901. Online Stock Span (Medium)
// https://leetcode.com/problems/online-stock-span/
// Run (Java 11+): java P0901OnlineStockSpan.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0901OnlineStockSpan {
    public static void main(String[] args) {
        StockSpanner sp = new StockSpanner();
        int[] out = new int[7], prices = {100, 80, 60, 70, 60, 75, 85};
        for (int i = 0; i < 7; i++) out[i] = sp.next(prices[i]);
        check(out, new int[]{1, 1, 1, 2, 1, 4, 6});
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

class StockSpanner {
    private final Deque<int[]> st = new ArrayDeque<>(); // {price, span}

    public int next(int price) {
        int span = 1;
        while (!st.isEmpty() && st.peek()[0] <= price) span += st.pop()[1];
        st.push(new int[]{price, span});
        return span;
    }
}
