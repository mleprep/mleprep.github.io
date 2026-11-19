// 274. H-Index (Medium)
// https://leetcode.com/problems/h-index/
// Run (Java 11+): java P0274HIndex.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0274HIndex {
    public static void main(String[] args) {
        check(new Solution().hIndex(new int[]{3, 0, 6, 1, 5}), 3);
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
    public int hIndex(int[] citations) {
        int n = citations.length;
        int[] cnt = new int[n + 1];
        for (int c : citations) cnt[Math.min(c, n)]++;
        for (int h = n, papers = 0; h >= 0; h--) {
            papers += cnt[h];
            if (papers >= h) return h;
        }
        return 0;
    }
}
