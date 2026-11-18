// 378. Kth Smallest Element in a Sorted Matrix (Medium)
// https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/
// Run (Java 11+): java P0378KthSmallestElementInASortedMatrix.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0378KthSmallestElementInASortedMatrix {
    public static void main(String[] args) {
        check(new Solution().kthSmallest(new int[][]{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8), 13);
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
    public int kthSmallest(int[][] matrix, int k) {
        int n = matrix.length;
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[0] - b[0]); // {val, r, c}
        for (int r = 0; r < Math.min(n, k); r++) pq.add(new int[]{matrix[r][0], r, 0});
        for (int i = 0; i < k - 1; i++) {
            int[] t = pq.poll();
            if (t[2] + 1 < n) pq.add(new int[]{matrix[t[1]][t[2] + 1], t[1], t[2] + 1});
        }
        return pq.peek()[0];
    }
}
