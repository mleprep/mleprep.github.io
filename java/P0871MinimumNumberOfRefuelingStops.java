// 871. Minimum Number of Refueling Stops (Hard)
// https://leetcode.com/problems/minimum-number-of-refueling-stops/
// Run (Java 11+): java P0871MinimumNumberOfRefuelingStops.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0871MinimumNumberOfRefuelingStops {
    public static void main(String[] args) {
        check(new Solution().minRefuelStops(100, 10, new int[][]{{10, 60}, {20, 30}, {30, 30}, {60, 40}}), 2);
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
    public int minRefuelStops(int target, int startFuel, int[][] stations) {
        PriorityQueue<Integer> pq = new PriorityQueue<>(Collections.reverseOrder());
        long reach = startFuel;
        int i = 0, stops = 0;
        while (reach < target) {
            while (i < stations.length && stations[i][0] <= reach) pq.add(stations[i++][1]);
            if (pq.isEmpty()) return -1;
            reach += pq.poll();
            stops++;
        }
        return stops;
    }
}
