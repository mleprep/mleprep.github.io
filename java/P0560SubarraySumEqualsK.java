// 560. Subarray Sum Equals K (Medium)
// https://leetcode.com/problems/subarray-sum-equals-k/
// Run (Java 11+): java P0560SubarraySumEqualsK.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0560SubarraySumEqualsK {
    public static void main(String[] args) {
        check(new Solution().subarraySum(new int[]{1, 1, 1}, 2), 2);
        check(new Solution().subarraySum(new int[]{1, 2, 3}, 3), 2);
        check(new Solution().subarraySum(new int[]{1, -1, 0}, 0), 3);
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
    public int subarraySum(int[] nums, int k) {
        Map<Integer, Integer> seen = new HashMap<>();
        seen.put(0, 1);
        int count = 0, pre = 0;
        for (int x : nums) {
            pre += x;
            count += seen.getOrDefault(pre - k, 0);
            seen.merge(pre, 1, Integer::sum);
        }
        return count;
    }
}
