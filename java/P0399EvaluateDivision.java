// 399. Evaluate Division (Medium)
// https://leetcode.com/problems/evaluate-division/
// Run (Java 11+): java P0399EvaluateDivision.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0399EvaluateDivision {
    public static void main(String[] args) {
        List<List<String>> eq = List.of(List.of("a", "b"), List.of("b", "c"));
        List<List<String>> qs = List.of(List.of("a", "c"), List.of("b", "a"), List.of("a", "e"), List.of("a", "a"), List.of("x", "x"));
        check(new Solution().calcEquation(eq, new double[]{2.0, 3.0}, qs), new double[]{6.0, 0.5, -1.0, 1.0, -1.0});
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
    public double[] calcEquation(List<List<String>> equations, double[] values, List<List<String>> queries) {
        Map<String, Map<String, Double>> g = new HashMap<>();
        for (int i = 0; i < values.length; i++) {
            String a = equations.get(i).get(0), b = equations.get(i).get(1);
            g.computeIfAbsent(a, k -> new HashMap<>()).put(b, values[i]);
            g.computeIfAbsent(b, k -> new HashMap<>()).put(a, 1 / values[i]);
        }
        double[] res = new double[queries.size()];
        for (int i = 0; i < res.length; i++) {
            String x = queries.get(i).get(0), y = queries.get(i).get(1);
            res[i] = g.containsKey(x) && g.containsKey(y) ? dfs(g, x, y, new HashSet<>()) : -1.0;
        }
        return res;
    }

    private double dfs(Map<String, Map<String, Double>> g, String x, String y, Set<String> seen) {
        if (x.equals(y)) return 1.0;
        seen.add(x);
        for (Map.Entry<String, Double> e : g.get(x).entrySet()) {
            if (seen.contains(e.getKey())) continue;
            double r = dfs(g, e.getKey(), y, seen);
            if (r != -1.0) return e.getValue() * r;
        }
        return -1.0;
    }
}
