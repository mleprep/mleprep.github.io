// 278. First Bad Version (Easy)
// https://leetcode.com/problems/first-bad-version/
// Run (Java 11+): java P0278FirstBadVersion.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0278FirstBadVersion {
    public static void main(String[] args) {
        VersionControl.bad = 1702766719;
        check(new Solution().firstBadVersion(2126753390), 1702766719);
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

class VersionControl { // provided by LeetCode
    static int bad;
    boolean isBadVersion(int v) { return v >= bad; }
}

class Solution extends VersionControl {
    public int firstBadVersion(int n) {
        int lo = 1, hi = n;
        while (lo < hi) {
            int mid = lo + (hi - lo) / 2;      // avoids int overflow for n near 2^31
            if (isBadVersion(mid)) hi = mid; else lo = mid + 1;
        }
        return lo;
    }
}
