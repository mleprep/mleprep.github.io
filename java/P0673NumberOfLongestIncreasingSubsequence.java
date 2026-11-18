// 673. Number of Longest Increasing Subsequence (Medium)
// https://leetcode.com/problems/number-of-longest-increasing-subsequence/
// Run (Java 11+): java P0673NumberOfLongestIncreasingSubsequence.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0673NumberOfLongestIncreasingSubsequence {
    public static void main(String[] args) {
        check(new Solution().findNumberOfLIS(new int[]{1, 3, 5, 4, 7}), 2);
        check(new Solution().findNumberOfLIS(new int[]{2, 2, 2, 2, 2}), 5);
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
    public int findNumberOfLIS(int[] nums) {
        int n = nums.length, best = 0, res = 0;
        int[] len = new int[n], cnt = new int[n];
        for (int i = 0; i < n; i++) {
            len[i] = cnt[i] = 1;
            for (int j = 0; j < i; j++) {
                if (nums[j] >= nums[i]) continue;
                if (len[j] + 1 > len[i]) { len[i] = len[j] + 1; cnt[i] = cnt[j]; }
                else if (len[j] + 1 == len[i]) cnt[i] += cnt[j];
            }
            if (len[i] > best) { best = len[i]; res = cnt[i]; }
            else if (len[i] == best) res += cnt[i];
        }
        return res;
    }
}
