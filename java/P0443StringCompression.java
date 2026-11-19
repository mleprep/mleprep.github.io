// 443. String Compression (Medium)
// https://leetcode.com/problems/string-compression/
// Run (Java 11+): java P0443StringCompression.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0443StringCompression {
    public static void main(String[] args) {
        char[] a = {'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'};
        int k = new Solution().compress(a);
        check(new String(a, 0, k), "ab12");
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
    public int compress(char[] chars) {
        int w = 0;
        for (int i = 0, j; i < chars.length; i = j) {
            for (j = i; j < chars.length && chars[j] == chars[i]; j++) ;
            chars[w++] = chars[i];
            if (j - i > 1) for (char d : String.valueOf(j - i).toCharArray()) chars[w++] = d;
        }
        return w;
    }
}
