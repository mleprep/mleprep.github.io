// 753. Cracking the Safe (Hard)
// https://leetcode.com/problems/cracking-the-safe/
// Run (Java 11+): java P0753CrackingTheSafe.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0753CrackingTheSafe {
    public static void main(String[] args) {
        String s = new Solution().crackSafe(3, 3);
        Set<String> all = new HashSet<>();
        for (int i = 0; i + 3 <= s.length(); i++) all.add(s.substring(i, i + 3));
        check(all.size() == 27 && s.length() == 29, true);
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
    private final Set<String> seen = new HashSet<>();
    private final StringBuilder out = new StringBuilder();

    public String crackSafe(int n, int k) {
        String start = "0".repeat(n - 1);
        dfs(start, k);
        return out.append(start).toString();
    }

    private void dfs(String node, int k) {
        for (int d = 0; d < k; d++) {
            String edge = node + d;
            if (seen.add(edge)) {
                dfs(edge.substring(1), k);
                out.append(d);
            }
        }
    }
}
