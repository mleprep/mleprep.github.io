// 287. Find the Duplicate Number (Medium)
// https://leetcode.com/problems/find-the-duplicate-number/
// Run (Java 11+): java P0287FindTheDuplicateNumber.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0287FindTheDuplicateNumber {
    public static void main(String[] args) {
        check(new Solution().findDuplicate(new int[]{1, 3, 4, 2, 2}), 2);
        check(new Solution().findDuplicate(new int[]{3, 1, 3, 4, 2}), 3);
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
    public int findDuplicate(int[] nums) {
        int slow = 0, fast = 0;
        do { slow = nums[slow]; fast = nums[nums[fast]]; } while (slow != fast);
        int p = 0;
        while (p != slow) { p = nums[p]; slow = nums[slow]; }
        return p;
    }
}
