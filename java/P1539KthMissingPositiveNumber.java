// 1539. Kth Missing Positive Number (Easy)
// https://leetcode.com/problems/kth-missing-positive-number/
// Run (Java 11+): java P1539KthMissingPositiveNumber.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1539KthMissingPositiveNumber {
    public static void main(String[] args) {
        check(new Solution().findKthPositive(new int[]{2, 3, 4, 7, 11}, 5), 9);
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
    public int findKthPositive(int[] arr, int k) {
        int lo = 0, hi = arr.length;
        while (lo < hi) {
            int mid = (lo + hi) / 2;
            if (arr[mid] - (mid + 1) >= k) hi = mid; else lo = mid + 1;
        }
        return lo + k;
    }
}
