// 219. Contains Duplicate II (Easy)
// https://leetcode.com/problems/contains-duplicate-ii/
// Run (Java 11+): java P0219ContainsDuplicateIi.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0219ContainsDuplicateIi {
    public static void main(String[] args) {
        check(new Solution().containsNearbyDuplicate(new int[]{1, 0, 1, 1}, 1), true);
        check(new Solution().containsNearbyDuplicate(new int[]{1, 2, 3, 1, 2, 3}, 2), false);
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
    public boolean containsNearbyDuplicate(int[] nums, int k) {
        Set<Integer> window = new HashSet<>();
        for (int i = 0; i < nums.length; i++) {
            if (!window.add(nums[i])) return true;
            if (window.size() > k) window.remove(nums[i - k]);
        }
        return false;
    }
}
