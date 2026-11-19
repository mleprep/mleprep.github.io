// 1095. Find in Mountain Array (Hard)
// https://leetcode.com/problems/find-in-mountain-array/
// Run (Java 11+): java P1095FindInMountainArray.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1095FindInMountainArray {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5, 3, 1};
        MountainArray ma = new MountainArray() {
            public int get(int i) { return arr[i]; }
            public int length() { return arr.length; }
        };
        check(new Solution().findInMountainArray(3, ma), 2);
        check(new Solution().findInMountainArray(6, ma), -1);
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

interface MountainArray { // provided by LeetCode
    int get(int index);
    int length();
}

class Solution {
    public int findInMountainArray(int target, MountainArray m) {
        int n = m.length(), lo = 0, hi = n - 1;
        while (lo < hi) {
            int mid = (lo + hi) / 2;
            if (m.get(mid) < m.get(mid + 1)) lo = mid + 1; else hi = mid;
        }
        int peak = lo, i = search(m, target, 0, peak, true);
        return i != -1 ? i : search(m, target, peak + 1, n - 1, false);
    }

    private int search(MountainArray m, int target, int lo, int hi, boolean asc) {
        while (lo <= hi) {
            int mid = (lo + hi) / 2, v = m.get(mid);
            if (v == target) return mid;
            if ((v < target) == asc) lo = mid + 1; else hi = mid - 1;
        }
        return -1;
    }
}
