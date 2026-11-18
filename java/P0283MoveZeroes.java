// 283. Move Zeroes (Easy)
// https://leetcode.com/problems/move-zeroes/
// Run (Java 11+): java P0283MoveZeroes.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0283MoveZeroes {
    public static void main(String[] args) {
        int[] a = {0, 1, 0, 3, 12};
        new Solution().moveZeroes(a);
        check(a, new int[]{1, 3, 12, 0, 0});
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
    public void moveZeroes(int[] nums) {
        int k = 0;
        for (int i = 0; i < nums.length; i++)
            if (nums[i] != 0) { int t = nums[k]; nums[k++] = nums[i]; nums[i] = t; }
    }
}
