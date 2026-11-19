// 886. Possible Bipartition (Medium)
// https://leetcode.com/problems/possible-bipartition/
// Run (Java 11+): java P0886PossibleBipartition.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0886PossibleBipartition {
    public static void main(String[] args) {
        check(new Solution().possibleBipartition(3, new int[][]{{1, 2}, {1, 3}, {2, 3}}), false);
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
    public boolean possibleBipartition(int n, int[][] dislikes) {
        List<List<Integer>> g = new ArrayList<>();
        for (int i = 0; i <= n; i++) g.add(new ArrayList<>());
        for (int[] d : dislikes) { g.get(d[0]).add(d[1]); g.get(d[1]).add(d[0]); }
        int[] color = new int[n + 1];
        Arrays.fill(color, -1);
        for (int s = 1; s <= n; s++) {
            if (color[s] != -1) continue;
            color[s] = 0;
            Deque<Integer> q = new ArrayDeque<>(List.of(s));
            while (!q.isEmpty()) {
                int u = q.poll();
                for (int v : g.get(u)) {
                    if (color[v] == -1) { color[v] = 1 - color[u]; q.add(v); }
                    else if (color[v] == color[u]) return false;
                }
            }
        }
        return true;
    }
}
