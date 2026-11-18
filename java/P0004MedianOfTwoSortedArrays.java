// 4. Median of Two Sorted Arrays (Hard)
// https://leetcode.com/problems/median-of-two-sorted-arrays/
// Run (Java 11+): java P0004MedianOfTwoSortedArrays.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0004MedianOfTwoSortedArrays {
    public static void main(String[] args) {
        check(new Solution().findMedianSortedArrays(new int[]{1, 3}, new int[]{2}), 2.0);
        check(new Solution().findMedianSortedArrays(new int[]{1, 2}, new int[]{3, 4}), 2.5);
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
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int[] A = nums1.length <= nums2.length ? nums1 : nums2, B = A == nums1 ? nums2 : nums1;
        int m = A.length, n = B.length, half = (m + n + 1) / 2, lo = 0, hi = m;
        while (true) {
            int i = (lo + hi) / 2, j = half - i;
            int al = i > 0 ? A[i - 1] : Integer.MIN_VALUE, ar = i < m ? A[i] : Integer.MAX_VALUE;
            int bl = j > 0 ? B[j - 1] : Integer.MIN_VALUE, br = j < n ? B[j] : Integer.MAX_VALUE;
            if (al <= br && bl <= ar) {
                if ((m + n) % 2 == 1) return Math.max(al, bl);
                return (Math.max(al, bl) + (double) Math.min(ar, br)) / 2;
            }
            if (al > br) hi = i - 1; else lo = i + 1;
        }
    }
}
