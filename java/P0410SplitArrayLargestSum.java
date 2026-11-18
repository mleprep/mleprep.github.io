// 410. Split Array Largest Sum (Hard)
// https://leetcode.com/problems/split-array-largest-sum/
// Run (Java 11+): java P0410SplitArrayLargestSum.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0410SplitArrayLargestSum {
    public static void main(String[] args) {
        check(new Solution().splitArray(new int[]{7, 2, 5, 10, 8}, 2), 18);
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
    public int splitArray(int[] nums, int k) {
        long lo = 0, hi = 0;
        for (int x : nums) { lo = Math.max(lo, x); hi += x; }
        while (lo < hi) {
            long mid = (lo + hi) / 2;
            int cnt = 1;
            long s = 0;
            for (int x : nums) {
                if (s + x > mid) { cnt++; s = 0; }
                s += x;
            }
            if (cnt <= k) hi = mid; else lo = mid + 1;
        }
        return (int) lo;
    }
}
