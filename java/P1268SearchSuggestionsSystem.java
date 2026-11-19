// 1268. Search Suggestions System (Medium)
// https://leetcode.com/problems/search-suggestions-system/
// Run (Java 11+): java P1268SearchSuggestionsSystem.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1268SearchSuggestionsSystem {
    public static void main(String[] args) {
        check(new Solution().suggestedProducts(new String[]{"havana"}, "tatiana"), "[[], [], [], [], [], [], []]");
        check(new Solution().suggestedProducts(new String[]{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mou").get(2), "[mouse, mousepad]");
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
    public List<List<String>> suggestedProducts(String[] products, String searchWord) {
        Arrays.sort(products);
        List<List<String>> res = new ArrayList<>();
        String prefix = "";
        for (char c : searchWord.toCharArray()) {
            prefix += c;
            int lo = 0, hi = products.length;
            while (lo < hi) {
                int mid = (lo + hi) / 2;
                if (products[mid].compareTo(prefix) >= 0) hi = mid; else lo = mid + 1;
            }
            List<String> cur = new ArrayList<>();
            for (int i = lo; i < Math.min(lo + 3, products.length) && products[i].startsWith(prefix); i++) cur.add(products[i]);
            res.add(cur);
        }
        return res;
    }
}
