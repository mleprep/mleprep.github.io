// 1192. Critical Connections in a Network (Hard)
// https://leetcode.com/problems/critical-connections-in-a-network/
// Run (Java 11+): java P1192CriticalConnectionsInANetwork.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1192CriticalConnectionsInANetwork {
    public static void main(String[] args) {
        check(new Solution().criticalConnections(4, List.of(List.of(0, 1), List.of(1, 2), List.of(2, 0), List.of(1, 3))), "[[1, 3]]");
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
    private int timer = 0;
    private int[] tin, low;
    private List<List<Integer>> g;
    private final List<List<Integer>> res = new ArrayList<>();

    public List<List<Integer>> criticalConnections(int n, List<List<Integer>> connections) {
        g = new ArrayList<>();
        for (int i = 0; i < n; i++) g.add(new ArrayList<>());
        for (List<Integer> e : connections) { g.get(e.get(0)).add(e.get(1)); g.get(e.get(1)).add(e.get(0)); }
        tin = new int[n];
        low = new int[n];
        Arrays.fill(tin, -1);
        dfs(0, -1);
        return res;
    }

    private void dfs(int u, int parent) {
        tin[u] = low[u] = timer++;
        for (int v : g.get(u)) {
            if (v == parent) continue;
            if (tin[v] == -1) {
                dfs(v, u);
                low[u] = Math.min(low[u], low[v]);
                if (low[v] > tin[u]) res.add(List.of(u, v));
            } else low[u] = Math.min(low[u], tin[v]);
        }
    }
}
