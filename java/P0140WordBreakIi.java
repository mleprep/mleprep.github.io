// 140. Word Break II (Hard)
// https://leetcode.com/problems/word-break-ii/
// Run (Java 11+): java P0140WordBreakIi.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0140WordBreakIi {
    public static void main(String[] args) {
        List<String> r = new Solution().wordBreak("catsanddog", List.of("cat", "cats", "and", "sand", "dog"));
        Collections.sort(r);
        check(r, "[cat sand dog, cats and dog]");
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
    private final Map<Integer, List<String>> memo = new HashMap<>();

    public List<String> wordBreak(String s, List<String> wordDict) {
        return go(s, 0, new HashSet<>(wordDict));
    }

    private List<String> go(String s, int i, Set<String> words) {
        if (memo.containsKey(i)) return memo.get(i);
        List<String> out = new ArrayList<>();
        if (i == s.length()) out.add("");
        for (int j = i + 1; j <= s.length(); j++) {
            String w = s.substring(i, j);
            if (!words.contains(w)) continue;
            for (String rest : go(s, j, words)) out.add(rest.isEmpty() ? w : w + " " + rest);
        }
        memo.put(i, out);
        return out;
    }
}
