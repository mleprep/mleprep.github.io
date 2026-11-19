// 336. Palindrome Pairs (Hard)
// https://leetcode.com/problems/palindrome-pairs/
// Run (Java 11+): java P0336PalindromePairs.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0336PalindromePairs {
    public static void main(String[] args) {
        List<List<Integer>> pp = new ArrayList<>(new Solution().palindromePairs(new String[]{"abcd", "dcba", "lls", "s", "sssll"}));
        pp.sort(Comparator.comparing((List<Integer> x) -> x.get(0)).thenComparing(x -> x.get(1)));
        check(pp, List.of(List.of(0, 1), List.of(1, 0), List.of(2, 4), List.of(3, 2)));
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
    public List<List<Integer>> palindromePairs(String[] words) {
        Map<String, Integer> rev = new HashMap<>();
        for (int i = 0; i < words.length; i++) rev.put(new StringBuilder(words[i]).reverse().toString(), i);
        List<List<Integer>> res = new ArrayList<>();
        for (int i = 0; i < words.length; i++) {
            String w = words[i];
            for (int j = 0; j <= w.length(); j++) {
                String pre = w.substring(0, j), suf = w.substring(j);
                Integer k = rev.get(pre);
                if (k != null && k != i && isPal(suf)) res.add(List.of(i, k));
                k = rev.get(suf);
                if (j > 0 && k != null && k != i && isPal(pre)) res.add(List.of(k, i));
            }
        }
        return res;
    }

    private boolean isPal(String s) {
        for (int i = 0, j = s.length() - 1; i < j; i++, j--) if (s.charAt(i) != s.charAt(j)) return false;
        return true;
    }
}
