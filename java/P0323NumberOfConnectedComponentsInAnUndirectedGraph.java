// 323. Number of Connected Components in an Undirected Graph (Medium)
// https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/
// Run (Java 11+): java P0323NumberOfConnectedComponentsInAnUndirectedGraph.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0323NumberOfConnectedComponentsInAnUndirectedGraph {
    public static void main(String[] args) {
        check(new Solution().countComponents(5, new int[][]{{0, 1}, {1, 2}, {3, 4}}), 2);
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
    public int countComponents(int n, int[][] edges) {
        int[] parent = new int[n];
        for (int i = 0; i < n; i++) parent[i] = i;
        int comps = n;
        for (int[] e : edges) {
            int ra = find(parent, e[0]), rb = find(parent, e[1]);
            if (ra != rb) { parent[ra] = rb; comps--; }
        }
        return comps;
    }

    private int find(int[] p, int x) {
        while (p[x] != x) { p[x] = p[p[x]]; x = p[x]; }
        return x;
    }
}
