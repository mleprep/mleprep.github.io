// 724. Find Pivot Index (Easy)
// https://leetcode.com/problems/find-pivot-index/
// Run (Java 11+): java P0724FindPivotIndex.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0724FindPivotIndex {
    public static void main(String[] args) {
        check(new Solution().pivotIndex(new int[]{1, 7, 3, 6, 5, 6}), 3);
        check(new Solution().pivotIndex(new int[]{2, 1, -1}), 0);
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
    public int pivotIndex(int[] nums) {
        int total = 0, left = 0;
        for (int x : nums) total += x;
        for (int i = 0; i < nums.length; i++) {
            if (left == total - left - nums[i]) return i;
            left += nums[i];
        }
        return -1;
    }
}
