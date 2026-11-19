// 80. Remove Duplicates from Sorted Array II (Medium)
// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
// Run (Java 11+): java P0080RemoveDuplicatesFromSortedArrayIi.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0080RemoveDuplicatesFromSortedArrayIi {
    public static void main(String[] args) {
        int[] a = {1, 1, 1, 2, 2, 3};
        int k = new Solution().removeDuplicates(a);
        check(Arrays.copyOf(a, k), new int[]{1, 1, 2, 2, 3});
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
    public int removeDuplicates(int[] nums) {
        int k = 0;
        for (int x : nums) if (k < 2 || x != nums[k - 2]) nums[k++] = x;
        return k;
    }
}
