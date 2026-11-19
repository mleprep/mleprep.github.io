// 528. Random Pick with Weight (Medium)
// https://leetcode.com/problems/random-pick-with-weight/
// Run (Java 11+): java P0528RandomPickWithWeight.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0528RandomPickWithWeight {
    public static void main(String[] args) {
        Solution s = new Solution(new int[]{1, 3});
        int ones = 0;
        for (int i = 0; i < 4000; i++) ones += s.pickIndex();
        check(ones > 2800 && ones < 3200, true);
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
    private final int[] pre;
    private final Random rnd = new Random(1);

    public Solution(int[] w) {
        pre = new int[w.length];
        for (int i = 0, s = 0; i < w.length; i++) pre[i] = s += w[i];
    }

    public int pickIndex() {
        int r = rnd.nextInt(pre[pre.length - 1]);          // r in [0, total)
        int lo = 0, hi = pre.length - 1;
        while (lo < hi) {                                   // first prefix > r
            int mid = (lo + hi) / 2;
            if (pre[mid] > r) hi = mid; else lo = mid + 1;
        }
        return lo;
    }
}
