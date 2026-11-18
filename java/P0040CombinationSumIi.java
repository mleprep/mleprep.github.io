// 40. Combination Sum II (Medium)
// https://leetcode.com/problems/combination-sum-ii/
// Run (Java 11+): java P0040CombinationSumIi.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0040CombinationSumIi {
    public static void main(String[] args) {
        check(new Solution().combinationSum2(new int[]{10, 1, 2, 7, 6, 1, 5}, 8), "[[1, 1, 6], [1, 2, 5], [1, 7], [2, 6]]");
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
    public List<List<Integer>> combinationSum2(int[] candidates, int target) {
        Arrays.sort(candidates);
        List<List<Integer>> res = new ArrayList<>();
        bt(candidates, 0, target, new ArrayList<>(), res);
        return res;
    }

    private void bt(int[] c, int start, int rest, List<Integer> path, List<List<Integer>> res) {
        if (rest == 0) { res.add(new ArrayList<>(path)); return; }
        for (int i = start; i < c.length && c[i] <= rest; i++) {
            if (i > start && c[i] == c[i - 1]) continue;
            path.add(c[i]);
            bt(c, i + 1, rest - c[i], path, res);
            path.remove(path.size() - 1);
        }
    }
}
