// 1976. Number of Ways to Arrive at Destination (Medium)
// https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/
// Run (Java 11+): java P1976NumberOfWaysToArriveAtDestination.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1976NumberOfWaysToArriveAtDestination {
    public static void main(String[] args) {
        int[][] R = {{0, 6, 7}, {0, 1, 2}, {1, 2, 3}, {1, 3, 3}, {6, 3, 3}, {3, 5, 1}, {6, 5, 1}, {2, 5, 1}, {0, 4, 5}, {4, 6, 2}};
        check(new Solution().countPaths(7, R), 4);
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
    public int countPaths(int n, int[][] roads) {
        final int MOD = 1_000_000_007;
        List<List<int[]>> g = new ArrayList<>();
        for (int i = 0; i < n; i++) g.add(new ArrayList<>());
        for (int[] r : roads) { g.get(r[0]).add(new int[]{r[1], r[2]}); g.get(r[1]).add(new int[]{r[0], r[2]}); }
        long[] dist = new long[n];
        Arrays.fill(dist, Long.MAX_VALUE);
        long[] ways = new long[n];
        dist[0] = 0; ways[0] = 1;
        PriorityQueue<long[]> pq = new PriorityQueue<>((a, b) -> Long.compare(a[0], b[0]));
        pq.add(new long[]{0, 0});
        while (!pq.isEmpty()) {
            long[] cur = pq.poll();
            int u = (int) cur[1];
            if (cur[0] > dist[u]) continue;
            for (int[] e : g.get(u)) {
                long nd = cur[0] + e[1];
                if (nd < dist[e[0]]) { dist[e[0]] = nd; ways[e[0]] = ways[u]; pq.add(new long[]{nd, e[0]}); }
                else if (nd == dist[e[0]]) ways[e[0]] = (ways[e[0]] + ways[u]) % MOD;
            }
        }
        return (int) ways[n - 1];
    }
}
