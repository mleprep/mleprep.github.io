// 204. Count Primes (Medium)
// https://leetcode.com/problems/count-primes/
// Run (Java 11+): java P0204CountPrimes.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0204CountPrimes {
    public static void main(String[] args) {
        check(new Solution().countPrimes(10), 4);
        check(new Solution().countPrimes(100), 25);
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
    public int countPrimes(int n) {
        if (n < 3) return 0;
        boolean[] composite = new boolean[n];
        int count = 0;
        for (int p = 2; p < n; p++) {
            if (composite[p]) continue;
            count++;
            for (long q = (long) p * p; q < n; q += p) composite[(int) q] = true;
        }
        return count;
    }
}
