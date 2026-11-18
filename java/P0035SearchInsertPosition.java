// 35. Search Insert Position (Easy)
// https://leetcode.com/problems/search-insert-position/
// Run (Java 11+): java P0035SearchInsertPosition.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0035SearchInsertPosition {
    public static void main(String[] args) {
        check(new Solution().searchInsert(new int[]{1, 3, 5, 6}, 2), 1);
        check(new Solution().searchInsert(new int[]{1, 3, 5, 6}, 7), 4);
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
    public int searchInsert(int[] nums, int target) {
        int lo = 0, hi = nums.length;
        while (lo < hi) {
            int mid = lo + (hi - lo) / 2;
            if (nums[mid] >= target) hi = mid; else lo = mid + 1;
        }
        return lo;
    }
}
