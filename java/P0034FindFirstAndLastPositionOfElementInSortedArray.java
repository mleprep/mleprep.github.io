// 34. Find First and Last Position of Element in Sorted Array (Medium)
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// Run (Java 11+): java P0034FindFirstAndLastPositionOfElementInSortedArray.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0034FindFirstAndLastPositionOfElementInSortedArray {
    public static void main(String[] args) {
        check(new Solution().searchRange(new int[]{5, 7, 7, 8, 8, 10}, 8), new int[]{3, 4});
        check(new Solution().searchRange(new int[]{5, 7, 7, 8, 8, 10}, 6), new int[]{-1, -1});
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
    public int[] searchRange(int[] nums, int target) {
        int a = lower(nums, target);
        if (a == nums.length || nums[a] != target) return new int[]{-1, -1};
        return new int[]{a, lower(nums, target + 1) - 1};
    }

    private int lower(int[] nums, int t) {
        int lo = 0, hi = nums.length;
        while (lo < hi) {
            int mid = lo + (hi - lo) / 2;
            if (nums[mid] >= t) hi = mid; else lo = mid + 1;
        }
        return lo;
    }
}
