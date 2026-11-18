// 85. Maximal Rectangle (Hard)
// https://leetcode.com/problems/maximal-rectangle/
// Run (Java 11+): java P0085MaximalRectangle.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0085MaximalRectangle {
    public static void main(String[] args) {
        char[][] m = Stream.of("10100", "10111", "11111", "10010").map(String::toCharArray).toArray(char[][]::new);
        check(new Solution().maximalRectangle(m), 6);
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
    public int maximalRectangle(char[][] matrix) {
        int cols = matrix[0].length, best = 0;
        int[] h = new int[cols];
        for (char[] row : matrix) {
            for (int c = 0; c < cols; c++) h[c] = row[c] == '1' ? h[c] + 1 : 0;
            Deque<Integer> st = new ArrayDeque<>();
            for (int i = 0; i <= cols; i++) {
                int x = i == cols ? 0 : h[i];
                while (!st.isEmpty() && h[st.peek()] >= x) {
                    int top = h[st.pop()];
                    best = Math.max(best, top * (i - (st.isEmpty() ? -1 : st.peek()) - 1));
                }
                st.push(i);
            }
        }
        return best;
    }
}
