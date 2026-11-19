// 480. Sliding Window Median (Hard)
// https://leetcode.com/problems/sliding-window-median/
// Run (Java 11+): java P0480SlidingWindowMedian.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0480SlidingWindowMedian {
    public static void main(String[] args) {
        check(new Solution().medianSlidingWindow(new int[]{1, 3, -1, -3, 5, 3, 6, 7}, 3), new double[]{1.0, -1.0, -1.0, 3.0, 5.0, 6.0});
        check(new Solution().medianSlidingWindow(new int[]{1, 4, 2, 3}, 4), new double[]{2.5});
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
    public double[] medianSlidingWindow(int[] nums, int k) {
        // indices ordered by (value, index) so duplicates are distinct; lo holds the smaller half (size ceil(k/2))
        Comparator<Integer> cmp = (a, b) -> nums[a] != nums[b] ? Integer.compare(nums[a], nums[b]) : a - b;
        TreeSet<Integer> lo = new TreeSet<>(cmp), hi = new TreeSet<>(cmp);
        double[] res = new double[nums.length - k + 1];
        for (int i = 0; i < nums.length; i++) {
            if (i >= k && !lo.remove(i - k)) hi.remove(i - k);
            lo.add(i);
            hi.add(lo.pollLast());                       // keep every lo <= every hi
            if (hi.size() > lo.size()) lo.add(hi.pollFirst());
            if (i >= k - 1)
                res[i - k + 1] = k % 2 == 1 ? nums[lo.last()] : ((double) nums[lo.last()] + nums[hi.first()]) / 2;
        }
        return res;
    }
}
