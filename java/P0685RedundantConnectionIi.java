// 685. Redundant Connection II (Hard)
// https://leetcode.com/problems/redundant-connection-ii/
// Run (Java 11+): java P0685RedundantConnectionIi.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0685RedundantConnectionIi {
    public static void main(String[] args) {
        check(new Solution().findRedundantDirectedConnection(new int[][]{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}}), new int[]{4, 1});
        check(new Solution().findRedundantDirectedConnection(new int[][]{{2, 1}, {3, 1}, {4, 2}, {1, 4}}), new int[]{2, 1});
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
    private int[] root;

    private int find(int x) {
        while (root[x] != x) { root[x] = root[root[x]]; x = root[x]; }
        return x;
    }

    public int[] findRedundantDirectedConnection(int[][] edges) {
        int n = edges.length;
        int[] parentOf = new int[n + 1];
        int[] cand1 = null, cand2 = null;
        for (int[] e : edges) {
            if (parentOf[e[1]] != 0) { cand1 = new int[]{parentOf[e[1]], e[1]}; cand2 = e; }
            else parentOf[e[1]] = e[0];
        }
        root = new int[n + 1];
        for (int i = 0; i <= n; i++) root[i] = i;
        for (int[] e : edges) {
            if (e == cand2) continue;
            int ru = find(e[0]), rv = find(e[1]);
            if (ru == rv) return cand1 != null ? cand1 : e;
            root[rv] = ru;
        }
        return cand2;
    }
}
