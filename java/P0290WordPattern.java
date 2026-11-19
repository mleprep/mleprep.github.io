// 290. Word Pattern (Easy)
// https://leetcode.com/problems/word-pattern/
// Run (Java 11+): java P0290WordPattern.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0290WordPattern {
    public static void main(String[] args) {
        check(new Solution().wordPattern("abba", "dog cat cat dog"), true);
        check(new Solution().wordPattern("abba", "dog dog dog dog"), false);
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
    public boolean wordPattern(String pattern, String s) {
        String[] words = s.split(" ");
        if (words.length != pattern.length()) return false;
        Map<Character, String> p2w = new HashMap<>();
        Map<String, Character> w2p = new HashMap<>();
        for (int i = 0; i < words.length; i++) {
            char p = pattern.charAt(i);
            String w = words[i];
            if (!p2w.computeIfAbsent(p, x -> w).equals(w)) return false;
            if (w2p.computeIfAbsent(w, x -> p).charValue() != p) return false;
        }
        return true;
    }
}
