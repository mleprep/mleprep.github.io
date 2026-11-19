// 218. The Skyline Problem (Hard)
// https://leetcode.com/problems/the-skyline-problem/
// Run (Java 11+): java P0218TheSkylineProblem.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0218TheSkylineProblem {
    public static void main(String[] args) {
        int[][] B = {{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}};
        check(new Solution().getSkyline(B), "[[2, 10], [3, 15], [7, 12], [12, 0], [15, 10], [20, 8], [24, 0]]");
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
    public List<List<Integer>> getSkyline(int[][] buildings) {
        TreeSet<Integer> xs = new TreeSet<>();
        for (int[] b : buildings) { xs.add(b[0]); xs.add(b[1]); }
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> b[0] - a[0]);   // {height, right}
        pq.add(new int[]{0, Integer.MAX_VALUE});
        List<List<Integer>> res = new ArrayList<>();
        int i = 0;
        for (int x : xs) {
            while (i < buildings.length && buildings[i][0] <= x) { pq.add(new int[]{buildings[i][2], buildings[i][1]}); i++; }
            while (pq.peek()[1] <= x) pq.poll();
            int top = pq.peek()[0];
            if (res.isEmpty() || res.get(res.size() - 1).get(1) != top) res.add(List.of(x, top));
        }
        return res;
    }
}
