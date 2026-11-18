// 743. Network Delay Time (Medium)
// https://leetcode.com/problems/network-delay-time/
// Run (Java 11+): java P0743NetworkDelayTime.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0743NetworkDelayTime {
    public static void main(String[] args) {
        check(new Solution().networkDelayTime(new int[][]{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2), 2);
        check(new Solution().networkDelayTime(new int[][]{{1, 2, 1}}, 2, 2), -1);
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
    public int networkDelayTime(int[][] times, int n, int k) {
        List<List<int[]>> g = new ArrayList<>();
        for (int i = 0; i <= n; i++) g.add(new ArrayList<>());
        for (int[] t : times) g.get(t[0]).add(new int[]{t[1], t[2]});
        int[] dist = new int[n + 1];
        Arrays.fill(dist, Integer.MAX_VALUE);
        dist[k] = 0;
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[0] - b[0]); // {dist, node}
        pq.add(new int[]{0, k});
        while (!pq.isEmpty()) {
            int[] cur = pq.poll();
            int d = cur[0], u = cur[1];
            if (d > dist[u]) continue;
            for (int[] e : g.get(u))
                if (d + e[1] < dist[e[0]]) { dist[e[0]] = d + e[1]; pq.add(new int[]{dist[e[0]], e[0]}); }
        }
        int best = 0;
        for (int i = 1; i <= n; i++) best = Math.max(best, dist[i]);
        return best == Integer.MAX_VALUE ? -1 : best;
    }
}
