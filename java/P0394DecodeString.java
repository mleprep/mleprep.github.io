// 394. Decode String (Medium)
// https://leetcode.com/problems/decode-string/
// Run (Java 11+): java P0394DecodeString.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0394DecodeString {
    public static void main(String[] args) {
        check(new Solution().decodeString("3[a2[c]]"), "accaccacc");
        check(new Solution().decodeString("2[abc]3[cd]ef"), "abcabccdcdcdef");
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
    public String decodeString(String s) {
        Deque<StringBuilder> strs = new ArrayDeque<>();
        Deque<Integer> counts = new ArrayDeque<>();
        StringBuilder cur = new StringBuilder();
        int k = 0;
        for (char c : s.toCharArray()) {
            if (Character.isDigit(c)) k = k * 10 + (c - '0');
            else if (c == '[') { strs.push(cur); counts.push(k); cur = new StringBuilder(); k = 0; }
            else if (c == ']') {
                StringBuilder prev = strs.pop();
                prev.append(cur.toString().repeat(counts.pop()));
                cur = prev;
            } else cur.append(c);
        }
        return cur.toString();
    }
}
