// 344. Reverse String (Easy)
// https://leetcode.com/problems/reverse-string/
// Run (Java 11+): java P0344ReverseString.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0344ReverseString {
    public static void main(String[] args) {
        char[] a = "hello".toCharArray();
        new Solution().reverseString(a);
        check(new String(a), "olleh");
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
    public void reverseString(char[] s) {
        for (int l = 0, r = s.length - 1; l < r; l++, r--) { char t = s[l]; s[l] = s[r]; s[r] = t; }
    }
}
