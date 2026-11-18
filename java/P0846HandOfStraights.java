// 846. Hand of Straights (Medium)
// https://leetcode.com/problems/hand-of-straights/
// Run (Java 11+): java P0846HandOfStraights.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0846HandOfStraights {
    public static void main(String[] args) {
        check(new Solution().isNStraightHand(new int[]{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3), true);
        check(new Solution().isNStraightHand(new int[]{1, 2, 3, 4, 5}, 4), false);
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
    public boolean isNStraightHand(int[] hand, int groupSize) {
        if (hand.length % groupSize != 0) return false;
        TreeMap<Integer, Integer> cnt = new TreeMap<>();
        for (int x : hand) cnt.merge(x, 1, Integer::sum);
        for (int x : new ArrayList<>(cnt.keySet())) {
            int c = cnt.get(x);
            if (c == 0) continue;
            for (int y = x; y < x + groupSize; y++) {
                int have = cnt.getOrDefault(y, 0);
                if (have < c) return false;
                cnt.put(y, have - c);
            }
        }
        return true;
    }
}
