// 153. Find Minimum in Rotated Sorted Array (Medium)
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
// Run (Java 11+): java P0153FindMinimumInRotatedSortedArray.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0153FindMinimumInRotatedSortedArray {
    public static void main(String[] args) {
        check(new Solution().findMin(new int[]{4, 5, 6, 7, 0, 1, 2}), 0);
        check(new Solution().findMin(new int[]{11, 13, 15, 17}), 11);
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
    public int findMin(int[] nums) {
        int lo = 0, hi = nums.length - 1;
        while (lo < hi) {
            int mid = lo + (hi - lo) / 2;
            if (nums[mid] > nums[hi]) lo = mid + 1; else hi = mid;
        }
        return nums[lo];
    }
}
