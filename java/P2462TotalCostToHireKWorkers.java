// 2462. Total Cost to Hire K Workers (Medium)
// https://leetcode.com/problems/total-cost-to-hire-k-workers/
// Run (Java 11+): java P2462TotalCostToHireKWorkers.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P2462TotalCostToHireKWorkers {
    public static void main(String[] args) {
        check(new Solution().totalCost(new int[]{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4), 11L);
        check(new Solution().totalCost(new int[]{1, 2, 4, 1}, 3, 3), 4L);
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
    public long totalCost(int[] costs, int k, int candidates) {
        PriorityQueue<Integer> front = new PriorityQueue<>(), back = new PriorityQueue<>();
        int i = 0, j = costs.length - 1;
        long total = 0;
        while (k-- > 0) {
            while (front.size() < candidates && i <= j) front.add(costs[i++]);
            while (back.size() < candidates && i <= j) back.add(costs[j--]);
            if (back.isEmpty() || (!front.isEmpty() && front.peek() <= back.peek())) total += front.poll();
            else total += back.poll();
        }
        return total;
    }
}
