// 217. Contains Duplicate (Easy)
// https://leetcode.com/problems/contains-duplicate/
// Run (Java 11+): java P0217ContainsDuplicate.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0217ContainsDuplicate {
    public static void main(String[] args) {
        check(new Solution().containsDuplicate(new int[]{1, 2, 3, 1}), true);
        check(new Solution().containsDuplicate(new int[]{1, 2, 3, 4}), false);
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
    public boolean containsDuplicate(int[] nums) {
        Set<Integer> seen = new HashSet<>();
        for (int x : nums) if (!seen.add(x)) return true; // add() returns false if present
        return false;
    }
}
