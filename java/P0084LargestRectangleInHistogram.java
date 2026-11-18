// 84. Largest Rectangle in Histogram (Hard)
// https://leetcode.com/problems/largest-rectangle-in-histogram/
// Run (Java 11+): java P0084LargestRectangleInHistogram.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0084LargestRectangleInHistogram {
    public static void main(String[] args) {
        check(new Solution().largestRectangleArea(new int[]{2, 1, 5, 6, 2, 3}), 10);
        check(new Solution().largestRectangleArea(new int[]{2, 4}), 4);
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
    public int largestRectangleArea(int[] heights) {
        Deque<Integer> st = new ArrayDeque<>();
        int best = 0, n = heights.length;
        for (int i = 0; i <= n; i++) {
            int h = i == n ? 0 : heights[i];
            while (!st.isEmpty() && heights[st.peek()] >= h) {
                int top = heights[st.pop()];
                int left = st.isEmpty() ? -1 : st.peek();
                best = Math.max(best, top * (i - left - 1));
            }
            st.push(i);
        }
        return best;
    }
}
