// 992. Subarrays with K Different Integers (Hard)
// https://leetcode.com/problems/subarrays-with-k-different-integers/
// Run (Java 11+): java P0992SubarraysWithKDifferentIntegers.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0992SubarraysWithKDifferentIntegers {
    public static void main(String[] args) {
        check(new Solution().subarraysWithKDistinct(new int[]{1, 2, 1, 2, 3}, 2), 7);
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
    public int subarraysWithKDistinct(int[] nums, int k) {
        return atMost(nums, k) - atMost(nums, k - 1);
    }

    private int atMost(int[] nums, int k) {
        int[] cnt = new int[nums.length + 1];
        int l = 0, distinct = 0, res = 0;
        for (int r = 0; r < nums.length; r++) {
            if (cnt[nums[r]]++ == 0) distinct++;
            while (distinct > k) if (--cnt[nums[l++]] == 0) distinct--;
            res += r - l + 1;
        }
        return res;
    }
}
