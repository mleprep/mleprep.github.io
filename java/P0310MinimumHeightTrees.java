// 310. Minimum Height Trees (Medium)
// https://leetcode.com/problems/minimum-height-trees/
// Run (Java 11+): java P0310MinimumHeightTrees.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0310MinimumHeightTrees {
    public static void main(String[] args) {
        check(new Solution().findMinHeightTrees(6, new int[][]{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}), "[3, 4]");
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
    public List<Integer> findMinHeightTrees(int n, int[][] edges) {
        if (n <= 2) { List<Integer> r = new ArrayList<>(); for (int i = 0; i < n; i++) r.add(i); return r; }
        List<Set<Integer>> g = new ArrayList<>();
        for (int i = 0; i < n; i++) g.add(new HashSet<>());
        for (int[] e : edges) { g.get(e[0]).add(e[1]); g.get(e[1]).add(e[0]); }
        List<Integer> leaves = new ArrayList<>();
        for (int i = 0; i < n; i++) if (g.get(i).size() == 1) leaves.add(i);
        int remaining = n;
        while (remaining > 2) {
            remaining -= leaves.size();
            List<Integer> nxt = new ArrayList<>();
            for (int leaf : leaves) {
                int nb = g.get(leaf).iterator().next();
                g.get(nb).remove(leaf);
                if (g.get(nb).size() == 1) nxt.add(nb);
            }
            leaves = nxt;
        }
        Collections.sort(leaves);
        return leaves;
    }
}
