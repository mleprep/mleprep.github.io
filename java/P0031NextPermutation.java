// 31. Next Permutation (Medium)
// https://leetcode.com/problems/next-permutation/
// Run (Java 11+): java P0031NextPermutation.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0031NextPermutation {
    public static void main(String[] args) {
        int[] a = {1, 3, 5, 4, 2};
        new Solution().nextPermutation(a);
        check(a, new int[]{1, 4, 2, 3, 5});
        int[] b = {3, 2, 1};
        new Solution().nextPermutation(b);
        check(b, new int[]{1, 2, 3});
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
    public void nextPermutation(int[] nums) {
        int i = nums.length - 2;
        while (i >= 0 && nums[i] >= nums[i + 1]) i--;
        if (i >= 0) {
            int j = nums.length - 1;
            while (nums[j] <= nums[i]) j--;
            swap(nums, i, j);
        }
        for (int l = i + 1, r = nums.length - 1; l < r; l++, r--) swap(nums, l, r);
    }

    private void swap(int[] a, int i, int j) { int t = a[i]; a[i] = a[j]; a[j] = t; }
}
