// 493. Reverse Pairs (Hard)
// https://leetcode.com/problems/reverse-pairs/
// Run (Java 11+): java P0493ReversePairs.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0493ReversePairs {
    public static void main(String[] args) {
        check(new Solution().reversePairs(new int[]{2, 4, 3, 5, 1}), 3);
        check(new Solution().reversePairs(new int[]{2147483647, 2147483647, 2147483647}), 0);
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
    public int reversePairs(int[] nums) { return sort(nums, 0, nums.length - 1, new int[nums.length]); }

    private int sort(int[] a, int lo, int hi, int[] tmp) {
        if (lo >= hi) return 0;
        int mid = (lo + hi) / 2;
        int count = sort(a, lo, mid, tmp) + sort(a, mid + 1, hi, tmp);
        for (int i = lo, j = mid + 1; i <= mid; i++) {
            while (j <= hi && a[i] > 2L * a[j]) j++;
            count += j - (mid + 1);
        }
        int i = lo, j = mid + 1, k = lo;
        while (i <= mid && j <= hi) tmp[k++] = a[i] <= a[j] ? a[i++] : a[j++];
        while (i <= mid) tmp[k++] = a[i++];
        while (j <= hi) tmp[k++] = a[j++];
        System.arraycopy(tmp, lo, a, lo, hi - lo + 1);
        return count;
    }
}
