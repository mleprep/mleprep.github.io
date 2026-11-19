// 1472. Design Browser History (Medium)
// https://leetcode.com/problems/design-browser-history/
// Run (Java 11+): java P1472DesignBrowserHistory.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1472DesignBrowserHistory {
    public static void main(String[] args) {
        BrowserHistory b = new BrowserHistory("leetcode.com");
        b.visit("google.com"); b.visit("facebook.com"); b.visit("youtube.com");
        check(b.back(2), "google.com");
        b.visit("linkedin.com");
        check(b.forward(2), "linkedin.com");
        check(b.back(7), "leetcode.com");
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

class BrowserHistory {
    private final List<String> pages = new ArrayList<>();
    private int cur = 0, last = 0;

    public BrowserHistory(String homepage) { pages.add(homepage); }

    public void visit(String url) {
        cur++;
        if (cur == pages.size()) pages.add(url); else pages.set(cur, url);
        last = cur;
    }

    public String back(int steps) { cur = Math.max(0, cur - steps); return pages.get(cur); }
    public String forward(int steps) { cur = Math.min(last, cur + steps); return pages.get(cur); }
}
