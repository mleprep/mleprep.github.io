// 229. Majority Element II (Medium)
// https://leetcode.com/problems/majority-element-ii/
// Run (Java 11+): java P0229MajorityElementIi.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0229MajorityElementIi {
    public static void main(String[] args) {
        List<Integer> me = new Solution().majorityElement(new int[]{2, 2, 1, 1, 1, 2, 2});
        Collections.sort(me);
        check(me, List.of(1, 2));
        check(new Solution().majorityElement(new int[]{3, 2, 3}), List.of(3));
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
    public List<Integer> majorityElement(int[] nums) {
        int c1 = 0, c2 = 1, n1 = 0, n2 = 0;           // distinct placeholders; verified below
        for (int x : nums) {
            if (x == c1) n1++;
            else if (x == c2) n2++;
            else if (n1 == 0) { c1 = x; n1 = 1; }
            else if (n2 == 0) { c2 = x; n2 = 1; }
            else { n1--; n2--; }
        }
        List<Integer> res = new ArrayList<>();
        for (int c : new int[]{c1, c2}) {
            int cnt = 0;
            for (int x : nums) if (x == c) cnt++;
            if (cnt > nums.length / 3) res.add(c);
        }
        return res;
    }
}
