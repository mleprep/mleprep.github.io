// 1863. Sum of All Subset XOR Totals (Easy)
// https://leetcode.com/problems/sum-of-all-subset-xor-totals/
// Run (Java 11+): java P1863SumOfAllSubsetXorTotals.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1863SumOfAllSubsetXorTotals {
    public static void main(String[] args) {
        check(new Solution().subsetXORSum(new int[]{5, 1, 6}), 28);
        check(new Solution().subsetXORSum(new int[]{3, 4, 5, 6, 7, 8}), 480);
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
    public int subsetXORSum(int[] nums) {
        int or = 0;
        for (int x : nums) or |= x;
        return or << (nums.length - 1);
    }
}
