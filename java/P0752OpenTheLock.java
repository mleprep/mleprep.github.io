// 752. Open the Lock (Medium)
// https://leetcode.com/problems/open-the-lock/
// Run (Java 11+): java P0752OpenTheLock.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0752OpenTheLock {
    public static void main(String[] args) {
        check(new Solution().openLock(new String[]{"0201", "0101", "0102", "1212", "2002"}, "0202"), 6);
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
    public int openLock(String[] deadends, String target) {
        Set<String> seen = new HashSet<>(Arrays.asList(deadends));
        if (!seen.add("0000")) return -1;
        Deque<String> q = new ArrayDeque<>(List.of("0000"));
        for (int steps = 0; !q.isEmpty(); steps++) {
            for (int k = q.size(); k > 0; k--) {
                String cur = q.poll();
                if (cur.equals(target)) return steps;
                char[] c = cur.toCharArray();
                for (int i = 0; i < 4; i++) {
                    char orig = c[i];
                    for (int delta : new int[]{1, 9}) {
                        c[i] = (char) ('0' + (orig - '0' + delta) % 10);
                        String nxt = new String(c);
                        if (seen.add(nxt)) q.add(nxt);
                    }
                    c[i] = orig;
                }
            }
        }
        return -1;
    }
}
