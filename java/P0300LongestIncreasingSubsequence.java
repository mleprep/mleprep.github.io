// 300. Longest Increasing Subsequence (Medium)
// https://leetcode.com/problems/longest-increasing-subsequence/
// Run (Java 11+): java P0300LongestIncreasingSubsequence.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0300LongestIncreasingSubsequence {
    public static void main(String[] args) {
        check(new Solution().lengthOfLIS(new int[]{10, 9, 2, 5, 3, 7, 101, 18}), 4);
        check(new Solution().lengthOfLIS(new int[]{7, 7, 7, 7}), 1);
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
    public int lengthOfLIS(int[] nums) {
        int[] tails = new int[nums.length];
        int len = 0;
        for (int x : nums) {
            int lo = 0, hi = len;
            while (lo < hi) {
                int mid = (lo + hi) / 2;
                if (tails[mid] >= x) hi = mid; else lo = mid + 1;
            }
            tails[lo] = x;
            if (lo == len) len++;
        }
        return len;
    }
}
