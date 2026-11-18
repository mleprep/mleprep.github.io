// 167. Two Sum II - Input Array Is Sorted (Medium)
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
// Run (Java 11+): java P0167TwoSumIiInputArrayIsSorted.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0167TwoSumIiInputArrayIsSorted {
    public static void main(String[] args) {
        check(new Solution().twoSum(new int[]{2, 7, 11, 15}, 9), new int[]{1, 2});
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
    public int[] twoSum(int[] numbers, int target) {
        int l = 0, r = numbers.length - 1;
        while (l < r) {
            int s = numbers[l] + numbers[r];
            if (s == target) return new int[]{l + 1, r + 1};
            if (s < target) l++; else r--;
        }
        return new int[0];
    }
}
