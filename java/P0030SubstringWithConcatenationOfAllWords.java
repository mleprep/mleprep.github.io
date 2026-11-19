// 30. Substring with Concatenation of All Words (Hard)
// https://leetcode.com/problems/substring-with-concatenation-of-all-words/
// Run (Java 11+): java P0030SubstringWithConcatenationOfAllWords.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0030SubstringWithConcatenationOfAllWords {
    public static void main(String[] args) {
        check(new Solution().findSubstring("barfoofoobarthefoobarman", new String[]{"bar", "foo", "the"}), "[6, 9, 12]");
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
    public List<Integer> findSubstring(String s, String[] words) {
        int L = words[0].length(), k = words.length;
        Map<String, Integer> need = new HashMap<>();
        for (String w : words) need.merge(w, 1, Integer::sum);
        List<Integer> res = new ArrayList<>();
        for (int off = 0; off < L; off++) {
            Map<String, Integer> cnt = new HashMap<>();
            int l = off, used = 0;
            for (int r = off; r + L <= s.length(); r += L) {
                String w = s.substring(r, r + L);
                if (!need.containsKey(w)) { cnt.clear(); used = 0; l = r + L; continue; }
                cnt.merge(w, 1, Integer::sum);
                used++;
                while (cnt.get(w) > need.get(w)) {
                    cnt.merge(s.substring(l, l + L), -1, Integer::sum);
                    used--;
                    l += L;
                }
                if (used == k) res.add(l);
            }
        }
        Collections.sort(res);
        return res;
    }
}
