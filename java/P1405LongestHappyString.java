// 1405. Longest Happy String (Medium)
// https://leetcode.com/problems/longest-happy-string/
// Run (Java 11+): java P1405LongestHappyString.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1405LongestHappyString {
    public static void main(String[] args) {
        String hs = new Solution().longestDiverseString(1, 1, 7);
        check(hs.length() == 8 && !hs.contains("ccc"), true);
        check(new Solution().longestDiverseString(7, 1, 0).length(), 5);
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
    public String longestDiverseString(int a, int b, int c) {
        PriorityQueue<int[]> pq = new PriorityQueue<>((x, y) -> y[0] - x[0]);   // {count, letter}
        if (a > 0) pq.add(new int[]{a, 'a'});
        if (b > 0) pq.add(new int[]{b, 'b'});
        if (c > 0) pq.add(new int[]{c, 'c'});
        StringBuilder sb = new StringBuilder();
        while (!pq.isEmpty()) {
            int[] top = pq.poll();
            int len = sb.length();
            if (len >= 2 && sb.charAt(len - 1) == top[1] && sb.charAt(len - 2) == top[1]) {
                if (pq.isEmpty()) break;
                int[] second = pq.poll();
                sb.append((char) second[1]);
                if (--second[0] > 0) pq.add(second);
                pq.add(top);
            } else {
                sb.append((char) top[1]);
                if (--top[0] > 0) pq.add(top);
            }
        }
        return sb.toString();
    }
}
