// 1834. Single-Threaded CPU (Medium)
// https://leetcode.com/problems/single-threaded-cpu/
// Run (Java 11+): java P1834SingleThreadedCpu.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1834SingleThreadedCpu {
    public static void main(String[] args) {
        check(new Solution().getOrder(new int[][]{{1, 2}, {2, 4}, {3, 2}, {4, 1}}), new int[]{0, 2, 3, 1});
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
    public int[] getOrder(int[][] tasks) {
        int n = tasks.length;
        Integer[] order = new Integer[n];
        for (int i = 0; i < n; i++) order[i] = i;
        Arrays.sort(order, (a, b) -> tasks[a][0] - tasks[b][0]);
        PriorityQueue<Integer> pq = new PriorityQueue<>((a, b) -> tasks[a][1] != tasks[b][1] ? tasks[a][1] - tasks[b][1] : a - b);
        int[] res = new int[n];
        long time = 0;
        for (int k = 0, i = 0; k < n; k++) {
            if (pq.isEmpty() && time < tasks[order[i]][0]) time = tasks[order[i]][0];
            while (i < n && tasks[order[i]][0] <= time) pq.add(order[i++]);
            int j = pq.poll();
            time += tasks[j][1];
            res[k] = j;
        }
        return res;
    }
}
