// 448. Find All Numbers Disappeared in an Array (Easy)
// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
// Run (Java 11+): java P0448FindAllNumbersDisappearedInAnArray.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0448FindAllNumbersDisappearedInAnArray {
    public static void main(String[] args) {
        check(new Solution().findDisappearedNumbers(new int[]{4, 3, 2, 7, 8, 2, 3, 1}), "[5, 6]");
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
    public List<Integer> findDisappearedNumbers(int[] nums) {
        for (int v : nums) {
            int i = Math.abs(v) - 1;
            nums[i] = -Math.abs(nums[i]);
        }
        List<Integer> res = new ArrayList<>();
        for (int i = 0; i < nums.length; i++) if (nums[i] > 0) res.add(i + 1);
        return res;
    }
}
