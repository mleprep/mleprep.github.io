// 2392. Build a Matrix With Conditions (Hard)
// https://leetcode.com/problems/build-a-matrix-with-conditions/
// Run (Java 11+): java P2392BuildAMatrixWithConditions.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P2392BuildAMatrixWithConditions {
    public static void main(String[] args) {
        int[][] bm = new Solution().buildMatrix(3, new int[][]{{1, 2}, {3, 2}}, new int[][]{{2, 1}, {3, 2}});
        int[][] at = new int[4][];
        for (int i = 0; i < 3; i++) for (int j = 0; j < 3; j++) if (bm[i][j] > 0) at[bm[i][j]] = new int[]{i, j};
        check(at[1][0] < at[2][0] && at[3][0] < at[2][0] && at[2][1] < at[1][1] && at[3][1] < at[2][1], true);
        check(new Solution().buildMatrix(3, new int[][]{{1, 2}, {2, 3}, {3, 1}}, new int[][]{{2, 1}}).length, 0);
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
    public int[][] buildMatrix(int k, int[][] rowConditions, int[][] colConditions) {
        int[] row = topo(k, rowConditions), col = topo(k, colConditions);
        if (row == null || col == null) return new int[0][0];
        int[][] m = new int[k][k];
        for (int v = 1; v <= k; v++) m[row[v]][col[v]] = v;
        return m;
    }

    private int[] topo(int k, int[][] conds) {                 // position of each value, or null on a cycle
        List<List<Integer>> adj = new ArrayList<>();
        for (int i = 0; i <= k; i++) adj.add(new ArrayList<>());
        int[] indeg = new int[k + 1], pos = new int[k + 1];
        for (int[] c : conds) { adj.get(c[0]).add(c[1]); indeg[c[1]]++; }
        Deque<Integer> q = new ArrayDeque<>();
        for (int v = 1; v <= k; v++) if (indeg[v] == 0) q.add(v);
        int idx = 0;
        while (!q.isEmpty()) {
            int v = q.poll();
            pos[v] = idx++;
            for (int w : adj.get(v)) if (--indeg[w] == 0) q.add(w);
        }
        return idx == k ? pos : null;
    }
}
