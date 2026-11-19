// 96. Unique Binary Search Trees (Medium)
// https://leetcode.com/problems/unique-binary-search-trees/
// Run (Java 11+): java P0096UniqueBinarySearchTrees.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0096UniqueBinarySearchTrees {
    public static void main(String[] args) {
        check(new Solution().numTrees(3), 5);
        check(new Solution().numTrees(19), 1767263190);
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
    public int numTrees(int n) {
        long[] G = new long[n + 1];
        G[0] = 1;
        for (int m = 1; m <= n; m++)
            for (int r = 1; r <= m; r++) G[m] += G[r - 1] * G[m - r];
        return (int) G[n];
    }
}
