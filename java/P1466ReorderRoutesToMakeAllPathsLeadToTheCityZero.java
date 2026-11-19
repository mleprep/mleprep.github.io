// 1466. Reorder Routes to Make All Paths Lead to the City Zero (Medium)
// https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/
// Run (Java 11+): java P1466ReorderRoutesToMakeAllPathsLeadToTheCityZero.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1466ReorderRoutesToMakeAllPathsLeadToTheCityZero {
    public static void main(String[] args) {
        check(new Solution().minReorder(6, new int[][]{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}), 3);
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
    public int minReorder(int n, int[][] connections) {
        List<List<int[]>> g = new ArrayList<>();
        for (int i = 0; i < n; i++) g.add(new ArrayList<>());
        for (int[] c : connections) { g.get(c[0]).add(new int[]{c[1], 1}); g.get(c[1]).add(new int[]{c[0], 0}); }
        boolean[] seen = new boolean[n];
        seen[0] = true;
        Deque<Integer> st = new ArrayDeque<>(List.of(0));
        int flips = 0;
        while (!st.isEmpty())
            for (int[] e : g.get(st.pop()))
                if (!seen[e[0]]) { seen[e[0]] = true; flips += e[1]; st.push(e[0]); }
        return flips;
    }
}
