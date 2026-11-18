// 767. Reorganize String (Medium)
// https://leetcode.com/problems/reorganize-string/
// Run (Java 11+): java P0767ReorganizeString.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0767ReorganizeString {
    public static void main(String[] args) {
        check(new Solution().reorganizeString("aab"), "aba");
        check(new Solution().reorganizeString("aaab"), "");
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
    public String reorganizeString(String s) {
        int[] cnt = new int[26];
        for (char c : s.toCharArray()) cnt[c - 'a']++;
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> b[1] - a[1]); // {char, count}
        for (int i = 0; i < 26; i++) if (cnt[i] > 0) pq.add(new int[]{i, cnt[i]});
        StringBuilder sb = new StringBuilder();
        int[] prev = null;
        while (!pq.isEmpty()) {
            int[] cur = pq.poll();
            sb.append((char) ('a' + cur[0]));
            if (prev != null) pq.add(prev);
            prev = --cur[1] > 0 ? cur : null;
        }
        return prev == null ? sb.toString() : "";
    }
}
