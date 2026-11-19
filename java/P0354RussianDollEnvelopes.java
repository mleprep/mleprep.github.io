// 354. Russian Doll Envelopes (Hard)
// https://leetcode.com/problems/russian-doll-envelopes/
// Run (Java 11+): java P0354RussianDollEnvelopes.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0354RussianDollEnvelopes {
    public static void main(String[] args) {
        check(new Solution().maxEnvelopes(new int[][]{{5, 4}, {6, 4}, {6, 7}, {2, 3}}), 3);
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
    public int maxEnvelopes(int[][] envelopes) {
        Arrays.sort(envelopes, (a, b) -> a[0] != b[0] ? a[0] - b[0] : b[1] - a[1]);
        int[] tails = new int[envelopes.length];
        int len = 0;
        for (int[] e : envelopes) {
            int i = Arrays.binarySearch(tails, 0, len, e[1]);
            if (i < 0) i = -i - 1;
            tails[i] = e[1];
            if (i == len) len++;
        }
        return len;
    }
}
