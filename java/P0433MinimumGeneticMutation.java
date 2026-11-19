// 433. Minimum Genetic Mutation (Medium)
// https://leetcode.com/problems/minimum-genetic-mutation/
// Run (Java 11+): java P0433MinimumGeneticMutation.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0433MinimumGeneticMutation {
    public static void main(String[] args) {
        check(new Solution().minMutation("AACCGGTT", "AAACGGTA", new String[]{"AACCGGTA", "AACCGCTA", "AAACGGTA"}), 2);
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
    public int minMutation(String startGene, String endGene, String[] bank) {
        Set<String> valid = new HashSet<>(Arrays.asList(bank)), seen = new HashSet<>(List.of(startGene));
        Deque<String> q = new ArrayDeque<>(List.of(startGene));
        for (int steps = 0; !q.isEmpty(); steps++)
            for (int k = q.size(); k > 0; k--) {
                String g = q.poll();
                if (g.equals(endGene)) return steps;
                char[] c = g.toCharArray();
                for (int i = 0; i < c.length; i++) {
                    char orig = c[i];
                    for (char ch : "ACGT".toCharArray()) {
                        c[i] = ch;
                        String nxt = new String(c);
                        if (valid.contains(nxt) && seen.add(nxt)) q.add(nxt);
                    }
                    c[i] = orig;
                }
            }
        return -1;
    }
}
