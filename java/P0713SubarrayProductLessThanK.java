// 713. Subarray Product Less Than K (Medium)
// https://leetcode.com/problems/subarray-product-less-than-k/
// Run (Java 11+): java P0713SubarrayProductLessThanK.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0713SubarrayProductLessThanK {
    public static void main(String[] args) {
        check(new Solution().numSubarrayProductLessThanK(new int[]{10, 5, 2, 6}, 100), 8);
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
    public int numSubarrayProductLessThanK(int[] nums, int k) {
        if (k <= 1) return 0;
        int prod = 1, l = 0, count = 0;
        for (int r = 0; r < nums.length; r++) {
            prod *= nums[r];
            while (prod >= k) prod /= nums[l++];
            count += r - l + 1;
        }
        return count;
    }
}
