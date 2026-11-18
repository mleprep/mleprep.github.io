// 787. Cheapest Flights Within K Stops (Medium)
// https://leetcode.com/problems/cheapest-flights-within-k-stops/
// Run (Java 11+): java P0787CheapestFlightsWithinKStops.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0787CheapestFlightsWithinKStops {
    public static void main(String[] args) {
        int[][] F = {{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}};
        check(new Solution().findCheapestPrice(4, F, 0, 3, 1), 700);
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
    public int findCheapestPrice(int n, int[][] flights, int src, int dst, int k) {
        int INF = Integer.MAX_VALUE;
        int[] dist = new int[n];
        Arrays.fill(dist, INF);
        dist[src] = 0;
        for (int i = 0; i <= k; i++) {
            int[] nxt = dist.clone();
            for (int[] f : flights)
                if (dist[f[0]] != INF && dist[f[0]] + f[2] < nxt[f[1]]) nxt[f[1]] = dist[f[0]] + f[2];
            dist = nxt;
        }
        return dist[dst] == INF ? -1 : dist[dst];
    }
}
