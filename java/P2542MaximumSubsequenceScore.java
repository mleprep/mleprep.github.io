// 2542. Maximum Subsequence Score (Medium)
// https://leetcode.com/problems/maximum-subsequence-score/
// Run (Java 11+): java P2542MaximumSubsequenceScore.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P2542MaximumSubsequenceScore {
    public static void main(String[] args) {
        check(new Solution().maxScore(new int[]{1, 3, 3, 2}, new int[]{2, 1, 3, 4}, 3), 12L);
        check(new Solution().maxScore(new int[]{4, 2, 3, 1, 1}, new int[]{7, 5, 10, 9, 6}, 1), 30L);
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
    public long maxScore(int[] nums1, int[] nums2, int k) {
        int n = nums1.length;
        Integer[] idx = new Integer[n];
        for (int i = 0; i < n; i++) idx[i] = i;
        Arrays.sort(idx, (x, y) -> nums2[y] - nums2[x]);
        PriorityQueue<Integer> heap = new PriorityQueue<>();
        long total = 0, best = 0;
        for (int i : idx) {
            heap.add(nums1[i]);
            total += nums1[i];
            if (heap.size() > k) total -= heap.poll();
            if (heap.size() == k) best = Math.max(best, total * nums2[i]);
        }
        return best;
    }
}
