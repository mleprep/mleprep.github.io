// 334. Increasing Triplet Subsequence (Medium)
// https://leetcode.com/problems/increasing-triplet-subsequence/
// Run (Java 11+): java P0334IncreasingTripletSubsequence.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0334IncreasingTripletSubsequence {
    public static void main(String[] args) {
        check(new Solution().increasingTriplet(new int[]{2, 1, 5, 0, 4, 6}), true);
        check(new Solution().increasingTriplet(new int[]{5, 4, 3, 2, 1}), false);
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
    public boolean increasingTriplet(int[] nums) {
        int first = Integer.MAX_VALUE, second = Integer.MAX_VALUE;
        for (int x : nums) {
            if (x <= first) first = x;
            else if (x <= second) second = x;
            else return true;
        }
        return false;
    }
}
