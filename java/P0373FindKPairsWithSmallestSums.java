// 373. Find K Pairs with Smallest Sums (Medium)
// https://leetcode.com/problems/find-k-pairs-with-smallest-sums/
// Run (Java 11+): java P0373FindKPairsWithSmallestSums.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0373FindKPairsWithSmallestSums {
    public static void main(String[] args) {
        check(new Solution().kSmallestPairs(new int[]{1, 7, 11}, new int[]{2, 4, 6}, 3), "[[1, 2], [1, 4], [1, 6]]");
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
    public List<List<Integer>> kSmallestPairs(int[] nums1, int[] nums2, int k) {
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[0] - b[0]); // {sum, i, j}
        for (int i = 0; i < Math.min(k, nums1.length); i++) pq.add(new int[]{nums1[i] + nums2[0], i, 0});
        List<List<Integer>> res = new ArrayList<>();
        while (!pq.isEmpty() && res.size() < k) {
            int[] t = pq.poll();
            int i = t[1], j = t[2];
            res.add(List.of(nums1[i], nums2[j]));
            if (j + 1 < nums2.length) pq.add(new int[]{nums1[i] + nums2[j + 1], i, j + 1});
        }
        return res;
    }
}
