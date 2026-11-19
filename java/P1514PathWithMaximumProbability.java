// 1514. Path with Maximum Probability (Medium)
// https://leetcode.com/problems/path-with-maximum-probability/
// Run (Java 11+): java P1514PathWithMaximumProbability.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1514PathWithMaximumProbability {
    public static void main(String[] args) {
        check(new Solution().maxProbability(3, new int[][]{{0, 1}, {1, 2}, {0, 2}}, new double[]{0.5, 0.5, 0.2}, 0, 2), 0.25);
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
    public double maxProbability(int n, int[][] edges, double[] succProb, int start, int end) {
        List<List<double[]>> g = new ArrayList<>();
        for (int i = 0; i < n; i++) g.add(new ArrayList<>());
        for (int i = 0; i < edges.length; i++) {
            g.get(edges[i][0]).add(new double[]{edges[i][1], succProb[i]});
            g.get(edges[i][1]).add(new double[]{edges[i][0], succProb[i]});
        }
        double[] best = new double[n];
        best[start] = 1;
        PriorityQueue<double[]> pq = new PriorityQueue<>((a, b) -> Double.compare(b[0], a[0]));
        pq.add(new double[]{1, start});
        while (!pq.isEmpty()) {
            double[] cur = pq.poll();
            int u = (int) cur[1];
            if (u == end) return cur[0];
            if (cur[0] < best[u]) continue;
            for (double[] e : g.get(u)) {
                int v = (int) e[0];
                if (cur[0] * e[1] > best[v]) { best[v] = cur[0] * e[1]; pq.add(new double[]{best[v], v}); }
            }
        }
        return 0;
    }
}
