// 213. House Robber II (Medium)
// https://leetcode.com/problems/house-robber-ii/
// Run (Java 11+): java P0213HouseRobberIi.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0213HouseRobberIi {
    public static void main(String[] args) {
        check(new Solution().rob(new int[]{2, 3, 2}), 3);
        check(new Solution().rob(new int[]{1, 2, 3, 1}), 4);
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
    public int rob(int[] nums) {
        int n = nums.length;
        if (n == 1) return nums[0];
        return Math.max(line(nums, 0, n - 2), line(nums, 1, n - 1));
    }

    private int line(int[] a, int lo, int hi) {
        int prev = 0, cur = 0;
        for (int i = lo; i <= hi; i++) { int nxt = Math.max(cur, prev + a[i]); prev = cur; cur = nxt; }
        return cur;
    }
}
