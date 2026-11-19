// 659. Split Array into Consecutive Subsequences (Medium)
// https://leetcode.com/problems/split-array-into-consecutive-subsequences/
// Run (Java 11+): java P0659SplitArrayIntoConsecutiveSubsequences.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0659SplitArrayIntoConsecutiveSubsequences {
    public static void main(String[] args) {
        check(new Solution().isPossible(new int[]{1, 2, 3, 3, 4, 4, 5, 5}), true);
        check(new Solution().isPossible(new int[]{1, 2, 3, 4, 4, 5}), false);
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
    public boolean isPossible(int[] nums) {
        Map<Integer, Integer> left = new HashMap<>(), ends = new HashMap<>();
        for (int x : nums) left.merge(x, 1, Integer::sum);
        for (int x : nums) {
            if (left.getOrDefault(x, 0) == 0) continue;
            left.merge(x, -1, Integer::sum);
            if (ends.getOrDefault(x - 1, 0) > 0) {
                ends.merge(x - 1, -1, Integer::sum);
                ends.merge(x, 1, Integer::sum);
            } else if (left.getOrDefault(x + 1, 0) > 0 && left.getOrDefault(x + 2, 0) > 0) {
                left.merge(x + 1, -1, Integer::sum);
                left.merge(x + 2, -1, Integer::sum);
                ends.merge(x + 2, 1, Integer::sum);
            } else return false;
        }
        return true;
    }
}
