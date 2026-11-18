// 77. Combinations (Medium)
// https://leetcode.com/problems/combinations/
// Run (Java 11+): java P0077Combinations.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0077Combinations {
    public static void main(String[] args) {
        check(new Solution().combine(4, 2), "[[1, 2], [1, 3], [1, 4], [2, 3], [2, 4], [3, 4]]");
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
    public List<List<Integer>> combine(int n, int k) {
        List<List<Integer>> res = new ArrayList<>();
        bt(n, k, 1, new ArrayList<>(), res);
        return res;
    }

    private void bt(int n, int k, int start, List<Integer> path, List<List<Integer>> res) {
        if (path.size() == k) { res.add(new ArrayList<>(path)); return; }
        for (int i = start; i <= n - (k - path.size()) + 1; i++) {
            path.add(i);
            bt(n, k, i + 1, path, res);
            path.remove(path.size() - 1);
        }
    }
}
