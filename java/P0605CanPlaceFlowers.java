// 605. Can Place Flowers (Easy)
// https://leetcode.com/problems/can-place-flowers/
// Run (Java 11+): java P0605CanPlaceFlowers.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0605CanPlaceFlowers {
    public static void main(String[] args) {
        check(new Solution().canPlaceFlowers(new int[]{1, 0, 0, 0, 1}, 1), true);
        check(new Solution().canPlaceFlowers(new int[]{1, 0, 0, 0, 1}, 2), false);
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
    public boolean canPlaceFlowers(int[] flowerbed, int n) {
        int len = flowerbed.length;
        for (int i = 0; i < len && n > 0; i++) {
            boolean leftFree = i == 0 || flowerbed[i - 1] == 0, rightFree = i == len - 1 || flowerbed[i + 1] == 0;
            if (flowerbed[i] == 0 && leftFree && rightFree) { flowerbed[i] = 1; n--; }
        }
        return n <= 0;
    }
}
