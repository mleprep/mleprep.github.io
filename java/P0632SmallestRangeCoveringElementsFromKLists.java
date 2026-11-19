// 632. Smallest Range Covering Elements from K Lists (Hard)
// https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/
// Run (Java 11+): java P0632SmallestRangeCoveringElementsFromKLists.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0632SmallestRangeCoveringElementsFromKLists {
    public static void main(String[] args) {
        check(new Solution().smallestRange(List.of(List.of(4, 10, 15, 24, 26), List.of(0, 9, 12, 20), List.of(5, 18, 22, 30))), new int[]{20, 24});
        check(new Solution().smallestRange(List.of(List.of(1, 2, 3), List.of(1, 2, 3), List.of(1, 2, 3))), new int[]{1, 1});
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
    public int[] smallestRange(List<List<Integer>> nums) {
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> Integer.compare(a[0], b[0])); // {value, list, index}
        int hi = Integer.MIN_VALUE;
        for (int i = 0; i < nums.size(); i++) {
            pq.add(new int[]{nums.get(i).get(0), i, 0});
            hi = Math.max(hi, nums.get(i).get(0));
        }
        int[] best = {pq.peek()[0], hi};
        while (true) {
            int[] top = pq.poll();
            if (hi - top[0] < best[1] - best[0]) best = new int[]{top[0], hi};
            List<Integer> row = nums.get(top[1]);
            if (top[2] + 1 == row.size()) return best;
            int nxt = row.get(top[2] + 1);
            hi = Math.max(hi, nxt);
            pq.add(new int[]{nxt, top[1], top[2] + 1});
        }
    }
}
