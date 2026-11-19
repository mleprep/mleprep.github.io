// 1971. Find if Path Exists in Graph (Easy)
// https://leetcode.com/problems/find-if-path-exists-in-graph/
// Run (Java 11+): java P1971FindIfPathExistsInGraph.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1971FindIfPathExistsInGraph {
    public static void main(String[] args) {
        check(new Solution().validPath(6, new int[][]{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5), false);
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
    public boolean validPath(int n, int[][] edges, int source, int destination) {
        int[] p = new int[n];
        for (int i = 0; i < n; i++) p[i] = i;
        for (int[] e : edges) p[find(p, e[0])] = find(p, e[1]);
        return find(p, source) == find(p, destination);
    }

    private int find(int[] p, int x) {
        while (p[x] != x) { p[x] = p[p[x]]; x = p[x]; }
        return x;
    }
}
