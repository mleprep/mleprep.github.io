// 990. Satisfiability of Equality Equations (Medium)
// https://leetcode.com/problems/satisfiability-of-equality-equations/
// Run (Java 11+): java P0990SatisfiabilityOfEqualityEquations.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0990SatisfiabilityOfEqualityEquations {
    public static void main(String[] args) {
        check(new Solution().equationsPossible(new String[]{"a==b", "b!=a"}), false);
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
    private final int[] p = new int[26];

    private int find(int x) {
        while (p[x] != x) { p[x] = p[p[x]]; x = p[x]; }
        return x;
    }

    public boolean equationsPossible(String[] equations) {
        for (int i = 0; i < 26; i++) p[i] = i;
        for (String e : equations) if (e.charAt(1) == '=') p[find(e.charAt(0) - 'a')] = find(e.charAt(3) - 'a');
        for (String e : equations) if (e.charAt(1) == '!' && find(e.charAt(0) - 'a') == find(e.charAt(3) - 'a')) return false;
        return true;
    }
}
