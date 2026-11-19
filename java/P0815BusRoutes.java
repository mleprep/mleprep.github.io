// 815. Bus Routes (Hard)
// https://leetcode.com/problems/bus-routes/
// Run (Java 11+): java P0815BusRoutes.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0815BusRoutes {
    public static void main(String[] args) {
        check(new Solution().numBusesToDestination(new int[][]{{1, 2, 7}, {3, 6, 7}}, 1, 6), 2);
        check(new Solution().numBusesToDestination(new int[][]{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, 15, 12), -1);
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
    public int numBusesToDestination(int[][] routes, int source, int target) {
        if (source == target) return 0;
        Map<Integer, List<Integer>> at = new HashMap<>();
        for (int i = 0; i < routes.length; i++) for (int s : routes[i]) at.computeIfAbsent(s, k -> new ArrayList<>()).add(i);
        boolean[] used = new boolean[routes.length];
        Set<Integer> seen = new HashSet<>(List.of(source));
        Deque<Integer> q = new ArrayDeque<>(List.of(source));
        for (int buses = 1; !q.isEmpty(); buses++) {
            for (int i = q.size(); i > 0; i--) {
                for (int r : at.getOrDefault(q.poll(), List.of())) {
                    if (used[r]) continue;
                    used[r] = true;
                    for (int s : routes[r]) {
                        if (s == target) return buses;
                        if (seen.add(s)) q.add(s);
                    }
                }
            }
        }
        return -1;
    }
}
