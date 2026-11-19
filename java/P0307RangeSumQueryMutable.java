// 307. Range Sum Query - Mutable (Medium)
// https://leetcode.com/problems/range-sum-query-mutable/
// Run (Java 11+): java P0307RangeSumQueryMutable.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0307RangeSumQueryMutable {
    public static void main(String[] args) {
        NumArray na = new NumArray(new int[]{1, 3, 5});
        check(na.sumRange(0, 2), 9);
        na.update(1, 2);
        check(na.sumRange(0, 2), 8);
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

class NumArray {
    private final int[] t, nums;
    private final int n;

    public NumArray(int[] nums) {
        n = nums.length;
        this.nums = nums.clone();
        t = new int[n + 1];
        for (int i = 0; i < n; i++) add(i + 1, nums[i]);
    }

    private void add(int i, int v) { for (; i <= n; i += i & -i) t[i] += v; }

    private int pre(int i) { int s = 0; for (; i > 0; i -= i & -i) s += t[i]; return s; }

    public void update(int index, int val) { add(index + 1, val - nums[index]); nums[index] = val; }

    public int sumRange(int left, int right) { return pre(right + 1) - pre(left); }
}
