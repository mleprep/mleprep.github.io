// 52. N-Queens II (Hard)
// https://leetcode.com/problems/n-queens-ii/
// Run (Java 11+): java P0052NQueensIi.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0052NQueensIi {
    public static void main(String[] args) {
        check(new Solution().totalNQueens(8), 92);
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
    private int full;

    public int totalNQueens(int n) {
        full = (1 << n) - 1;
        return bt(0, 0, 0);
    }

    private int bt(int cols, int d1, int d2) {
        if (cols == full) return 1;
        int count = 0;
        for (int free = ~(cols | d1 | d2) & full; free != 0; ) {
            int bit = free & -free;
            free -= bit;
            count += bt(cols | bit, ((d1 | bit) << 1) & full, (d2 | bit) >> 1);
        }
        return count;
    }
}
