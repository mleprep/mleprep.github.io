// 1679. Max Number of K-Sum Pairs (Medium)
// https://leetcode.com/problems/max-number-of-k-sum-pairs/
// Run (Java 11+): java P1679MaxNumberOfKSumPairs.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1679MaxNumberOfKSumPairs {
    public static void main(String[] args) {
        check(new Solution().maxOperations(new int[]{1, 2, 3, 4}, 5), 2);
        check(new Solution().maxOperations(new int[]{3, 1, 3, 4, 3}, 6), 1);
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
    public int maxOperations(int[] nums, int k) {
        Map<Integer, Integer> waiting = new HashMap<>();
        int ops = 0;
        for (int x : nums) {
            int c = waiting.getOrDefault(k - x, 0);
            if (c > 0) { waiting.put(k - x, c - 1); ops++; }
            else waiting.merge(x, 1, Integer::sum);
        }
        return ops;
    }
}
