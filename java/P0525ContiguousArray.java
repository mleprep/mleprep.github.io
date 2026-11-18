// 525. Contiguous Array (Medium)
// https://leetcode.com/problems/contiguous-array/
// Run (Java 11+): java P0525ContiguousArray.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0525ContiguousArray {
    public static void main(String[] args) {
        check(new Solution().findMaxLength(new int[]{0, 1}), 2);
        check(new Solution().findMaxLength(new int[]{0, 1, 1, 0, 1, 1}), 4);
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
    public int findMaxLength(int[] nums) {
        Map<Integer, Integer> first = new HashMap<>();
        first.put(0, -1);
        int pre = 0, best = 0;
        for (int i = 0; i < nums.length; i++) {
            pre += nums[i] == 1 ? 1 : -1;
            Integer j = first.putIfAbsent(pre, i);
            if (j != null) best = Math.max(best, i - j);
        }
        return best;
    }
}
