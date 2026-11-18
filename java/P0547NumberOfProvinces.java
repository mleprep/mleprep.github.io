// 547. Number of Provinces (Medium)
// https://leetcode.com/problems/number-of-provinces/
// Run (Java 11+): java P0547NumberOfProvinces.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0547NumberOfProvinces {
    public static void main(String[] args) {
        check(new Solution().findCircleNum(new int[][]{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}), 2);
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
    public int findCircleNum(int[][] isConnected) {
        int n = isConnected.length, count = 0;
        boolean[] seen = new boolean[n];
        for (int u = 0; u < n; u++)
            if (!seen[u]) { count++; dfs(isConnected, u, seen); }
        return count;
    }

    private void dfs(int[][] g, int u, boolean[] seen) {
        seen[u] = true;
        for (int v = 0; v < g.length; v++) if (g[u][v] == 1 && !seen[v]) dfs(g, v, seen);
    }
}
