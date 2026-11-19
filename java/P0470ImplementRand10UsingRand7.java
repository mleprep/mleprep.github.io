// 470. Implement Rand10() Using Rand7() (Medium)
// https://leetcode.com/problems/implement-rand10-using-rand7/
// Run (Java 11+): java P0470ImplementRand10UsingRand7.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0470ImplementRand10UsingRand7 {
    public static void main(String[] args) {
        Solution s = new Solution();
        int[] cnt = new int[11];
        for (int i = 0; i < 20000; i++) cnt[s.rand10()]++;
        boolean ok = true;
        for (int d = 1; d <= 10; d++) ok &= cnt[d] > 1800 && cnt[d] < 2200;
        check(ok, true);
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

class SolBase { // provided by LeetCode
    private final Random rnd = new Random(7);
    int rand7() { return rnd.nextInt(7) + 1; }
}

class Solution extends SolBase {
    public int rand10() {
        while (true) {
            int v = (rand7() - 1) * 7 + rand7();
            if (v <= 40) return (v - 1) % 10 + 1;
        }
    }
}
