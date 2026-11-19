// 844. Backspace String Compare (Easy)
// https://leetcode.com/problems/backspace-string-compare/
// Run (Java 11+): java P0844BackspaceStringCompare.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0844BackspaceStringCompare {
    public static void main(String[] args) {
        check(new Solution().backspaceCompare("ab##", "c#d#"), true);
        check(new Solution().backspaceCompare("a#c", "b"), false);
        check(new Solution().backspaceCompare("bxj##tw", "bxo#j##tw"), true);
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
    public boolean backspaceCompare(String s, String t) {
        int i = s.length() - 1, j = t.length() - 1;
        while (true) {
            i = next(s, i);
            j = next(t, j);
            if (i < 0 || j < 0) return i < 0 && j < 0;
            if (s.charAt(i--) != t.charAt(j--)) return false;
        }
    }

    private int next(String x, int i) {   // index of the next surviving char at or before i, or -1
        for (int skip = 0; i >= 0; i--) {
            if (x.charAt(i) == '#') skip++;
            else if (skip > 0) skip--;
            else break;
        }
        return i;
    }
}
