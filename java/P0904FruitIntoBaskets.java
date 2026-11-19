// 904. Fruit Into Baskets (Medium)
// https://leetcode.com/problems/fruit-into-baskets/
// Run (Java 11+): java P0904FruitIntoBaskets.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0904FruitIntoBaskets {
    public static void main(String[] args) {
        check(new Solution().totalFruit(new int[]{1, 2, 3, 2, 2}), 4);
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
    public int totalFruit(int[] fruits) {
        Map<Integer, Integer> cnt = new HashMap<>();
        int l = 0, best = 0;
        for (int r = 0; r < fruits.length; r++) {
            cnt.merge(fruits[r], 1, Integer::sum);
            while (cnt.size() > 2) {
                if (cnt.merge(fruits[l], -1, Integer::sum) == 0) cnt.remove(fruits[l]);
                l++;
            }
            best = Math.max(best, r - l + 1);
        }
        return best;
    }
}
