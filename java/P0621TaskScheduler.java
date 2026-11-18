// 621. Task Scheduler (Medium)
// https://leetcode.com/problems/task-scheduler/
// Run (Java 11+): java P0621TaskScheduler.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0621TaskScheduler {
    public static void main(String[] args) {
        check(new Solution().leastInterval("AAABBB".toCharArray(), 2), 8);
        check(new Solution().leastInterval("AAABBB".toCharArray(), 3), 10);
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
    public int leastInterval(char[] tasks, int n) {
        int[] cnt = new int[26];
        int maxf = 0, tied = 0;
        for (char t : tasks) maxf = Math.max(maxf, ++cnt[t - 'A']);
        for (int c : cnt) if (c == maxf) tied++;
        return Math.max(tasks.length, (maxf - 1) * (n + 1) + tied);
    }
}
