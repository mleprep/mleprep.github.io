// 2300. Successful Pairs of Spells and Potions (Medium)
// https://leetcode.com/problems/successful-pairs-of-spells-and-potions/
// Run (Java 11+): java P2300SuccessfulPairsOfSpellsAndPotions.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P2300SuccessfulPairsOfSpellsAndPotions {
    public static void main(String[] args) {
        check(new Solution().successfulPairs(new int[]{5, 1, 3}, new int[]{1, 2, 3, 4, 5}, 7), new int[]{4, 0, 3});
        check(new Solution().successfulPairs(new int[]{100000}, new int[]{100000}, 10000000000L), new int[]{1});
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
    public int[] successfulPairs(int[] spells, int[] potions, long success) {
        Arrays.sort(potions);
        int m = potions.length;
        int[] res = new int[spells.length];
        for (int i = 0; i < spells.length; i++) {
            long need = (success + spells[i] - 1) / spells[i];
            int lo = 0, hi = m;                       // first potion >= need
            while (lo < hi) {
                int mid = (lo + hi) / 2;
                if (potions[mid] >= need) hi = mid; else lo = mid + 1;
            }
            res[i] = m - lo;
        }
        return res;
    }
}
