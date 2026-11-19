// 456. 132 Pattern (Medium)
// https://leetcode.com/problems/132-pattern/
// Run (Java 11+): java P0456132Pattern.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0456132Pattern {
    public static void main(String[] args) {
        check(new Solution().find132pattern(new int[]{3, 1, 4, 2}), true);
        check(new Solution().find132pattern(new int[]{1, 2, 3, 4}), false);
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
    public boolean find132pattern(int[] nums) {
        Deque<Integer> st = new ArrayDeque<>();
        long two = Long.MIN_VALUE;
        for (int i = nums.length - 1; i >= 0; i--) {
            if (nums[i] < two) return true;
            while (!st.isEmpty() && st.peek() < nums[i]) two = st.pop();
            st.push(nums[i]);
        }
        return false;
    }
}
