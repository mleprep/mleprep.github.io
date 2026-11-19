// 1871. Jump Game VII (Medium)
// https://leetcode.com/problems/jump-game-vii/
// Run (Java 11+): java P1871JumpGameVii.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1871JumpGameVii {
    public static void main(String[] args) {
        check(new Solution().canReach("011010", 2, 3), true);
        check(new Solution().canReach("01101110", 2, 3), false);
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
    public boolean canReach(String s, int minJump, int maxJump) {
        int n = s.length(), window = 0;
        boolean[] ok = new boolean[n];
        ok[0] = true;
        for (int i = 1; i < n; i++) {
            if (i >= minJump && ok[i - minJump]) window++;
            if (i > maxJump && ok[i - maxJump - 1]) window--;
            ok[i] = window > 0 && s.charAt(i) == '0';
        }
        return ok[n - 1];
    }
}
