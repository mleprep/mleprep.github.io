// 1489. Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree (Hard)
// https://leetcode.com/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree/
// Run (Java 11+): java P1489FindCriticalAndPseudoCriticalEdgesInMinimumSpanningTree.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1489FindCriticalAndPseudoCriticalEdgesInMinimumSpanningTree {
    public static void main(String[] args) {
        check(new Solution().findCriticalAndPseudoCriticalEdges(5, new int[][]{{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {0, 3, 2}, {0, 4, 3}, {3, 4, 3}, {1, 4, 6}}),
              List.of(List.of(0, 1), List.of(2, 3, 4, 5)));
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
    private int n;
    private int[][] edges;
    private Integer[] order;

    public List<List<Integer>> findCriticalAndPseudoCriticalEdges(int n, int[][] edges) {
        this.n = n;
        this.edges = edges;
        order = new Integer[edges.length];
        for (int i = 0; i < edges.length; i++) order[i] = i;
        Arrays.sort(order, (x, y) -> edges[x][2] - edges[y][2]);
        int base = mst(-1, -1);
        List<Integer> critical = new ArrayList<>(), pseudo = new ArrayList<>();
        for (int i = 0; i < edges.length; i++) {
            if (mst(i, -1) > base) critical.add(i);
            else if (mst(-1, i) == base) pseudo.add(i);
        }
        return List.of(critical, pseudo);
    }

    private int mst(int skip, int force) {
        int[] parent = new int[n];
        for (int i = 0; i < n; i++) parent[i] = i;
        int weight = 0, used = 0;
        if (force >= 0) { union(parent, edges[force][0], edges[force][1]); weight = edges[force][2]; used = 1; }
        for (int i : order) if (i != skip && union(parent, edges[i][0], edges[i][1])) { weight += edges[i][2]; used++; }
        return used == n - 1 ? weight : Integer.MAX_VALUE;
    }

    private boolean union(int[] p, int a, int b) {
        int ra = find(p, a), rb = find(p, b);
        if (ra == rb) return false;
        p[ra] = rb;
        return true;
    }

    private int find(int[] p, int x) {
        while (p[x] != x) x = p[x] = p[p[x]];
        return x;
    }
}
