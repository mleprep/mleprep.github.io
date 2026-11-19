// 1493. Longest Subarray of 1's After Deleting One Element (Medium)
// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/
// Run (Java 11+): java P1493LongestSubarrayOf1sAfterDeletingOneElement.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1493LongestSubarrayOf1sAfterDeletingOneElement {
    public static void main(String[] args) {
        check(new Solution().longestSubarray(new int[]{0, 1, 1, 1, 0, 1, 1, 0, 1}), 5);
        check(new Solution().longestSubarray(new int[]{1, 1, 1}), 2);
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
    public int longestSubarray(int[] nums) {
        int lo = 0, zeros = 0, best = 0;
        for (int hi = 0; hi < nums.length; hi++) {
            if (nums[hi] == 0) zeros++;
            while (zeros > 1) if (nums[lo++] == 0) zeros--;
            best = Math.max(best, hi - lo);
        }
        return best;
    }
}
