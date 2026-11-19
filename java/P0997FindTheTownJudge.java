// 997. Find the Town Judge (Easy)
// https://leetcode.com/problems/find-the-town-judge/
// Run (Java 11+): java P0997FindTheTownJudge.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0997FindTheTownJudge {
    public static void main(String[] args) {
        check(new Solution().findJudge(3, new int[][]{{1, 3}, {2, 3}}), 3);
        check(new Solution().findJudge(3, new int[][]{{1, 3}, {2, 3}, {3, 1}}), -1);
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
    public int findJudge(int n, int[][] trust) {
        int[] score = new int[n + 1];
        for (int[] t : trust) { score[t[0]]--; score[t[1]]++; }
        for (int p = 1; p <= n; p++) if (score[p] == n - 1) return p;
        return -1;
    }
}
