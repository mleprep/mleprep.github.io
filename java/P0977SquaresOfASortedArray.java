// 977. Squares of a Sorted Array (Easy)
// https://leetcode.com/problems/squares-of-a-sorted-array/
// Run (Java 11+): java P0977SquaresOfASortedArray.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0977SquaresOfASortedArray {
    public static void main(String[] args) {
        check(new Solution().sortedSquares(new int[]{-7, -3, 2, 3, 11}), new int[]{4, 9, 9, 49, 121});
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
    public int[] sortedSquares(int[] nums) {
        int n = nums.length, l = 0, r = n - 1;
        int[] res = new int[n];
        for (int k = n - 1; k >= 0; k--) {
            if (Math.abs(nums[l]) > Math.abs(nums[r])) { res[k] = nums[l] * nums[l]; l++; }
            else { res[k] = nums[r] * nums[r]; r--; }
        }
        return res;
    }
}
