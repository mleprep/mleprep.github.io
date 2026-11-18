// 269. Alien Dictionary (Hard)
// https://leetcode.com/problems/alien-dictionary/
// Run (Java 11+): java P0269AlienDictionary.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0269AlienDictionary {
    public static void main(String[] args) {
        check(new Solution().alienOrder(new String[]{"wrt", "wrf", "er", "ett", "rftt"}), "wertf");
        check(new Solution().alienOrder(new String[]{"z", "x", "z"}), "");
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
    public String alienOrder(String[] words) {
        Map<Character, Set<Character>> g = new TreeMap<>();
        Map<Character, Integer> indeg = new TreeMap<>();
        for (String w : words) for (char c : w.toCharArray()) { g.putIfAbsent(c, new TreeSet<>()); indeg.putIfAbsent(c, 0); }
        for (int i = 0; i + 1 < words.length; i++) {
            String a = words[i], b = words[i + 1];
            int j = 0;
            while (j < Math.min(a.length(), b.length()) && a.charAt(j) == b.charAt(j)) j++;
            if (j == Math.min(a.length(), b.length())) {
                if (a.length() > b.length()) return "";
                continue;
            }
            if (g.get(a.charAt(j)).add(b.charAt(j))) indeg.merge(b.charAt(j), 1, Integer::sum);
        }
        Deque<Character> q = new ArrayDeque<>();
        for (Map.Entry<Character, Integer> e : indeg.entrySet()) if (e.getValue() == 0) q.add(e.getKey());
        StringBuilder sb = new StringBuilder();
        while (!q.isEmpty()) {
            char u = q.poll();
            sb.append(u);
            for (char v : g.get(u)) if (indeg.merge(v, -1, Integer::sum) == 0) q.add(v);
        }
        return sb.length() == g.size() ? sb.toString() : "";
    }
}
