// 315. Count of Smaller Numbers After Self (Hard)
// https://leetcode.com/problems/count-of-smaller-numbers-after-self/
// Run (Java 11+): java P0315CountOfSmallerNumbersAfterSelf.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0315CountOfSmallerNumbersAfterSelf {
    public static void main(String[] args) {
        check(new Solution().countSmaller(new int[]{5, 2, 6, 1}), "[2, 1, 1, 0]");
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
    public List<Integer> countSmaller(int[] nums) {
        int[] sorted = Arrays.stream(nums).distinct().sorted().toArray();
        int m = sorted.length;
        int[] t = new int[m + 1];
        Integer[] res = new Integer[nums.length];
        for (int k = nums.length - 1; k >= 0; k--) {
            int r = Arrays.binarySearch(sorted, nums[k]) + 1, s = 0;
            for (int i = r - 1; i > 0; i -= i & -i) s += t[i];
            res[k] = s;
            for (int i = r; i <= m; i += i & -i) t[i]++;
        }
        return Arrays.asList(res);
    }
}
