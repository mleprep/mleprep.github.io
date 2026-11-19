// 228. Summary Ranges (Easy)
// https://leetcode.com/problems/summary-ranges/
// Run (Java 11+): java P0228SummaryRanges.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0228SummaryRanges {
    public static void main(String[] args) {
        check(new Solution().summaryRanges(new int[]{0, 1, 2, 4, 5, 7}), "[0->2, 4->5, 7]");
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
    public List<String> summaryRanges(int[] nums) {
        List<String> res = new ArrayList<>();
        for (int i = 0, j; i < nums.length; i = j + 1) {
            j = i;
            while (j + 1 < nums.length && nums[j + 1] == nums[j] + 1) j++;
            res.add(i == j ? String.valueOf(nums[i]) : nums[i] + "->" + nums[j]);
        }
        return res;
    }
}
