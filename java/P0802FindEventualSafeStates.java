// 802. Find Eventual Safe States (Medium)
// https://leetcode.com/problems/find-eventual-safe-states/
// Run (Java 11+): java P0802FindEventualSafeStates.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0802FindEventualSafeStates {
    public static void main(String[] args) {
        check(new Solution().eventualSafeNodes(new int[][]{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}), "[2, 4, 5, 6]");
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
    public List<Integer> eventualSafeNodes(int[][] graph) {
        int[] color = new int[graph.length];
        List<Integer> res = new ArrayList<>();
        for (int u = 0; u < graph.length; u++) if (safe(graph, u, color)) res.add(u);
        return res;
    }

    private boolean safe(int[][] g, int u, int[] color) {
        if (color[u] != 0) return color[u] == 2;
        color[u] = 1;
        for (int v : g[u]) if (!safe(g, v, color)) return false;
        color[u] = 2;
        return true;
    }
}
