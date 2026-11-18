// 1851. Minimum Interval to Include Each Query (Hard)
// https://leetcode.com/problems/minimum-interval-to-include-each-query/
// Run (Java 11+): java P1851MinimumIntervalToIncludeEachQuery.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1851MinimumIntervalToIncludeEachQuery {
    public static void main(String[] args) {
        check(new Solution().minInterval(new int[][]{{1, 4}, {2, 4}, {3, 6}, {4, 4}}, new int[]{2, 3, 4, 5}), new int[]{3, 3, 1, 4});
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
    public int[] minInterval(int[][] intervals, int[] queries) {
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));
        Integer[] qi = new Integer[queries.length];
        for (int k = 0; k < qi.length; k++) qi[k] = k;
        Arrays.sort(qi, (a, b) -> Integer.compare(queries[a], queries[b]));
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[0] - b[0]); // {size, end}
        int[] res = new int[queries.length];
        int i = 0;
        for (int k : qi) {
            int q = queries[k];
            while (i < intervals.length && intervals[i][0] <= q) {
                pq.add(new int[]{intervals[i][1] - intervals[i][0] + 1, intervals[i][1]});
                i++;
            }
            while (!pq.isEmpty() && pq.peek()[1] < q) pq.poll();
            res[k] = pq.isEmpty() ? -1 : pq.peek()[0];
        }
        return res;
    }
}
