// 438. Find All Anagrams in a String (Medium)
// https://leetcode.com/problems/find-all-anagrams-in-a-string/
// Run (Java 11+): java P0438FindAllAnagramsInAString.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0438FindAllAnagramsInAString {
    public static void main(String[] args) {
        check(new Solution().findAnagrams("cbaebabacd", "abc"), "[0, 6]");
        check(new Solution().findAnagrams("abab", "ab"), "[0, 1, 2]");
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
    public List<Integer> findAnagrams(String s, String p) {
        int k = p.length();
        int[] need = new int[26], win = new int[26];
        for (char c : p.toCharArray()) need[c - 'a']++;
        List<Integer> res = new ArrayList<>();
        for (int r = 0; r < s.length(); r++) {
            win[s.charAt(r) - 'a']++;
            if (r >= k) win[s.charAt(r - k) - 'a']--;
            if (Arrays.equals(win, need)) res.add(r - k + 1);
        }
        return res;
    }
}
