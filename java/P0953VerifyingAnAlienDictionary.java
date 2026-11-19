// 953. Verifying an Alien Dictionary (Easy)
// https://leetcode.com/problems/verifying-an-alien-dictionary/
// Run (Java 11+): java P0953VerifyingAnAlienDictionary.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0953VerifyingAnAlienDictionary {
    public static void main(String[] args) {
        check(new Solution().isAlienSorted(new String[]{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"), true);
        check(new Solution().isAlienSorted(new String[]{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"), false);
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
    public boolean isAlienSorted(String[] words, String order) {
        int[] rank = new int[26];
        for (int i = 0; i < 26; i++) rank[order.charAt(i) - 'a'] = i;
        for (int w = 0; w + 1 < words.length; w++) {
            String a = words[w], b = words[w + 1];
            int i = 0;
            while (i < a.length() && i < b.length() && a.charAt(i) == b.charAt(i)) i++;
            if (i == a.length() || i == b.length()) { if (a.length() > b.length()) return false; }
            else if (rank[a.charAt(i) - 'a'] > rank[b.charAt(i) - 'a']) return false;
        }
        return true;
    }
}
