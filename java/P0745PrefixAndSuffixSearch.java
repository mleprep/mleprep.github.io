// 745. Prefix and Suffix Search (Hard)
// https://leetcode.com/problems/prefix-and-suffix-search/
// Run (Java 11+): java P0745PrefixAndSuffixSearch.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0745PrefixAndSuffixSearch {
    public static void main(String[] args) {
        WordFilter wf = new WordFilter(new String[]{"apple", "ape", "maple"});
        check(wf.f("a", "e"), 1);
        check(wf.f("", "le"), 2);
        check(wf.f("b", ""), -1);
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

class WordFilter {
    private final Map<String, Integer> idx = new HashMap<>();

    public WordFilter(String[] words) {
        for (int i = 0; i < words.length; i++) {
            String w = words[i];
            for (int p = 0; p <= w.length(); p++)
                for (int s = 0; s <= w.length(); s++) idx.put(w.substring(0, p) + "#" + w.substring(s), i);
        }
    }

    public int f(String pref, String suff) { return idx.getOrDefault(pref + "#" + suff, -1); }
}
