// 216. Combination Sum III (Medium)
// https://leetcode.com/problems/combination-sum-iii/
// Run (Java 11+): java P0216CombinationSumIii.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0216CombinationSumIii {
    public static void main(String[] args) {
        check(new Solution().combinationSum3(3, 9), "[[1, 2, 6], [1, 3, 5], [2, 3, 4]]");
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
    public List<List<Integer>> combinationSum3(int k, int n) {
        List<List<Integer>> res = new ArrayList<>();
        bt(k, 1, n, new ArrayList<>(), res);
        return res;
    }

    private void bt(int k, int start, int rest, List<Integer> path, List<List<Integer>> res) {
        if (path.size() == k) { if (rest == 0) res.add(new ArrayList<>(path)); return; }
        for (int d = start; d <= 9 && d <= rest; d++) {
            path.add(d);
            bt(k, d + 1, rest - d, path, res);
            path.remove(path.size() - 1);
        }
    }
}
