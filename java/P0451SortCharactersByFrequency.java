// 451. Sort Characters By Frequency (Medium)
// https://leetcode.com/problems/sort-characters-by-frequency/
// Run (Java 11+): java P0451SortCharactersByFrequency.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0451SortCharactersByFrequency {
    public static void main(String[] args) {
        String r = new Solution().frequencySort("tree");
        check(r.equals("eert") || r.equals("eetr"), true);
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
    public String frequencySort(String s) {
        Map<Character, Integer> cnt = new HashMap<>();
        for (char c : s.toCharArray()) cnt.merge(c, 1, Integer::sum);
        List<Character>[] bucket = new List[s.length() + 1];
        for (Map.Entry<Character, Integer> e : cnt.entrySet()) {
            if (bucket[e.getValue()] == null) bucket[e.getValue()] = new ArrayList<>();
            bucket[e.getValue()].add(e.getKey());
        }
        StringBuilder sb = new StringBuilder();
        for (int f = s.length(); f > 0; f--)
            if (bucket[f] != null) for (char c : bucket[f]) sb.append(String.valueOf(c).repeat(f));
        return sb.toString();
    }
}
