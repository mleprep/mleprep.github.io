// 205. Isomorphic Strings (Easy)
// https://leetcode.com/problems/isomorphic-strings/
// Run (Java 11+): java P0205IsomorphicStrings.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0205IsomorphicStrings {
    public static void main(String[] args) {
        check(new Solution().isIsomorphic("egg", "add"), true);
        check(new Solution().isIsomorphic("foo", "bar"), false);
        check(new Solution().isIsomorphic("badc", "baba"), false);
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
    public boolean isIsomorphic(String s, String t) {
        int[] st = new int[256], ts = new int[256]; // store last index+1 seen
        for (int i = 0; i < s.length(); i++) {
            char a = s.charAt(i), b = t.charAt(i);
            if (st[a] != ts[b]) return false;   // must have been paired together before (or never)
            st[a] = ts[b] = i + 1;
        }
        return true;
    }
}
