// 523. Continuous Subarray Sum (Medium)
// https://leetcode.com/problems/continuous-subarray-sum/
// Run (Java 11+): java P0523ContinuousSubarraySum.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0523ContinuousSubarraySum {
    public static void main(String[] args) {
        check(new Solution().checkSubarraySum(new int[]{23, 2, 4, 6, 7}, 6), true);
        check(new Solution().checkSubarraySum(new int[]{23, 2, 6, 4, 7}, 13), false);
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
    public boolean checkSubarraySum(int[] nums, int k) {
        Map<Integer, Integer> first = new HashMap<>();
        first.put(0, -1);
        int pre = 0;
        for (int i = 0; i < nums.length; i++) {
            pre = (pre + nums[i]) % k;
            Integer j = first.putIfAbsent(pre, i);
            if (j != null && i - j >= 2) return true;
        }
        return false;
    }
}
