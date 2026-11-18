// 71. Simplify Path (Medium)
// https://leetcode.com/problems/simplify-path/
// Run (Java 11+): java P0071SimplifyPath.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0071SimplifyPath {
    public static void main(String[] args) {
        check(new Solution().simplifyPath("/a/./b/../../c/"), "/c");
        check(new Solution().simplifyPath("/../"), "/");
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
    public String simplifyPath(String path) {
        Deque<String> st = new ArrayDeque<>();
        for (String p : path.split("/")) {
            if (p.equals("..")) { if (!st.isEmpty()) st.pollLast(); }
            else if (!p.isEmpty() && !p.equals(".")) st.addLast(p);
        }
        return "/" + String.join("/", st);
    }
}
